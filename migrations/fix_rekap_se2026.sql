-- ============================================================
-- Perbaikan Rekap SE2026: hapus extra, tambah pengganti
-- ============================================================

-- ====== HAPUS 4 REKAP PPL YANG TIDAK ADA DI SPK ======
DELETE FROM rekap WHERE idsobat='520522100119' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
DELETE FROM rekap WHERE idsobat='520523110113' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
DELETE FROM rekap WHERE idsobat='520525110004' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
DELETE FROM rekap WHERE idsobat='520525110197' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';

-- ====== HAPUS MITRA 2026 YANG TIDAK ADA DI SPK ======
DELETE FROM mitra WHERE idsobat='520522100119' AND tahun='2026';
DELETE FROM mitra WHERE idsobat='520523110113' AND tahun='2026';
DELETE FROM mitra WHERE idsobat='520525110004' AND tahun='2026';
DELETE FROM mitra WHERE idsobat='520525110197' AND tahun='2026';

-- ====== FIX AMELIA NARASTA FARATIKA (#N/A → ID benar) ======
UPDATE rekap SET idsobat='520525110090' WHERE idsobat='#N/A' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
UPDATE mitra SET idsobat='520525110090' WHERE idsobat='#N/A' AND tahun='2026';

-- ====== TAMBAH SRI RAHMAWATI (pengganti Nurwahidah G2 SPK-110) ======
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110079', 'Sri Rahmawati', '', 'Desa Tanju', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110079' AND tahun='2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110079', 'ipds5205', 'Sri Rahmawati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-110/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110079' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk='B-110/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523110079' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
