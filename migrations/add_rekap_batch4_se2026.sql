-- ============================================================
-- Tambah baris `rekap` (kegiatan Pendataan Sensus Ekonomi 2026) untuk 3
-- petugas pengganti batch 4 - BELUM PERNAH ada baris rekap kegiatan ini
-- sama sekali utk mereka (beda dari batch 1-3 yang tinggal assign nomor
-- ke baris yang sudah ada). Sekaligus assign nomor SPK (baru pertama kali
-- dapat SPK) - format & sequence SAMA seperti id_spk existing lain
-- (B-NNN/SPK-SE2026/5205/PL.200/2026, url urut, max existing = B-265).
--
-- Honor 3 orang ini Rp 4.629.000 (arahan user - BEDA dari honor PPL
-- reguler Rp 11.561.525), field lain (volume/satuan/hsatuan/mak/waktu)
-- ikut pola SPK PPL reguler lainnya, hsatuan disesuaikan supaya
-- volume x hsatuan = honor (2.5 x 1.851.600 = 4.629.000).
--
-- Idempotent: INSERT...SELECT...WHERE NOT EXISTS (idsobat+kegiatan+tahun).
-- id_pernyataan1/id_bapp1 sengaja NULL disini - diisi terpisah oleh
-- migrations/assign_nomor_surat_se2026_batch4.sql.
-- ============================================================

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, id_spk, realisasi_sls)
SELECT '520526050109', 'ipds5205', 'IIN NURROHMADANIL', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, '4629000', 2.5, 'O-B', '1851600', '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 'B-266/SPK-SE2026/5205/PL.200/2026', 0
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat = '520526050109' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, id_spk, realisasi_sls)
SELECT '520526050093', 'ipds5205', 'Moh Naufal', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, '4629000', 2.5, 'O-B', '1851600', '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 'B-267/SPK-SE2026/5205/PL.200/2026', 0
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat = '520526050093' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, id_spk, realisasi_sls)
SELECT '520526070001', 'ipds5205', 'Sarwan', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, '4629000', 2.5, 'O-B', '1851600', '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 'B-268/SPK-SE2026/5205/PL.200/2026', 0
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat = '520526070001' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026');

SELECT id, idsobat, namamitra, kegiatan, honor, id_spk FROM rekap WHERE idsobat IN ('520526050109','520526050093','520526070001') AND kegiatan = 'Pendataan Sensus Ekonomi 2026';
