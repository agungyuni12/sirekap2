-- ============================================================
-- Periode kegiatan Penilaian Mitra: dikelola admin (Kelola Petugas), bukan
-- diketik bebas oleh penilai di form Input Penilaian.
--
-- Sebelumnya "Periode"/"Tahun" di form Input Penilaian cuma teks/dropdown
-- bebas yang diisi siapa saja yang login, dan roster (penilaian_kegiatan_petugas)
-- di-scope ke kegiatan_id saja — tidak ada cara admin membedakan roster
-- Triwulan I vs Triwulan II utk kegiatan yang sama.
--
-- Sekaligus membenahi bug laten: unique key evaluasi_petugas cuma
-- (kegiatan_id, yang_dinilai_idsobat, tahap) — tidak melibatkan periode, jadi
-- menilai orang yang sama di kegiatan yang sama pada periode berikutnya
-- MENIMPA baris periode sebelumnya alih-alih jadi baris baru. Sekarang
-- di-scope ke periode_id.
--
-- CATATAN SEBELUM MENJALANKAN: cek `SELECT COUNT(*) FROM evaluasi_petugas`.
-- Baris lama (kalau ada) akan punya periode_id NULL — tetap tersimpan, tapi
-- tidak akan muncul lagi di Daftar/Detail/Dashboard karena semua query baru
-- mensyaratkan periode_id. Putuskan cara backfill-nya kalau tabel tidak kosong.
-- ============================================================

CREATE TABLE IF NOT EXISTS penilaian_kegiatan_periode (
    id INT AUTO_INCREMENT PRIMARY KEY,
    kegiatan_id INT NOT NULL, -- merujuk ke penilaian_kegiatan.id
    periode VARCHAR(100) NOT NULL,
    tahun VARCHAR(5) NOT NULL,
    added_by INT NOT NULL, -- user.id admin yang membuat
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uniq_kegiatan_periode_tahun (kegiatan_id, periode, tahun),
    KEY idx_kegiatan (kegiatan_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Roster (dari add_penilaian_kegiatan_petugas_roster.sql +
-- add_penilai_organik_kegiatan.sql) di-scope ulang ke periode, bukan kegiatan.
ALTER TABLE penilaian_kegiatan_petugas
    DROP INDEX uniq_kegiatan_idsobat,
    DROP INDEX uniq_kegiatan_user,
    DROP INDEX idx_kegiatan_peran,
    CHANGE COLUMN kegiatan_id periode_id INT NOT NULL,
    ADD UNIQUE KEY uniq_periode_idsobat (periode_id, idsobat),
    ADD UNIQUE KEY uniq_periode_user (periode_id, user_id),
    ADD INDEX idx_periode_peran (periode_id, peran);

-- Skor di-scope ke periode juga — nullable dulu supaya baris lama (kalau ada)
-- tidak hilang, cuma jadi "belum terkait periode manapun" (lihat catatan di atas).
ALTER TABLE evaluasi_petugas
    ADD COLUMN periode_id INT NULL AFTER kegiatan_id,
    DROP INDEX uniq_kegiatan_petugas_tahap,
    ADD UNIQUE KEY uniq_periode_petugas_tahap (periode_id, yang_dinilai_idsobat, tahap),
    ADD INDEX idx_periode (periode_id);
