-- ============================================================
-- LK PPK Termin 1 BATCH 3 - Kabupaten Dompu (Realisasi Fasih)
-- Sumber: LK PPK Termin 1 Batch 3 - Kabupaten Dompu (Realisasi Fasih).xlsx
-- Batch 3 = 5 PPL yang sebelumnya dikecualikan dari batch 2 (Aminullah,
-- Ari Apriadi, Waliey Muftihakam Marza Karyadi, Ophiyansyah Achrul Putra,
-- Adetya Anhar) - sekarang datanya lengkap, sehingga 4 PML mereka
-- (JULKARNAIN, SALAHUDDIN, Syahril Sidik, Yulianti - yang sebelumnya
-- dikecualikan di batch 1 & 2 karena PPL binaan ini belum lengkap) juga
-- bisa dibayar.
--
-- PENTING: kolom 'realisasi' di batch ini = Realisasi Jumlah (fasih_total -
-- fasih_open - jumlah_draft), SAMA basis dengan batch 1 (BEDA dari batch 2
-- yang pakai Realisasi Usaha BKU saja). 'target' = Target Keluarga + Target
-- Usaha (sama semua batch).
--
-- 43 baris SLS, 19 prioritas. Data ditarik dari se2026 (8).sql per 2026-07-25.
--
-- APPEND ke tabel yang sama dgn batch 1 & 2 (idsobat beda, tidak ada bentrok) -
-- TIDAK di-TRUNCATE, supaya data batch 1 & 2 tidak hilang. PML di batch ini
-- sudah punya baris SLS lain dari batch 1 (PPL binaan lain yg sudah payable) -
-- target/realisasi PML dihitung otomatis kumulatif dari SEMUA baris (lama+baru).
-- ============================================================

INSERT INTO lk_ppk_termin1_se2026 (pml_idsobat, pml_nama, ppl_idsobat, ppl_nama, kode_kec, kode_desa, kode_sls, target, realisasi, prioritas) VALUES
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','000100',80,83,1),
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','000200',109,115,1),
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','000300',78,81,1),
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','000400',77,9,0),
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','000600',78,1,0),
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','000900',137,0,0),
('520523050009','Syahril Sidik','520525110180','OPHIYANSYAH ACHRUL PUTRA','051','003','100200',4,0,0),
('520522100175','Yulianti','520525110257','Adetya Anhar','010','003','000200',73,0,0),
('520522100175','Yulianti','520525110257','Adetya Anhar','010','003','001000',248,295,1),
('520522100175','Yulianti','520525110257','Adetya Anhar','010','007','001400',31,37,1),
('520522100175','Yulianti','520525110257','Adetya Anhar','010','007','001600',51,0,0),
('520522100175','Yulianti','520525110257','Adetya Anhar','010','007','200300',2,5,1),
('520522100175','Yulianti','520525110257','Adetya Anhar','010','008','000100',87,1,0),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','000300',110,131,1),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','000400',120,139,1),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','000500',55,13,0),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','000600',55,3,0),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','000900',33,38,1),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','001000',32,10,0),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','100100',0,1,0),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','015','100200',1,1,1),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','016','000600',76,0,0),
('520522100030','JULKARNAIN','520526050181','ARI APRIADI','050','016','300300',2,2,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','013','200200',2,2,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','000200',217,12,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','000500',66,68,1),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','000600',102,110,1),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','000700',38,42,1),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','000800',35,40,1),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','001100',29,20,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','001200',80,34,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','100100',10,10,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','100200',50,50,0),
('520522100030','JULKARNAIN','520526050249','Aminullah','050','014','200300',49,49,1),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','000500',142,161,1),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','001200',95,112,1),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','001300',89,22,0),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','001400',52,68,1),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','001500',91,9,0),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','001600',45,37,0),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','100200',2,2,0),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','100300',7,7,1),
('520522030029','SALAHUDDIN','520526050266','Waliey muftihakam marza karyadi','050','009','200100',14,14,0);

SELECT 'Migration add_lk_ppk_termin1_batch3_se2026 selesai.' AS status, COUNT(*) AS total_baris_baru, SUM(prioritas) AS jumlah_prioritas
FROM lk_ppk_termin1_se2026
WHERE pml_idsobat IN ('520523050009','520522100175','520522100030','520522030029')
  AND ppl_idsobat IN ('520525110180','520525110257','520526050181','520526050249','520526050266');
