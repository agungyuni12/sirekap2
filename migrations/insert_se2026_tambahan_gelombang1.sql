-- ============================================================
-- Tambah 3 petugas gelombang 1 yang tidak ada di export awal
-- ============================================================

-- MITRA
INSERT IGNORE INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun) VALUES ('520526050204', 'YULIATI', '', 'Jl. Lintas Lakey, Desa Adu, Dusun Woro', 'Hu`U', '2026');
INSERT IGNORE INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun) VALUES ('520526050015', 'ARIS MUNANDAR', '', 'Dusun woro jaya Rt/001 Rw/000', 'Manggalewa', '2026');
INSERT IGNORE INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun) VALUES ('520526050266', 'WALIEY MUFTIHAKAM MARZA KARYADI', '', 'Dusun Permata Hijau RT/RW 001/000', 'Kempo', '2026');

-- REKAP (langsung dengan nomor surat)
INSERT IGNORE INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk) VALUES ('520526050204', 'ipds5205', 'YULIATI', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 4629000, 1, 'O-B', 4629000, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-054/SPK-SE2026/3201/PL.200/2026');
INSERT IGNORE INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk) VALUES ('520526050015', 'ipds5205', 'ARIS MUNANDAR', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 4629000, 1, 'O-B', 4629000, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-086/SPK-SE2026/3201/PL.200/2026');
INSERT IGNORE INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk) VALUES ('520526050266', 'ipds5205', 'WALIEY MUFTIHAKAM MARZA KARYADI', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 4629000, 1, 'O-B', 4629000, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-011/SPK-SE2026/3201/PL.200/2026');