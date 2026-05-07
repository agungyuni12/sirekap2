package handlers

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"sirekap/internal/models"
)

func parsePostedForm(r *http.Request) error {
	if err := r.ParseMultipartForm(10 << 20); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return nil
}

// Indonesian month names
var bulanIndonesia = map[int]string{
	1: "Januari", 2: "Februari", 3: "Maret", 4: "April",
	5: "Mei", 6: "Juni", 7: "Juli", 8: "Agustus",
	9: "September", 10: "Oktober", 11: "November", 12: "Desember",
}

var hariIndonesia = map[string]string{
	"Sunday": "Minggu", "Monday": "Senin", "Tuesday": "Selasa",
	"Wednesday": "Rabu", "Thursday": "Kamis", "Friday": "Jumat", "Saturday": "Sabtu",
}

type spkBulkOption struct {
	IDSobat   string `json:"idsobat"`
	NamaMitra string `json:"namamitra"`
	Bulan     string `json:"bulan"`
	Tahun     string `json:"tahun"`
	SudahAda  bool   `json:"sudahAda"`
	NSurat    string `json:"nsurat,omitempty"`
	TglSurat  string `json:"tglsurat,omitempty"`
}

type bastBulkOption struct {
	ID        int    `json:"id"`
	IDSobat   string `json:"idsobat"`
	NamaMitra string `json:"namamitra"`
	Kegiatan  string `json:"kegiatan"`
	Bulan     string `json:"bulan"`
	Tahun     string `json:"tahun"`
	SPKAda    bool   `json:"spkAda"`
	SPKSurat  string `json:"spkSurat,omitempty"`
	SudahAda  bool   `json:"sudahAda"`
	NSurat    string `json:"nsurat,omitempty"`
	TglSurat  string `json:"tglsurat,omitempty"`
}

type bulkSPKCreatePayload struct {
	TanggalSurat string `json:"tanggal_surat"`
	Items        []struct {
		IDSobat string `json:"idsobat"`
		Bulan   string `json:"bulan"`
		Tahun   string `json:"tahun"`
	} `json:"items"`
}

type bulkBASTCreatePayload struct {
	TanggalSurat string `json:"tanggal_surat"`
	Items        []struct {
		ID int `json:"id"`
	} `json:"items"`
}

// Terbilang converts number to Indonesian words
func Terbilang(n int64) string {
	if n == 0 {
		return "nol"
	}

	huruf := []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

	if n < 12 {
		return huruf[n]
	} else if n < 20 {
		return Terbilang(n-10) + " belas"
	} else if n < 100 {
		return Terbilang(n/10) + " puluh " + Terbilang(n%10)
	} else if n < 200 {
		return "seratus " + Terbilang(n-100)
	} else if n < 1000 {
		return Terbilang(n/100) + " ratus " + Terbilang(n%100)
	} else if n < 2000 {
		return "seribu " + Terbilang(n-1000)
	} else if n < 1000000 {
		return Terbilang(n/1000) + " ribu " + Terbilang(n%1000)
	} else if n < 1000000000 {
		return Terbilang(n/1000000) + " juta " + Terbilang(n%1000000)
	} else {
		return Terbilang(n/1000000000) + " milyar " + Terbilang(n%1000000000)
	}
}

// escapeXML escapes special characters for XML compliance
func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}

// FormatRupiah formats number as Indonesian currency
func FormatRupiah(amount float64) string {
	intVal := int64(amount)
	str := fmt.Sprintf("%d", intVal)
	result := ""
	length := len(str)
	for i, c := range str {
		if i > 0 && (length-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return "Rp " + result + ",00"
}

// generateWordTableXML creates a Word XML table for the kegiatan list
func generateWordTableXML(kegiatanList []models.MitraKegiatan, passedBulan, passedTahun string) string {
	if len(kegiatanList) == 0 {
		return ""
	}

	var rowXML strings.Builder
	totalHonor := 0.0

	// Helper to format date
	formatDate := func(t time.Time) (string, string, string) {
		// Get Indonesian month name
		months := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
		m := t.Month()
		return fmt.Sprintf("%d", t.Day()), months[m], fmt.Sprintf("%d", t.Year())
	}

	for i, k := range kegiatanList {
		// Default to passed args if semantic dates missing
		// We use passedBulan and passedTahun as fallback string representation

		var jpelaksanaan string
		if !k.PMWaktu.IsZero() && !k.PSWaktu.IsZero() {
			tpm, bpm, thpm := formatDate(k.PMWaktu)
			tps, bps, thps := formatDate(k.PSWaktu)

			if tpm != "" && tps != "" {
				if bpm == bps {
					if thpm == thps {
						jpelaksanaan = fmt.Sprintf("%s-%s %s %s", tpm, tps, bpm, thpm)
					} else {
						jpelaksanaan = fmt.Sprintf("%s %s %s - %s %s %s", tpm, bpm, thpm, tps, bps, thps)
					}
				} else {
					if thpm == thps {
						jpelaksanaan = fmt.Sprintf("%s %s - %s %s %s", tpm, bpm, tps, bps, thpm)
					} else {
						jpelaksanaan = fmt.Sprintf("%s %s %s - %s %s %s", tpm, bpm, thpm, tps, bps, thps)
					}
				}
			} else {
				// Fallback
				jpelaksanaan = fmt.Sprintf("%s %s", passedBulan, passedTahun)
			}
		} else {
			jpelaksanaan = fmt.Sprintf("%s %s", passedBulan, passedTahun)
		}

		// Escape all dynamic values to prevent XML corruption
		cleanKegiatan := escapeXML(k.Kegiatan)
		cleanJadwal := escapeXML(jpelaksanaan)
		cleanSatuan := escapeXML(k.Satuan)
		cleanMak := escapeXML(k.Mak)

		// Volume, Harga, Honor are numbers/formatted strings
		cleanVolume := fmt.Sprintf("%.0f", k.Volume)
		if k.Volume != float64(int64(k.Volume)) {
			cleanVolume = fmt.Sprintf("%.2f", k.Volume)
		}

		cleanHSatuan := FormatRupiah(k.HSatuan)
		cleanHonor := FormatRupiah(k.Honor)

		rowXML.WriteString(fmt.Sprintf(`
		<w:tr>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%d</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="left"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%s</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%s</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%s</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%s</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="right"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%s</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="right"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>%s</w:t></w:r></w:p>
			</w:tc>
			<w:tc>
				<w:tcPr><w:tcW w:w="1750" w:type="dxa"/></w:tcPr>
				<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:rPr><w:rFonts w:ascii="Bookman Old Style" w:hAnsi="Bookman Old Style"/><w:sz w:val="22"/></w:rPr><w:t>[%s]</w:t></w:r></w:p>
			</w:tc>
		</w:tr>`, i+1, cleanKegiatan, cleanJadwal, cleanVolume, cleanSatuan, cleanHSatuan, cleanHonor, cleanMak))

		totalHonor += k.Honor
	}

	// Start table
	tableXML := `<w:tbl>
<w:tblPr>
<w:tblStyle w:val="TableGrid"/>
<w:tblW w:w="0" w:type="auto"/>
<w:tblBorders>
<w:top w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:left w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:bottom w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:right w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideH w:val="single" w:sz="4" w:space="0" w:color="000000"/>
<w:insideV w:val="single" w:sz="4" w:space="0" w:color="000000"/>
</w:tblBorders>
</w:tblPr>
<w:tr>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>No</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Nama Survei</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Jadwal</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Volume</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Satuan</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Harga Satuan</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>Nilai</w:t></w:r></w:p></w:tc>
<w:tc><w:p><w:r><w:rPr><w:b/></w:rPr><w:t>MAK</w:t></w:r></w:p></w:tc>
</w:tr>
`

	// Append rows
	tableXML += rowXML.String()

	// Total row
	tableXML += fmt.Sprintf(`<w:tr>
<w:tc>
        <w:tcPr>
            <w:tcW w:w="7200" w:type="dxa"/>
            <w:gridSpan w:val="6"/>
            <w:vAlign w:val="center"/>
        </w:tcPr>
        <w:p>
            <w:pPr>
                <w:jc w:val="center"/>
            </w:pPr>
            <w:r>
                <w:rPr><w:b/></w:rPr>
                <w:t>%s Rupiah</w:t>
            </w:r>
        </w:p>
    </w:tc>

    <w:tc>
        <w:tcPr>
            <w:tcW w:w="2400" w:type="dxa"/>
            <w:gridSpan w:val="2"/>
            <w:vAlign w:val="center"/>
        </w:tcPr>
        <w:p>
            <w:pPr>
                <w:jc w:val="right"/>
            </w:pPr>
            <w:r>
                <w:rPr><w:b/></w:rPr>
                <w:t>%s</w:t>
            </w:r>
        </w:p>
    </w:tc>
</w:tr>
`, strings.Title(Terbilang(int64(totalHonor))), FormatRupiah(totalHonor))

	tableXML += `</w:tbl>`
	return tableXML
}

// cleanXMLSplitPlaceholders removes XML tags that split placeholders like ${bulan</w:t>...</w:r>...<w:t>1}
func cleanXMLSplitPlaceholders(content string) string {
	// First, handle case where $ and { are split by tags
	// Pattern: $ followed by XML tags, then {
	reHead := regexp.MustCompile(`\$((?:</w:t>.*?<w:t[^>]*>)+)\{`)
	for reHead.MatchString(content) {
		content = reHead.ReplaceAllString(content, "${")
	}

	// Pattern to find ${...} where the content may be split by XML tags
	// We'll find all ${...} patterns and rebuild them without internal XML
	// Pattern: ${ followed by text and XML tags until we find }
	re := regexp.MustCompile(`\$\{([^}]*?)(</w:t>.*?<w:t[^>]*>)([^}]*?)\}`)

	// Keep replacing until no more splits found
	for re.MatchString(content) {
		content = re.ReplaceAllStringFunc(content, func(match string) string {
			// Extract the placeholder name by removing all XML tags
			cleanMatch := regexp.MustCompile(`<[^>]+>`).ReplaceAllString(match, "")
			return cleanMatch
		})
	}

	return content
}

// replaceInDocx opens a DOCX, replaces placeholders, and returns the modified content
func replaceInDocx(templatePath string, replacements map[string]string) ([]byte, error) {
	// Open the template file
	reader, err := zip.OpenReader(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open template: %v", err)
	}
	defer reader.Close()

	// Create a buffer for the output
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)

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

		// Replace placeholders in document.xml and other XML files
		if strings.HasSuffix(file.Name, ".xml") || strings.HasSuffix(file.Name, ".rels") {
			contentStr := string(content)

			// Clean up split placeholders first
			contentStr = cleanXMLSplitPlaceholders(contentStr)

			for placeholder, value := range replacements {
				if placeholder == "table" {
					// Special handling for table: replace the containing paragraph
					// Find the placeholder: ${table} (allowing spaces)
					re := regexp.MustCompile(`\$\{\s*` + regexp.QuoteMeta(placeholder) + `\s*\}`)
					loc := re.FindStringIndex(contentStr)
					if loc != nil {
						// Found placeholder. Now find enclosing <w:p>
						// Be careful: strings.LastIndex searches for exact string match.
						// "<w:p" would match "<w:pPr>", which is wrong.
						// We must match "<w:p>" or "<w:p " (with space).

						startP1 := strings.LastIndex(contentStr[:loc[0]], "<w:p>")
						startP2 := strings.LastIndex(contentStr[:loc[0]], "<w:p ")

						startP := startP1
						if startP2 > startP {
							startP = startP2
						}

						endP := strings.Index(contentStr[loc[1]:], "</w:p>")

						if startP != -1 && endP != -1 {
							endP += loc[1] + 6 // Include </w:p> length
							// Replace the paragraph with the table XML
							contentStr = contentStr[:startP] + value + contentStr[endP:]
						} else {
							// Fallback to simple replace if paragraph not found
							contentStr = re.ReplaceAllString(contentStr, value)
						}
					}
				} else {
					// Normal replacement with regex to handle spaces like ${nsurat }
					re := regexp.MustCompile(`\$\{\s*` + regexp.QuoteMeta(placeholder) + `\s*\}`)
					contentStr = re.ReplaceAllString(contentStr, value)

					// Also try without $ just in case {placeholder} format is used
					re2 := regexp.MustCompile(`\{\s*` + regexp.QuoteMeta(placeholder) + `\s*\}`)
					contentStr = re2.ReplaceAllString(contentStr, value)
				}
			}
			content = []byte(contentStr)
		}

		// Create the file in the new zip
		header, err := zip.FileInfoHeader(file.FileInfo())
		if err != nil {
			return nil, err
		}
		header.Name = file.Name
		header.Method = file.Method

		w, err := writer.CreateHeader(header)
		if err != nil {
			return nil, err
		}

		_, err = w.Write(content)
		if err != nil {
			return nil, err
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func parseRekapFilterFromRequest(r *http.Request) models.RekapFilter {
	return models.RekapFilter{
		TahunAnggaran: strings.TrimSpace(r.URL.Query().Get("tahunAnggaran")),
		Bulan:         strings.TrimSpace(r.URL.Query().Get("bulan")),
		Search:        strings.TrimSpace(r.URL.Query().Get("search")),
	}
}

func ListBulkSPKOptionsHandler(w http.ResponseWriter, r *http.Request) {
	filter := parseRekapFilterFromRequest(r)

	groups, err := models.GetFilteredMitraGroups(filter)
	if err != nil {
		log.Printf("Error getting SPK bulk options: %v", err)
		http.Error(w, "Gagal memuat kandidat SPK", http.StatusInternalServerError)
		return
	}

	options := make([]spkBulkOption, 0, len(groups))
	for _, item := range groups {
		option := spkBulkOption{
			IDSobat:   item.IDSobat,
			NamaMitra: item.NamaMitra,
			Bulan:     item.Bulan,
			Tahun:     item.Tahun,
		}

		surat, err := models.GetSPKSurat(item.IDSobat, item.Bulan, item.Tahun)
		if err != nil {
			log.Printf("Error checking SPK surat for %s: %v", item.IDSobat, err)
		} else if surat != nil {
			option.SudahAda = true
			option.NSurat = surat.NSurat
			option.TglSurat = surat.TglSurat
		}

		options = append(options, option)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  options,
		"total": len(options),
	})
}

func ListBulkBASTOptionsHandler(w http.ResponseWriter, r *http.Request) {
	filter := parseRekapFilterFromRequest(r)

	items, err := models.GetFilteredBASTDownloadItems(filter)
	if err != nil {
		log.Printf("Error getting BAST bulk options: %v", err)
		http.Error(w, "Gagal memuat kandidat BAST", http.StatusInternalServerError)
		return
	}

	options := make([]bastBulkOption, 0, len(items))
	for _, item := range items {
		option := bastBulkOption{
			ID:        item.ID,
			IDSobat:   item.IDSobat,
			NamaMitra: item.NamaMitra,
			Kegiatan:  item.Kegiatan,
			Bulan:     item.Bulan,
			Tahun:     item.Tahun,
		}

		spkSurat, err := models.GetSPKSurat(item.IDSobat, item.Bulan, item.Tahun)
		if err != nil {
			log.Printf("Error checking SPK surat for BAST %d: %v", item.ID, err)
		} else if spkSurat != nil {
			option.SPKAda = true
			option.SPKSurat = spkSurat.NSurat
		}

		bastSurat, err := models.GetBASTSurat(item.IDSobat, item.Bulan, item.Tahun, item.Kegiatan)
		if err != nil {
			log.Printf("Error checking BAST surat %d: %v", item.ID, err)
		} else if bastSurat != nil {
			option.SudahAda = true
			option.NSurat = bastSurat.NSurat
			option.TglSurat = bastSurat.TglSurat
		}

		options = append(options, option)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  options,
		"total": len(options),
	})
}

func CreateBulkSPKHandler(w http.ResponseWriter, r *http.Request) {
	var payload bulkSPKCreatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Payload tidak valid", http.StatusBadRequest)
		return
	}

	payload.TanggalSurat = strings.TrimSpace(payload.TanggalSurat)
	if payload.TanggalSurat == "" {
		http.Error(w, "Tanggal surat wajib diisi", http.StatusBadRequest)
		return
	}
	if len(payload.Items) == 0 {
		http.Error(w, "Pilih minimal satu mitra", http.StatusBadRequest)
		return
	}

	successCount := 0
	createdNumbers := []string{}
	failedItems := []string{}

	for _, item := range payload.Items {
		mitra, err := models.GetMitraByIDSobat(strings.TrimSpace(item.IDSobat))
		if err != nil {
			failedItems = append(failedItems, item.IDSobat)
			continue
		}

		nsurat, err := models.CreateSurat(item.IDSobat, mitra.NamaMitra, item.Bulan, item.Tahun, "SPK", payload.TanggalSurat, "")
		if err != nil {
			log.Printf("Error creating bulk SPK surat for %s: %v", item.IDSobat, err)
			failedItems = append(failedItems, mitra.NamaMitra)
			continue
		}

		docBytes, filename, err := generateSPKDocument(item.IDSobat, item.Bulan, item.Tahun)
		if err != nil {
			log.Printf("Error generating bulk SPK document for %s: %v", item.IDSobat, err)
			failedItems = append(failedItems, mitra.NamaMitra)
			continue
		}

		if err := saveGeneratedDocument(filename, docBytes); err != nil {
			log.Printf("Error saving bulk SPK document for %s: %v", item.IDSobat, err)
			failedItems = append(failedItems, mitra.NamaMitra)
			continue
		}

		successCount++
		createdNumbers = append(createdNumbers, nsurat)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": successCount > 0,
		"created": successCount,
		"failed":  len(failedItems),
		"numbers": createdNumbers,
		"errors":  failedItems,
	})
}

func CreateBulkBASTHandler(w http.ResponseWriter, r *http.Request) {
	var payload bulkBASTCreatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Payload tidak valid", http.StatusBadRequest)
		return
	}

	payload.TanggalSurat = strings.TrimSpace(payload.TanggalSurat)
	if payload.TanggalSurat == "" {
		http.Error(w, "Tanggal surat wajib diisi", http.StatusBadRequest)
		return
	}
	if len(payload.Items) == 0 {
		http.Error(w, "Pilih minimal satu kegiatan", http.StatusBadRequest)
		return
	}

	successCount := 0
	createdNumbers := []string{}
	failedItems := []string{}

	for _, item := range payload.Items {
		rekap, err := models.GetRekapRecordByID(item.ID)
		if err != nil {
			failedItems = append(failedItems, strconv.Itoa(item.ID))
			continue
		}

		spkSurat, err := models.GetSPKSurat(rekap.IDSobat, rekap.Bulan, rekap.Tahun)
		if err != nil || spkSurat == nil {
			failedItems = append(failedItems, fmt.Sprintf("%s - %s", rekap.NamaMitra, rekap.Kegiatan))
			continue
		}

		mitra, err := models.GetMitraByIDSobat(rekap.IDSobat)
		if err != nil {
			failedItems = append(failedItems, fmt.Sprintf("%s - %s", rekap.IDSobat, rekap.Kegiatan))
			continue
		}

		nsurat, err := models.CreateSurat(rekap.IDSobat, mitra.NamaMitra, rekap.Bulan, rekap.Tahun, "BAST", payload.TanggalSurat, rekap.Kegiatan)
		if err != nil {
			log.Printf("Error creating bulk BAST surat for %d: %v", item.ID, err)
			failedItems = append(failedItems, fmt.Sprintf("%s - %s", mitra.NamaMitra, rekap.Kegiatan))
			continue
		}

		docBytes, filename, err := generateBASTDocument(item.ID)
		if err != nil {
			log.Printf("Error generating bulk BAST document for %d: %v", item.ID, err)
			failedItems = append(failedItems, fmt.Sprintf("%s - %s", mitra.NamaMitra, rekap.Kegiatan))
			continue
		}

		if err := saveGeneratedDocument(filename, docBytes); err != nil {
			log.Printf("Error saving bulk BAST document for %d: %v", item.ID, err)
			failedItems = append(failedItems, fmt.Sprintf("%s - %s", mitra.NamaMitra, rekap.Kegiatan))
			continue
		}

		successCount++
		createdNumbers = append(createdNumbers, nsurat)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": successCount > 0,
		"created": successCount,
		"failed":  len(failedItems),
		"numbers": createdNumbers,
		"errors":  failedItems,
	})
}

// CreateSPKHandler handles POST /api/rekap/spk/create
func CreateSPKHandler(w http.ResponseWriter, r *http.Request) {
	if err := parsePostedForm(r); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	idsobat := r.FormValue("idsobat")
	bulan := r.FormValue("bulan")
	tahun := r.FormValue("tahun")
	tglSurat := r.FormValue("tanggal_surat")

	var missing []string
	if idsobat == "" {
		missing = append(missing, "idsobat")
	}
	if bulan == "" {
		missing = append(missing, "bulan")
	}
	if tahun == "" {
		missing = append(missing, "tahun")
	}
	if tglSurat == "" {
		missing = append(missing, "tanggal_surat")
	}

	if len(missing) > 0 {
		http.Error(w, "Missing required fields: "+strings.Join(missing, ", "), http.StatusBadRequest)
		return
	}

	// Get mitra info
	mitra, err := models.GetMitraByIDSobat(idsobat)
	if err != nil {
		log.Printf("Error getting mitra: %v", err)
		http.Error(w, "Mitra tidak ditemukan", http.StatusNotFound)
		return
	}

	// Create letter number
	nsurat, err := models.CreateSurat(idsobat, mitra.NamaMitra, bulan, tahun, "SPK", tglSurat, "")
	if err != nil {
		log.Printf("Error creating surat: %v", err)
		http.Error(w, "Gagal membuat nomor surat", http.StatusInternalServerError)
		return
	}

	docBytes, filename, err := generateSPKDocument(idsobat, bulan, tahun)
	if err != nil {
		log.Printf("Error generating SPK: %v", err)
		http.Error(w, "Gagal membuat dokumen SPK", http.StatusInternalServerError)
		return
	}

	// Save to output folder
	if err := saveGeneratedDocument(filename, docBytes); err != nil {
		log.Printf("Error saving SPK: %v", err)
		http.Error(w, "Gagal menyimpan dokumen", http.StatusInternalServerError)
		return
	}

	// Return success
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true, "message": "SPK berhasil dibuat", "nsurat": "` + nsurat + `"}`))
}

// DownloadSPKHandler handles GET /api/rekap/spk/download
func DownloadSPKHandler(w http.ResponseWriter, r *http.Request) {
	idsobat := r.URL.Query().Get("idsobat")
	bulan := r.URL.Query().Get("bulan")
	tahun := r.URL.Query().Get("tahun")

	docBytes, filename, err := generateSPKDocument(idsobat, bulan, tahun)
	if err != nil {
		http.Error(w, "SPK belum bisa diunduh. Pastikan nomor surat SPK sudah dibuat terlebih dahulu.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", strconv.Itoa(len(docBytes)))
	_, _ = w.Write(docBytes)
}

func DownloadFilteredSPKZipHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.RekapFilter{
		TahunAnggaran: strings.TrimSpace(r.URL.Query().Get("tahunAnggaran")),
		Bulan:         strings.TrimSpace(r.URL.Query().Get("bulan")),
		Search:        strings.TrimSpace(r.URL.Query().Get("search")),
	}

	groups, err := models.GetFilteredMitraGroups(filter)
	if err != nil {
		log.Printf("Error getting filtered SPK groups: %v", err)
		http.Error(w, "Gagal memuat data rekap", http.StatusInternalServerError)
		return
	}

	if len(groups) == 0 {
		http.Error(w, "Tidak ada data yang sesuai dengan filter", http.StatusNotFound)
		return
	}

	zipBytes, addedCount, missingNames, err := buildRekapZip(func(addFile func(name string, content []byte) error) ([]string, error) {
		missingNames := []string{}
		for _, item := range groups {
			docBytes, filename, genErr := generateSPKDocument(item.IDSobat, item.Bulan, item.Tahun)
			if genErr != nil {
				missingNames = append(missingNames, expectedSPKFilename(item.NamaMitra, item.Bulan, item.Tahun))
				continue
			}
			if err := addFile(filename, docBytes); err != nil {
				return nil, err
			}
		}
		return missingNames, nil
	})
	if err != nil {
		log.Printf("Error building SPK zip: %v", err)
		http.Error(w, "Gagal menyiapkan file unduhan", http.StatusInternalServerError)
		return
	}

	if addedCount == 0 {
		http.Error(w, "Belum ada file SPK yang bisa diunduh untuk filter ini", http.StatusNotFound)
		return
	}
	if len(missingNames) > 0 {
		log.Printf("SPK bulk download skipped %d missing files", len(missingNames))
	}

	filename := fmt.Sprintf("SPK_Filtered_%s.zip", time.Now().Format("20060102_150405"))
	serveZipDownload(w, zipBytes, filename)
}

// CreateBASTHandler handles POST /api/rekap/bast/create
func CreateBASTHandler(w http.ResponseWriter, r *http.Request) {
	if err := parsePostedForm(r); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	idRekapStr := r.FormValue("id_rekap")
	idsobat := r.FormValue("idsobat")
	bulan := r.FormValue("bulan")
	tahun := r.FormValue("tahun")
	tglSurat := r.FormValue("tanggal_surat")

	idRekap, _ := strconv.Atoi(idRekapStr)

	var missing []string
	if idRekap == 0 {
		missing = append(missing, "id_rekap")
	}
	if idsobat == "" {
		missing = append(missing, "idsobat")
	}
	if bulan == "" {
		missing = append(missing, "bulan")
	}
	if tahun == "" {
		missing = append(missing, "tahun")
	}
	if tglSurat == "" {
		missing = append(missing, "tanggal_surat")
	}

	if len(missing) > 0 {
		http.Error(w, "Missing required fields: "+strings.Join(missing, ", "), http.StatusBadRequest)
		return
	}

	// Get rekap/kegiatan info
	kegiatan, err := models.GetRekapById(idRekap)
	if err != nil {
		log.Printf("Error getting rekap: %v", err)
		http.Error(w, "Kegiatan tidak ditemukan", http.StatusNotFound)
		return
	}

	// Get mitra info
	mitra, err := models.GetMitraByIDSobat(idsobat)
	if err != nil {
		log.Printf("Error getting mitra: %v", err)
		http.Error(w, "Mitra tidak ditemukan", http.StatusNotFound)
		return
	}

	// Create letter number
	nsurat, err := models.CreateSurat(idsobat, mitra.NamaMitra, bulan, tahun, "BAST", tglSurat, kegiatan.Kegiatan)
	if err != nil {
		log.Printf("Error creating surat: %v", err)
		http.Error(w, "Gagal membuat nomor surat", http.StatusInternalServerError)
		return
	}

	docBytes, filename, err := generateBASTDocument(idRekap)
	if err != nil {
		log.Printf("Error generating BAST: %v", err)
		http.Error(w, "Gagal membuat dokumen BAST", http.StatusInternalServerError)
		return
	}

	// Save to output folder
	if err := saveGeneratedDocument(filename, docBytes); err != nil {
		log.Printf("Error saving BAST: %v", err)
		http.Error(w, "Gagal menyimpan dokumen", http.StatusInternalServerError)
		return
	}

	// Return success
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true, "message": "BAST berhasil dibuat", "nsurat": "` + nsurat + `"}`))
}

// DownloadBASTHandler handles GET /api/rekap/bast/download
func DownloadBASTHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if id == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	docBytes, filename, err := generateBASTDocument(id)
	if err != nil {
		http.Error(w, "BAST belum bisa diunduh. Pastikan nomor surat BAST sudah dibuat terlebih dahulu.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", strconv.Itoa(len(docBytes)))
	_, _ = w.Write(docBytes)
}

func DownloadFilteredBASTZipHandler(w http.ResponseWriter, r *http.Request) {
	filter := models.RekapFilter{
		TahunAnggaran: strings.TrimSpace(r.URL.Query().Get("tahunAnggaran")),
		Bulan:         strings.TrimSpace(r.URL.Query().Get("bulan")),
		Search:        strings.TrimSpace(r.URL.Query().Get("search")),
	}

	items, err := models.GetFilteredBASTDownloadItems(filter)
	if err != nil {
		log.Printf("Error getting filtered BAST data: %v", err)
		http.Error(w, "Gagal memuat data rekap", http.StatusInternalServerError)
		return
	}

	if len(items) == 0 {
		http.Error(w, "Tidak ada data yang sesuai dengan filter", http.StatusNotFound)
		return
	}

	zipBytes, addedCount, missingNames, err := buildRekapZip(func(addFile func(name string, content []byte) error) ([]string, error) {
		missingNames := []string{}
		for _, item := range items {
			docBytes, filename, genErr := generateBASTDocument(item.ID)
			if genErr != nil {
				missingNames = append(missingNames, expectedBASTFilename(item.NamaMitra, item.Bulan, item.Tahun, item.Kegiatan))
				continue
			}
			zipName := fmt.Sprintf("%03d_%s", item.ID, filename)
			if err := addFile(zipName, docBytes); err != nil {
				return nil, err
			}
		}
		return missingNames, nil
	})
	if err != nil {
		log.Printf("Error building BAST zip: %v", err)
		http.Error(w, "Gagal menyiapkan file unduhan", http.StatusInternalServerError)
		return
	}

	if addedCount == 0 {
		http.Error(w, "Belum ada file BAST yang bisa diunduh untuk filter ini", http.StatusNotFound)
		return
	}
	if len(missingNames) > 0 {
		log.Printf("BAST bulk download skipped %d missing files", len(missingNames))
	}

	filename := fmt.Sprintf("BAST_Filtered_%s.zip", time.Now().Format("20060102_150405"))
	serveZipDownload(w, zipBytes, filename)
}

// Helper function to get month index from Indonesian month name
func getBulanIndex(bulan string) int {
	monthMap := map[string]int{
		"Januari": 1, "Februari": 2, "Maret": 3, "April": 4,
		"Mei": 5, "Juni": 6, "Juli": 7, "Agustus": 8,
		"September": 9, "Oktober": 10, "November": 11, "Desember": 12,
	}
	return monthMap[bulan]
}

// Helper function to get last day of month
func getLastDayOfMonth(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

func generateSPKDocument(idsobat, bulan, tahun string) ([]byte, string, error) {
	surat, err := models.GetSPKSurat(idsobat, bulan, tahun)
	if err != nil {
		return nil, "", err
	}
	if surat == nil {
		return nil, "", os.ErrNotExist
	}

	mitra, err := models.GetMitraByIDSobat(idsobat)
	if err != nil {
		return nil, "", err
	}

	kegiatan, err := models.GetMitraKegiatan(idsobat, bulan, tahun)
	if err != nil {
		return nil, "", err
	}

	t, err := time.Parse("2006-01-02", surat.TglSurat)
	if err != nil {
		return nil, "", err
	}

	tanggalKata := strings.Title(Terbilang(int64(t.Day())))
	bulanKata := bulanIndonesia[int(t.Month())]
	tahunKata := strings.Title(Terbilang(int64(t.Year())))
	hariKata := hariIndonesia[t.Weekday().String()]

	bulanKerja := getBulanIndex(bulan)
	tahunKerjaInt, _ := strconv.Atoi(tahun)
	lastDay := getLastDayOfMonth(tahunKerjaInt, bulanKerja)

	replacements := map[string]string{
		"nsurat":    surat.NSurat,
		"tanggal":   tanggalKata,
		"hari":      hariKata,
		"bulan1":    strings.ToUpper(bulan),
		"bulan2":    bulan,
		"bulan3":    bulanKata,
		"nama":      mitra.NamaMitra,
		"alamat":    mitra.Alamat,
		"kecamatan": mitra.Kecamatan,
		"tahun11":   tahun,
		"tahun12":   tahunKata,
		"tanggal2":  fmt.Sprintf("%d %s %s", lastDay, bulan, tahun),
		"table":     generateWordTableXML(kegiatan, bulan, tahun),
	}

	docBytes, err := replaceInDocx("template/SPK bulanan.docx", replacements)
	if err != nil {
		return nil, "", err
	}

	return docBytes, expectedSPKFilename(mitra.NamaMitra, bulan, tahun), nil
}

func generateBASTDocument(rekapID int) ([]byte, string, error) {
	rekap, err := models.GetRekapRecordByID(rekapID)
	if err != nil {
		return nil, "", err
	}

	surat, err := models.GetBASTSurat(rekap.IDSobat, rekap.Bulan, rekap.Tahun, rekap.Kegiatan)
	if err != nil {
		return nil, "", err
	}
	if surat == nil {
		return nil, "", os.ErrNotExist
	}

	mitra, err := models.GetMitraByIDSobat(rekap.IDSobat)
	if err != nil {
		return nil, "", err
	}

	spkSurat, err := models.GetSPKSurat(rekap.IDSobat, rekap.Bulan, rekap.Tahun)
	if err != nil {
		return nil, "", err
	}
	spkNsurat := ""
	if spkSurat != nil {
		spkNsurat = spkSurat.NSurat
	}

	t, err := time.Parse("2006-01-02", surat.TglSurat)
	if err != nil {
		return nil, "", err
	}

	tanggalKata := strings.Title(Terbilang(int64(t.Day())))
	bulanKata := bulanIndonesia[int(t.Month())]
	hariKata := hariIndonesia[t.Weekday().String()]
	tahunKata := strings.Title(Terbilang(int64(t.Year())))

	replacements := map[string]string{
		"nsurat":      surat.NSurat,
		"namatanggal": tanggalKata,
		"namahari":    hariKata,
		"bulan1":      bulanKata,
		"bulan2":      rekap.Bulan,
		"nama":        mitra.NamaMitra,
		"nik":         mitra.NIK,
		"alamat":      mitra.Alamat,
		"tahun21":     rekap.Tahun,
		"tahun12":     tahunKata,
		"kegiatan":    rekap.Kegiatan,
		"volume":      fmt.Sprintf("%.0f", rekap.Volume),
		"satuan":      rekap.Satuan,
		"nsurat11":    spkNsurat,
	}

	docBytes, err := replaceInDocx("template/BAST bulanan.docx", replacements)
	if err != nil {
		return nil, "", err
	}

	return docBytes, expectedBASTFilename(mitra.NamaMitra, rekap.Bulan, rekap.Tahun, rekap.Kegiatan), nil
}

func expectedSPKFilename(namaMitra, bulan, tahun string) string {
	return sanitizeGeneratedFilename(fmt.Sprintf("SPK %s %s %s.docx", namaMitra, bulan, tahun))
}

func expectedBASTFilename(namaMitra, bulan, tahun, kegiatan string) string {
	cleanKegiatan := strings.ReplaceAll(kegiatan, "/", "_")
	return sanitizeGeneratedFilename(fmt.Sprintf("BAST %s %s %s %s.docx", namaMitra, bulan, tahun, cleanKegiatan))
}

func saveGeneratedDocument(filename string, docBytes []byte) error {
	if err := os.MkdirAll("output", 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join("output", filename), docBytes, 0644)
}

func sanitizeGeneratedFilename(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "dokumen.docx"
	}

	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "-",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "",
		"\n", " ",
		"\r", " ",
		"\t", " ",
	)
	name = replacer.Replace(name)
	name = strings.Join(strings.Fields(name), " ")

	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	if ext == "" {
		ext = ".docx"
	}

	const maxBaseLen = 180
	if len(base) > maxBaseLen {
		base = strings.TrimSpace(base[:maxBaseLen])
	}
	if base == "" {
		base = "dokumen"
	}

	return base + ext
}

func buildRekapZip(populate func(addFile func(name string, content []byte) error) ([]string, error)) ([]byte, int, []string, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	addedNames := map[string]int{}
	addedCount := 0

	addFile := func(name string, contents []byte) error {
		zipName := uniqueZipEntryName(name, addedNames)
		entryWriter, err := zipWriter.Create(zipName)
		if err != nil {
			return err
		}
		if _, err := entryWriter.Write(contents); err != nil {
			return err
		}
		addedCount++
		return nil
	}

	missingNames, err := populate(addFile)
	if err != nil {
		zipWriter.Close()
		return nil, 0, nil, err
	}

	if len(missingNames) > 0 {
		sort.Strings(missingNames)
		noteWriter, err := zipWriter.Create("_catatan_file_tidak_ditemukan.txt")
		if err != nil {
			zipWriter.Close()
			return nil, 0, nil, err
		}
		note := "File berikut belum ditemukan di folder output:\n\n" + strings.Join(missingNames, "\n")
		if _, err := noteWriter.Write([]byte(note)); err != nil {
			zipWriter.Close()
			return nil, 0, nil, err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, 0, nil, err
	}

	return buf.Bytes(), addedCount, missingNames, nil
}

func uniqueZipEntryName(name string, existing map[string]int) string {
	if count, ok := existing[name]; ok {
		count++
		existing[name] = count
		ext := filepath.Ext(name)
		base := strings.TrimSuffix(name, ext)
		return fmt.Sprintf("%s (%d)%s", base, count, ext)
	}

	existing[name] = 1
	return name
}

func serveZipDownload(w http.ResponseWriter, zipBytes []byte, filename string) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", strconv.Itoa(len(zipBytes)))
	_, _ = w.Write(zipBytes)
}

// Unused import prevention
var _ = mux.Vars
var _ = regexp.Compile
