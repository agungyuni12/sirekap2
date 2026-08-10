-- ============================================================
-- Assign nomor Surat Pernyataan PPL, BAPP, dan Surat Pernyataan Kepala
-- BPS - SE2026 BATCH 4, utk 3 petugas pengganti (lihat
-- migrations/add_lk_ppk_payable_batch4_se2026.sql). Tidak ada Pernyataan
-- PML batch ini (PML mereka sudah payable & sudah punya nomor sejak
-- batch 1/2).
--
-- Jalankan SETELAH migrations/add_mitra_batch4_se2026.sql,
-- migrations/add_rekap_batch4_se2026.sql,
-- migrations/add_lk_ppk_termin1_batch4_se2026.sql, dan
-- migrations/add_lk_ppk_payable_batch4_se2026.sql.
--
-- Tanggal: Pernyataan PPL = 8 Agustus 2026, BAPP & Pernyataan Kepala BPS
-- = 9 Agustus 2026. Nomor urut LANJUT dari batch 1-3 (PPL 231-233,
-- seq_all 263-265). Idempotent: tidak menimpa nomor yang sudah pernah dibuat.
-- ============================================================

-- ================== Surat Pernyataan PPL (3 orang, seq 231-233, 8 Agustus 2026) ==================
UPDATE rekap SET id_pernyataan1 = 'B-08.08.231/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-08-08' WHERE idsobat = '520526050109' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 231. IIN NURROHMADANIL
UPDATE rekap SET id_pernyataan1 = 'B-08.08.232/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-08-08' WHERE idsobat = '520526050093' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 232. Moh Naufal
UPDATE rekap SET id_pernyataan1 = 'B-08.08.233/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-08-08' WHERE idsobat = '520526070001' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 233. Sarwan

-- ================== BAPP Batch 4 (PPL saja, seq_all 263-265, 9 Agustus 2026) ==================
UPDATE rekap SET id_bapp1 = 'B-08.09.263/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-08-09' WHERE idsobat = '520526050109' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 263. IIN NURROHMADANIL (PPL)
UPDATE rekap SET id_bapp1 = 'B-08.09.264/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-08-09' WHERE idsobat = '520526050093' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 264. Moh Naufal (PPL)
UPDATE rekap SET id_bapp1 = 'B-08.09.265/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-08-09' WHERE idsobat = '520526070001' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 265. Sarwan (PPL)

-- ================== Surat Pernyataan Kepala BPS Batch 4 (dokumen tunggal, KUMULATIF batch 1+2+3+4, 9 Agustus 2026) ==================
INSERT INTO surat_kepala_se2026 (termin, tahun, batch, nomor, tanggal) SELECT 1, '2026', 4, 'B-08.09.004/SE2026/5205/Super.KPL/2026', '2026-08-09' WHERE NOT EXISTS (SELECT 1 FROM surat_kepala_se2026 WHERE termin=1 AND tahun='2026' AND batch=4);

SELECT
  (SELECT COUNT(*) FROM rekap WHERE id_pernyataan1 LIKE 'B-08.08.%/Super.PPL/2026') AS pernyataan_ppl_batch4,
  (SELECT COUNT(*) FROM rekap WHERE id_bapp1 LIKE 'B-08.09.%/BAPP-I-SE2026/%') AS bapp_batch4,
  (SELECT COUNT(*) FROM surat_kepala_se2026 WHERE termin=1 AND tahun='2026' AND batch=4) AS kepala_batch4;
