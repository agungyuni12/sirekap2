-- ============================================================
-- 7 PPL di bawah PML HAERUDDIN DIKELUARKAN dari LK PPK Termin 1 (dibayar
-- terpisah tgl 20 - lihat catatan di LK PPK Termin 1 xlsx versi terbaru).
--
-- Migrasi ini:
--   1. Menghapus mereka dari lk_ppk_payable_se2026 (kalau masih ada dari versi lama -
--      seharusnya sudah otomatis hilang stlh menjalankan add_lk_ppk_payable_se2026.sql
--      yang baru, TRUNCATE dulu sblm INSERT ulang - ini cuma jaga-jaga).
--   2. Mengosongkan nomor Pernyataan (id_pernyataan1) & BAPP (id_bapp1) yang SUDAH
--      TERLANJUR di-assign ke mereka lewat migrations/assign_nomor_surat_se2026_termin1.sql
--      versi lama (yg masih pakai daftar 229 orang, termasuk 7 orang ini) - supaya nanti
--      saat mereka dibayar terpisah tgl 20, nomornya dibuatkan baru dgn skema penomoran
--      sendiri, bukan numpang skema termin I tanggal 16-17 Juli.
-- ============================================================

DELETE FROM lk_ppk_payable_se2026 WHERE idsobat IN (
  '520522030024', -- Rachmad Soebari
  '520522100013', -- Bahril Qamar
  '520525110108', -- Nurlailah
  '520525110198', -- Moh Ma'ruf
  '520526050218', -- Sri Handayani
  '520526050243', -- Triana Amalia ND
  '520526050273'  -- Husnul Wahyu Lestari
) AND role = 'ppl';

UPDATE rekap SET id_pernyataan1 = NULL, tgl_pernyataan1 = NULL
WHERE idsobat IN ('520522030024','520522100013','520525110108','520525110198','520526050218','520526050243','520526050273')
  AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026'
  AND id_pernyataan1 LIKE 'B-07.16%';

UPDATE rekap SET id_bapp1 = NULL, tgl_bapp1 = NULL
WHERE idsobat IN ('520522030024','520522100013','520525110108','520525110198','520526050218','520526050243','520526050273')
  AND kegiatan = 'Pendataan Sensus Ekonomi 2026' AND tahun = '2026'
  AND id_bapp1 LIKE 'B-07.17%';

SELECT 'Migration reset_excluded_ppl_se2026 selesai.' AS status;
