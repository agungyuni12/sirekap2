package handlers

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"sirekap/internal/models"
	"sirekap/internal/storage"
)

type docxImageReplacement struct {
	Source  string
	WidthPx int
}

type docxImageAsset struct {
	Placeholder string
	Bytes       []byte
	RelID       string
	Target      string
	Ext         string
	WidthEMU    int64
	HeightEMU   int64
	DocPrID     int
}

func DownloadMatriksTransportReportHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	data, err := models.GetMatriksTransportPrintData(id)
	if err != nil {
		handlePrintLoadError(w, err, "laporan translok")
		return
	}

	replacements := map[string]string{
		"nama":      docxText(data.Base.Nama),
		"namagelar": docxText(firstNonEmpty(data.Base.NamaGelar, data.Base.Nama)),
		"tanggal":   docxText(formatTanggalIndonesia(data.Base.TanggalMulai)),
		"kec":       docxText(strings.TrimPrefix(data.Base.Tujuan, "Kecamatan ")),
		"mak":       docxText(data.Base.Mak),
		"perihal":   docxText(data.Base.PerihalFP),
		"nipnik":    docxText(getNipLabel(data.Base.Golongan)),
		"nip":       docxText(getPrintableNIP(data.Base.NIP, data.Base.Golongan)),
		"laporan":   docxMultilineText(data.LaporanText),
		"foto":      "",
	}

	images := buildDocxImages(map[string]docxImageReplacement{
		"foto": {
			Source:  data.LaporanFoto,
			WidthPx: 300,
		},
	})

	docBytes, err := replaceInDocxWithImages("template/template_translok.docx", replacements, images)
	if err != nil {
		log.Printf("Error generating transport report: %v", err)
		http.Error(w, "Gagal membuat dokumen laporan translok", http.StatusInternalServerError)
		return
	}

	serveDocxDownload(w, r, docBytes, fmt.Sprintf("Laporan_Translok_%s.docx", sanitizeFilenamePart(data.Base.Nama)))
}

func DownloadMatriksSPPDReportHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	data, err := models.GetMatriksSPPDPrintData(id)
	if err != nil {
		handlePrintLoadError(w, err, "laporan SPPD")
		return
	}

	replacements := map[string]string{
		"nama":        docxText(data.Base.Nama),
		"namagelar":   docxText(firstNonEmpty(data.Base.NamaGelar, data.Base.Nama)),
		"nip":         docxText(data.Base.NIP),
		"tanggal":     docxText(formatTanggalIndonesia(data.Base.TanggalMulai)),
		"kec":         docxText(strings.TrimPrefix(data.Base.Tujuan, "Kecamatan ")),
		"mak":         docxText(data.Base.Mak),
		"perihal":     docxText(data.Base.PerihalFP),
		"brgkt_m":     docxText(data.WaktuBerangkatMulai),
		"brgkt_s":     docxText(data.WaktuBerangkatSelesai),
		"text_brgkt":  docxMultilineText(data.IsianBerangkat),
		"peng1_m":     docxText(data.WaktuPengawasan1Mulai),
		"peng1_s":     docxText(data.WaktuPengawasan1Selesai),
		"text_peng1":  docxMultilineText(data.IsianPengawasan1),
		"isoma_m":     docxText(data.WaktuIsomaMulai),
		"isoma_s":     docxText(data.WaktuIsomaSelesai),
		"peng2_m":     docxText(data.WaktuPengawasan2Mulai),
		"peng2_s":     docxText(data.WaktuPengawasan2Selesai),
		"text_peng2":  docxMultilineText(data.IsianPengawasan2),
		"pulang_m":    docxText(data.WaktuPulangMulai),
		"pulang_s":    docxText(data.WaktuPulangSelesai),
		"text_pulang": docxMultilineText(data.IsianPulang),
		"foto":        "",
	}

	images := buildDocxImages(map[string]docxImageReplacement{
		"foto": {
			Source:  data.LaporanFoto,
			WidthPx: 450,
		},
	})

	docBytes, err := replaceInDocxWithImages("template/template_laporan_perjadin.docx", replacements, images)
	if err != nil {
		log.Printf("Error generating SPPD report: %v", err)
		http.Error(w, "Gagal membuat dokumen laporan SPPD", http.StatusInternalServerError)
		return
	}

	serveDocxDownload(w, r, docBytes, fmt.Sprintf("Laporan_SPPD_%s.docx", sanitizeFilenamePart(data.Base.Nama)))
}

func DownloadMatriksPernyataanKendaraanHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	data, err := models.GetMatriksPrintBase(id)
	if err != nil {
		handlePrintLoadError(w, err, "surat pernyataan")
		return
	}

	replacements := map[string]string{
		"nama":      docxText(data.Nama),
		"nip":       docxText(data.NIP),
		"gol":       docxText(getGolonganLengkap(data.Golongan)),
		"jabatan":   docxText(firstNonEmpty(data.Jabatan, "Mitra BPS")),
		"no_surtug": docxText(data.NoSurtug),
		"perihal":   docxText(data.Kegiatan),
		"tujuan":    docxText(strings.TrimPrefix(strings.TrimPrefix(data.Tujuan, "Kabupaten "), "Kecamatan ")),
		"tgl_mulai": docxText(formatTanggalIndonesia(data.TanggalMulai)),
		"tgl_skrg":  docxText(formatTanggalIndonesia(time.Now())),
	}

	docBytes, err := replaceInDocxWithImages("template/Template_Surat Pernyataan_tidak_menggunakan_kendaraan_dinas.docx", replacements, nil)
	if err != nil {
		log.Printf("Error generating kendaraan statement: %v", err)
		http.Error(w, "Gagal membuat surat pernyataan", http.StatusInternalServerError)
		return
	}

	serveDocxDownload(w, r, docBytes, fmt.Sprintf("Pernyataan_Kendaraan_%s.docx", sanitizeFilenamePart(data.Nama)))
}

func DownloadMatriksKwitansiTranslokHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	data, err := models.GetMatriksPrintBase(id)
	if err != nil {
		handlePrintLoadError(w, err, "kwitansi translok")
		return
	}

	tanggalKwitansi, ok := parseKwitansiDate(w, r)
	if !ok {
		return
	}

	replacements := map[string]string{
		"tahun":       docxText(strconv.Itoa(data.TanggalMulai.Year())),
		"mak":         docxText(data.Mak),
		"tgl_mulai":   docxText(formatTanggalIndonesia(data.TanggalMulai)),
		"tgl_selesai": docxText(formatTanggalIndonesia(data.TanggalSelesai)),
		"asal":        docxText(data.Asal),
		"tujuan":      docxText(data.Tujuan),
		"no_surtug":   docxText(data.NoSurtug),
		"tgl_surtug":  docxText(formatTanggalIndonesia(data.TglSurtugUtama)),
		"tgl_skrg":    docxText(formatTanggalIndonesia(tanggalKwitansi)),
		"nama":        docxText(data.Nama),
		"nipnik":      docxText(getNipLabel(data.Golongan)),
		"nip":         docxText(data.NIP),
		"no_spd":      docxText(data.NoSpd),
		"tgl_spd":     docxText(formatTanggalIndonesia(data.TglSpdUtama)),
		"perihal":     docxText(data.PerihalFP),
		"jabatan":     docxText(firstNonEmpty(data.Jabatan, "Mitra BPS")),
	}

	docBytes, err := replaceInDocxWithImages("template/Template_Kwitansi_Translok.docx", replacements, nil)
	if err != nil {
		log.Printf("Error generating transport kwitansi: %v", err)
		http.Error(w, "Gagal membuat kwitansi translok", http.StatusInternalServerError)
		return
	}

	serveDocxDownload(w, r, docBytes, fmt.Sprintf("Kwitansi_Translok_%s.docx", sanitizeFilenamePart(data.Nama)))
}

func DownloadMatriksKwitansiSPPDHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDQuery(w, r)
	if !ok {
		return
	}

	data, err := models.GetMatriksPrintBase(id)
	if err != nil {
		handlePrintLoadError(w, err, "kwitansi SPPD")
		return
	}

	tanggalKwitansi, ok := parseKwitansiDate(w, r)
	if !ok {
		return
	}

	replacements := map[string]string{
		"tahun":       docxText(strconv.Itoa(data.TanggalMulai.Year())),
		"mak":         docxText(data.Mak),
		"tgl_mulai":   docxText(formatTanggalIndonesia(data.TanggalMulai)),
		"tgl_selesai": docxText(formatTanggalIndonesia(data.TanggalSelesai)),
		"asal":        docxText(data.Asal),
		"tujuan":      docxText(data.Tujuan),
		"no_surtug":   docxText(data.NoSurtug),
		"tgl_surtug":  docxText(formatTanggalIndonesia(data.TglSurtugUtama)),
		"tgl_skrg":    docxText(formatTanggalIndonesia(tanggalKwitansi)),
		"nama":        docxText(data.Nama),
		"nip":         docxText(data.NIP),
		"no_spd":      docxText(data.NoSpd),
		"tgl_spd":     docxText(formatTanggalIndonesia(data.TglSpdUtama)),
		"perihal":     docxText(data.PerihalFP),
		"jabatan":     docxText(firstNonEmpty(data.Jabatan, "Mitra BPS")),
	}

	docBytes, err := replaceInDocxWithImages("template/Template_Kwitansi_SPPD.docx", replacements, nil)
	if err != nil {
		log.Printf("Error generating SPPD kwitansi: %v", err)
		http.Error(w, "Gagal membuat kwitansi SPPD", http.StatusInternalServerError)
		return
	}

	serveDocxDownload(w, r, docBytes, fmt.Sprintf("Kwitansi_SPPD_%s.docx", sanitizeFilenamePart(data.Nama)))
}

func parseIDQuery(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(strings.TrimSpace(r.URL.Query().Get("id")))
	if err != nil || id <= 0 {
		http.Error(w, "ID matriks tidak valid", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func parseKwitansiDate(w http.ResponseWriter, r *http.Request) (time.Time, bool) {
	tanggalInput := strings.TrimSpace(r.URL.Query().Get("tgl"))
	if tanggalInput == "" {
		return time.Now(), true
	}

	tanggal, err := models.ParseMatriksDate(tanggalInput)
	if err != nil {
		http.Error(w, "Format tanggal kwitansi tidak valid", http.StatusBadRequest)
		return time.Time{}, false
	}

	return tanggal, true
}

func handlePrintLoadError(w http.ResponseWriter, err error, label string) {
	if err == sql.ErrNoRows {
		http.Error(w, strings.Title(label)+" tidak ditemukan", http.StatusNotFound)
		return
	}

	log.Printf("Error loading %s: %v", label, err)
	http.Error(w, "Gagal memuat data dokumen", http.StatusInternalServerError)
}

func serveDocxDownload(w http.ResponseWriter, r *http.Request, docBytes []byte, filename string) {
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", strconv.Itoa(len(docBytes)))
	http.ServeContent(w, r, filename, time.Now(), bytes.NewReader(docBytes))
}

func buildDocxImages(source map[string]docxImageReplacement) map[string]docxImageReplacement {
	images := make(map[string]docxImageReplacement)
	for placeholder, image := range source {
		if strings.TrimSpace(image.Source) == "" {
			continue
		}
		images[placeholder] = image
	}
	return images
}

func replaceInDocxWithImages(templatePath string, replacements map[string]string, images map[string]docxImageReplacement) ([]byte, error) {
	reader, err := zip.OpenReader(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open template: %w", err)
	}
	defer reader.Close()

	fileContents := make(map[string][]byte, len(reader.File))
	fileMethods := make(map[string]uint16, len(reader.File))
	fileOrder := make([]string, 0, len(reader.File))
	maxRelID := 0
	maxImageIndex := 0

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}

		fileContents[file.Name] = content
		fileMethods[file.Name] = file.Method
		fileOrder = append(fileOrder, file.Name)

		if file.Name == "word/_rels/document.xml.rels" {
			re := regexp.MustCompile(`Id="rId(\d+)"`)
			for _, match := range re.FindAllStringSubmatch(string(content), -1) {
				value, _ := strconv.Atoi(match[1])
				if value > maxRelID {
					maxRelID = value
				}
			}
		}

		if strings.HasPrefix(file.Name, "word/media/image") {
			re := regexp.MustCompile(`word/media/image(\d+)`)
			match := re.FindStringSubmatch(file.Name)
			if len(match) == 2 {
				value, _ := strconv.Atoi(match[1])
				if value > maxImageIndex {
					maxImageIndex = value
				}
			}
		}
	}

	for name, content := range fileContents {
		if !(strings.HasSuffix(name, ".xml") || strings.HasSuffix(name, ".rels")) {
			continue
		}

		contentStr := cleanXMLSplitPlaceholders(string(content))
		for placeholder, value := range replacements {
			if _, isImage := images[placeholder]; isImage {
				continue
			}
			contentStr = replaceDocxPlaceholder(contentStr, placeholder, value)
		}
		fileContents[name] = []byte(contentStr)
	}

	imageAssets := make([]docxImageAsset, 0, len(images))
	docPrID := 1000
	for placeholder, image := range images {
		resolvedSource := strings.TrimSpace(image.Source)
		if resolvedSource == "" {
			continue
		}

		imageBytes, widthEMU, heightEMU, ext, err := prepareDocxImage(resolvedSource, image.WidthPx)
		if err != nil {
			log.Printf("Skipping DOCX image for placeholder %q from %q: %v", placeholder, resolvedSource, err)
			continue
		}

		maxRelID++
		maxImageIndex++
		docPrID++

		imageAssets = append(imageAssets, docxImageAsset{
			Placeholder: placeholder,
			Bytes:       imageBytes,
			RelID:       fmt.Sprintf("rId%d", maxRelID),
			Target:      fmt.Sprintf("word/media/image%d%s", maxImageIndex, ext),
			Ext:         ext,
			WidthEMU:    widthEMU,
			HeightEMU:   heightEMU,
			DocPrID:     docPrID,
		})
	}

	documentXML := string(fileContents["word/document.xml"])
	for _, asset := range imageAssets {
		documentXML = replaceDocxPlaceholderParagraph(documentXML, asset.Placeholder, buildImageDrawingXML(asset))
	}
	for placeholder := range images {
		documentXML = replaceDocxPlaceholderParagraph(documentXML, placeholder, "")
	}
	fileContents["word/document.xml"] = []byte(documentXML)

	if len(imageAssets) > 0 {
		rels := string(fileContents["word/_rels/document.xml.rels"])
		for _, asset := range imageAssets {
			rels = strings.Replace(
				rels,
				"</Relationships>",
				fmt.Sprintf(`<Relationship Id="%s" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="%s"/></Relationships>`, asset.RelID, strings.TrimPrefix(asset.Target, "word/")),
				1,
			)
			fileContents[asset.Target] = asset.Bytes
			fileMethods[asset.Target] = zip.Deflate
			fileOrder = append(fileOrder, asset.Target)
		}
		fileContents["word/_rels/document.xml.rels"] = []byte(rels)

		contentTypes := string(fileContents["[Content_Types].xml"])
		for _, asset := range imageAssets {
			contentType := imageContentType(asset.Ext)
			extension := strings.TrimPrefix(asset.Ext, ".")
			defaultTag := fmt.Sprintf(`<Default Extension="%s" ContentType="%s"/>`, extension, contentType)
			if !strings.Contains(contentTypes, defaultTag) {
				contentTypes = strings.Replace(contentTypes, "</Types>", defaultTag+"</Types>", 1)
			}
		}
		fileContents["[Content_Types].xml"] = []byte(contentTypes)
	}

	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	for _, name := range fileOrder {
		content, exists := fileContents[name]
		if !exists {
			continue
		}
		header := &zip.FileHeader{Name: name, Method: fileMethods[name]}
		w, err := writer.CreateHeader(header)
		if err != nil {
			return nil, err
		}
		if _, err := w.Write(content); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func replaceDocxPlaceholder(content, placeholder, value string) string {
	re := regexp.MustCompile(`\$\{\s*` + regexp.QuoteMeta(placeholder) + `\s*\}`)
	content = re.ReplaceAllString(content, value)

	rePlain := regexp.MustCompile(`\{\s*` + regexp.QuoteMeta(placeholder) + `\s*\}`)
	return rePlain.ReplaceAllString(content, value)
}

func replaceDocxPlaceholderParagraph(content, placeholder, replacement string) string {
	re := regexp.MustCompile(`\$\{\s*` + regexp.QuoteMeta(placeholder) + `\s*\}`)
	loc := re.FindStringIndex(content)
	if loc == nil {
		return content
	}

	startP1 := strings.LastIndex(content[:loc[0]], "<w:p>")
	startP2 := strings.LastIndex(content[:loc[0]], "<w:p ")
	startP := startP1
	if startP2 > startP {
		startP = startP2
	}

	endP := strings.Index(content[loc[1]:], "</w:p>")
	if startP == -1 || endP == -1 {
		return re.ReplaceAllString(content, replacement)
	}

	endP += loc[1] + len("</w:p>")
	return content[:startP] + replacement + content[endP:]
}

func prepareDocxImage(source string, widthPx int) ([]byte, int64, int64, string, error) {
	imageBytes, contentType, err := storage.LoadReportPhoto(source)
	if err != nil {
		return nil, 0, 0, "", fmt.Errorf("failed to read image %s: %w", source, err)
	}

	cfg, _, err := image.DecodeConfig(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, 0, 0, "", fmt.Errorf("failed to inspect image %s: %w", source, err)
	}

	if widthPx <= 0 {
		widthPx = cfg.Width
	}

	heightPx := cfg.Height
	if cfg.Width > 0 && widthPx > 0 {
		heightPx = int(float64(cfg.Height) * (float64(widthPx) / float64(cfg.Width)))
	}

	ext := contentTypeToDocxExt(contentType)
	if ext == "" {
		ext = strings.ToLower(filepath.Ext(source))
	}
	if ext == ".jpg" {
		ext = ".jpeg"
	}

	return imageBytes, int64(widthPx) * 9525, int64(heightPx) * 9525, ext, nil
}

func buildImageDrawingXML(asset docxImageAsset) string {
	return fmt.Sprintf(
		`<w:p><w:r><w:drawing><wp:inline distT="0" distB="0" distL="0" distR="0" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"><wp:extent cx="%d" cy="%d"/><wp:docPr id="%d" name="%s"/><a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"><a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:nvPicPr><pic:cNvPr id="%d" name="%s"/><pic:cNvPicPr/></pic:nvPicPr><pic:blipFill><a:blip r:embed="%s" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"/><a:stretch><a:fillRect/></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="%d" cy="%d"/></a:xfrm><a:prstGeom prst="rect"><a:avLst/></a:prstGeom></pic:spPr></pic:pic></a:graphicData></a:graphic></wp:inline></w:drawing></w:r></w:p>`,
		asset.WidthEMU,
		asset.HeightEMU,
		asset.DocPrID,
		asset.Placeholder,
		asset.DocPrID,
		asset.Placeholder,
		asset.RelID,
		asset.WidthEMU,
		asset.HeightEMU,
	)
}

func imageContentType(ext string) string {
	switch ext {
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "image/jpeg"
	}
}

func formatTanggalIndonesia(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	bulan := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	return fmt.Sprintf("%02d %s %04d", t.Day(), bulan[int(t.Month())], t.Year())
}

func getGolonganLengkap(kode string) string {
	golongan := map[string]string{
		"IV/e":  "Pembina Utama (IV/e)",
		"IV/d":  "Pembina Utama Madya (IV/d)",
		"IV/c":  "Pembina Utama Muda (IV/c)",
		"IV/b":  "Pembina Tingkat I (IV/b)",
		"IV/a":  "Pembina (IV/a)",
		"III/d": "Penata Tingkat I (III/d)",
		"III/c": "Penata (III/c)",
		"III/b": "Penata Muda Tingkat I (III/b)",
		"III/a": "Penata Muda (III/a)",
		"II/d":  "Pengatur Tingkat I (II/d)",
		"II/c":  "Pengatur (II/c)",
		"II/b":  "Pengatur Muda Tingkat I (II/b)",
		"II/a":  "Pengatur Muda (II/a)",
		"I/d":   "Juru Tingkat I (I/d)",
		"I/c":   "Juru (I/c)",
		"I/b":   "Juru Muda Tingkat I (I/b)",
		"I/a":   "Juru Muda (I/a)",
	}

	if strings.TrimSpace(kode) == "" {
		return "-"
	}
	if value, ok := golongan[kode]; ok {
		return value
	}
	return kode
}

func getNipLabel(golongan string) string {
	if strings.TrimSpace(golongan) != "" {
		return "NIP."
	}
	return "NIK"
}

func getPrintableNIP(nip, golongan string) string {
	if strings.TrimSpace(golongan) == "" {
		return strings.TrimSpace(nip)
	}
	return strings.TrimSpace(nip)
}

func docxText(value string) string {
	return escapeXML(strings.TrimSpace(value))
}

func docxMultilineText(value string) string {
	value = strings.ReplaceAll(strings.TrimSpace(value), "\r\n", "\n")
	if value == "" {
		return ""
	}
	return strings.ReplaceAll(escapeXML(value), "\n", "</w:t><w:br/><w:t>")
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func sanitizeFilenamePart(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "dokumen"
	}

	replacer := strings.NewReplacer("/", "_", "\\", "_", ":", "_", "*", "_", "?", "_", "\"", "_", "<", "_", ">", "_", "|", "_")
	return replacer.Replace(strings.ReplaceAll(value, " ", "_"))
}

func contentTypeToDocxExt(contentType string) string {
	switch strings.ToLower(strings.TrimSpace(contentType)) {
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/jpeg", "image/jpg":
		return ".jpeg"
	default:
		return ""
	}
}
