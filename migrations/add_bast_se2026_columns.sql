-- ============================================================
-- rekap: kolom id_bast_se2026/tgl_bast_se2026 - BAST khusus SE2026, terpisah
-- dari alur BAST generik (spk_bast.go pakai tabel `surat`, tidak cocok krn
-- mitra SE2026 nomor SPK-nya diimport langsung ke rekap.id_spk, bukan lewat
-- models.CreateSurat). Satu BAST per orang utk SELURUH kontrak ("gabungan
-- total" - Termin 1+2 digabung), makanya kolomnya terpisah dari id_bapp1/2
-- yang per-termin. Lihat internal/handlers/bast_se2026.go.
-- ============================================================

ALTER TABLE rekap ADD COLUMN id_bast_se2026 VARCHAR(255) NULL;
ALTER TABLE rekap ADD COLUMN tgl_bast_se2026 DATE NULL;
