-- ============================================================
-- Penilaian Kinerja Petugas Lapangan (PPL & PML)
-- Tahap 1: PML (organik/mitra) menilai PPL yang diawasi.
-- Tahap 2: Subject Matter (PJ Kegiatan) memfinalisasi nilai PPL
--          dan menilai langsung PML Mitra.
-- ============================================================

CREATE TABLE IF NOT EXISTS evaluasi_petugas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    kegiatan_id INT NOT NULL,
    penilai_id INT NOT NULL,
    -- Referensi via idsobat (bukan id auto-increment) karena tabel `mitra`
    -- di aplikasi ini memakai idsobat sebagai kunci alami, konsisten dengan
    -- `rekap.idsobat`, `surat.idsobat`, dst.
    yang_dinilai_idsobat VARCHAR(20) NOT NULL,
    peran_yang_dinilai ENUM('ppl','pml') NOT NULL,
    tahap TINYINT NOT NULL,
    skor_kualitas DECIMAL(5,2) NOT NULL,
    skor_ketepatan_waktu DECIMAL(5,2) NOT NULL,
    skor_etika DECIMAL(5,2) NOT NULL,
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- Satu skor per (kegiatan, petugas dinilai, tahap): submit ulang = update nilai.
    UNIQUE KEY uniq_kegiatan_petugas_tahap (kegiatan_id, yang_dinilai_idsobat, tahap),
    KEY idx_kegiatan (kegiatan_id),
    KEY idx_yang_dinilai (yang_dinilai_idsobat),
    KEY idx_penilai (penilai_id),
    CONSTRAINT fk_evaluasi_kegiatan FOREIGN KEY (kegiatan_id) REFERENCES kegiatan(id),
    CONSTRAINT fk_evaluasi_penilai FOREIGN KEY (penilai_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Perluas kolom level user agar menampung nilai baru 'pml_mitra' (peran mitra
-- yang bertugas sebagai PML dan hanya berhak mengakses menu Penilaian Kinerja
-- PPL + Lapor Translok Mandiri). Aman dijalankan baik `level` sudah berupa
-- VARCHAR maupun ENUM lama: nilai data yang sudah ada tidak berubah.
ALTER TABLE user MODIFY COLUMN level VARCHAR(20) NOT NULL DEFAULT 'user';
