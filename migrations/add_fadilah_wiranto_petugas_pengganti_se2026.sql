-- ============================================================
-- Onboarding FADILAH & Wiranto sbg mitra SE2026 (petugas pengganti,
-- baru dibayarkan mulai Termin 2) - kohort sama dgn IIN NURROHMADANIL/
-- Moh Naufal/Sarwan (SPK B-266..270, honor 4629000, batch 4).
-- Data dari user: NIK, ID Sobat, nomor SPK, tanggal SPK 24 Juli 2026.
-- ============================================================

INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun) VALUES
('520526070003','FADILAH','5205010703000002','Dusun Sorisakolo Timur','Dompu','2026'),
('520526070002','Wiranto','5205073009990001','Dusun Ta,a Paju RT 07/RW 03','Manggalewa','2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor,
    volume, satuan, hsatuan, mak, pmwaktu, pswaktu, id_spk, realisasi_sls) VALUES
('520526070003','ipds5205','FADILAH','Pendataan Sensus Ekonomi 2026','Juni-Agustus','2026',2026,'4629000',
    '2.5','O-B','1851600','2902.BMA.006.005.B.521213','2026-06-15','2026-08-31','B-269/SPK-SE2026/5205/PL.200/2026',0),
('520526070002','ipds5205','Wiranto','Pendataan Sensus Ekonomi 2026','Juni-Agustus','2026',2026,'4629000',
    '2.5','O-B','1851600','2902.BMA.006.005.B.521213','2026-06-15','2026-08-31','B-270/SPK-SE2026/5205/PL.200/2026',0);

INSERT INTO lk_ppk_payable_se2026 (idsobat, nama, role, seq, seq_all, batch, excluded_termin2) VALUES
('520526070003','FADILAH','ppl',234,266,4,0),
('520526070002','Wiranto','ppl',235,267,4,0);
