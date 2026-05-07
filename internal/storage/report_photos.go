package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"sirekap/internal/config"
)

const (
	reportPhotoPrefix       = "minio://"
	legacyReportPhotoPrefix = "minio:"
)

type reportPhotoStorage struct {
	client         *s3.Client
	bucket         string
	useObjectStore bool
}

var defaultReportPhotoStorage = &reportPhotoStorage{}

func Init(cfg config.StorageConfig) error {
	defaultReportPhotoStorage = &reportPhotoStorage{}

	if strings.TrimSpace(cfg.EndpointURL) == "" ||
		strings.TrimSpace(cfg.AccessKeyID) == "" ||
		strings.TrimSpace(cfg.SecretAccessKey) == "" ||
		strings.TrimSpace(cfg.Bucket) == "" {
		return nil
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithRegion(cfg.Region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		),
	)
	if err != nil {
		return fmt.Errorf("load object storage config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = cfg.UsePathStyle
		o.BaseEndpoint = &cfg.EndpointURL
	})

	defaultReportPhotoStorage = &reportPhotoStorage{
		client:         client,
		bucket:         cfg.Bucket,
		useObjectStore: true,
	}
	return nil
}

func SaveReportPhoto(content []byte, ext string) (string, error) {
	return defaultReportPhotoStorage.saveReportPhoto(content, ext)
}

func DeleteReportPhoto(ref string) error {
	return defaultReportPhotoStorage.deleteReportPhoto(ref)
}

func LoadReportPhoto(ref string) ([]byte, string, error) {
	return defaultReportPhotoStorage.loadReportPhoto(ref)
}

func (s *reportPhotoStorage) saveReportPhoto(content []byte, ext string) (string, error) {
	if !s.useObjectStore {
		return saveReportPhotoLocally(content, ext)
	}

	key := path.Join(
		"laporan-foto",
		time.Now().Format("2006/01/02"),
		fmt.Sprintf("lapor_%d%s", time.Now().UnixNano(), normalizeImageExt(ext)),
	)
	contentType := contentTypeFromExt(ext)

	_, err := s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      &s.bucket,
		Key:         &key,
		Body:        bytes.NewReader(content),
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}

	return formatObjectStorageRef(s.bucket, key), nil
}

func (s *reportPhotoStorage) deleteReportPhoto(ref string) error {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return nil
	}

	bucket, key, isObjectRef := parseObjectStorageRef(ref, s.bucket)
	if !isObjectRef {
		localPath := resolveLegacyLocalPath(ref)
		if localPath == "" {
			return nil
		}
		if err := os.Remove(localPath); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return nil
	}

	if !s.useObjectStore {
		return nil
	}

	key = strings.TrimSpace(key)
	if strings.TrimSpace(key) == "" {
		return nil
	}

	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	return err
}

func (s *reportPhotoStorage) loadReportPhoto(ref string) ([]byte, string, error) {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return nil, "", os.ErrNotExist
	}

	bucket, key, isObjectRef := parseObjectStorageRef(ref, s.bucket)
	if !isObjectRef {
		localPath := resolveLegacyLocalPath(ref)
		if localPath == "" {
			return nil, "", os.ErrNotExist
		}

		content, err := os.ReadFile(localPath)
		if err != nil {
			return nil, "", err
		}
		return content, contentTypeFromExt(filepath.Ext(localPath)), nil
	}

	if !s.useObjectStore {
		return nil, "", os.ErrNotExist
	}

	key = strings.TrimSpace(key)
	output, err := s.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, "", err
	}
	defer output.Body.Close()

	content, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, "", err
	}

	contentType := contentTypeFromExt(filepath.Ext(key))
	if output.ContentType != nil && strings.TrimSpace(*output.ContentType) != "" {
		contentType = strings.TrimSpace(*output.ContentType)
	}

	return content, contentType, nil
}

func saveReportPhotoLocally(content []byte, ext string) (string, error) {
	if err := os.MkdirAll("static/uploads", 0755); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("lapor_translok_%d%s", time.Now().UnixNano(), normalizeImageExt(ext))
	destination := filepath.Join("static", "uploads", filename)
	if err := os.WriteFile(destination, content, 0644); err != nil {
		return "", err
	}

	return "/" + filepath.ToSlash(destination), nil
}

func resolveLegacyLocalPath(ref string) string {
	candidates := make([]string, 0, 5)
	if filepath.IsAbs(ref) {
		candidates = append(candidates, ref, strings.TrimPrefix(ref, "/"))
	} else {
		candidates = append(candidates,
			ref,
			filepath.Join("static", "uploads", filepath.Base(ref)),
			filepath.Join("sirekap2", "dashboard", ref),
			filepath.Join("dashboard", ref),
			filepath.Join("uploads", ref),
		)
	}

	for _, candidate := range candidates {
		candidate = strings.TrimSpace(candidate)
		if candidate == "" {
			continue
		}
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	return ""
}

func normalizeImageExt(ext string) string {
	ext = strings.ToLower(strings.TrimSpace(ext))
	switch ext {
	case ".jpeg":
		return ".jpg"
	case ".jpg", ".png", ".gif":
		return ext
	default:
		return ".jpg"
	}
}

func contentTypeFromExt(ext string) string {
	switch strings.ToLower(strings.TrimSpace(ext)) {
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "image/jpeg"
	}
}

func formatObjectStorageRef(bucket, key string) string {
	bucket = strings.Trim(strings.TrimSpace(bucket), "/")
	key = strings.Trim(strings.TrimSpace(key), "/")
	if bucket == "" || key == "" {
		return legacyReportPhotoPrefix + key
	}
	return reportPhotoPrefix + bucket + "/" + key
}

func parseObjectStorageRef(ref, defaultBucket string) (string, string, bool) {
	ref = strings.TrimSpace(ref)
	defaultBucket = strings.Trim(strings.TrimSpace(defaultBucket), "/")

	if strings.HasPrefix(ref, reportPhotoPrefix) {
		trimmed := strings.TrimPrefix(ref, reportPhotoPrefix)
		parts := strings.SplitN(strings.Trim(trimmed, "/"), "/", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return "", "", false
		}
		return parts[0], parts[1], true
	}

	if strings.HasPrefix(ref, legacyReportPhotoPrefix) {
		key := strings.Trim(strings.TrimPrefix(ref, legacyReportPhotoPrefix), "/")
		if defaultBucket == "" || key == "" {
			return "", "", false
		}
		return defaultBucket, key, true
	}

	return "", "", false
}
