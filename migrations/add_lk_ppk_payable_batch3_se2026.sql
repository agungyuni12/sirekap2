-- ============================================================
-- Tambah 4 PML + 5 PPL batch 3 ke lk_ppk_payable_se2026 - seq & seq_all
-- LANJUT dari batch 1+2 (tidak reset). Kolom 'batch' sudah ada dari
-- migrations/add_lk_ppk_payable_batch2_se2026.sql.
--
-- Batch 3 = 5 PPL yang sebelumnya dikecualikan dari batch 2 (Aminullah,
-- Ari Apriadi, Waliey Muftihakam Marza Karyadi, Ophiyansyah Achrul Putra,
-- Adetya Anhar), yang membuat 4 PML mereka (JULKARNAIN, SALAHUDDIN,
-- Syahril Sidik, Yulianti) sekarang juga bisa dibayar.
-- ============================================================

INSERT INTO lk_ppk_payable_se2026 (idsobat, nama, role, seq, seq_all, batch) VALUES
('520522100030','JULKARNAIN','pml',29,254,3),
('520522030029','SALAHUDDIN','pml',30,255,3),
('520523050009','Syahril Sidik','pml',31,256,3),
('520522100175','Yulianti','pml',32,257,3),
('520526050181','ARI APRIADI','ppl',226,258,3),
('520525110257','Adetya Anhar','ppl',227,259,3),
('520526050249','Aminullah','ppl',228,260,3),
('520525110180','OPHIYANSYAH ACHRUL PUTRA','ppl',229,261,3),
('520526050266','Waliey muftihakam marza karyadi','ppl',230,262,3);

SELECT batch, role, COUNT(*) AS jumlah, MIN(seq) AS seq_min, MAX(seq) AS seq_max, MIN(seq_all) AS seq_all_min, MAX(seq_all) AS seq_all_max FROM lk_ppk_payable_se2026 GROUP BY batch, role ORDER BY batch, role;
