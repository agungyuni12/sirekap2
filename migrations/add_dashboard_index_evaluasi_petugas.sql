-- Dashboard & Rekap Penilaian Kinerja Mitra: mempercepat query agregasi
-- (GROUP BY kegiatan_id + yang_dinilai_idsobat, filter per tahap) yang dipakai
-- GET /api/penilaian/dashboard dan GET /api/penilaian/rekap-dashboard.
ALTER TABLE evaluasi_petugas
    ADD INDEX idx_evaluasi_kegiatan_tahap (kegiatan_id, tahap, yang_dinilai_idsobat);
