-- ============================================================
-- SQL untuk Mitra, Rekap & SPK Gelombang 2
-- Serta penambahan SPK untuk Taufikurramadhan (Gelombang 1)
-- ============================================================

-- ================== MITRA GELOMBANG 2 ==================
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522020015', 'Zordian', '', 'Jalan Datuk anggrat dusun suka jaya desa kadindi', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522020015' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030058', 'Adiman', '', 'Jalan lintas calabai pekat dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030058' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050003', 'Rani Annisa Sari', '', 'Jalan Lintas Pasar Senin, Dusun Bagek Payung, Desa Kadindi Barat, Kecamatan Pekat, Kabupaten Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050003' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050026', 'Robiatul Adawiyah', '', 'Dusun Sorikalate, Des. Kadindi Barat', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050026' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050027', 'Sandikawati', '', 'Desa Doropeti Dusun Doropeti Rt 002 rw 000', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050027' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050035', 'Sari Ulandari', '', 'Dusun pade maju', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050035' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050004', 'Sri Andriani', '', 'Jln.lintas labuan kananga', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050004' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050023', 'Suci Purnaningsih', '', 'RT/RW: 001/001-Dusun Pekat II', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050023' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050133', 'Susi Susanti', '', 'Pertigaan jalan menuju pulau satonda', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050133' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030109', 'Jumrah', '', 'Jl Lintas Wisata Taman Nasional Tambora Desa Tambora Dusun Pancasila Kec Pekat Kab Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030109' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110145', 'Windi', '', 'Jln datuk anggrat dusun pusaka kadindi', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110145' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520524090001', 'Meri Mariani', '', 'Jalan Lintas Calabai', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520524090001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100223', 'Leni Karlina', '', 'RT 02 Dusun Suka Jaya, Desa Kadindi, Kecamatan Pekat, Kabupaten Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100223' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100141', 'Mita Sayuti', '', 'Jl. Datuk Anggrat Dusun Suka Jaya,Desa Kadindi, kecematan Pekat, Kab. Dompu NTB', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100141' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090013', 'Muslimin Akbar Ars', '', 'Jalan Lintas Sumbawa Gang Lingkar', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090013' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100046', 'Baiq Nurfitriani Rahmawati', '', 'Jalan lintas sumbawa dompu', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100046' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110016', 'Ega Nur Munziatunnas', '', 'Dusun Mada oi u,a RT 001 RW 000', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110016' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050131', 'Lastry', '', 'Dusun Jatimengi Desa Teka sire Kecamatan manggelewa Kabupaten Dompu', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050131' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110086', 'Mita Rahmatullah', '', 'Dusun jatibaru, desa tekasire', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110086' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100027', 'Nurjanah', '', 'Jalan Lintas Sumbawa', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100027' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030028', 'Nurwahidah', '', 'Dusun Bolo Baka', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030028' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050215', 'BAIQ ERINE WAHYU CHANDRA NIRWANA', '', 'Desa Tanju', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050215' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050238', 'Sumiati', '', 'Dusun meci angi desa soriutu kecematan Manggelewa', 'MANGGALEWA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050238' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522020017', 'Suhardin Putra', '', 'Jalan.Datuk Anggrat Lintas Tambora', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522020017' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100209', 'Abubakar Ismail', '', 'Jalan lintas calabai dusun sortatanga', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100209' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110076', 'Imran', '', 'Dusun Ladore Desa Karombo', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110076' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050025', 'M. Arwenda Prayogi', '', 'Jln. Lintas Gunung Tambora', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050025' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110255', 'Muamar Surajudin', '', 'Dusun Pancasila, Desa Tambora', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110255' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050187', 'Muhammad Nur Azhari', '', 'Jalan Datuk Anggrat Dusun Karang Juli RT 001 RW 001', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050187' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110115', 'Nugie Akbar', '', 'Dusun pancasila, Desa Tambora', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110115' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110080', 'Restu Subroto', '', 'Dusun Bagik Payung RT/RW 001/001 Desa Kadindi Barat Kecamatan Pekat Kabupaten Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110080' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110208', 'Subhan', '', 'Jalan Lintas Calabai', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110208' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050007', 'Veri Irawan', '', 'Jalan : Calabai-Dompu Desan Pekat. Kec Pekat. RT 001/001 Kab.Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050007' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050015', 'Wahyudin', '', 'Dusun Pancasila Desa Tambora Rt 002 Rw 000', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050015' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050066', 'Amelia', '', 'Dusun karang anyar, Desa Beringin Jaya, Kecamatan Pekat, kab. Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050066' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110228', 'Anggun Anabela Yustika Putri', '', 'Desa Calabai kec pekat kab dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110228' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110048', 'Annisa', '', 'Jalan Datu anggrat', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110048' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110074', 'Astuti', '', 'Jln Lintas labuan kananga', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110074' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050042', 'Etty Moto Logis', '', 'Dusun Jonggat RT 001 RW 005', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050042' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050032', 'Fani', '', 'Dusun Tanjung Pasir Desa Calabai', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050032' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050048', 'Haifa Paradisa', '', 'Jalan lintas calabai dusun karang juli desa kadindi', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050048' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050089', 'Levy Viola Ovaliani', '', 'Jl. Datuk Anggrat Dusun Dasan Nangka', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050089' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050002', 'Titian Martini', '', 'Jalan lintas calabai desa pekat dusun pekat 2', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050002' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110009', 'Lisnawati', '', 'Jl. Datuk Anggrat, Dusun Suka Jaya, Desa Kadindi', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110009' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050097', 'Yulia Safitri', '', 'Dusun Karang Juli', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050097' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050088', 'Miftahul Riyadah', '', 'Jalan datuk anggrat dusun pusaka Kadindi', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050088' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050044', 'Nawasari Dwi Sulistiawati', '', 'Jalan Datuk Anggrat, Desa Kadindi, Kecamatan Pekat, Kabupaten Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050044' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030100', 'Nilmawani', '', 'Dusun Calabai Atas RT/RW 003/001 Desa Calabai', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030100' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050264', 'Nurrul Fadillah', '', 'Jl Calabai Dompu', 'PEKAT', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050264' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522020005', 'A''An Aditya Munandar', '', 'Jln. Lintas Mbawi Desa Wawonduru', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522020005' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090045', 'Adi Hidayat', '', 'Jalan lintas desa saneo', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090045' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110107', 'Ade Arahma', '', 'Jalan Diponegoro, Lingkungan V Kelurahan Montabaru', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110107' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050045', 'Adrian', '', 'Lingk. Simpasai', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050045' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110171', 'Anshari Nawawi', '', 'Jl. Teuku Umar RT 15 RW 06 Lingkungan Renda', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110171' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050272', 'Dedi Rachmat', '', 'Lingkungan Simpasai rt 006 rw 003', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050272' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110050', 'Adya Ragita Cahyani', '', 'Lingkungan Bali Bunga. Kelurahan Kandai Dua Kecamatan Woja', 'DOMPU', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110050' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110262', 'Aini', '', 'Jalan Lintas Ria, Dusun Tirtamengi Mekar Desa Riwo Kec.Woja Kab.Dompu', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110262' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110202', 'Ainun Nasirah', '', 'Jalan lintas mbawi dusun rabatumpu, RT 011 RW 000', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110202' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522010023', 'Atikasari', '', 'Buncu utara Rt. 02 Masuk gang mesjid suhada', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522010023' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050039', 'Nurul Fitri', '', 'DUSUN SELAPARANG RT 003', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050039' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110054', 'Nurradian', '', '', 'Kecamatan Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110054' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050218', 'Sri Handayani', '', 'Lingkungan RT 008RW 004', 'DOMPU', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050218' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110199', 'Dirta Uari Darniati', '', 'Jalan Ompu Buncu, Buncu Utara, Desa Matua, Kecamatan Woja, Kabupaten Dompu', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110199' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050192', 'Dyah Rahayu', '', 'Jalan Ompu Beko, Dusun Selaparang, Desa Matua, RT/RW 02, Dompu, Nusa Tenggara Barat', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050192' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110149', 'Ersa Nur Ulfa Islamia', '', 'Jl. Lintas mbawi, Dusun Rato Baka, RT/RW. 001/001', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110149' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050001', 'Esa Ariani', '', 'Jln. Lintas Sumbawa Bima', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520524090003', 'Feni Alfaonita', '', 'ling.bukit larema', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520524090003' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110259', 'Fithria Anggraeni', '', 'Jalan Lintas Sumbawa RT/RW 001/002', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110259' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100206', 'Fitriani', '', 'Dusun Sigi RT.04', 'WOJA', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100206' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110012', 'Afdiansyah', '', 'Jl. Lintas Lakey Dusun Lawiti Desa Tembalae Kecamatan Pajo', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110012' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050261', 'Dermawan', '', 'Jalan lintas felojanga', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050261' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030009', 'Mia Agustina', '', 'Dusun Wera RT/RW 010/005 Desa Lepadi Kecamatan Pajo Kabupaten Dompu', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030009' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100152', 'Tri Nuril Safitriani', '', 'Jl.lintas lakey RT 01 RW 01 Dusun Rasabou', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100152' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050033', 'Dhia Istiqomah', '', 'Jln Lintas Lakey, Dusun Fupu, RT 002 RW 003, Desa Ranggo, Kecamatan Pajo, Kabupaten Dompu, NTB', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050033' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110216', 'Dwi Fathir', '', 'jalan lintas lakey, dusun restu, desa ranggo, Rt 01, Rw 04, kec pajo, kab dompu', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110216' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110194', 'Ersa Riga Puspita', '', 'Jlan Lintas Lakey Dusun Pajo Permai Rt/Rw 004/002', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110194' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050254', 'Fathuljannah', '', 'Jalan lintas felo janga, dusun nata desa lune kec. Pajo kab. Dompu', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050254' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '#N/A', 'Amelia Narasta Faratika', '', '', 'Kecamatan Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='#N/A' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110213', 'Jihan Tri Hapsari', '', 'jalan lintas lakey, dusun mangga dua, desa ranggo, rt 003, rw 001', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110213' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110212', 'Jumiati', '', 'Jalan lintas Lakey', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110212' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050068', 'Kurniati', '', 'Dusun Dorotoi', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050068' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050034', 'Nur Fitrah', '', 'Jln lintas lakey dusun fupu Desa Ranggo kecamatan pajo Dompu.', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050034' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110192', 'Nur Fuziatun Islamiah', '', 'Jalan lintas lakey dusun rasa bou', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110192' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030014', 'Nurhidayati', '', 'Jl. Lintas Jambu Dusun Dorotoi', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050013', 'Nursani', '', 'Ranggo', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050013' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050012', 'Nurul Wahida', '', 'Jl. Lintas Lakey Dusun Wera', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050012' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050014', 'Ulvia Mardiani', '', 'Jalan lintas lakey Dusun Mangga Dua', 'PAJO', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100175', 'Yulianti', '', 'Dusun sawe rt 07 rw 02', 'HU`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100175' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100016', 'Ajwar Anas', '', 'dusun pandai RT/Rw 007/004 desa kareke', 'HU`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100016' AND tahun='2026');

-- ================== REKAP GELOMBANG 2 ==================
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522020015', 'ipds5205', 'Zordian', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-090/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522020015' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-090/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522020015' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030058', 'ipds5205', 'Adiman', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-091/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030058' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-091/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522030058' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050003', 'ipds5205', 'Rani Annisa Sari', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-092/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-092/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050026', 'ipds5205', 'Robiatul Adawiyah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-093/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050026' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-093/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050026' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050027', 'ipds5205', 'Sandikawati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-094/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050027' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-094/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050027' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050035', 'ipds5205', 'Sari Ulandari', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-095/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050035' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-095/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050035' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050004', 'ipds5205', 'Sri Andriani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-096/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050004' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-096/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050004' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050023', 'ipds5205', 'Suci Purnaningsih', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-097/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050023' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-097/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050023' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050133', 'ipds5205', 'Susi Susanti', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-098/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050133' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-098/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050133' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030109', 'ipds5205', 'Jumrah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-099/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030109' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-099/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523030109' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110145', 'ipds5205', 'Windi', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-100/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110145' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-100/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110145' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520524090001', 'ipds5205', 'Meri Mariani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-101/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520524090001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-101/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520524090001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100223', 'ipds5205', 'Leni Karlina', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-102/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100223' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-102/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100223' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100141', 'ipds5205', 'Mita Sayuti', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-103/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100141' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-103/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100141' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522090013', 'ipds5205', 'Muslimin Akbar Ars', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-104/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522090013' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-104/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522090013' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100046', 'ipds5205', 'Baiq Nurfitriani Rahmawati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-105/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100046' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-105/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100046' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110016', 'ipds5205', 'Ega Nur Munziatunnas', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-106/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110016' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-106/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523110016' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050131', 'ipds5205', 'Lastry', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-107/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050131' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-107/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050131' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110086', 'ipds5205', 'Mita Rahmatullah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-108/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110086' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-108/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110086' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100027', 'ipds5205', 'Nurjanah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-109/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100027' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-109/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100027' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030028', 'ipds5205', 'Nurwahidah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-110/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030028' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-110/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522030028' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050215', 'ipds5205', 'BAIQ ERINE WAHYU CHANDRA NIRWANA', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-111/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050215' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-111/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050215' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050238', 'ipds5205', 'Sumiati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-112/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050238' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-112/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050238' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522020017', 'ipds5205', 'Suhardin Putra', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-113/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522020017' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-113/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522020017' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100209', 'ipds5205', 'Abubakar Ismail', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-114/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100209' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-114/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100209' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110076', 'ipds5205', 'Imran', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-115/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110076' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-115/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523110076' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050025', 'ipds5205', 'M. Arwenda Prayogi', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-116/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050025' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-116/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050025' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110255', 'ipds5205', 'Muamar Surajudin', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-117/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110255' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-117/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110255' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050187', 'ipds5205', 'Muhammad Nur Azhari', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-118/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050187' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-118/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050187' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110115', 'ipds5205', 'Nugie Akbar', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-119/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110115' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-119/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523110115' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110080', 'ipds5205', 'Restu Subroto', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-120/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110080' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-120/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110080' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110208', 'ipds5205', 'Subhan', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-121/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110208' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-121/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110208' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050007', 'ipds5205', 'Veri Irawan', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-122/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050007' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-122/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050007' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050015', 'ipds5205', 'Wahyudin', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-123/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050015' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-123/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050015' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050066', 'ipds5205', 'Amelia', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-124/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050066' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-124/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050066' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110228', 'ipds5205', 'Anggun Anabela Yustika Putri', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-125/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110228' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-125/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110228' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110048', 'ipds5205', 'Annisa', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-126/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110048' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-126/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110048' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110074', 'ipds5205', 'Astuti', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-127/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110074' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-127/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110074' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050042', 'ipds5205', 'Etty Moto Logis', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-128/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050042' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-128/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050042' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050032', 'ipds5205', 'Fani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-129/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050032' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-129/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050032' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050048', 'ipds5205', 'Haifa Paradisa', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-130/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050048' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-130/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050048' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050089', 'ipds5205', 'Levy Viola Ovaliani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-131/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050089' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-131/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050089' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050002', 'ipds5205', 'Titian Martini', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-132/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050002' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-132/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050002' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110009', 'ipds5205', 'Lisnawati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-133/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110009' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-133/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110009' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050097', 'ipds5205', 'Yulia Safitri', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-134/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050097' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-134/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050097' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050088', 'ipds5205', 'Miftahul Riyadah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-135/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050088' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-135/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050088' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050044', 'ipds5205', 'Nawasari Dwi Sulistiawati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-136/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050044' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-136/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050044' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030100', 'ipds5205', 'Nilmawani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-137/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030100' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-137/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523030100' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050264', 'ipds5205', 'Nurrul Fadillah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-138/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050264' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-138/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050264' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522020005', 'ipds5205', 'A''An Aditya Munandar', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-139/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522020005' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-139/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522020005' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522090045', 'ipds5205', 'Adi Hidayat', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-140/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522090045' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-140/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522090045' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110107', 'ipds5205', 'Ade Arahma', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-141/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110107' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-141/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523110107' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050045', 'ipds5205', 'Adrian', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-142/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050045' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-142/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050045' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110171', 'ipds5205', 'Anshari Nawawi', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-143/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110171' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-143/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110171' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050272', 'ipds5205', 'Dedi Rachmat', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-144/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050272' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-144/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050272' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110050', 'ipds5205', 'Adya Ragita Cahyani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-145/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110050' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-145/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523110050' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110262', 'ipds5205', 'Aini', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-146/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110262' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-146/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110262' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110202', 'ipds5205', 'Ainun Nasirah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-147/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110202' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-147/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110202' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522010023', 'ipds5205', 'Atikasari', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-148/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522010023' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-148/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522010023' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050039', 'ipds5205', 'Nurul Fitri', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-149/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050039' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-149/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050039' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110054', 'ipds5205', 'Nurradian', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-150/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110054' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-150/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110054' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050218', 'ipds5205', 'Sri Handayani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-151/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050218' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-151/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050218' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110199', 'ipds5205', 'Dirta Uari Darniati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-152/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110199' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-152/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110199' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050192', 'ipds5205', 'Dyah Rahayu', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-153/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050192' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-153/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050192' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110149', 'ipds5205', 'Ersa Nur Ulfa Islamia', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-154/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110149' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-154/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110149' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050001', 'ipds5205', 'Esa Ariani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-155/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-155/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520524090003', 'ipds5205', 'Feni Alfaonita', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-156/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520524090003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-156/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520524090003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110259', 'ipds5205', 'Fithria Anggraeni', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-157/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110259' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-157/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110259' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100206', 'ipds5205', 'Fitriani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-158/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100206' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-158/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100206' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110012', 'ipds5205', 'Afdiansyah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-159/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110012' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-159/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110012' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050261', 'ipds5205', 'Dermawan', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-160/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050261' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-160/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050261' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030009', 'ipds5205', 'Mia Agustina', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-161/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030009' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-161/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522030009' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100152', 'ipds5205', 'Tri Nuril Safitriani', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-162/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100152' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-162/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100152' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050033', 'ipds5205', 'Dhia Istiqomah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-163/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050033' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-163/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050033' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110216', 'ipds5205', 'Dwi Fathir', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-164/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110216' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-164/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110216' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110194', 'ipds5205', 'Ersa Riga Puspita', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-165/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110194' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-165/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110194' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050254', 'ipds5205', 'Fathuljannah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-166/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050254' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-166/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050254' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '#N/A', 'ipds5205', 'Amelia Narasta Faratika', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-167/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='#N/A' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-167/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='#N/A' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110213', 'ipds5205', 'Jihan Tri Hapsari', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-168/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110213' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-168/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110213' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110212', 'ipds5205', 'Jumiati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-169/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110212' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-169/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110212' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050068', 'ipds5205', 'Kurniati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-170/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050068' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-170/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050068' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050034', 'ipds5205', 'Nur Fitrah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-171/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050034' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-171/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050034' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110192', 'ipds5205', 'Nur Fuziatun Islamiah', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-172/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110192' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-172/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525110192' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030014', 'ipds5205', 'Nurhidayati', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-173/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-173/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520523030014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050013', 'ipds5205', 'Nursani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-174/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050013' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-174/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520526050013' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050012', 'ipds5205', 'Nurul Wahida', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-175/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050012' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-175/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050012' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050014', 'ipds5205', 'Ulvia Mardiani', 'Pendataan Sensus Ekonomi 2026', 'Juni-Agustus', '2026', 2026, 11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213', '2026-06-15', '2026-08-31', 0, 'B-176/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-176/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520525050014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100175', 'ipds5205', 'Yulianti', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-177/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100175' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-177/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100175' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100016', 'ipds5205', 'Ajwar Anas', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-178/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100016' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-178/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522100016' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';

-- ================== TAUFIKURRAMADHAN MENGGANTIKAN IKHWAN KURNIAWAN ==================
DELETE FROM rekap WHERE idsobat='520526050110' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026';
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran, honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030011', 'ipds5205', 'TAUFIKURRAMADHAN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 'Juni-Agustus', '2026', 2026, 12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213', '2026-06-15', '2026-08-31', 0, 'B-087/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030011' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');

UPDATE rekap SET id_spk = 'B-087/SPK-SE2026/5205/PL.200/2026' WHERE idsobat='520522030011' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026';
