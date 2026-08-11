-- create_evaluasi_petugas_table.sql tidak menyebutkan COLLATE secara eksplisit,
-- jadi tabel ini mengikuti default collation server (utf8mb4_0900_ai_ci di server
-- produksi ini), sementara tabel-tabel lain di aplikasi (rekap, kegiatan, mitra,
-- surat, dst) memakai utf8mb4_general_ci. Perbedaan collation ini menyebabkan
-- error 1267 "Illegal mix of collations" setiap kali evaluasi_petugas di-JOIN
-- langsung ke kolom string tabel lain (mis. yang_dinilai_idsobat vs rekap.idsobat).
-- Menyamakan ke utf8mb4_general_ci di sini (bukan mengubah seluruh database)
-- supaya konsisten dengan konvensi collation yang sudah dipakai semua tabel lain.
ALTER TABLE evaluasi_petugas CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
