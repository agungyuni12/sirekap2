-- ============================================================
-- Tambah kolom 'batch' ke lk_ppk_payable_se2026 (backfill batch=1 utk
-- data yg sudah ada), lalu tambahkan 14 PML + 17 PPL batch 2 - seq & seq_all
-- LANJUT dari batch 1 (tidak reset), supaya nomor batch 1 yg sudah
-- di-assign & dipakai di dokumen tidak berubah.
--
-- PML batch 2: SEMUA PML yg belum pernah dibayar di batch 1 SEKARANG
-- dibayar, KECUALI 4 yg dikecualikan: Yulianti (520522100175), JULKARNAIN
-- (520522100030), Syahril Sidik (520523050009), SALAHUDDIN (520522030029 -
-- PPL binaannya, Waliey Muftihakam Marza Karyadi 520526050266, belum ada
-- realisasi SLS sama sekali). 8 dari 14 PML ini tidak muncul di file Excel
-- Batch 2 (PPL binaannya belum lengkap batch 2) - target/realisasi mereka
-- dihitung dari data SLS batch 1 yg sudah ada di lk_ppk_termin1_se2026
-- (tabel itu tidak dipisah per-batch).
-- ============================================================

ALTER TABLE lk_ppk_payable_se2026 ADD COLUMN IF NOT EXISTS batch INT NOT NULL DEFAULT 1;
UPDATE lk_ppk_payable_se2026 SET batch = 1 WHERE batch IS NULL OR batch = 0;

INSERT INTO lk_ppk_payable_se2026 (idsobat, nama, role, seq, seq_all, batch) VALUES
('520522020005','A''an Aditya Munandar','pml',15,223,2),
('520522010029','ADI NEGORO HERLAMBANG','pml',16,224,2),
('520522030058','ADIMAN','pml',17,225,2),
('520522090045','Adi Hidayat','pml',18,226,2),
('520522030054','HAERUDDIN','pml',19,227,2),
('520522030018','Hanafi','pml',20,228,2),
('520522050001','Kurniawan','pml',21,229,2),
('520522030009','Mia Agustina','pml',22,230,2),
('520523010001','NILA UTAMA','pml',23,231,2),
('520522030014','SUHERMAN','pml',24,232,2),
('520522020017','Suhardin Putra','pml',25,233,2),
('520522030011','TAUFIKURRAMADHAN','pml',26,234,2),
('520522030006','YUSUF','pml',27,235,2),
('520522020015','ZORDIAN','pml',28,236,2),
('520526050193','Amrizal','ppl',209,237,2),
('520522100011','ANWAR RIFAID','ppl',210,238,2),
('520525110011','Ayu wandira','ppl',211,239,2),
('520522100013','Bahril Qamar','ppl',212,240,2),
('520522030017','EKO ANSHARI','ppl',213,241,2),
('520526050273','Husnul Wahyu Lestari','ppl',214,242,2),
('520526050031','IFA NURUL AISYAH','ppl',215,243,2),
('520522100132','MIRJAN ALHALIK','ppl',216,244,2),
('520525110198','Moh Ma''ruf','ppl',217,245,2),
('520526050050','Muhammad Arkham','ppl',218,246,2),
('520525110108','Nurlailah','ppl',219,247,2),
('520523060005','Nurul Khotimah','ppl',220,248,2),
('520522030024','Rachmad Soebari','ppl',221,249,2),
('520526050218','Sri Handayani','ppl',222,250,2),
('520525110028','SRI SUSANTI','ppl',223,251,2),
('520522030021','Titin Fatmawati','ppl',224,252,2),
('520526050243','Triana Amalia ND','ppl',225,253,2);

SELECT batch, role, COUNT(*) AS jumlah, MIN(seq) AS seq_min, MAX(seq) AS seq_max, MIN(seq_all) AS seq_all_min, MAX(seq_all) AS seq_all_max FROM lk_ppk_payable_se2026 GROUP BY batch, role ORDER BY batch, role;
