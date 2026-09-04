-- ============================================================
-- Roster petugas (PPL & PML Mitra) per kegiatan Penilaian Mitra.
--
-- Sebelumnya "Nama Petugas" di form Input Penilaian mencari bebas ke seluruh
-- tabel `mitra` tanpa terikat kegiatan — berisiko salah pilih orang yang
-- sebenarnya tidak bertugas di kegiatan tsb. Sekarang: hanya admin (Subject
-- Matter) yang boleh menambah/menghapus petugas ke roster kegiatan lewat
-- menu "Kelola Petugas"; Korwil dan PML (organik/mitra) hanya memilih dari
-- roster yang sudah ada saat menilai — tidak bisa menambah nama baru sendiri.
-- ============================================================

CREATE TABLE IF NOT EXISTS penilaian_kegiatan_petugas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    kegiatan_id INT NOT NULL, -- merujuk ke penilaian_kegiatan.id
    idsobat VARCHAR(20) NOT NULL, -- merujuk ke mitra.idsobat
    peran ENUM('ppl','pml') NOT NULL,
    added_by INT NOT NULL, -- user.id admin yang menambahkan
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uniq_kegiatan_idsobat (kegiatan_id, idsobat),
    KEY idx_kegiatan_peran (kegiatan_id, peran)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
