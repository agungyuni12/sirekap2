CREATE TABLE IF NOT EXISTS surat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    idsobat VARCHAR(20) NOT NULL,
    nmitra VARCHAR(255),
    sbulan VARCHAR(20),
    stahun VARCHAR(10),
    stujuan VARCHAR(10),
    ksurat INT DEFAULT 1,
    nsurat VARCHAR(100),
    tglsurat DATE,
    nmrsurat VARCHAR(50),
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
