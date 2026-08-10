-- ============================================================
-- Tambah baris `mitra` untuk 2 dari 3 petugas pengganti batch 4 SE2026
-- (IIN NURROHMADANIL, Sarwan) - belum pernah terdaftar di tabel mitra
-- sama sekali. Moh Naufal SUDAH ada (idsobat 520526050093, NIK sudah
-- benar 5205012706950003), jadi tidak disentuh.
--
-- Data (alamat, kecamatan) dari SE2026_66_5205_exportmitra_2026-08-10_073809.xlsx.
-- NIK dari input manual user: Sarwan 5272023112870042, Iin 5205016411020003.
--
-- Idempotent: pakai INSERT...SELECT...WHERE NOT EXISTS supaya aman dijalankan ulang.
-- ============================================================

INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050109', 'IIN NURROHMADANIL', '5205016411020003', 'Kelurahan Kandai Satu RT 006 RW 003', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat = '520526050109');

INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526070001', 'Sarwan', '5272023112870042', 'Dusun Sori Mangge RT 001 RW 001', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat = '520526070001');

SELECT idsobat, nmitra, nik, alamat, kecamatan FROM mitra WHERE idsobat IN ('520526050093','520526050109','520526070001');
