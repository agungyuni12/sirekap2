-- ============================================================
-- Penugasan peran penilai organik (PML / Korwil) per kegiatan.
--
-- Sebelumnya peran seorang penilai (menilai PPL atau menilai PML Mitra)
-- ditentukan murni dari user.level GLOBAL ("korwil" vs level lain) —
-- tidak bisa beda per kegiatan, dan tidak ada UI untuk mengaturnya (harus
-- edit level di database langsung). Sekarang: penilaian_kegiatan_petugas
-- (dulu cuma roster mitra "yang dinilai") diperluas agar bisa juga
-- menampung akun organik (tabel `user`) sebagai entri "penilai" dengan
-- peran 'pml' (menilai PPL) atau 'korwil' (menilai PML Mitra), diatur
-- admin lewat menu Kelola Petugas — per kegiatan, bukan global.
--
-- Baris mitra (peran ppl/pml "yang dinilai", existing): idsobat terisi,
-- user_id NULL. Baris organik (peran pml/korwil "penilai", baru):
-- user_id terisi, idsobat NULL. MySQL unique key mengizinkan banyak NULL
-- berdampingan, jadi kedua unique key ini aman.
-- ============================================================

ALTER TABLE penilaian_kegiatan_petugas
    MODIFY COLUMN idsobat VARCHAR(20) NULL,
    ADD COLUMN user_id INT NULL AFTER idsobat,
    MODIFY COLUMN peran ENUM('ppl','pml','korwil') NOT NULL,
    ADD UNIQUE KEY uniq_kegiatan_user (kegiatan_id, user_id),
    ADD INDEX idx_user (user_id);
