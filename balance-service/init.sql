CREATE DATABASE IF NOT EXISTS wallet;

USE wallet;

CREATE TABLE IF NOT EXISTS accounts (
    id CHAR(36) PRIMARY KEY,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- inserts accounts
INSERT INTO accounts (id, balance, updated_at)
SELECT 'f1e4b8b2-6a44-4d5b-83e8-111c8f4dcbeb', 1000.00, NOW()
WHERE NOT EXISTS (SELECT 1 FROM accounts WHERE id = 'f1e4b8b2-6a44-4d5b-83e8-111c8f4dcbeb');

INSERT INTO accounts (id, balance, updated_at)
SELECT '3c79aede-12a4-4f65-9307-6d4d1a35b1c8', 2500.00, NOW()
WHERE NOT EXISTS (SELECT 1 FROM accounts WHERE id = '3c79aede-12a4-4f65-9307-6d4d1a35b1c8');
--