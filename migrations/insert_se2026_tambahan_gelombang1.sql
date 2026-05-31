-- ============================================================
-- Tambah 4 petugas gelombang 1 + nomor surat (idempotent)
-- Aman dijalankan berulang kali - tidak duplikat
-- ============================================================

-- MITRA
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050204', 'YULIATI', '', 'Jl. Lintas Lakey, Desa Adu, Dusun Woro', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050204' AND tahun='2026');

INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050015', 'ARIS MUNANDAR', '', 'Dusun woro jaya Rt/001 Rw/000', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050015' AND tahun='2026');

INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050266', 'WALIEY MUFTIHAKAM MARZA KARYADI', '', 'Dusun Permata Hijau RT/RW 001/000', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050266' AND tahun='2026');

INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050108', 'VERA NOVITASARI', '', 'Jalan lintas malaju desa Kramat.', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050108' AND tahun='2026');

-- REKAP (skip jika sudah ada)
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050204', 'ipds5205', 'YULIATI', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-054/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050204' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050015', 'ipds5205', 'ARIS MUNANDAR', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-086/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050015' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050266', 'ipds5205', 'WALIEY MUFTIHAKAM MARZA KARYADI', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-011/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050266' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050108', 'ipds5205', 'VERA NOVITASARI', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-079/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050108' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
