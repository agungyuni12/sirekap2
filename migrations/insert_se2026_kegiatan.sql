-- ============================================================
-- INSERT kegiatan Sensus Ekonomi 2026
-- OB sementara: 170 PCL x 2.5 = 425 OB, 32 PML x 2.5 = 80 OB
-- Update setelah data petugas final
-- ============================================================

INSERT INTO kegiatan (nama, vol, satuan, harga, total, mak, tanggaran)
VALUES
  ('Pendataan Sensus Ekonomi 2026',
   425, 'O-B', 4629000, 425 * 4629000,
   '2902.BMA.006.005.B.521213', 2026),

  ('Pemeriksa Lapangan Sensus Ekonomi 2026 (PML)',
   80, 'O-B', 4877000, 80 * 4877000,
   '2902.FAN.ZZ1.051.A.521213', 2026);
