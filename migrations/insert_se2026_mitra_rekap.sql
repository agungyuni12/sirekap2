-- ============================================================
-- SE2026 Mitra + Rekap (idempotent - WHERE NOT EXISTS)
-- PCL: 2.5 x 4.629.000 = 11.561.525
-- PML: 2.5 x 4.877.000 = 12.192.500
-- ============================================================

-- MITRA
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522010006', 'Siti Nurlaila', '', 'Lingkunga Rasabou Rt.12 Rw.07', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522010006' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522010007', 'HUMAIRAH', '', 'Jln. Lintas Calabai, Dusun Kajenje, RT/RW 004/001', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522010007' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522010016', 'Sahlan', '', 'Jl. Lintas Lanci Dusun Rinjani Desa Nusa Jaya', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522010016' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522010029', 'ADI NEGORO HERLAMBANG', '', 'KANDAI SATU', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522010029' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522010030', 'Sri nurnaningsih', '', 'Dusun rora barat', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522010030' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522020005', 'A''an Aditya Munandar', '', 'Jln. Lintas Mbawi Desa Wawonduru', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522020005' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522020015', 'ZORDIAN', '', 'Jalan Datuk anggrat dusun suka jaya desa kadindi', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522020015' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522020017', 'Suhardin Putra', '', 'Jalan.Datuk Anggrat Lintas Tambora', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522020017' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030002', 'Isdyansih', '', 'Dusun Sawe Baru Desa Sawe Jln Lintas Lakey Hu''u', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030002' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030002', 'ISDYANSIH', '', 'Dusun Sawe Baru Desa Sawe Jln Lintas Lakey Hu''u', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030002' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030006', 'YUSUF', '', 'dusun pelita 1 Rt 14', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030006' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030008', 'BUHARI SALAM', '', 'Jl.Lintas Pelita Lingk Salama RT011RW005 Bada Dompu', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030008' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030009', 'Mia Agustina', '', 'Dusun Wera RT/RW 010/005 Desa Lepadi Kecamatan Pajo Kabupaten Dompu', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030009' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030010', 'Sri Mardianingsih', '', 'Lingkungan kota baru rt 016 rw 006 Kelurahan Bada', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030010' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030011', 'TAUFIKURRAMADHAN', '', 'Jalan Kartini no 24', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030011' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030014', 'SUHERMAN', '', 'Dusun dorebara utara rt 007\\004', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030014', 'SUHERMAN', '', 'Dusun dorebara utara RT 007 RW 004', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030017', 'EKO ANSHARI', '', 'Rt 11 Rw 04 lingkungan rato', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030017' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030018', 'Hanafi', '', 'Dusun fo''o kompo', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030018' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030018', 'HANAFI', '', 'Dusun fo''o kompo', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030018' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030024', 'Rachmad Soebari', '', 'Lingkungan karijawa RT 003/ RW 002', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030024' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030026', 'HASANUDDIN', '', 'Dusun Dermaga RT03', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030026' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030028', 'Nurwahidah', '', 'Dusun Bolo Baka', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030028' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030029', 'SALAHUDDIN', '', 'dusun saleko rt 003', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030029' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030042', 'NOERHALIMAH', '', 'Dusun Wawo timur RT.08/00 Desa Nowa kec.woja Dompu', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030042' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030054', 'HAERUDDIN', '', 'DUSUN MAULANA SELATAN RT 008 RW 003', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030054' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522030058', 'ADIMAN', '', 'Jalan lintas calabai pekat dompu', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522030058' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522050001', 'Kurniawan', '', 'Woja', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522050001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090003', 'MAYA AYUNDARI', '', 'Jln lintas calabai', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090003' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090013', 'MUSLIMIN AKBAR', '', 'Jalan Lintas Sumbawa Gang Lingkar', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090013' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090013', 'MUSLIMIN AKBAR ARS', '', 'Jalan Lintas Sumbawa Gang Lingkar', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090013' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090014', 'HADIJAH', '', 'Dusun Kalate Rt010/rw004. Kempo', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522090045', 'Adi Hidayat', '', 'Jalan lintas desa saneo', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522090045' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100006', 'ABD. HADI IRAWAN', '', 'Dusun. Rinjani Desa. Nusa Jaya', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100006' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100011', 'ANWAR RIFAID', '', 'Jln. Lanci jaya dusun mujur rt 002-000', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100011' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100013', 'Bahril Qamar', '', 'Lingkungan sawete barat rt011 rw005 kec. Dompu kab. Dompu', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100013' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100016', 'ajwar anas', '', 'dusun pandai RT/Rw 007/004 desa kareke', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100016' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100027', 'NURJANAH', '', 'Jalan Lintas Sumbawa', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100027' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100030', 'JULKARNAIN', '', 'Jln. Diponegoro Dusun Rasabou', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100030' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100046', 'BAIQ NURFITRIANI RAHMAWATI', '', 'Jalan lintas sumbawa dompu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100046' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100054', 'Muji Burrahman Putra', '', 'Jalan lintas calabai', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100054' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100054', 'MUJIBURRAHMAN PUTRA', '', 'Jalan lintas calabai', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100054' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100072', 'Syamsudin', '', 'Dusun Dorebara utara', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100072' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100093', 'NURUL HIDAYATULLAH', '', 'DUSUN WORO BAKA RT/RW 009/000 DESA BAKA JAYA', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100093' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100093', 'NURUL HIDAYATULLAH', '', 'Dusun Worobaka', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100093' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100095', 'Emi Faridah', '', 'Lingkungan Kandai satu Kelurahan kandai satu RT 006 RW 003', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100095' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100110', 'M Agil Al Husna', '', 'Jalan lintas Mbawi Dusun Raba Tumpu Rt. 011 Rw. 000', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100110' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100119', 'Delfi wulandari', '', 'jalan sultan hasanudin jl. lintas sumbawa bima', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100119' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100132', 'MIRJAN ALHALIK', '', 'Jln lintas sumbawa dusun napa rt 03', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100132' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100140', 'KIKI RIZKI AMELIA', '', 'Jln. Teuku umar No. 29 RT 004 RW 002 lingkungan simpasai Woja', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100140' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100141', 'Mita Sayuti', '', 'Jl. Datuk Anggrat Dusun Suka Jaya,Desa Kadindi, kecematan Pekat, Kab. Dompu NTB', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100141' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100142', 'ABDUL KODIR JAELANI', '', 'Dusun madya labali RT 018 RW 005', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100142' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100152', 'Tri Nuril Safitriani', '', 'Jl.lintas lakey RT 01 RW 01 Dusun Rasabou', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100152' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100165', 'RASMINI', '', 'Dusun jambu mente', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100165' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100174', 'PARLI PURMASIDI', '', 'Dusun Muhajir desa nusa jaya', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100174' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100175', 'Yulianti', '', 'Dusun sawe rt 07 rw 02', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100175' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100175', 'YULIANTI', '', 'Dusun sawe rt 07 rw 02', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100175' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100182', 'M.Syahrir', '', 'Jl.lintas calabai desa soro dusun wodi rt 001 rw 001', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100182' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100186', 'Indah putri sari', '', 'Dusun Kesi', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100186' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100206', 'Fitriani', '', 'Dusun Sigi RT.04', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100206' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100209', 'Abubakar Ismail', '', 'Jalan lintas calabai dusun sortatanga', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100209' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520522100223', 'Leni Karlina', '', 'RT 02 Dusun Suka Jaya, Desa Kadindi, Kecamatan Pekat, Kabupaten Dompu ', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520522100223' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523010001', 'nila utama', '', 'jalan Kiai haji Ahmad dahlan', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523010001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030007', 'KHALIFAH ADRIANI PUTRI', '', 'Dusun tirtamengi RT 04, Desa Riwo, Kecamatan Woja, Kabupaten Dompu', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030007' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030014', 'NURHIDAYATI', '', 'Jl. Lintas Jambu Dusun Dorotoi', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030028', 'Yana', '', 'Lingkungan 4, Rt.11/Rw.05', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030028' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030048', 'DEWI PURWATI', '', 'Jl dompu sumbawa dsn kwangk Rt/Rw 003/001', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030048' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030051', 'Nurjahra', '', 'Jl. Lintas Calabai RT 06 RW 01 Dusun Karama', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030051' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030100', 'Nilmawani', '', 'Dusun Calabai Atas RT/RW 003/001 Desa Calabai', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030100' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523030109', 'Jumrah', '', 'Jl Lintas Wisata Taman Nasional Tambora Desa Tambora Dusun Pancasila Kec Pekat Kab Dompu', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523030109' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523050009', 'Syahril Sidik', '', 'Dusun Jembatan Ndano Desa Banggo', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523050009' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523050009', 'SYAHRIL SIDIK', '', 'Dusun Jembatan Ndano', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523050009' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523060001', 'Nur Sholehah', '', 'Jln. Jeruk, Lingkungan Rasabou, RT 12/RW 07', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523060001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523060005', 'Nurul Khotimah', '', 'Jl. Soriwono Rt 008 Rw 003', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523060005' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523060007', 'AHMAD MUHAZZIR', '', 'JL. A. YANI . DOROTOI II', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523060007' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523070002', 'TIRANI APRILIA', '', 'Jl. Dr. Sutomo', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523070002' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110016', 'EGA NUR MUNZIATUNNAS', '', 'Dusun Mada oi u,a RT 001 RW 000', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110016' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110033', 'amin khairi', '', 'dusun muhajrin RT 007 rw 001', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110033' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110050', 'Adya Ragita Cahyani', '', 'Lingkungan Bali Bunga. Kelurahan Kandai Dua Kecamatan Woja', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110050' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110057', 'ABDUL RAHMAN', '', 'Dusun muhajirin rt.01 rw 02', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110057' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110061', 'Wira Nurmayadi', '', 'Dusun Suka Maju', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110061' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110076', 'IMRAN', '', 'Dusun Ladore Desa Karombo', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110076' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110093', 'Irman Fitriani', '', 'Dusun sigi desa soriutu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110093' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110098', 'SOFIAN HALIDIN', '', 'Dusun Sigi, Desa Soriutu, Kec. Manggelewa, Kab. Dompu NTB', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110098' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110101', 'Tri Satria Darmawan, SH', '', 'RT/RW 012/004, Dusun Bajo Baru ', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110101' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110107', 'ADE ARAHMA', '', 'Jalan Diponegoro, Lingkungan V Kelurahan Montabaru', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110107' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110113', 'Ansari', '', 'Jalan lintas Lakey desa Hu''u dusun finis RT 006 RW 003', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110113' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110113', 'ANSARI', '', 'Jalan lintas Lakey desa Hu''u dusun finis RT 006 RW 003', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110113' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110114', 'MUH. YUSRIL', '', 'Jalan lintas calabai dusun Doromelo', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110114' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110115', 'NUGIE AKBAR', '', 'Dusun pancasila, Desa Tambora', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110115' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520523110115', 'Nugie akbar', '', 'DUSUN PANCASILA DESA TAMBORA', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520523110115' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520524090001', 'MERI MARIANI', '', 'Jalan Lintas Calabai', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520524090001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520524090003', 'FENI ALFAONITA', '', 'ling.bukit larema', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520524090003' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050001', 'ESA ARIANI', '', 'Jln. Lintas Sumbawa Bima', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050001' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050002', 'Titian Martini', '', 'Jalan lintas calabai desa pekat dusun pekat 2', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050002' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050003', 'Rani Annisa Sari', '', 'Jalan Lintas Pasar Senin, Dusun Bagek Payung, Desa Kadindi Barat, Kecamatan Pekat, Kabupaten Dompu', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050003' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050004', 'Sri Andriani', '', 'Jln.lintas labuan kananga', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050004' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050005', 'Qori prisita', '', 'Dusun ragi Rt 02 Rw 00 Desa Mbawi Kec. Dompu', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050005' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050006', 'aenul wahtan', '', 'Dusun ragi', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050006' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050007', 'Veri Irawan', '', 'Jalan : Calabai-Dompu Desan Pekat. Kec Pekat. RT 001/001 Kab.Dompu', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050007' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050008', 'Mutmainah', '', 'RT 001 RW 000, Dusun Ncoha', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050008' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050012', 'Nurul Wahida', '', 'Jl. Lintas Lakey Dusun Wera', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050012' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050014', 'Ulvia Mardiani', '', 'Jalan lintas lakey Dusun Mangga Dua', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050014' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050023', 'Suci Purnaningsih', '', 'Dusun Pekat II-Desa Pekat-Kecamatan Pekat', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050023' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050023', 'SUCI PURNANINGSIH', '', 'RT/RW: 001/001-Dusun Pekat II', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050023' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050025', 'M. Arwenda Prayogi', '', 'Jln. Lintas Gunung Tambora', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050025' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050026', 'ROBIATUL ADAWIYAH', '', 'Dusun Sorikalate, Des. Kadindi Barat', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050026' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050027', 'Sandikawati', '', 'Desa Doropeti Dusun Doropeti Rt 002 rw 000', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050027' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050028', 'Melinda Susilarini', '', 'Jl. Lintas Malaju, Dusun Ncoha, Desa Malaju, Rt/Rw 002/000, Kecamatan Kilo, Kabupaten Dompu, Nusa Tenggara Barat', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050028' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050031', 'AULIA APRILIA', '', 'Lingkungan I Kelurahan Montabaru', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050031' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050032', 'FANI', '', 'Dusun Tanjung Pasir Desa Calabai', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050032' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050035', 'Sari ulandari', '', 'Dusun pade maju', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050035' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050036', 'SRI HERAWATI', '', 'Lingkungan Salama RT 011 RW 005', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050036' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050039', 'Nurul Fitri', '', 'DUSUN SELAPARANG RT 003', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050039' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525050045', 'Adrian', '', 'Lingk. Simpasai', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525050045' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110004', 'Rabil Nawara', '', 'Jl. Sultan Hasanuddin Lingkungan Renda Kelurahan Simpasai Kec. Woja Kab', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110004' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110011', 'Ayu wandira', '', 'Dusun anamina', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110011' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110012', 'Afdiansyah', '', 'Jl. Lintas Lakey Dusun Lawiti Desa Tembalae Kecamatan Pajo', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110012' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110013', 'RISKA AMELIA ADE PUTRI', '', 'Jalan lintas calabai, RT 003/ RW 001, Dusun kajenje Desa Soro Barat', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110013' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110018', 'Yeni Desiliani', '', 'Monta baru RT 011 RW 005', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110018' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110024', 'Kusnadin', '', 'Dusun Nanga to''i', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110024' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110025', 'TRI PUSPA KARTININGSIH', '', 'Jalan Lintas Malaju', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110025' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110026', 'ST. RAIHAN', '', 'Dusun Saka', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110026' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110029', 'Maghfiratul khaerani', '', 'Jln.soriwono kelurahan potu, rt.007.rw 004', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110029' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110035', 'ANGGI ANGGRIANI', '', 'Jalan Lintas Calalai Desa Ta''a Kecamatan Kempo Kabupaten Dompu', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110035' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110038', 'IKA SUCIYARTI', '', 'Jln lintas sumbawa', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110038' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110040', 'NOVITTA ISLAMIYAH', '', 'Jalan lintas calabai. Dusun rasa bou. Desa ta’a kec. Kempo kab. Dompu NTB', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110040' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110046', 'Bela Safira', '', 'Jalan Lintas malaju RT 001', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110046' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110048', 'ANNISA', '', 'Jalan Datu anggrat', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110048' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110050', 'ERIK ADITYA PRATAMA', '', 'Jalan Lintas Dermaga Paropa, Dusun Paropa', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110050' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110064', 'Junaiddin', '', 'Dusun Sawe', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110064' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110064', 'JUNAIDDIN', '', 'Dusun Sawe', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110064' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110066', 'Indra Wulan', '', 'RT 003 RW 000 Dusun Malaju Desa Malaju', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110066' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110067', 'Jumliati', '', 'Jalan lintas kilo', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110067' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110067', 'JUMLIATI', '', 'Jln Lintas Kilo Dusun patuh karya RT/RW : 004/000 Desa Lanci Jaya Kecamatan Manggelewa Kab. Dompu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110067' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110074', 'Astuti', '', 'Jln Lintas labuan kananga', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110074' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110080', 'Restu Subroto', '', 'Dusun Bagik Payung RT/RW 001/001', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110080' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110080', 'RESTU SUBROTO', '', 'Dusun Bagik Payung RT/RW 001/001 Desa Kadindi Barat Kecamatan Pekat Kabupaten Dompu', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110080' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110086', 'Mita Rahmatullah', '', 'Dusun jati baru desa tekasire', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110086' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110086', 'MITA RAHMATULLAH', '', 'Dusun jatibaru, desa tekasire', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110086' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110102', 'JUL ASFI WARAIHAN', '', 'Jalan Proros Taropo Dusun Taropo, RT/RW 002/000 Desa Taropo KAB. DOMPU, KILO, NUSA TENGGARA BARAT (NTB), ID, 84252', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110102' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110108', 'Nurlailah', '', 'Desa o''o', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110108' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110108', 'NURLAILAH', '', 'Dusun Muhajirin desa o''o', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110108' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110113', 'Arneliana', '', 'Dusun daha timur', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110113' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110113', 'ARNELIANA', '', 'Dusun daha timur', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110113' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110118', 'Desi Ratnasari', '', 'lingkungan renda kelurahan simpasai', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110118' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110126', 'Patrialis Akbar', '', 'RT 001 RW 000 Dsn. Patula Ds. Malaju', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110126' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110141', 'Nuramanda Yuniar Hartono', '', 'Jalan Syech Muhammad No.16 RT 09 RW 04 Jado, Dorotangga', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110141' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110145', 'Windi', '', 'Jln datuk anggrat rumah putih dusun pusaka', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110145' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110145', 'WINDI', '', 'Jln datuk anggrat dusun pusaka kadindi', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110145' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110149', 'Ersa Nur Ulfa Islamia', '', 'Jl. Lintas mbawi, Dusun Rato Baka, RT/RW. 001/001', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110149' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110157', 'Fayza Shabilla', '', 'Perumahan Mekarsari Dusun Mekarsari Desa Malaju', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110157' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110159', 'R. Isfahul Husna', '', 'Desa Matua Dusun Selaparang Barat RT 4 RW 0 Kecamatan Woja Kabupaten Dompu NTB', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110159' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110164', 'Naufal Rifdal Fadillah', '', 'Ling.Balibunga Kel. Kandai Dua', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110164' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110171', 'Anshari Nawawi', '', 'Jl. Teuku Umar RT 15 RW 06 Lingkungan Renda', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110171' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110176', 'ARDIANSYAH', '', 'DUSUN DAHA TIMUR RT 001 RW 000', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110176' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110176', 'ARDIANSYAH', '', 'DUSUN DAHA TIMUR RT 001 RW 000', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110176' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110177', 'M Iksan', '', 'Jln Lintas Sumbawa-Dompu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110177' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110180', 'OPHIYANSYAH ACHRUL PUTRA', '', 'jalan lintas sumbawa-bima dusun mpongge rt.001 rw.001', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110180' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110181', 'Aminah', '', 'Dusun rangga mbolo rt006/rw000 desa karamabura kecamatan dompu', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110181' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110181', 'AMINAH', '', 'Dusun rangga mbolo desa karamabura', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110181' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110192', 'Nur Fuziatun Islamiah', '', 'Jalan lintas lakey dusun rasa bou', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110192' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110194', 'ERSA RIGA PUSPITA', '', 'Jlan Lintas Lakey Dusun Pajo Permai Rt/Rw 004/002', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110194' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110195', 'ST. Nurhasnah', '', 'Jl. Akasia No. 25 Lingkungan Kotabaru RT.015/RW.006', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110195' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110197', 'Fitriyani', '', 'Jln. Lintas Lakey Dusun Patuh Pada Kena Desa Woko Kecamatan Pajo Kabupaten Dompu', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110197' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110199', 'Dirta Uari Darniati', '', 'Jalan Ompu Buncu, Buncu Utara, Desa Matua, Kecamatan Woja, Kabupaten Dompu ', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110199' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110201', 'NURUL RAHMANIA', '', 'Lingkungan Renda RT 001 RW 005', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110201' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110202', 'Ainun Nasirah', '', 'Jalan lintas mbawi dusun rabatumpu, RT 011 RW 000', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110202' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110203', 'WAWAN SETIAWAN', '', 'Jln. Ahmad Yani-Sera Talaka RT/RW 003/001', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110203' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110204', 'Yul safirah', '', 'Dusun ncoha, Desa malaju, Kecematan Kilo', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110204' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110205', 'LIDYA SRI RAHAYU', '', 'Jalan Lintas Mbawi', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110205' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110208', 'Subhan', '', 'Jalan Lintas Calabai', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110208' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110213', 'Jihan Tri Hapsari', '', 'jalan lintas lakey, dusun mangga dua, desa ranggo, rt 003, rw 001 ', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110213' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110216', 'Dwi Fathir', '', 'jalan lintas lakey, dusun restu, desa ranggo, Rt 01, Rw 04, kec pajo, kab dompu', 'Pajo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110216' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110221', 'Syehlin Rauhar', '', 'Lingkungan Drompana Kelurahan Kandai Satu RW 002', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110221' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110221', 'SYEHLIN RAUHAR', '', 'Desa Nowa Dusun dermaga  RT 03', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110221' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110228', 'Anggun anabela yustika putri', '', 'Desa Calabai kec pekat kab dompu', 'Pekat', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110228' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110234', 'MEGA LESTARI', '', 'jalan lintas sumbawa,desa nowa,dusun dermaga Rt01,Rw01', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110234' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110241', 'Maraatun Hasanah', '', 'Dusun Ragi RT, 002 RW 000 Desa Mbawi KEC.DOMPU', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110241' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110244', 'Vira oktaviana, S. Pd. ', '', 'Jalan lintas pantai ria', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110244' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110244', 'VIRA OKTAVIANA', '', 'Jalan lintas pantai ria, Dusun mumbu, Rt 012, Rw 004, Desa mumbu', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110244' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110245', 'MOCH. RYADI HUSNA', '', 'Dusun Rasanae, RT 003, RW 000, Desa Baka Jaya, Kec. Woja', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110245' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110252', 'SRI MARYAM ULFAH', '', 'Dusun Malaju RT 002 / RW 002', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110252' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110257', 'Adetya Anhar', '', 'Jln. Lintas Lakey, Dusun Lanta, RT 008/RW 000', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110257' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110257', 'ADETYA ANHAR', '', 'Jln. Lintas Lakey, Dusun Lanta, RT 008/RW 000', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110257' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110259', 'Fithria Anggraeni', '', 'Jalan Lintas Sumbawa RT/RW 001/002', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110259' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110262', 'AINI', '', 'Jalan Lintas Ria, Dusun Tirtamengi Mekar Desa Riwo Kec.Woja Kab.Dompu', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110262' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110267', 'Ivid Muhibullah, SE', '', 'Lingkungan Karijawa Utara RT/RW 004/002', 'Dompu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110267' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110270', 'Nabila Rohmatul Ulya', '', 'Jln lintas dermaga', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110270' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110280', 'HANAFIA', '', 'Jalan lintas Lakey Desa Daha Kecematan Hu''u', 'Huu', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110280' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520525110280', 'HANAFIA', '', 'Jln.Lintas Lakey Desa Daha Kecematan Huu Kabupaten Dompu NTB', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520525110280' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050015', 'ARIS MUNANDAR', '', 'Dusun woro jaya Rt/001 Rw/000', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050015' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050017', 'MULIANI', '', 'Jln. Lakey desa rasabou kecamatan hu''u', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050017' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050022', 'HERMAN', '', 'Dusun Daha Timur', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050022' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050023', 'RANGGA BARANI SATRIA', '', 'Jalan lintas Sumbawa-Dompu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050023' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050047', 'WENI RAHAYU', '', 'Jln. Lintas Lakey', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050047' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050050', 'MUHAMMAD ARKHAM', '', 'Jalan Lintas Kilo-Kore, RT 001, RW 000, Dusun Lanci II, Desa Lanci Jaya', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050050' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050055', 'ANGGUN PURNAMA', '', 'Jalan Lintas Calabai RT/001 RW/001 Dusun Kesi', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050055' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050057', 'DHIA ULHAQ', '', 'Jln. Lintas calabai', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050057' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050059', 'IRAWAN DANDI', '', 'Dusun Samakai', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050059' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050070', 'JUMRIATI', '', 'Jalan Lintas Taropo RT 001,', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050070' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050071', 'KURATUL AINI', '', 'Jalan lintas calabai', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050071' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050110', 'IKHWAN KURNIAWAN', '', 'Jlan Lintas Sumbawa Desa Anamina Kecamatan Manggelewa Kabupaten Domou', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050110' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050120', 'DIAN APRIANI', '', 'Jalan Lintas Calabai, Desa Konte, Kecamatan Kempo, Kabupaten Dompu, Nusa Tenggara Barat', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050120' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050121', 'ISKANDAR JULKARNAIN', '', 'Dusun ruhu ruma RT 04 RW 000', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050121' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050125', 'EMI YULIANA', '', 'Dusun sorinaru', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050125' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050126', 'IMANSYAH', '', 'JL. LINTAS LAKEY DUSUN ADU', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050126' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050138', 'MUSLIMATUN FITRIAH', '', 'Dusun Malaju, Desa Malaju, Kecamatan Kilo', 'Kilo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050138' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050144', 'SUHADA', '', 'Jl lintas PT sira/Cempi jaya', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050144' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050147', 'AIDA RAHAYU', '', 'Jln lintas lakey dusun adu desa adu kecematan hu’u', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050147' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050151', 'DENTI ALIFIA RAMDHOANI', '', 'Jalan Lintas Lakey, Dusun Rasabou', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050151' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050157', 'ANDI MUHAMMAD ALI', '', 'Dusun kajenje Rt.004 Re.001', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050157' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050159', 'SULFAH MEIASTRI', '', 'Dusun Padamara, RT/RW 003/001, Desa Kempo', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050159' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050176', 'SULAIMAN', '', 'Dusun dasanbaru', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050176' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050178', 'A. ARIF MUNANDAR', '', 'Jalan lintas calabai, RT/RW 005/001', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050178' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050181', 'ARI APRIADI', '', 'Dusun Permata Hijau RT 001 / RW 002', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050181' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050193', 'AMRIZAL', '', 'Jln. Lintas lakey desa Adu kecamatan Hu''u kabupaten Dompu', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050193' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050204', 'YULIATI', '', 'Jl. Lintas Lakey, Desa Adu, Dusun Woro', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050204' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050208', 'AL BIMA', '', 'Jln lintas calabai desa doromelo kecamatan manggelewa kabupaten Dompu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050208' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050223', 'NURDIN', '', 'Jln Lintas Sumbawa Anamin Manggelewa Dompu', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050223' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050231', 'MUHAMAD ERWINSYAH PUTRA', '', 'Dusun nanga na''e desa jala, kecamatan hu''u kabupaten dompu', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050231' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050235', 'NURMIFTAHUL JANNAH', '', 'Dusun lodo, Desa Sawe, kecamatan Hu''u.', 'Hu`U', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050235' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050244', 'NASRUDIN', '', 'Dusun rinjani', 'Manggalewa', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050244' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050249', 'AMINULLAH', '', 'Dusun Kesi RT/RW 004/002', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050249' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '520526050266', 'WALIEY MUFTIHAKAM MARZA KARYADI', '', 'Dusun Permata Hijau RT/RW 001/000', 'Kempo', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='520526050266' AND tahun='2026');
INSERT INTO mitra (idsobat, nmitra, nik, alamat, kecamatan, tahun)
SELECT '527223110270', 'RAUDATUL JANNAH', '', 'Jln.lintas sumbawa desa matua,dusun rasanggaro barat ', 'Woja', '2026'
WHERE NOT EXISTS (SELECT 1 FROM mitra WHERE idsobat='527223110270' AND tahun='2026');

-- REKAP
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522010006', 'ipds5205', 'SITI NURLAILA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522010006' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522010007', 'ipds5205', 'HUMAIRAH', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-012/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522010007' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522010016', 'ipds5205', 'SAHLAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-034/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522010016' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522010029', 'ipds5205', 'ADI NEGORO HERLAMBANG', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522010029' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522010030', 'ipds5205', 'SRI NURNANINGSIH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522010030' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522020005', 'ipds5205', 'A''AN ADITYA MUNANDAR', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522020005' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522020015', 'ipds5205', 'ZORDIAN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522020015' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522020017', 'ipds5205', 'SUHARDIN PUTRA', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522020017' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030002', 'ipds5205', 'ISDYANSIH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-059/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030002' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030006', 'ipds5205', 'YUSUF', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030006' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030008', 'ipds5205', 'BUHARI SALAM', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030008' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030009', 'ipds5205', 'MIA AGUSTINA', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030009' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030010', 'ipds5205', 'SRI MARDIANINGSIH', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030010' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030011', 'ipds5205', 'TAUFIKURRAMADHAN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030011' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030014', 'ipds5205', 'SUHERMAN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030014' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030017', 'ipds5205', 'EKO ANSHARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030017' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030018', 'ipds5205', 'HANAFI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-044/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030018' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030024', 'ipds5205', 'RACHMAD SOEBARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030024' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030026', 'ipds5205', 'HASANUDDIN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030026' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030028', 'ipds5205', 'NURWAHIDAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030028' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030029', 'ipds5205', 'SALAHUDDIN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-002/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030029' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030042', 'ipds5205', 'NOERHALIMAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030042' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030054', 'ipds5205', 'HAERUDDIN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030054' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522030058', 'ipds5205', 'ADIMAN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522030058' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522050001', 'ipds5205', 'KURNIAWAN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522050001' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522090003', 'ipds5205', 'MAYA AYUNDARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-019/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522090003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522090013', 'ipds5205', 'MUSLIMIN AKBAR ARS', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522090013' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522090014', 'ipds5205', 'HADIJAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-016/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522090014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522090045', 'ipds5205', 'ADI HIDAYAT', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522090045' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100006', 'ipds5205', 'ABD. HADI IRAWAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-029/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100006' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100011', 'ipds5205', 'ANWAR RIFAID', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-085/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100011' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100013', 'ipds5205', 'BAHRIL QAMAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100013' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100016', 'ipds5205', 'AJWAR ANAS', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100016' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100027', 'ipds5205', 'NURJANAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100027' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100030', 'ipds5205', 'JULKARNAIN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-001/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100030' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100046', 'ipds5205', 'BAIQ NURFITRIANI RAHMAWATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100046' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100054', 'ipds5205', 'MUJIBURRAHMAN PUTRA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-010/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100054' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100072', 'ipds5205', 'SYAMSUDIN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100072' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100093', 'ipds5205', 'NURUL HIDAYATULLAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100093' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100095', 'ipds5205', 'EMI FARIDAH', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100095' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100110', 'ipds5205', 'M AGIL AL HUSNA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100110' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100119', 'ipds5205', 'DELFI WULANDARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100119' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100132', 'ipds5205', 'MIRJAN ALHALIK', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-067/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100132' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100140', 'ipds5205', 'KIKI RIZKI AMELIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100140' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100141', 'ipds5205', 'MITA SAYUTI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100141' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100142', 'ipds5205', 'ABDUL KODIR JAELANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-004/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100142' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100152', 'ipds5205', 'TRI NURIL SAFITRIANI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100152' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100165', 'ipds5205', 'RASMINI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-033/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100165' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100174', 'ipds5205', 'PARLI PURMASIDI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-031/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100174' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100175', 'ipds5205', 'YULIANTI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100175' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100182', 'ipds5205', 'M.SYAHRIR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-009/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100182' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100186', 'ipds5205', 'INDAH PUTRI SARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-017/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100186' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100206', 'ipds5205', 'FITRIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100206' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100209', 'ipds5205', 'ABUBAKAR ISMAIL', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100209' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520522100223', 'ipds5205', 'LENI KARLINA', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520522100223' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523010001', 'ipds5205', 'NILA UTAMA', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-081/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523010001' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030007', 'ipds5205', 'KHALIFAH ADRIANI PUTRI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030007' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030014', 'ipds5205', 'NURHIDAYATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030028', 'ipds5205', 'YANA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030028' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030048', 'ipds5205', 'DEWI PURWATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-039/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030048' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030051', 'ipds5205', 'NURJAHRA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-021/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030051' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030100', 'ipds5205', 'NILMAWANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030100' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523030109', 'ipds5205', 'JUMRAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523030109' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523050009', 'ipds5205', 'SYAHRIL SIDIK', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-024/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523050009' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523060001', 'ipds5205', 'NUR SHOLEHAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523060001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523060005', 'ipds5205', 'NURUL KHOTIMAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523060005' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523060007', 'ipds5205', 'AHMAD MUHAZZIR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523060007' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523070002', 'ipds5205', 'TIRANI APRILIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523070002' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110016', 'ipds5205', 'EGA NUR MUNZIATUNNAS', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110016' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110033', 'ipds5205', 'AMIN KHAIRI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-066/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110033' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110050', 'ipds5205', 'ADYA RAGITA CAHYANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110050' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110057', 'ipds5205', 'ABDUL RAHMAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-083/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110057' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110061', 'ipds5205', 'WIRA NURMAYADI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-025/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110061' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110076', 'ipds5205', 'IMRAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110076' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110093', 'ipds5205', 'IRMAN FITRIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-042/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110093' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110098', 'ipds5205', 'SOFIAN HALIDIN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-035/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110098' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110101', 'ipds5205', 'TRI SATRIA DARMAWAN, SH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-037/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110101' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110107', 'ipds5205', 'ADE ARAHMA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110107' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110113', 'ipds5205', 'ANSARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110113' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110114', 'ipds5205', 'MUH. YUSRIL', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-026/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110114' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520523110115', 'ipds5205', 'NUGIE AKBAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520523110115' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520524090001', 'ipds5205', 'MERI MARIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520524090001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520524090003', 'ipds5205', 'FENI ALFAONITA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520524090003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050001', 'ipds5205', 'ESA ARIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050001' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050002', 'ipds5205', 'TITIAN MARTINI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050002' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050003', 'ipds5205', 'RANI ANNISA SARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050003' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050004', 'ipds5205', 'SRI ANDRIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050004' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050005', 'ipds5205', 'QORI PRISITA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050005' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050006', 'ipds5205', 'AENUL WAHTAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050006' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050007', 'ipds5205', 'VERI IRAWAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050007' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050008', 'ipds5205', 'MUTMAINAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-075/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050008' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050012', 'ipds5205', 'NURUL WAHIDA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050012' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050014', 'ipds5205', 'ULVIA MARDIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050014' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050023', 'ipds5205', 'SUCI PURNANINGSIH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050023' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050025', 'ipds5205', 'M. ARWENDA PRAYOGI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050025' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050026', 'ipds5205', 'ROBIATUL ADAWIYAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050026' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050027', 'ipds5205', 'SANDIKAWATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050027' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050028', 'ipds5205', 'MELINDA SUSILARINI', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-068/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050028' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050031', 'ipds5205', 'AULIA APRILIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050031' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050032', 'ipds5205', 'FANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050032' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050035', 'ipds5205', 'SARI ULANDARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050035' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050036', 'ipds5205', 'SRI HERAWATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050036' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050039', 'ipds5205', 'NURUL FITRI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050039' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525050045', 'ipds5205', 'ADRIAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525050045' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110004', 'ipds5205', 'RABIL NAWARA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110004' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110011', 'ipds5205', 'AYU WANDIRA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-038/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110011' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110012', 'ipds5205', 'AFDIANSYAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110012' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110013', 'ipds5205', 'RISKA AMELIA ADE PUTRI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-022/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110013' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110018', 'ipds5205', 'YENI DESILIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110018' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110024', 'ipds5205', 'KUSNADIN', 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
       'Juni-Agustus', '2026', 2026,
       12192500, '2.5', 'O-B', 4877000, '2902.FAN.ZZ1.051.A.521213',
       '2026-06-15', '2026-08-31', 0, 'B-064/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110024' AND kegiatan='Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110025', 'ipds5205', 'TRI PUSPA KARTININGSIH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-078/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110025' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110026', 'ipds5205', 'ST. RAIHAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110026' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110029', 'ipds5205', 'MAGHFIRATUL KHAERANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110029' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110035', 'ipds5205', 'ANGGI ANGGRIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-013/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110035' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110038', 'ipds5205', 'IKA SUCIYARTI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-041/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110038' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110040', 'ipds5205', 'NOVITTA ISLAMIYAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-020/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110040' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110046', 'ipds5205', 'BELA SAFIRA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-069/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110046' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110048', 'ipds5205', 'ANNISA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110048' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110050', 'ipds5205', 'ERIK ADITYA PRATAMA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-065/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110050' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110064', 'ipds5205', 'JUNAIDDIN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-052/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110064' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110066', 'ipds5205', 'INDRA WULAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-071/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110066' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110067', 'ipds5205', 'JUMLIATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-043/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110067' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110074', 'ipds5205', 'ASTUTI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110074' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110080', 'ipds5205', 'RESTU SUBROTO', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110080' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110086', 'ipds5205', 'MITA RAHMATULLAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110086' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110102', 'ipds5205', 'JUL ASFI WARAIHAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-072/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110102' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110108', 'ipds5205', 'NURLAILAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110108' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110113', 'ipds5205', 'ARNELIANA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-056/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110113' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110118', 'ipds5205', 'DESI RATNASARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110118' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110126', 'ipds5205', 'PATRIALIS AKBAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-089/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110126' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110141', 'ipds5205', 'NURAMANDA YUNIAR HARTONO', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110141' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110145', 'ipds5205', 'WINDI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110145' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110149', 'ipds5205', 'ERSA NUR ULFA ISLAMIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110149' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110157', 'ipds5205', 'FAYZA SHABILLA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-070/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110157' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110159', 'ipds5205', 'R. ISFAHUL HUSNA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110159' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110164', 'ipds5205', 'NAUFAL RIFDAL FADILLAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110164' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110171', 'ipds5205', 'ANSHARI NAWAWI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110171' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110176', 'ipds5205', 'ARDIANSYAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-047/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110176' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110177', 'ipds5205', 'M IKSAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-088/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110177' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110180', 'ipds5205', 'OPHIYANSYAH ACHRUL PUTRA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-030/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110180' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110181', 'ipds5205', 'AMINAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110181' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110192', 'ipds5205', 'NUR FUZIATUN ISLAMIAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110192' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110194', 'ipds5205', 'ERSA RIGA PUSPITA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110194' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110195', 'ipds5205', 'ST. NURHASNAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110195' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110197', 'ipds5205', 'FITRIYANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110197' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110199', 'ipds5205', 'DIRTA UARI DARNIATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110199' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110201', 'ipds5205', 'NURUL RAHMANIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110201' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110202', 'ipds5205', 'AINUN NASIRAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110202' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110203', 'ipds5205', 'WAWAN SETIAWAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110203' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110204', 'ipds5205', 'YUL SAFIRAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-080/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110204' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110205', 'ipds5205', 'LIDYA SRI RAHAYU', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110205' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110208', 'ipds5205', 'SUBHAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110208' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110213', 'ipds5205', 'JIHAN TRI HAPSARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110213' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110216', 'ipds5205', 'DWI FATHIR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110216' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110221', 'ipds5205', 'SYEHLIN RAUHAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110221' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110228', 'ipds5205', 'ANGGUN ANABELA YUSTIKA PUTRI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110228' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110234', 'ipds5205', 'MEGA LESTARI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110234' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110241', 'ipds5205', 'MARAATUN HASANAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110241' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110244', 'ipds5205', 'VIRA OKTAVIANA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110244' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110245', 'ipds5205', 'MOCH. RYADI HUSNA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110245' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110252', 'ipds5205', 'SRI MARYAM ULFAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-077/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110252' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110257', 'ipds5205', 'ADETYA ANHAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-045/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110257' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110259', 'ipds5205', 'FITHRIA ANGGRAENI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110259' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110262', 'ipds5205', 'AINI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110262' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110267', 'ipds5205', 'IVID MUHIBULLAH, SE', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110267' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110270', 'ipds5205', 'NABILA ROHMATUL ULYA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-076/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110270' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520525110280', 'ipds5205', 'HANAFIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-058/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520525110280' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050015', 'ipds5205', 'ARIS MUNANDAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-086/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050015' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050017', 'ipds5205', 'MULIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-060/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050017' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050022', 'ipds5205', 'HERMAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-048/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050022' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050023', 'ipds5205', 'RANGGA BARANI SATRIA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-032/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050023' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050047', 'ipds5205', 'WENI RAHAYU', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-063/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050047' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050050', 'ipds5205', 'MUHAMMAD ARKHAM', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-027/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050050' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050055', 'ipds5205', 'ANGGUN PURNAMA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-014/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050055' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050057', 'ipds5205', 'DHIA ULHAQ', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-008/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050057' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050059', 'ipds5205', 'IRAWAN DANDI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-050/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050059' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050070', 'ipds5205', 'JUMRIATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-073/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050070' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050071', 'ipds5205', 'KURATUL AINI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-018/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050071' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050110', 'ipds5205', 'IKHWAN KURNIAWAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-087/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050110' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050120', 'ipds5205', 'DIAN APRIANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-015/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050120' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050121', 'ipds5205', 'ISKANDAR JULKARNAIN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-051/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050121' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050125', 'ipds5205', 'EMI YULIANA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-040/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050125' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050126', 'ipds5205', 'IMANSYAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-049/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050126' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050138', 'ipds5205', 'MUSLIMATUN FITRIAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-074/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050138' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050144', 'ipds5205', 'SUHADA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-062/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050144' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050147', 'ipds5205', 'AIDA RAHAYU', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-055/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050147' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050151', 'ipds5205', 'DENTI ALIFIA RAMDHOANI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-057/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050151' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050157', 'ipds5205', 'ANDI MUHAMMAD ALI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-006/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050157' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050159', 'ipds5205', 'SULFAH MEIASTRI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-023/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050159' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050176', 'ipds5205', 'SULAIMAN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-036/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050176' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050178', 'ipds5205', 'A. ARIF MUNANDAR', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-003/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050178' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050181', 'ipds5205', 'ARI APRIADI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-007/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050181' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050193', 'ipds5205', 'AMRIZAL', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-046/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050193' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050204', 'ipds5205', 'YULIATI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-054/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050204' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050208', 'ipds5205', 'AL BIMA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-084/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050208' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050223', 'ipds5205', 'NURDIN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-082/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050223' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050231', 'ipds5205', 'MUHAMAD ERWINSYAH PUTRA', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-053/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050231' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050235', 'ipds5205', 'NURMIFTAHUL JANNAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-061/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050235' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050244', 'ipds5205', 'NASRUDIN', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-028/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050244' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050249', 'ipds5205', 'AMINULLAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-005/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050249' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '520526050266', 'ipds5205', 'WALIEY MUFTIHAKAM MARZA KARYADI', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, 'B-011/SPK-SE2026/5205/PL.200/2026'
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='520526050266' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');
INSERT INTO rekap (idsobat, username, namamitra, kegiatan, bulan, tahun, tanggaran,
                   honor, volume, satuan, hsatuan, mak, pmwaktu, pswaktu, jumlah_sls, id_spk)
SELECT '527223110270', 'ipds5205', 'RAUDATUL JANNAH', 'Pendataan Sensus Ekonomi 2026',
       'Juni-Agustus', '2026', 2026,
       11561525, '2.5', 'O-B', 4624610, '2902.BMA.006.005.B.521213',
       '2026-06-15', '2026-08-31', 0, NULL
WHERE NOT EXISTS (SELECT 1 FROM rekap WHERE idsobat='527223110270' AND kegiatan='Pendataan Sensus Ekonomi 2026' AND tahun='2026');