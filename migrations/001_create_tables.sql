-- Create patients table
CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(50) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birthdate DATE NOT NULL,
    sex ENUM('male', 'female') NOT NULL,
    address TEXT,
    phone VARCHAR(20),
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_first_name (first_name),
    INDEX idx_last_name (last_name),
    INDEX idx_phone (phone),
    INDEX idx_email (email)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- Create work_orders table
CREATE TABLE IF NOT EXISTS work_orders (
    no_order VARCHAR(50) PRIMARY KEY,
    patient_id VARCHAR(50) NOT NULL,
    analyst VARCHAR(100) NOT NULL,
    doctor VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients (id) ON DELETE CASCADE,
    INDEX idx_patient_id (patient_id),
    INDEX idx_analyst (analyst),
    INDEX idx_doctor (doctor)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- Create work_order_test_codes table (for storing array of test codes)
CREATE TABLE IF NOT EXISTS work_order_test_codes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    no_order VARCHAR(50) NOT NULL,
    test_code VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (no_order) REFERENCES work_orders (no_order) ON DELETE CASCADE,
    INDEX idx_no_order (no_order),
    INDEX idx_test_code (test_code)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;