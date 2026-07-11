-- ============================================================
-- Kolom BAPP SE2026 (Berita Acara Pemeriksaan Pekerjaan) di tabel rekap
-- Mengikuti pola id_spk yang sudah ada. Aman dijalankan berulang.
-- ============================================================

DROP PROCEDURE IF EXISTS _add_bapp_se2026_columns;

DELIMITER //
CREATE PROCEDURE _add_bapp_se2026_columns()
BEGIN
    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'id_bapp1') THEN
        ALTER TABLE rekap ADD COLUMN id_bapp1 VARCHAR(255) NULL AFTER id_spk;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'tgl_bapp1') THEN
        ALTER TABLE rekap ADD COLUMN tgl_bapp1 DATE NULL AFTER id_bapp1;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'id_bapp2') THEN
        ALTER TABLE rekap ADD COLUMN id_bapp2 VARCHAR(255) NULL AFTER tgl_bapp1;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'tgl_bapp2') THEN
        ALTER TABLE rekap ADD COLUMN tgl_bapp2 DATE NULL AFTER id_bapp2;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'realisasi_sls') THEN
        ALTER TABLE rekap ADD COLUMN realisasi_sls INT NOT NULL DEFAULT 0 AFTER jumlah_sls;
    END IF;

    -- Nomor Surat Pernyataan Penyelesaian Lapangan (PML) - format beda dari BAPP,
    -- jadi kolom sendiri, bukan reuse id_bapp1.
    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'id_pernyataan1') THEN
        ALTER TABLE rekap ADD COLUMN id_pernyataan1 VARCHAR(255) NULL AFTER tgl_bapp2;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'rekap' AND COLUMN_NAME = 'tgl_pernyataan1') THEN
        ALTER TABLE rekap ADD COLUMN tgl_pernyataan1 DATE NULL AFTER id_pernyataan1;
    END IF;
END //
DELIMITER ;

CALL _add_bapp_se2026_columns();
DROP PROCEDURE IF EXISTS _add_bapp_se2026_columns;

SELECT 'Migration add_bapp_se2026 selesai.' AS status;
