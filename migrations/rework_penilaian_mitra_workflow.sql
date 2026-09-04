-- ============================================================
-- Rework Penilaian Mitra sesuai "SIREKAP Penilaian Mitra.docx":
--
-- 1) Daftar kegiatan khusus Penilaian Mitra, terpisah dari tabel
--    `kegiatan` (anggaran) yang dipakai modul REKAPITULASI — sesuai
--    permintaan agar daftar kegiatan penilaian tidak tercampur dengan
--    daftar kegiatan SIREKAP yang sudah ada. Diisi dari
--    Daftar_Nama_Kegiatan_Rapi.xlsx.
-- 2) Kolom tambahan pada evaluasi_petugas untuk field yang diminta di
--    form "Input Penilaian": Wilayah tugas (kecamatan), Periode, Tahun,
--    Tanggal penilaian.
-- 3) Alur konfirmasi Subject Matter untuk penilaian PML Mitra:
--    PML Mitra dinilai atasan langsung organik (PJK/Korwil) -> Subject
--    Matter (level 'admin') mengkonfirmasi (setuju/tidak) -> jika tidak
--    setuju, Subject Matter menilai ulang dan skor akhir = rata-rata
--    kedua nilai. Dibatasi ke peran_yang_dinilai='pml' karena volume
--    penilaian PPL terlalu besar untuk dikonfirmasi satu-satu (lihat
--    catatan di internal/models/penilaian.go).
--
-- CATATAN KOMPATIBILITAS: kolom evaluasi_petugas.kegiatan_id sebelum
-- migrasi ini merujuk ke id pada tabel `kegiatan` (anggaran). Setelah
-- migrasi ini, kegiatan_id pada baris BARU merujuk ke id pada tabel
-- `penilaian_kegiatan` (ruang id yang berbeda). Baris lama (jika ada)
-- tidak diubah — cukup aman karena fitur ini belum pernah dipakai di
-- produksi (belum ada data nyata).
-- ============================================================

CREATE TABLE IF NOT EXISTS penilaian_kegiatan (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    urutan INT NOT NULL DEFAULT 0,
    UNIQUE KEY uniq_nama (nama)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT IGNORE INTO penilaian_kegiatan (id, nama) VALUES
(1, 'Survei Penduduk Antar Sensus (Supas)'),
(2, 'Survei Angkatan Kerja Nasional (Sakernas)'),
(3, 'Survei Sosial Ekonomi Nasional (Susenas)'),
(4, 'Pengolahan Susenas'),
(5, 'Pengolahan Seruti'),
(6, 'Survei Ekonomi Rumah Tangga Triwulanan (Seruti)'),
(7, 'Survei Nasional Literasi dan Inklusi Keuangan (SNLIK)'),
(8, 'Pendataan Potensi Desa (Podes)'),
(9, 'Survei Khusus Triwulanan Neraca Produksi - Jasa'),
(10, 'Survei Khusus Neraca Produksi (SKNP)'),
(11, 'Survei Khusus Triwulanan Neraca Produksi - Barang'),
(12, 'Survei Khusus Lembaga Non Profit yang Melayani Rumah Tangga Triwulanan'),
(13, 'Survei Khusus Studi Penyusunan Perubahan Inventori'),
(14, 'Survei SKPS'),
(15, 'Survei SKSPPI'),
(16, 'Updating Direktori LNPRT Di Kab/Kota'),
(17, 'Pengolahan Survei SKNP'),
(18, 'Pengolahan Survei SKTNP Barang'),
(19, 'Pengolahan Survei SKTNP Jasa'),
(20, 'Survei Penyusunan Disagregasi PMTB'),
(21, 'Survei Jasa Penunjang Angkutan (Pergudangan dan Kurir)'),
(22, 'Survei Harga Perdesaan (SHPed)'),
(23, 'Survei Harga Perdagangan Besar (SHPB)'),
(24, 'Survei Harga Produsen (SHP)'),
(25, 'Pengolahan Survei Harga Kemahalan Konstruksi'),
(26, 'Survei Kemahalan Konstruksi (SHKK)'),
(27, 'Survei Perdagangan Barang Domestik'),
(28, 'Survei Pola Distribusi'),
(29, 'Survei Pola Usaha Non Pertanian'),
(30, 'Pengolahan Survei Perdagangan Barang Domestik'),
(31, 'Pengolahan Survei Pola Distribusi Barang'),
(32, 'Pengolahan Survei Pola Usaha Non Pertanian'),
(33, 'Pendataan Pemetaan Wilkerstat SE2026'),
(34, 'Survei Harga Konsumen Perdesaan (HKD)'),
(35, 'Survei Harga Perdagangan Besar (SHPB) Mingguan'),
(36, 'Survei Harga Produsen Identifikasi Komoditas Utama'),
(37, 'Pengolahan Survei Harga Konsumen Perdesaan (HKD)'),
(38, 'Pengolahan Survei Harga Perdagangan Besar'),
(39, 'Pengolahan Survei Harga Perdagangan Besar Mingguan'),
(40, 'Survei Harga Produsen Sektor Industri Pengolahan Dan Sektor Pertambangan Dan Penggalian (HP)'),
(41, 'Survei Harga Produsen Sektor Jasa (HPJ)'),
(42, 'Pengolahan Survei Harga Produsen Sektor Industri Pengolahan Dan Sektor Pertambangan Dan Penggalian (HP)'),
(43, 'Survei Harga Produsen Sektor Pertanian (HPT)'),
(44, 'Pengolahan Survei Harga Produsen Sektor Pertanian (HPT)'),
(45, 'Pengolahan Survei Harga Produsen Sektor Jasa (HPJ)'),
(46, 'Pengolahan Survei Harga Produsen Perdesaan (HD)'),
(47, 'Survei Industri Besar Sedang (IBS)'),
(48, 'Pemutakhiran Direktori Perusahaan Awal (DPA)'),
(49, 'Survei Tahunan Perusahaan Industri Manufaktur'),
(50, 'Listing Survei Industri Mikro dan Kecil (IMK)'),
(51, 'Survei Konstruksi Tahunan (SKTH)'),
(52, 'Survei Captive Power'),
(53, 'Survei Perusahaan Tahunan Air Bersih'),
(54, 'Survei Updating Direktori Perusahaan Awal'),
(55, 'Survei Air Bersih'),
(56, 'Survei Penggalian Badan Hukum'),
(57, 'Survei Konstruksi Triwulanan (SKTR)'),
(58, 'Updating Direktori Perusahaan Konstruksi'),
(59, 'Updating Direktori Perusahaan Pertambangan dan Energi'),
(60, 'Survei Tahunan Usaha Penggalian Bahan Industri dan Konstruksi'),
(61, 'Survei Statistik Keuangan Pemerintah Daerah (K3)'),
(62, 'Survei Statistik Lembaga Keuangan Koperasi Simpan Pinjam (SLK)'),
(63, 'Survei Perusahaan/Usaha Penyedia Makan Minum (VRESTUMB)'),
(64, 'Survei Usaha/Perusahaan Penyedia Jasa Akomodasi Bulanan'),
(65, 'Survei Usaha/Perusahaan Penyedia Jasa Akomodasi Tahunan'),
(66, 'Survei Bidang Usaha Pariwisata (Updating Direktori Survei Usaha Penyedia Jasa Pariwisata)'),
(67, 'Survei Statistik Keuangan BUMD'),
(68, 'Survei E-Commerce'),
(69, 'Pengolahan Survei E-commerce'),
(70, 'Pengolahan Listing Survei IMK'),
(71, 'Pengolahan Survei Air Bersih'),
(72, 'Pengolahan Survei Captive Power'),
(73, 'Pengolahan SKTH'),
(74, 'Pengolahan SKTR'),
(75, 'Pengolahan Survei IMK'),
(76, 'Survei Pelabuhan Perikanan dan Tempat Pelelangan Ikan'),
(77, 'Survei Pemotongan Ternak Bulanan'),
(78, 'Survei Pendaratan Ikan Tradisional'),
(79, 'Survei Perusahaan Kehutanan'),
(80, 'Survei Perusahaan Peternakan Tahunan'),
(81, 'Updating Direktori Perikanan'),
(82, 'Survei Amatan KSA Padi'),
(83, 'Survei Pertanian Tanaman Pangan/Ubinan Palawija'),
(84, 'Survei Perusahaan Perkebunan Bulanan'),
(85, 'Survei Perusahaan Perkebunan Tahunan'),
(86, 'Updating Direktori Perusahaan Pertanian (DPP) dan Direktori Usaha Pertanian Lainnya (DUTL)'),
(87, 'Survei Usaha Hortikultura lainnya (VN-Horti)'),
(88, 'Survei Amatan KSA Jagung'),
(89, 'Survei Konsumsi Bahan Pokok Non Rumah Tangga UMK'),
(90, 'Survei Konsumsi Bahan Pokok Non Rumah Tangga UMB'),
(91, 'Survei Konversi Gabah ke Beras'),
(92, 'Updating Survei Kesejahteraan Petani'),
(93, 'Pengolahan Updating Survei Kesejahteraan Petani'),
(94, 'Survei Kesejahteraan Petani'),
(95, 'Survei Produksi Hortikultura (UTL)'),
(96, 'Survei Produksi Hortikultura (UTP)'),
(97, 'Pendataan Lapangan Rumah Tangga Indeks Komstrat'),
(98, 'Survei Pertanian Tanaman Pangan/Ubinan Padi');

-- Field tambahan Input Penilaian (kecamatan/periode/tahun/tanggal) +
-- alur konfirmasi Subject Matter untuk penilaian PML Mitra.
ALTER TABLE evaluasi_petugas
    ADD COLUMN kecamatan VARCHAR(100) NULL AFTER peran_yang_dinilai,
    ADD COLUMN periode VARCHAR(50) NULL AFTER kecamatan,
    ADD COLUMN tahun VARCHAR(5) NULL AFTER periode,
    ADD COLUMN tanggal_penilaian DATE NULL AFTER tahun,
    ADD COLUMN status_konfirmasi VARCHAR(20) NOT NULL DEFAULT 'disetujui' AFTER catatan,
    ADD COLUMN confirmed_by INT NULL AFTER status_konfirmasi,
    ADD COLUMN confirmed_at TIMESTAMP NULL AFTER confirmed_by,
    ADD COLUMN catatan_konfirmasi TEXT NULL AFTER confirmed_at,
    ADD INDEX idx_status_konfirmasi (status_konfirmasi);

-- Perluas level user agar menampung 'korwil' (atasan langsung organik /
-- PJK-Korwil yang menilai PML Mitra Tahap 1). Kolom sudah VARCHAR(20)
-- sejak migrasi create_evaluasi_petugas_table.sql, jadi aman dijalankan.
ALTER TABLE user MODIFY COLUMN level VARCHAR(20) NOT NULL DEFAULT 'user';
