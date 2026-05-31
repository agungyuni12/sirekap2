-- ============================================================
-- INSERT kegiatan Sensus Ekonomi 2026 (idempotent - aman dijalankan ulang)
-- vol=1 = satuan per OB; total = anggaran keseluruhan (170 PCL + 32 PML x 2.5 OB)
-- ============================================================

INSERT INTO kegiatan (nama, vol, satuan, harga, total, mak, tanggaran)
SELECT 'Pendataan Sensus Ekonomi 2026', 1, 'O-B', 4629000, 425 * 4629000,
       '2902.BMA.006.005.B.521213', 2026
WHERE NOT EXISTS (
    SELECT 1 FROM kegiatan
    WHERE nama = 'Pendataan Sensus Ekonomi 2026' AND tanggaran = 2026
);

INSERT INTO kegiatan (nama, vol, satuan, harga, total, mak, tanggaran)
SELECT 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)', 1, 'O-B', 4877000, 80 * 4877000,
       '2902.FAN.ZZ1.051.A.521213', 2026
WHERE NOT EXISTS (
    SELECT 1 FROM kegiatan
    WHERE nama = 'Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)' AND tanggaran = 2026
);
