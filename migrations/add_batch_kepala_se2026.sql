-- ============================================================
-- Tambah kolom 'batch' ke surat_kepala_se2026 - Surat Pernyataan Kepala BPS
-- sekarang 1 dokumen TERPISAH per batch (batch 1 = 222 orang tgl 17 Juli,
-- batch 2 = 23 orang tgl 19 Juli), bukan 1 dokumen tunggal utk seluruh termin.
-- ============================================================

ALTER TABLE surat_kepala_se2026 ADD COLUMN IF NOT EXISTS batch INT NOT NULL DEFAULT 1;
UPDATE surat_kepala_se2026 SET batch = 1 WHERE batch IS NULL OR batch = 0;

-- Ganti unique key (termin,tahun) -> (termin,tahun,batch) supaya batch 2 bisa punya
-- baris nomor sendiri.
ALTER TABLE surat_kepala_se2026 DROP INDEX IF EXISTS uniq_termin_tahun;
ALTER TABLE surat_kepala_se2026 ADD UNIQUE KEY uniq_termin_tahun_batch (termin, tahun, batch);

SELECT 'Migration add_batch_kepala_se2026 selesai.' AS status;
