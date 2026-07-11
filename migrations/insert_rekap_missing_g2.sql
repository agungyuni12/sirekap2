-- Rekap missing Gelombang 2
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030028', 'ipds5205', 'Nurwahidah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-110/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030028' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-110/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522030028' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110054', 'ipds5205', 'Nurradian', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-150/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110054' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-150/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110054' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
