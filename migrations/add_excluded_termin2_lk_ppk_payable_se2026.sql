-- ============================================================
-- lk_ppk_payable_se2026: tambah kolom excluded_termin2
-- Arahan user: Moh Ma'ruf (520525110198) & M Iksan (520525110177)
-- TIDAK dapat BAPP/BAST di Termin 2 (masih valid di Termin 1).
-- ============================================================

ALTER TABLE lk_ppk_payable_se2026 ADD COLUMN excluded_termin2 TINYINT(1) NOT NULL DEFAULT 0;

UPDATE lk_ppk_payable_se2026 SET excluded_termin2 = 1
WHERE idsobat IN ('520525110177', '520525110198');
