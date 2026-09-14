-- ============================================================
-- rekap: kolom id_pernyataan2/tgl_pernyataan2 untuk Surat Pernyataan
-- Penyelesaian Lapangan Termin II (mirror id_pernyataan1/tgl_pernyataan1,
-- terpisah krn satu orang bisa dapat surat di kedua termin).
-- BAST tidak butuh kolom baru di rekap - tetap lewat tabel `surat` generik
-- (stujuan='BAST'), cuma proses pembuatannya sekarang ada endpoint SE2026-
-- specific (bast_se2026.go) krn alur generik butuh baris `surat` stujuan=SPK
-- yang tidak pernah ada untuk mitra SE2026 (SPK-nya di-import langsung ke
-- rekap.id_spk).
-- ============================================================

ALTER TABLE rekap ADD COLUMN id_pernyataan2 VARCHAR(255) NULL;
ALTER TABLE rekap ADD COLUMN tgl_pernyataan2 DATE NULL;
