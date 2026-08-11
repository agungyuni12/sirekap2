-- ============================================================
-- Penilaian Kinerja Petugas Lapangan (PPL & PML)
-- Tahap 1: PML (organik/mitra) menilai PPL yang diawasi.
-- Tahap 2: Subject Matter (PJ Kegiatan) memfinalisasi nilai PPL
--          dan menilai langsung PML Mitra.
--
-- 5 aspek penilaian: kualitas hasil kerja, ketepatan waktu,
-- kepatuhan terhadap SOP, komunikasi dan koordinasi, sikap dan perilaku.
-- ============================================================

CREATE TABLE IF NOT EXISTS evaluasi_petugas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    -- Tidak pakai FOREIGN KEY (konsisten dengan seluruh migrasi lain di
    -- aplikasi ini, mis. `rekap`/`surat` juga tidak menegakkan FK ke
    -- `kegiatan`/`user`) — menghindari error #1824 "Failed to open the
    -- referenced table" yang muncul saat tabel induk tidak kompatibel
    -- untuk constraint InnoDB (mis. engine/collation berbeda di server produksi).
    kegiatan_id INT NOT NULL,
    penilai_id INT NOT NULL,
    -- Referensi via idsobat (bukan id auto-increment) karena tabel `mitra`
    -- di aplikasi ini memakai idsobat sebagai kunci alami, konsisten dengan
    -- `rekap.idsobat`, `surat.idsobat`, dst.
    yang_dinilai_idsobat VARCHAR(20) NOT NULL,
    peran_yang_dinilai ENUM('ppl','pml') NOT NULL,
    tahap TINYINT NOT NULL,
    skor_kualitas DECIMAL(5,2) NOT NULL,        -- Kualitas Hasil Kerja
    skor_ketepatan_waktu DECIMAL(5,2) NOT NULL, -- Ketepatan Waktu
    skor_kepatuhan_sop DECIMAL(5,2) NOT NULL,   -- Kepatuhan terhadap SOP
    skor_komunikasi DECIMAL(5,2) NOT NULL,      -- Komunikasi dan Koordinasi
    skor_sikap DECIMAL(5,2) NOT NULL,           -- Sikap dan Perilaku
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- Satu skor per (kegiatan, petugas dinilai, tahap): submit ulang = update nilai.
    UNIQUE KEY uniq_kegiatan_petugas_tahap (kegiatan_id, yang_dinilai_idsobat, tahap),
    KEY idx_kegiatan (kegiatan_id),
    KEY idx_yang_dinilai (yang_dinilai_idsobat),
    KEY idx_penilai (penilai_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Perluas kolom level user agar menampung nilai baru 'pml_mitra' (peran mitra
-- yang bertugas sebagai PML dan hanya berhak mengakses menu Penilaian Kinerja
-- PPL + Lapor Translok Mandiri). Aman dijalankan baik `level` sudah berupa
-- VARCHAR maupun ENUM lama: nilai data yang sudah ada tidak berubah.
ALTER TABLE user MODIFY COLUMN level VARCHAR(20) NOT NULL DEFAULT 'user';
