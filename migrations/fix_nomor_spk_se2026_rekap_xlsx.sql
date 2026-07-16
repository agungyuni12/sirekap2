-- ============================================================
-- Perbaikan Nomor SPK SE2026 (PML & PPL)
-- Sumber: Rekap Nomor SPK PML dan PPL SE 2026.xlsx
-- Diverifikasi langsung terhadap dump database seleksimitra (3).sql
-- Dibuat otomatis oleh Claude Code - Juli 2026
--
-- Hasil pengecekan seluruh 265 petugas (32 PML + 233 PPL) di xlsx
-- terhadap tabel `mitra` dan `rekap` yang ada di dump:
--   - 263 baris SUDAH BENAR (id_spk di database sudah sama persis
--     dengan xlsx) -> tidak perlu diapa-apakan.
--   - 2 baris BELUM ADA baris rekap-nya sama sekali untuk kegiatan
--     SE2026 (idsobat & mitra-nya sudah terdaftar di tabel `mitra`,
--     tapi belum pernah di-insert ke `rekap` dengan kegiatan
--     'Pendataan Sensus Ekonomi 2026' tahun 2026) -> perlu INSERT baru.
--
-- SELALU BACKUP tabel `rekap` sebelum menjalankan script ini:
--   CREATE TABLE rekap_backup_20260716 AS SELECT * FROM rekap;
--
-- Pastikan sudah pilih database yang benar dulu:
--   USE seleksimitra;
-- ============================================================

-- ------------------------------------------------------------
-- 1. Rupi Rahmiati (PPL #163, SPK B-190)
--    idsobat 520525110042 sudah ada di tabel mitra (alamat: Rt. 007
--    Rw. 000 Dusun Wawonduru Timur Desa Wawonduru, Woja), tapi belum
--    pernah ada baris rekap SE2026 untuk dia.
-- ------------------------------------------------------------
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110042', 'ipds5205', 'RUPI RAHMIATI', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, '11561525', '2.5', 'O-B', '4624610', '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-190/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat = '520525110042' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026');

-- ------------------------------------------------------------
-- 2. Nurwahidah (PPL #196, SPK B-225) - beralamat Jln. Lintas Sumbawa
--    RT/RW 012/005, Desa Tekasire, Kec. Manggelewa (idsobat 520523030037).
--    INI BUKAN Nurwahidah yang sama dengan #208 (SPK B-239, alamat Dusun
--    Bolo Baka/Bakajaya, idsobat 520522030028) - yang itu id_spk-nya
--    SUDAH BENAR di database, jangan diubah.
--    Mitra-nya baru ada baris tahun 2025 (id=286), belum ada baris 2026,
--    jadi ditambahkan juga supaya konsisten dengan mitra lain.
-- ------------------------------------------------------------
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030037', 'NURWAHIDAH', '5205074404930003', 'Jln.Lintas Sumbawa.RT/RW.012/005', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat = '520523030037' AND tahun = '2026');

INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030037', 'ipds5205', 'NURWAHIDAH', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, '11561525', '2.5', 'O-B', '4624610', '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-225/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat = '520523030037' AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026');
