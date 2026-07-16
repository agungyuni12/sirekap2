-- ============================================================
-- Surat Pernyataan Kepala BPS SE2026 + reset nomor format lama
-- ============================================================

-- Tabel nomor Surat Pernyataan Kepala BPS - dokumen tunggal (bukan per-petugas),
-- 1 baris per termin/tahun.
CREATE TABLE IF NOT EXISTS surat_kepala_se2026 (
    id INT AUTO_INCREMENT PRIMARY KEY,
    termin INT NOT NULL,
    tahun VARCHAR(5) NOT NULL,
    nomor VARCHAR(255) NOT NULL,
    tanggal DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uniq_termin_tahun (termin, tahun)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Reset nomor BAPP/Surat Pernyataan yang sudah pernah dibuat pakai format lama
-- (B-XXX/BAPP-I-SE2026/5205/PL.200/2026 dst) supaya tombol "Buat" di /dashboard/se2026
-- generate ulang pakai format baru (B-BB.TT.XXX/.../Super.PPL|PML|KPL|BA/2026).
UPDATE rekap SET id_bapp1 = NULL, tgl_bapp1 = NULL WHERE id_bapp1 IS NOT NULL AND id_bapp1 != '';
UPDATE rekap SET id_bapp2 = NULL, tgl_bapp2 = NULL WHERE id_bapp2 IS NOT NULL AND id_bapp2 != '';
UPDATE rekap SET id_pernyataan1 = NULL, tgl_pernyataan1 = NULL WHERE id_pernyataan1 IS NOT NULL AND id_pernyataan1 != '';

SELECT 'Migration add_kepala_bps_se2026 selesai.' AS status;
