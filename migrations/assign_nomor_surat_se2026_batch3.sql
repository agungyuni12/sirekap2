-- ============================================================
-- Assign nomor Surat Pernyataan (PPL/PML), BAPP, dan Surat Pernyataan
-- Kepala BPS - SE2026 BATCH 3, utk 5 PPL + 4 PML yang bisa dicairkan
-- batch ini (lihat migrations/add_lk_ppk_payable_batch3_se2026.sql).
--
-- Jalankan SETELAH migrations/add_lk_ppk_termin1_batch3_se2026.sql dan
-- migrations/add_lk_ppk_payable_batch3_se2026.sql.
--
-- Tanggal: Pernyataan PPL & PML = 20 Juli 2026, BAPP & Pernyataan
-- Kepala BPS = 21 Juli 2026. Nomor urut LANJUT dari batch 1+2 (PPL 226-230,
-- PML 29-32, seq_all 254-262). Idempotent: tidak menimpa nomor yang
-- sudah pernah dibuat.
-- ============================================================

-- ================== Surat Pernyataan PML (4 orang, seq 29-32, 20 Juli 2026) ==================
UPDATE rekap SET id_pernyataan1 = 'B-07.20.029/SE2026/5205/Super.PML/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520522100030' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 29. JULKARNAIN
UPDATE rekap SET id_pernyataan1 = 'B-07.20.030/SE2026/5205/Super.PML/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520522030029' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 30. SALAHUDDIN
UPDATE rekap SET id_pernyataan1 = 'B-07.20.031/SE2026/5205/Super.PML/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520523050009' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 31. Syahril Sidik
UPDATE rekap SET id_pernyataan1 = 'B-07.20.032/SE2026/5205/Super.PML/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520522100175' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 32. Yulianti

-- ================== Surat Pernyataan PPL (5 orang, seq 226-230, 20 Juli 2026) ==================
UPDATE rekap SET id_pernyataan1 = 'B-07.20.226/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520526050181' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 226. ARI APRIADI
UPDATE rekap SET id_pernyataan1 = 'B-07.20.227/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520525110257' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 227. Adetya Anhar
UPDATE rekap SET id_pernyataan1 = 'B-07.20.228/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520526050249' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 228. Aminullah
UPDATE rekap SET id_pernyataan1 = 'B-07.20.229/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520525110180' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 229. OPHIYANSYAH ACHRUL PUTRA
UPDATE rekap SET id_pernyataan1 = 'B-07.20.230/SE2026/5205/Super.PPL/2026', tgl_pernyataan1 = '2026-07-20' WHERE idsobat = '520526050266' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_pernyataan1 IS NULL OR id_pernyataan1 = ''); -- 230. Waliey muftihakam marza karyadi

-- ================== BAPP Batch 3 (PML dulu, lalu PPL, seq_all 254-262, 21 Juli 2026) ==================
UPDATE rekap SET id_bapp1 = 'B-07.21.254/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520522100030' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 254. JULKARNAIN (PML)
UPDATE rekap SET id_bapp1 = 'B-07.21.255/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520522030029' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 255. SALAHUDDIN (PML)
UPDATE rekap SET id_bapp1 = 'B-07.21.256/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520523050009' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 256. Syahril Sidik (PML)
UPDATE rekap SET id_bapp1 = 'B-07.21.257/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520522100175' AND kegiatan = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 257. Yulianti (PML)
UPDATE rekap SET id_bapp1 = 'B-07.21.258/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520526050181' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 258. ARI APRIADI (PPL)
UPDATE rekap SET id_bapp1 = 'B-07.21.259/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520525110257' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 259. Adetya Anhar (PPL)
UPDATE rekap SET id_bapp1 = 'B-07.21.260/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520526050249' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 260. Aminullah (PPL)
UPDATE rekap SET id_bapp1 = 'B-07.21.261/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520525110180' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 261. OPHIYANSYAH ACHRUL PUTRA (PPL)
UPDATE rekap SET id_bapp1 = 'B-07.21.262/BAPP-I-SE2026/5205.PPK/BA/2026', tgl_bapp1 = '2026-07-21' WHERE idsobat = '520526050266' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026' AND (id_bapp1 IS NULL OR id_bapp1 = ''); -- 262. Waliey muftihakam marza karyadi (PPL)

-- ================== Surat Pernyataan Kepala BPS Batch 3 (dokumen tunggal, KUMULATIF batch 1+2+3, 21 Juli 2026) ==================
INSERT INTO surat_kepala_se2026 (termin, tahun, batch, nomor, tanggal) SELECT 1, '2026', 3, 'B-07.21.003/SE2026/5205/Super.KPL/2026', '2026-07-21' WHERE NOT EXISTS (SELECT 1 FROM surat_kepala_se2026 WHERE termin=1 AND tahun='2026' AND batch=3);

SELECT
  (SELECT COUNT(*) FROM rekap WHERE id_pernyataan1 LIKE 'B-07.20.%/Super.PPL/2026') AS pernyataan_ppl_batch3,
  (SELECT COUNT(*) FROM rekap WHERE id_pernyataan1 LIKE 'B-07.20.%/Super.PML/2026') AS pernyataan_pml_batch3,
  (SELECT COUNT(*) FROM rekap WHERE id_bapp1 LIKE 'B-07.21.%/BAPP-I-SE2026/%') AS bapp_batch3,
  (SELECT COUNT(*) FROM surat_kepala_se2026 WHERE termin=1 AND tahun='2026' AND batch=3) AS kepala_batch3;
