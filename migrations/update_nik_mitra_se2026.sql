-- ============================================================
-- Isi mitra.nik dari BPJS - Petugas SE.xlsx (join by idsobat/SOBAT ID,
-- bukan by nama - jauh lebih andal). mitra.nik sebelumnya kosong utk
-- hampir semua baris SE2026, dan lookup NIK di kode selama ini pakai
-- tabel pegawai by nama yang gagal utk 89 dari 229 petugas yang bisa
-- dicairkan termin ini (beda ejaan nama / tidak terdaftar).
--
-- Kode Go sudah diubah supaya ambil ${nik_petugas} dari mitra.nik by
-- idsobat, bukan dari pegawai by nama lagi.
--
-- Update semua baris mitra utk idsobat itu (2025 & 2026), krn NIK orang
-- tidak berubah per tahun.
-- ============================================================

UPDATE mitra SET nik = '5205031505970002' WHERE idsobat = '520525110257' AND (nik IS NULL OR nik = ''); -- Adetya Anhar
UPDATE mitra SET nik = '5205012407900002' WHERE idsobat = '520522010029' AND (nik IS NULL OR nik = ''); -- ADI NEGORO HERLAMBANG
UPDATE mitra SET nik = '5205011409750004' WHERE idsobat = '520522030008' AND (nik IS NULL OR nik = ''); -- BUHARI SALAM
UPDATE mitra SET nik = '5205014203870004' WHERE idsobat = '520522030010' AND (nik IS NULL OR nik = ''); -- Sri Mardianingsih
UPDATE mitra SET nik = '5205011002720007' WHERE idsobat = '520522030014' AND (nik IS NULL OR nik = ''); -- SUHERMAN
UPDATE mitra SET nik = '5205015210840005' WHERE idsobat = '520522100095' AND (nik IS NULL OR nik = ''); -- Emi Faridah
UPDATE mitra SET nik = '5205030210760002' WHERE idsobat = '520522030018' AND (nik IS NULL OR nik = ''); -- Hanafi
UPDATE mitra SET nik = '5272030802910002' WHERE idsobat = '520522100016' AND (nik IS NULL OR nik = ''); -- ajwar anas
UPDATE mitra SET nik = '5205035402810002' WHERE idsobat = '520522100175' AND (nik IS NULL OR nik = ''); -- Yulianti
UPDATE mitra SET nik = '5205024705000001' WHERE idsobat = '520522010007' AND (nik IS NULL OR nik = ''); -- HUMAIRAH
UPDATE mitra SET nik = '5205041706020001' WHERE idsobat = '520525110024' AND (nik IS NULL OR nik = ''); -- Kusnadin
UPDATE mitra SET nik = '5205013005850002' WHERE idsobat = '520522030011' AND (nik IS NULL OR nik = ''); -- TAUFIKURRAMADHAN
UPDATE mitra SET nik = '5205070107970317' WHERE idsobat = '520523110033' AND (nik IS NULL OR nik = ''); -- amin khairi
UPDATE mitra SET nik = '5205044708990002' WHERE idsobat = '520525050028' AND (nik IS NULL OR nik = ''); -- Melinda Susilarini
UPDATE mitra SET nik = '5205072102000003' WHERE idsobat = '520522090013' AND (nik IS NULL OR nik = ''); -- MUSLIMIN AKBAR ARS
UPDATE mitra SET nik = '5205012105710003' WHERE idsobat = '520523010001' AND (nik IS NULL OR nik = ''); -- NILA UTAMA
UPDATE mitra SET nik = '5205011811840001' WHERE idsobat = '520523050009' AND (nik IS NULL OR nik = ''); -- Syahril Sidik
UPDATE mitra SET nik = '5205064779900001' WHERE idsobat = '520522100223' AND (nik IS NULL OR nik = ''); -- Leni Karlina
UPDATE mitra SET nik = '5205085106960002' WHERE idsobat = '520522030009' AND (nik IS NULL OR nik = ''); -- Mia Agustina
UPDATE mitra SET nik = '5205063005950001' WHERE idsobat = '520522020017' AND (nik IS NULL OR nik = ''); -- Suhardin Putra
UPDATE mitra SET nik = '5205064801990001' WHERE idsobat = '520522100141' AND (nik IS NULL OR nik = ''); -- Mita Sayuti
UPDATE mitra SET nik = '5205060208980001' WHERE idsobat = '520522020015' AND (nik IS NULL OR nik = ''); -- ZORDIAN
UPDATE mitra SET nik = '5205052006000002' WHERE idsobat = '520522090045' AND (nik IS NULL OR nik = ''); -- Adi Hidayat
UPDATE mitra SET nik = '5205050612800002' WHERE idsobat = '520522030054' AND (nik IS NULL OR nik = ''); -- HAERUDDIN
UPDATE mitra SET nik = '5205051307910002' WHERE idsobat = '520522050001' AND (nik IS NULL OR nik = ''); -- Kurniawan
UPDATE mitra SET nik = '5205020107680145' WHERE idsobat = '520522030029' AND (nik IS NULL OR nik = ''); -- SALAHUDDIN
UPDATE mitra SET nik = '5205050605760002' WHERE idsobat = '520522030006' AND (nik IS NULL OR nik = ''); -- YUSUF
UPDATE mitra SET nik = '5205051206850004' WHERE idsobat = '520522030026' AND (nik IS NULL OR nik = ''); -- HASANUDDIN
UPDATE mitra SET nik = '5205034808030003' WHERE idsobat = '520525110280' AND (nik IS NULL OR nik = ''); -- HANAFIA
UPDATE mitra SET nik = '5205052701930003' WHERE idsobat = '520522020005' AND (nik IS NULL OR nik = ''); -- A''an Aditya Munandar
UPDATE mitra SET nik = '5205085812000001' WHERE idsobat = '520522100152' AND (nik IS NULL OR nik = ''); -- Tri Nuril Safitriani
UPDATE mitra SET nik = '5205060604840003' WHERE idsobat = '520522030058' AND (nik IS NULL OR nik = ''); -- ADIMAN
UPDATE mitra SET nik = '5205021408860001' WHERE idsobat = '520522100030' AND (nik IS NULL OR nik = ''); -- JULKARNAIN
UPDATE mitra SET nik = '5205070107950302' WHERE idsobat = '520523110061' AND (nik IS NULL OR nik = ''); -- Wira Nurmayadi
UPDATE mitra SET nik = '5205015002000002' WHERE idsobat = '520525110181' AND (nik IS NULL OR nik = ''); -- Aminah
UPDATE mitra SET nik = '5205060104770001' WHERE idsobat = '520522100209' AND (nik IS NULL OR nik = ''); -- Abubakar Ismail
UPDATE mitra SET nik = '5205056604940002' WHERE idsobat = '520525050031' AND (nik IS NULL OR nik = ''); -- AULIA APRILIA
UPDATE mitra SET nik = '5205050611980002' WHERE idsobat = '520523110107' AND (nik IS NULL OR nik = ''); -- ADE ARAHMA
UPDATE mitra SET nik = '5205010707900004' WHERE idsobat = '520522100013' AND (nik IS NULL OR nik = ''); -- Bahril Qamar
UPDATE mitra SET nik = '5205013108010003' WHERE idsobat = '520523060007' AND (nik IS NULL OR nik = ''); -- AHMAD MUHAZZIR
UPDATE mitra SET nik = '5205086207030002' WHERE idsobat = '520525110216' AND (nik IS NULL OR nik = ''); -- Dwi Fathir
UPDATE mitra SET nik = '5205075803930003' WHERE idsobat = '520522100046' AND (nik IS NULL OR nik = ''); -- BAIQ NURFITRIANI RAHMAWATI
UPDATE mitra SET nik = '5205046006970003' WHERE idsobat = '520525110046' AND (nik IS NULL OR nik = ''); -- Bela Safira
UPDATE mitra SET nik = '5205056111860004' WHERE idsobat = '520522100206' AND (nik IS NULL OR nik = ''); -- Fitriani
UPDATE mitra SET nik = '5205054408980002' WHERE idsobat = '520525110149' AND (nik IS NULL OR nik = ''); -- Ersa Nur Ulfa Islamia
UPDATE mitra SET nik = '5205045501030001' WHERE idsobat = '520525110102' AND (nik IS NULL OR nik = ''); -- JUL ASFI WARAIHAN
UPDATE mitra SET nik = '5205014303030006' WHERE idsobat = '520525110241' AND (nik IS NULL OR nik = ''); -- Maraatun Hasanah
UPDATE mitra SET nik = '5205065709010002' WHERE idsobat = '520524090001' AND (nik IS NULL OR nik = ''); -- MERI MARIANI
UPDATE mitra SET nik = '5205012008880003' WHERE idsobat = '520525110267' AND (nik IS NULL OR nik = ''); -- Ivid Muhibullah, SE
UPDATE mitra SET nik = '5205016005010006' WHERE idsobat = '520525110029' AND (nik IS NULL OR nik = ''); -- Maghfiratul khaerani
UPDATE mitra SET nik = '5205070310990002' WHERE idsobat = '520523110114' AND (nik IS NULL OR nik = ''); -- MUH. YUSRIL
UPDATE mitra SET nik = '5205044108010005' WHERE idsobat = '520525050008' AND (nik IS NULL OR nik = ''); -- Mutmainah
UPDATE mitra SET nik = '5205055012010002' WHERE idsobat = '520525110164' AND (nik IS NULL OR nik = ''); -- Naufal Rifdal Fadillah
UPDATE mitra SET nik = '5205085906980001' WHERE idsobat = '520523030014' AND (nik IS NULL OR nik = ''); -- NURHIDAYATI
UPDATE mitra SET nik = '5205017005000005' WHERE idsobat = '520523060005' AND (nik IS NULL OR nik = ''); -- Nurul Khotimah
UPDATE mitra SET nik = '5205075110980002' WHERE idsobat = '520522100027' AND (nik IS NULL OR nik = ''); -- NURJANAH
UPDATE mitra SET nik = '5205055405030003' WHERE idsobat = '520525110201' AND (nik IS NULL OR nik = ''); -- NURUL RAHMANIA
UPDATE mitra SET nik = '5205016801020003' WHERE idsobat = '520525050005' AND (nik IS NULL OR nik = ''); -- Qori prisita
UPDATE mitra SET nik = '5205055811900002' WHERE idsobat = '520522030028' AND (nik IS NULL OR nik = ''); -- Nurwahidah
UPDATE mitra SET nik = '5205074507990006' WHERE idsobat = '520523110079' AND (nik IS NULL OR nik = ''); -- Sri Rahmawati
UPDATE mitra SET nik = '5205056803040003' WHERE idsobat = '520525110042' AND (nik IS NULL OR nik = ''); -- RUPI RAHMIATI
UPDATE mitra SET nik = '5205016912900001' WHERE idsobat = '520522010030' AND (nik IS NULL OR nik = ''); -- Sri nurnaningsih
UPDATE mitra SET nik = '5205014105880004' WHERE idsobat = '520525050036' AND (nik IS NULL OR nik = ''); -- SRI HERAWATI
UPDATE mitra SET nik = '5205054305010005' WHERE idsobat = '520525110159' AND (nik IS NULL OR nik = ''); -- R. Isfahul Husna
UPDATE mitra SET nik = '5205066407030002' WHERE idsobat = '520525050004' AND (nik IS NULL OR nik = ''); -- Sri Andriani
UPDATE mitra SET nik = '5206035608030003' WHERE idsobat = '520525110244' AND (nik IS NULL OR nik = ''); -- Vira Oktaviana
UPDATE mitra SET nik = '5205011610940001' WHERE idsobat = '520522100072' AND (nik IS NULL OR nik = ''); -- Syamsudin
UPDATE mitra SET nik = '5205080311920001' WHERE idsobat = '520525110012' AND (nik IS NULL OR nik = ''); -- Afdiansyah
UPDATE mitra SET nik = '5205042506020001' WHERE idsobat = '520525110126' AND (nik IS NULL OR nik = ''); -- Patrialis Akbar
UPDATE mitra SET nik = '5205022107990002' WHERE idsobat = '520522100054' AND (nik IS NULL OR nik = ''); -- Mujiburrahman Putra
UPDATE mitra SET nik = '5205046802010001' WHERE idsobat = '520525110066' AND (nik IS NULL OR nik = ''); -- Indra Wulan
UPDATE mitra SET nik = '5205065509020004' WHERE idsobat = '520525110145' AND (nik IS NULL OR nik = ''); -- Windi
UPDATE mitra SET nik = '5205025208970002' WHERE idsobat = '520522100186' AND (nik IS NULL OR nik = ''); -- Indah putri sari
UPDATE mitra SET nik = '5205056007020003' WHERE idsobat = '520523030007' AND (nik IS NULL OR nik = ''); -- KHALIFAH ADRIANI PUTRI
UPDATE mitra SET nik = '5205024904010002' WHERE idsobat = '520525110035' AND (nik IS NULL OR nik = ''); -- ANGGI ANGGRIANI
UPDATE mitra SET nik = '5205086901950002' WHERE idsobat = '520525050014' AND (nik IS NULL OR nik = ''); -- Ulvia Mardiani
UPDATE mitra SET nik = '5205060206000002' WHERE idsobat = '520525050007' AND (nik IS NULL OR nik = ''); -- Veri Irawan
UPDATE mitra SET nik = '5205046803990001' WHERE idsobat = '520525110204' AND (nik IS NULL OR nik = ''); -- Yul safirah
UPDATE mitra SET nik = '5309061503980003' WHERE idsobat = '520525110203' AND (nik IS NULL OR nik = ''); -- WAWAN SETIAWAN
UPDATE mitra SET nik = '5205056312020001' WHERE idsobat = '520525110018' AND (nik IS NULL OR nik = ''); -- Yeni Desiliani
UPDATE mitra SET nik = '5205045001000003' WHERE idsobat = '520525110025' AND (nik IS NULL OR nik = ''); -- TRI PUSPA KARTININGSIH
UPDATE mitra SET nik = '5205072809960003' WHERE idsobat = '520523110101' AND (nik IS NULL OR nik = ''); -- Tri Satria Darmawan, SH
UPDATE mitra SET nik = '5205016101030002' WHERE idsobat = '520525110221' AND (nik IS NULL OR nik = ''); -- Syehlin Rauhar
UPDATE mitra SET nik = '5205066609970001' WHERE idsobat = '520525050002' AND (nik IS NULL OR nik = ''); -- Titian Martini
UPDATE mitra SET nik = '5205014306930003' WHERE idsobat = '520525110026' AND (nik IS NULL OR nik = ''); -- ST. RAIHAN
UPDATE mitra SET nik = '5205017001900003' WHERE idsobat = '520525110195' AND (nik IS NULL OR nik = ''); -- ST. Nurhasnah
UPDATE mitra SET nik = '5205066711000004' WHERE idsobat = '520525050023' AND (nik IS NULL OR nik = ''); -- Suci Purnaningsih
UPDATE mitra SET nik = '5205066505980001' WHERE idsobat = '520525050035' AND (nik IS NULL OR nik = ''); -- Sari ulandari
UPDATE mitra SET nik = '5205046002850001' WHERE idsobat = '520525110252' AND (nik IS NULL OR nik = ''); -- SRI MARYAM ULFAH
UPDATE mitra SET nik = '5205014602960001' WHERE idsobat = '520522010006' AND (nik IS NULL OR nik = ''); -- Siti Nurlaila
UPDATE mitra SET nik = '5205065501030004' WHERE idsobat = '520525050003' AND (nik IS NULL OR nik = ''); -- Rani Annisa Sari
UPDATE mitra SET nik = '5205060411000002' WHERE idsobat = '520525110208' AND (nik IS NULL OR nik = ''); -- Subhan
UPDATE mitra SET nik = '5205060411030001' WHERE idsobat = '520525110080' AND (nik IS NULL OR nik = ''); -- Restu Subroto
UPDATE mitra SET nik = '5205012711830002' WHERE idsobat = '520522030024' AND (nik IS NULL OR nik = ''); -- Rachmad Soebari
UPDATE mitra SET nik = '5205056604020003' WHERE idsobat = '527223110270' AND (nik IS NULL OR nik = ''); -- RAUDATUL JANNAH
UPDATE mitra SET nik = '5205073112870008' WHERE idsobat = '520522010016' AND (nik IS NULL OR nik = ''); -- Sahlan
UPDATE mitra SET nik = '5205086510030001' WHERE idsobat = '520525050012' AND (nik IS NULL OR nik = ''); -- Nurul Wahida
UPDATE mitra SET nik = '5205056907000004' WHERE idsobat = '520522100093' AND (nik IS NULL OR nik = ''); -- NURUL HIDAYATULLAH
UPDATE mitra SET nik = '5205055708010005' WHERE idsobat = '520525050039' AND (nik IS NULL OR nik = ''); -- Nurul Fitri
UPDATE mitra SET nik = '5205056002030003' WHERE idsobat = '520525110054' AND (nik IS NULL OR nik = ''); -- Nurradian
UPDATE mitra SET nik = '5205016410990003' WHERE idsobat = '520523060001' AND (nik IS NULL OR nik = ''); -- Nur Sholehah
UPDATE mitra SET nik = '5205015506020003' WHERE idsobat = '520525110141' AND (nik IS NULL OR nik = ''); -- Nuramanda Yuniar Hartono
UPDATE mitra SET nik = '5205024408910002' WHERE idsobat = '520523030051' AND (nik IS NULL OR nik = ''); -- Nurjahra
UPDATE mitra SET nik = '5205085203030004' WHERE idsobat = '520525110192' AND (nik IS NULL OR nik = ''); -- Nur Fuziatun Islamiah
UPDATE mitra SET nik = '5205061306990006' WHERE idsobat = '520523110115' AND (nik IS NULL OR nik = ''); -- Nugie akbar
UPDATE mitra SET nik = '5205054305850006' WHERE idsobat = '520522030042' AND (nik IS NULL OR nik = ''); -- NOERHALIMAH
UPDATE mitra SET nik = '5205065205000001' WHERE idsobat = '520523030100' AND (nik IS NULL OR nik = ''); -- Nilmawani
UPDATE mitra SET nik = '5205047001040001' WHERE idsobat = '520525110270' AND (nik IS NULL OR nik = ''); -- Nabila Rohmatul Ulya
UPDATE mitra SET nik = '5205054304940010' WHERE idsobat = '520524090003' AND (nik IS NULL OR nik = ''); -- FENI ALFAONITA
UPDATE mitra SET nik = '5205072703860001' WHERE idsobat = '520523110057' AND (nik IS NULL OR nik = ''); -- ABDUL RAHMAN
UPDATE mitra SET nik = '5205050809030002' WHERE idsobat = '520525050045' AND (nik IS NULL OR nik = ''); -- Adrian
UPDATE mitra SET nik = '5205074306990004' WHERE idsobat = '520523110093' AND (nik IS NULL OR nik = ''); -- Irman Fitriani
UPDATE mitra SET nik = '5205035209810003' WHERE idsobat = '520522030002' AND (nik IS NULL OR nik = ''); -- Isdyansih
UPDATE mitra SET nik = '5205014412010001' WHERE idsobat = '520523110050' AND (nik IS NULL OR nik = ''); -- Adya Ragita Cahyani
UPDATE mitra SET nik = '5205014607030002' WHERE idsobat = '520525050006' AND (nik IS NULL OR nik = ''); -- aenul wahtan
UPDATE mitra SET nik = '5205055303010007' WHERE idsobat = '520525110262' AND (nik IS NULL OR nik = ''); -- AINI
UPDATE mitra SET nik = '5205054505990009' WHERE idsobat = '520525110090' AND (nik IS NULL OR nik = ''); -- Amelia Narasta Faratika
UPDATE mitra SET nik = '5205066010020001' WHERE idsobat = '520525110228' AND (nik IS NULL OR nik = ''); -- Anggun anabela yustika putri
UPDATE mitra SET nik = '5205066801020002' WHERE idsobat = '520525110048' AND (nik IS NULL OR nik = ''); -- ANNISA
UPDATE mitra SET nik = '5206030712910001' WHERE idsobat = '520525110171' AND (nik IS NULL OR nik = ''); -- Anshari Nawawi
UPDATE mitra SET nik = '5205082311940001' WHERE idsobat = '520522100011' AND (nik IS NULL OR nik = ''); -- ANWAR RIFAID
UPDATE mitra SET nik = '5205084810030001' WHERE idsobat = '520525110213' AND (nik IS NULL OR nik = ''); -- Jihan Tri Hapsari
UPDATE mitra SET nik = '5205031502910001' WHERE idsobat = '520525110176' AND (nik IS NULL OR nik = ''); -- ARDIANSYAH
UPDATE mitra SET nik = '5205077103000001' WHERE idsobat = '520525110067' AND (nik IS NULL OR nik = ''); -- Jumliati
UPDATE mitra SET nik = '5205064108950001' WHERE idsobat = '520523030109' AND (nik IS NULL OR nik = ''); -- Jumrah
UPDATE mitra SET nik = '5205086905010001' WHERE idsobat = '520525110194' AND (nik IS NULL OR nik = ''); -- ERSA RIGA PUSPITA
UPDATE mitra SET nik = '5206054305010001' WHERE idsobat = '520525050032' AND (nik IS NULL OR nik = ''); -- FANI
UPDATE mitra SET nik = '5205054104040003' WHERE idsobat = '520525050001' AND (nik IS NULL OR nik = ''); -- ESA ARIANI
UPDATE mitra SET nik = '5205036906020002' WHERE idsobat = '520525110113' AND (nik IS NULL OR nik = ''); -- Arneliana
UPDATE mitra SET nik = '5205064107010274' WHERE idsobat = '520525110074' AND (nik IS NULL OR nik = ''); -- Astuti
UPDATE mitra SET nik = '5206186305971002' WHERE idsobat = '520525110011' AND (nik IS NULL OR nik = ''); -- Ayu wandira
UPDATE mitra SET nik = '5205025105960001' WHERE idsobat = '520525110118' AND (nik IS NULL OR nik = ''); -- Desi Ratnasari
UPDATE mitra SET nik = '5205076407010001' WHERE idsobat = '520523110016' AND (nik IS NULL OR nik = ''); -- EGA NUR MUNZIATUNNAS
UPDATE mitra SET nik = '5205012404740004' WHERE idsobat = '520522030017' AND (nik IS NULL OR nik = ''); -- EKO ANSHARI
UPDATE mitra SET nik = '5205074709990002' WHERE idsobat = '520525110038' AND (nik IS NULL OR nik = ''); -- IKA SUCIYARTI
UPDATE mitra SET nik = '5205062110940002' WHERE idsobat = '520523110076' AND (nik IS NULL OR nik = ''); -- IMRAN
UPDATE mitra SET nik = '5205031706790001' WHERE idsobat = '520525110064' AND (nik IS NULL OR nik = ''); -- Junaiddin
UPDATE mitra SET nik = '5205055505910006' WHERE idsobat = '520522100140' AND (nik IS NULL OR nik = ''); -- KIKI RIZKI AMELIA
UPDATE mitra SET nik = '5205016901030002' WHERE idsobat = '520525110205' AND (nik IS NULL OR nik = ''); -- LIDYA SRI RAHAYU
UPDATE mitra SET nik = '5205050505980005' WHERE idsobat = '520522100110' AND (nik IS NULL OR nik = ''); -- M Agil Al Husna
UPDATE mitra SET nik = '5205071310020001' WHERE idsobat = '520525110177' AND (nik IS NULL OR nik = ''); -- M Iksan
UPDATE mitra SET nik = '5205062607980001' WHERE idsobat = '520525050025' AND (nik IS NULL OR nik = ''); -- M. Arwenda Prayogi
UPDATE mitra SET nik = '5205020701950002' WHERE idsobat = '520522100182' AND (nik IS NULL OR nik = ''); -- M.Syahrir
UPDATE mitra SET nik = '5205076403980003' WHERE idsobat = '520522090003' AND (nik IS NULL OR nik = ''); -- MAYA AYUNDARI
UPDATE mitra SET nik = '5205054101030012' WHERE idsobat = '520525110234' AND (nik IS NULL OR nik = ''); -- MEGA LESTARI
UPDATE mitra SET nik = '5205071108910002' WHERE idsobat = '520522100132' AND (nik IS NULL OR nik = ''); -- MIRJAN ALHALIK
UPDATE mitra SET nik = '5205052708990006' WHERE idsobat = '520525110245' AND (nik IS NULL OR nik = ''); -- MOCH. RYADI HUSNA
UPDATE mitra SET nik = '5205045409030002' WHERE idsobat = '520525110157' AND (nik IS NULL OR nik = ''); -- Fayza Shabilla
UPDATE mitra SET nik = '5205040409990004' WHERE idsobat = '520525110050' AND (nik IS NULL OR nik = ''); -- ERIK ADITYA PRATAMA
UPDATE mitra SET nik = '5272035807020002' WHERE idsobat = '520525050027' AND (nik IS NULL OR nik = ''); -- Sandikawati
UPDATE mitra SET nik = '5205055906020002' WHERE idsobat = '520525110202' AND (nik IS NULL OR nik = ''); -- Ainun Nasirah
UPDATE mitra SET nik = '5205015204000002' WHERE idsobat = '520525110108' AND (nik IS NULL OR nik = ''); -- Nurlailah
UPDATE mitra SET nik = '5205055407020003' WHERE idsobat = '520525110199' AND (nik IS NULL OR nik = ''); -- Dirta Uari Darniati
UPDATE mitra SET nik = '5205055909000001' WHERE idsobat = '520523030028' AND (nik IS NULL OR nik = ''); -- Yana
UPDATE mitra SET nik = '5205016304000002' WHERE idsobat = '520523070002' AND (nik IS NULL OR nik = ''); -- TIRANI APRILIA
UPDATE mitra SET nik = '5205055512010004' WHERE idsobat = '520525110259' AND (nik IS NULL OR nik = ''); -- Fithria Anggraeni
UPDATE mitra SET nik = '5205064505990003' WHERE idsobat = '520525050026' AND (nik IS NULL OR nik = ''); -- ROBIATUL ADAWIYAH
UPDATE mitra SET nik = '5205074107990309' WHERE idsobat = '520525110086' AND (nik IS NULL OR nik = ''); -- Mita Rahmatullah
UPDATE mitra SET nik = '5205021009990003' WHERE idsobat = '520526050178' AND (nik IS NULL OR nik = ''); -- A. ARIF MUNANDAR
UPDATE mitra SET nik = '5205022808940001' WHERE idsobat = '520522100142' AND (nik IS NULL OR nik = ''); -- Abdul Kodir Jaelani
UPDATE mitra SET nik = '5205020505980001' WHERE idsobat = '520526050249' AND (nik IS NULL OR nik = ''); -- Aminullah
UPDATE mitra SET nik = '5205022206970002' WHERE idsobat = '520526050057' AND (nik IS NULL OR nik = ''); -- DHIA ULHAQ
UPDATE mitra SET nik = '5205020304990003' WHERE idsobat = '520526050181' AND (nik IS NULL OR nik = ''); -- ARI APRIADI
UPDATE mitra SET nik = '5205021103820002' WHERE idsobat = '520526050157' AND (nik IS NULL OR nik = ''); -- ANDI MUHAMMAD ALI
UPDATE mitra SET nik = '5205025304030001' WHERE idsobat = '520526050120' AND (nik IS NULL OR nik = ''); -- Dian Apriani
UPDATE mitra SET nik = '5205025503740002' WHERE idsobat = '520522090014' AND (nik IS NULL OR nik = ''); -- Hadijah
UPDATE mitra SET nik = '5205024210000002' WHERE idsobat = '520525110013' AND (nik IS NULL OR nik = ''); -- Riska amelia ade putri
UPDATE mitra SET nik = '5205025607000001' WHERE idsobat = '520526050071' AND (nik IS NULL OR nik = ''); -- Kuratul Aini
UPDATE mitra SET nik = '5205025501040001' WHERE idsobat = '520526050055' AND (nik IS NULL OR nik = ''); -- Anggun Purnama
UPDATE mitra SET nik = '5205024211990001' WHERE idsobat = '520525110040' AND (nik IS NULL OR nik = ''); -- Novitta islamiyah
UPDATE mitra SET nik = '5205024805010001' WHERE idsobat = '520526050159' AND (nik IS NULL OR nik = ''); -- SULFAH MEIASTRI
UPDATE mitra SET nik = '5205046401000001' WHERE idsobat = '520526050138' AND (nik IS NULL OR nik = ''); -- Muslimatun Fitriah
UPDATE mitra SET nik = '5205072604000002' WHERE idsobat = '520526050208' AND (nik IS NULL OR nik = ''); -- AL BIMA
UPDATE mitra SET nik = '5205046607020001' WHERE idsobat = '520526050070' AND (nik IS NULL OR nik = ''); -- JUMRIATI
UPDATE mitra SET nik = '5205070609830001' WHERE idsobat = '520526050223' AND (nik IS NULL OR nik = ''); -- Nurdin
UPDATE mitra SET nik = '5205070810960002' WHERE idsobat = '520526050244' AND (nik IS NULL OR nik = ''); -- nasrudin
UPDATE mitra SET nik = '5205071101850001' WHERE idsobat = '520522100174' AND (nik IS NULL OR nik = ''); -- Parli purmasidi
UPDATE mitra SET nik = '5205072808980001' WHERE idsobat = '520525110180' AND (nik IS NULL OR nik = ''); -- OPHIYANSYAH ACHRUL PUTRA
UPDATE mitra SET nik = '5205070811010004' WHERE idsobat = '520526050023' AND (nik IS NULL OR nik = ''); -- Rangga Barani Satria
UPDATE mitra SET nik = '5205070209860001' WHERE idsobat = '520522100006' AND (nik IS NULL OR nik = ''); -- ABD. HADI IRAWAN
UPDATE mitra SET nik = '5205071804040002' WHERE idsobat = '520526050050' AND (nik IS NULL OR nik = ''); -- Muhammad Arkham
UPDATE mitra SET nik = '5205072212940001' WHERE idsobat = '520522100165' AND (nik IS NULL OR nik = ''); -- Rasmini
UPDATE mitra SET nik = '5271012802880005' WHERE idsobat = '520523110098' AND (nik IS NULL OR nik = ''); -- Sofian Halidin
UPDATE mitra SET nik = '5205070504980003' WHERE idsobat = '520526050176' AND (nik IS NULL OR nik = ''); -- Sulaiman
UPDATE mitra SET nik = '5205074706920001' WHERE idsobat = '520523030048' AND (nik IS NULL OR nik = ''); -- DEWI PURWATI
UPDATE mitra SET nik = '5205024107990218' WHERE idsobat = '520526050125' AND (nik IS NULL OR nik = ''); -- Emi yuliana
UPDATE mitra SET nik = '5205030202990002' WHERE idsobat = '520526050126' AND (nik IS NULL OR nik = ''); -- Imansyah
UPDATE mitra SET nik = '5205031306910001' WHERE idsobat = '520526050022' AND (nik IS NULL OR nik = ''); -- Herman
UPDATE mitra SET nik = '5205030603980003' WHERE idsobat = '520526050121' AND (nik IS NULL OR nik = ''); -- Iskandar Julkarnain
UPDATE mitra SET nik = '5205032205980001' WHERE idsobat = '520526050231' AND (nik IS NULL OR nik = ''); -- Muhamad erwinsyah Putra
UPDATE mitra SET nik = '5205030905980001' WHERE idsobat = '520526050059' AND (nik IS NULL OR nik = ''); -- Irawan Dandi
UPDATE mitra SET nik = '5205032110960002' WHERE idsobat = '520526050193' AND (nik IS NULL OR nik = ''); -- Amrizal
UPDATE mitra SET nik = '5205035512000001' WHERE idsobat = '520526050151' AND (nik IS NULL OR nik = ''); -- Denti Alifia Ramdhoani
UPDATE mitra SET nik = '5205035301990001' WHERE idsobat = '520526050147' AND (nik IS NULL OR nik = ''); -- Aida rahayu
UPDATE mitra SET nik = '5205036502010003' WHERE idsobat = '520526050017' AND (nik IS NULL OR nik = ''); -- Muliani
UPDATE mitra SET nik = '5205035005020002' WHERE idsobat = '520526050144' AND (nik IS NULL OR nik = ''); -- Suhada
UPDATE mitra SET nik = '5205035003020008' WHERE idsobat = '520526050235' AND (nik IS NULL OR nik = ''); -- Nurmiftahul jannah
UPDATE mitra SET nik = '5205034801030001' WHERE idsobat = '520526050047' AND (nik IS NULL OR nik = ''); -- Weni Rahayu
UPDATE mitra SET nik = '5205034506010001' WHERE idsobat = '520526050204' AND (nik IS NULL OR nik = ''); -- YULIATI
UPDATE mitra SET nik = '5205071209950001' WHERE idsobat = '520526050015' AND (nik IS NULL OR nik = ''); -- ARIS MUNANDAR
UPDATE mitra SET nik = '5206096407920001' WHERE idsobat = '520526050108' AND (nik IS NULL OR nik = ''); -- Vera Novitasari
UPDATE mitra SET nik = '5205011001000006' WHERE idsobat = '520526050263' AND (nik IS NULL OR nik = ''); -- Miftach Khaerul Anas
UPDATE mitra SET nik = '5205010407930003' WHERE idsobat = '520525110198' AND (nik IS NULL OR nik = ''); -- Moh Ma''ruf
UPDATE mitra SET nik = '5205052004980004' WHERE idsobat = '520525110003' AND (nik IS NULL OR nik = ''); -- Alfi sahrir
UPDATE mitra SET nik = '5205016205950002' WHERE idsobat = '520526050045' AND (nik IS NULL OR nik = ''); -- NURRATUL FIDA
UPDATE mitra SET nik = '5205015203050002' WHERE idsobat = '520526050084' AND (nik IS NULL OR nik = ''); -- ISRATUL AINI
UPDATE mitra SET nik = '5205014503950001' WHERE idsobat = '520522100136' AND (nik IS NULL OR nik = ''); -- Nurul rofiah
UPDATE mitra SET nik = '5205016605010004' WHERE idsobat = '520523110049' AND (nik IS NULL OR nik = ''); -- Nurus Cahaya
UPDATE mitra SET nik = '5205015212010005' WHERE idsobat = '520526050240' AND (nik IS NULL OR nik = ''); -- Putri Safitri
UPDATE mitra SET nik = '5205016407030003' WHERE idsobat = '520526050259' AND (nik IS NULL OR nik = ''); -- Salsabillah Ananda
UPDATE mitra SET nik = '5205014911000005' WHERE idsobat = '520526050273' AND (nik IS NULL OR nik = ''); -- Husnul Wahyu Lestari
UPDATE mitra SET nik = '5205050502010001' WHERE idsobat = '520526050191' AND (nik IS NULL OR nik = ''); -- Nashril Muhammad Nur Aliandi
UPDATE mitra SET nik = '5205016911830001' WHERE idsobat = '520522030021' AND (nik IS NULL OR nik = ''); -- Titin Fatmawati
UPDATE mitra SET nik = '5205016701930003' WHERE idsobat = '520526050243' AND (nik IS NULL OR nik = ''); -- Triana Amalia ND
UPDATE mitra SET nik = '5205015209920001' WHERE idsobat = '520522090017' AND (nik IS NULL OR nik = ''); -- ATRI NURUL AFIAH
UPDATE mitra SET nik = '5205054609890002' WHERE idsobat = '520525110028' AND (nik IS NULL OR nik = ''); -- SRI SUSANTI
UPDATE mitra SET nik = '5205012309020002' WHERE idsobat = '520526050255' AND (nik IS NULL OR nik = ''); -- Taufikurahman
UPDATE mitra SET nik = '3508102201980005' WHERE idsobat = '520526050271' AND (nik IS NULL OR nik = ''); -- Al Busran
UPDATE mitra SET nik = '5205010606990003' WHERE idsobat = '520525110266' AND (nik IS NULL OR nik = ''); -- Demi wiliansyah
UPDATE mitra SET nik = '5205011609970005' WHERE idsobat = '520526050260' AND (nik IS NULL OR nik = ''); -- FERRY ZAINI FIRDHAUS
UPDATE mitra SET nik = '5205016003030003' WHERE idsobat = '520526050146' AND (nik IS NULL OR nik = ''); -- Aisiatul Juwariyah
UPDATE mitra SET nik = '5205016508980003' WHERE idsobat = '520525110158' AND (nik IS NULL OR nik = ''); -- Auliyatun Faatihah
UPDATE mitra SET nik = '5205016905010005' WHERE idsobat = '520525110148' AND (nik IS NULL OR nik = ''); -- Ida Puspita
UPDATE mitra SET nik = '5205014107980275' WHERE idsobat = '520526050031' AND (nik IS NULL OR nik = ''); -- IFA NURUL AISYAH
UPDATE mitra SET nik = '5205015103950004' WHERE idsobat = '520526050056' AND (nik IS NULL OR nik = ''); -- ISARATUSSARNI
UPDATE mitra SET nik = '5205050401830001' WHERE idsobat = '520526050269' AND (nik IS NULL OR nik = ''); -- Emros Brata ekas
UPDATE mitra SET nik = '5271020405860003' WHERE idsobat = '520526050257' AND (nik IS NULL OR nik = ''); -- Fathurrahman Hazairin
UPDATE mitra SET nik = '5205052209860001' WHERE idsobat = '520526050094' AND (nik IS NULL OR nik = ''); -- Hery zulhaid
UPDATE mitra SET nik = '5205051707040001' WHERE idsobat = '520526050030' AND (nik IS NULL OR nik = ''); -- Muhammad Rizalul Haq
UPDATE mitra SET nik = '5205052401940004' WHERE idsobat = '520526050001' AND (nik IS NULL OR nik = ''); -- Syafriadin Saputra
UPDATE mitra SET nik = '5205055009000002' WHERE idsobat = '520526050037' AND (nik IS NULL OR nik = ''); -- Gifa Nurfauziah
UPDATE mitra SET nik = '5205054102030006' WHERE idsobat = '520525110114' AND (nik IS NULL OR nik = ''); -- Mimi febriana
UPDATE mitra SET nik = '5205056407010001' WHERE idsobat = '520526050156' AND (nik IS NULL OR nik = ''); -- Nes watul jannah
UPDATE mitra SET nik = '5205054907880002' WHERE idsobat = '520526050039' AND (nik IS NULL OR nik = ''); -- Nining Yuliana
UPDATE mitra SET nik = '5205065309030006' WHERE idsobat = '520526050264' AND (nik IS NULL OR nik = ''); -- NURRUL FADILLAH
UPDATE mitra SET nik = '5205022505900002' WHERE idsobat = '520526050266' AND (nik IS NULL OR nik = ''); -- Waliey muftihakam marza karyadi
UPDATE mitra SET nik = '5205064510030002' WHERE idsobat = '520526050133' AND (nik IS NULL OR nik = ''); -- Susi susanti
UPDATE mitra SET nik = '5205075811050001' WHERE idsobat = '520526050131' AND (nik IS NULL OR nik = ''); -- Lastry
UPDATE mitra SET nik = '5203046903970005' WHERE idsobat = '520526050215' AND (nik IS NULL OR nik = ''); -- BAIQ ERINE WAHYU CHANDRA NIRWANA
UPDATE mitra SET nik = '5205074505990006' WHERE idsobat = '520526050238' AND (nik IS NULL OR nik = ''); -- SUMIATI
UPDATE mitra SET nik = '5205060709970003' WHERE idsobat = '520525050015' AND (nik IS NULL OR nik = ''); -- Wahyudin
UPDATE mitra SET nik = '5205052803010004' WHERE idsobat = '520525110255' AND (nik IS NULL OR nik = ''); -- MUAMAR SURAJUDIN
UPDATE mitra SET nik = '5205062209960002' WHERE idsobat = '520526050187' AND (nik IS NULL OR nik = ''); -- Muhammad Nur Azhari
UPDATE mitra SET nik = '5205066304040001' WHERE idsobat = '520526050066' AND (nik IS NULL OR nik = ''); -- AMELIA
UPDATE mitra SET nik = '5205065207040001' WHERE idsobat = '520525050042' AND (nik IS NULL OR nik = ''); -- Etty Moto Logis
UPDATE mitra SET nik = '5205065509030001' WHERE idsobat = '520526050048' AND (nik IS NULL OR nik = ''); -- Haifa Paradisa
UPDATE mitra SET nik = '5205064106010002' WHERE idsobat = '520526050089' AND (nik IS NULL OR nik = ''); -- Levy Viola Ovaliani
UPDATE mitra SET nik = '5205064110000003' WHERE idsobat = '520525110009' AND (nik IS NULL OR nik = ''); -- Lisnawati
UPDATE mitra SET nik = '5205066711000003' WHERE idsobat = '520526050097' AND (nik IS NULL OR nik = ''); -- Yulia Safitri
UPDATE mitra SET nik = '5205065803020001' WHERE idsobat = '520526050088' AND (nik IS NULL OR nik = ''); -- Miftahul Riyadah
UPDATE mitra SET nik = '5205064803010004' WHERE idsobat = '520526050044' AND (nik IS NULL OR nik = ''); -- Nawasari Dwi Sulistiawati
UPDATE mitra SET nik = '5205052412910004' WHERE idsobat = '520526050272' AND (nik IS NULL OR nik = ''); -- Dedi rachmat
UPDATE mitra SET nik = '5205054505900002' WHERE idsobat = '520522010023' AND (nik IS NULL OR nik = ''); -- ATIKASARI
UPDATE mitra SET nik = '5205016501890003' WHERE idsobat = '520526050218' AND (nik IS NULL OR nik = ''); -- Sri Handayani
UPDATE mitra SET nik = '5205056008990004' WHERE idsobat = '520526050192' AND (nik IS NULL OR nik = ''); -- Dyah Rahayu
UPDATE mitra SET nik = '5205081312930003' WHERE idsobat = '520526050261' AND (nik IS NULL OR nik = ''); -- Dermawan
UPDATE mitra SET nik = '5205085906000003' WHERE idsobat = '520526050033' AND (nik IS NULL OR nik = ''); -- Dhia Istiqomah
UPDATE mitra SET nik = '5205086606030002' WHERE idsobat = '520526050254' AND (nik IS NULL OR nik = ''); -- Fathuljannah
UPDATE mitra SET nik = '5205084107990089' WHERE idsobat = '520525110212' AND (nik IS NULL OR nik = ''); -- Jumiati
UPDATE mitra SET nik = '5205086802900002' WHERE idsobat = '520526050068' AND (nik IS NULL OR nik = ''); -- Kurniati
UPDATE mitra SET nik = '5205085712000002' WHERE idsobat = '520526050034' AND (nik IS NULL OR nik = ''); -- Nur Fitrah
UPDATE mitra SET nik = '5205055712010002' WHERE idsobat = '520526050013' AND (nik IS NULL OR nik = ''); -- Nursani
UPDATE mitra SET nik = '5205074404930003' WHERE idsobat = '520523030037' AND (nik IS NULL OR nik = ''); -- NURWAHIDAH
UPDATE mitra SET nik = '5205012301950003' WHERE idsobat = '520523030078' AND (nik IS NULL OR nik = ''); -- MUSTARI
UPDATE mitra SET nik = '5205052606960006' WHERE idsobat = '520526050105' AND (nik IS NULL OR nik = ''); -- ARIFIN

SELECT COUNT(*) AS total_nik_terisi FROM mitra WHERE nik IS NOT NULL AND nik != '' AND tahun='2026';