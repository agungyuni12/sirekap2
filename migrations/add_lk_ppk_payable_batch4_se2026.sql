-- ============================================================
-- Tambah 3 PPL batch 4 ke lk_ppk_payable_se2026 - seq & seq_all LANJUT
-- dari batch 1-3 (tidak reset). Tidak ada PML baru batch ini (PML mereka
-- - BUHARI SALAM, ADI NEGORO HERLAMBANG, ZORDIAN - sudah payable sejak
-- batch 1/2).
-- ============================================================

INSERT INTO lk_ppk_payable_se2026 (idsobat, nama, role, seq, seq_all, batch) VALUES
('520526050109','IIN NURROHMADANIL','ppl',231,263,4),
('520526050093','Moh Naufal','ppl',232,264,4),
('520526070001','Sarwan','ppl',233,265,4);

SELECT batch, role, COUNT(*) AS jumlah, MIN(seq) AS seq_min, MAX(seq) AS seq_max, MIN(seq_all) AS seq_all_min, MAX(seq_all) AS seq_all_max FROM lk_ppk_payable_se2026 GROUP BY batch, role ORDER BY batch, role;
