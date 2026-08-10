-- ============================================================
-- LK PPK Termin 1 BATCH 4 (PETUGAS PENGGANTI) - Kabupaten Dompu (Realisasi Fasih)
-- Sumber: LK PPK Termin 1 Batch 4 - Kabupaten Dompu (Realisasi Fasih).xlsx
-- Batch 4 = 3 petugas pengganti yang sebelumnya dikecualikan dari batch 3
-- (Moh Naufal, IIN NURROHMADANIL, Sarwan). PML mereka (BUHARI SALAM,
-- ADI NEGORO HERLAMBANG, ZORDIAN) SUDAH payable sejak batch 1/2 - tidak
-- perlu entry payable baru utk PML, cukup 3 PPL baru batch ini.
--
-- PENTING: sesuai catatan file sumber, kolom 'prioritas' DIABAIKAN utk
-- batch ini (semua 23 baris prioritas=0), berbeda dari batch 1-3 yang
-- pakai kolom sls.prioritas asli.
-- 'target' = Target Keluarga + Target Usaha, 'realisasi' = Realisasi
-- Jumlah (sama basis dgn batch 1 & 3).
--
-- 23 baris SLS, 0 prioritas. Data ditarik dari se2026 (10).sql per 2026-08-10.
--
-- APPEND ke tabel yang sama dgn batch 1-3 (idsobat beda, tidak ada bentrok) -
-- TIDAK di-TRUNCATE. PML di batch ini sudah punya baris SLS lain dari batch
-- sebelumnya - target/realisasi PML dihitung otomatis kumulatif.
-- ============================================================

INSERT INTO lk_ppk_termin1_se2026 (pml_idsobat, pml_nama, ppl_idsobat, ppl_nama, kode_kec, kode_desa, kode_sls, target, realisasi, prioritas) VALUES
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','000100',103,3,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','000200',104,119,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','000300',73,136,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','001200',39,4,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','001300',24,33,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','001400',91,0,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','001500',58,23,0),
('520522030008','BUHARI SALAM','520526050093','Moh Naufal','020','010','001700',115,133,0),
('520522010029','ADI NEGORO HERLAMBANG','520526050109','IIN NURROHMADANIL','030','002','001400',70,86,0),
('520522010029','ADI NEGORO HERLAMBANG','520526050109','IIN NURROHMADANIL','030','003','300600',0,1,0),
('520522010029','ADI NEGORO HERLAMBANG','520526050109','IIN NURROHMADANIL','030','004','000700',55,71,0),
('520522010029','ADI NEGORO HERLAMBANG','520526050109','IIN NURROHMADANIL','030','004','001000',114,135,0),
('520522010029','ADI NEGORO HERLAMBANG','520526050109','IIN NURROHMADANIL','030','004','001100',128,169,0),
('520522010029','ADI NEGORO HERLAMBANG','520526050109','IIN NURROHMADANIL','030','007','001100',134,152,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001000',106,64,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001100',72,75,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001200',60,61,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001300',17,17,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001400',28,14,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001600',42,43,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','001700',30,32,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','002100',7,7,0),
('520522020015','ZORDIAN','520526070001','Sarwan','060','011','002400',31,31,0);

SELECT 'Migration add_lk_ppk_termin1_batch4_se2026 selesai.' AS status, COUNT(*) AS total_baris_baru, SUM(prioritas) AS jumlah_prioritas
FROM lk_ppk_termin1_se2026
WHERE ppl_idsobat IN ('520526050093','520526050109','520526070001');
