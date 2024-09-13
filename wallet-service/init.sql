CREATE DATABASE IF NOT EXISTS wallet;

USE wallet;

CREATE TABLE IF NOT EXISTS clients (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL
);

CREATE TABLE IF NOT EXISTS accounts (
    id CHAR(36) PRIMARY KEY,
    client_id CHAR(36),
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT FK_accountClient FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transactions (
    id CHAR(36) PRIMARY KEY,
    account_id_from CHAR(36),
    account_id_to CHAR(36),
    amount DECIMAL(10, 2) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT FK_transactionAccountFrom FOREIGN KEY (account_id_from) REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT FK_transactionAccountTo FOREIGN KEY (account_id_to) REFERENCES accounts(id) ON DELETE CASCADE
);

-- insert clients
INSERT INTO clients (id, name, email, created_at)
SELECT 'a7573efa-3053-4b86-9ece-db7cb18d0d39', 'Jhon Doe', 'jhon@hotmail.com', NOW()
WHERE NOT EXISTS (SELECT 1 FROM clients WHERE id = 'a7573efa-3053-4b86-9ece-db7cb18d0d39');

INSERT INTO clients (id, name, email, created_at)
SELECT '05e3f336-770f-4747-8ffd-a9b666bf5706', 'Jane Doe', 'jane@hotmail.com', NOW()
WHERE NOT EXISTS (SELECT 1 FROM clients WHERE id = '05e3f336-770f-4747-8ffd-a9b666bf5706');
--

-- insert accounts
INSERT INTO accounts (id, client_id, balance, created_at)
SELECT 'f1e4b8b2-6a44-4d5b-83e8-111c8f4dcbeb', '05e3f336-770f-4747-8ffd-a9b666bf5706', 1000.00, NOW()
WHERE NOT EXISTS (SELECT 1 FROM accounts WHERE id = 'f1e4b8b2-6a44-4d5b-83e8-111c8f4dcbeb');

INSERT INTO accounts (id, client_id, balance, created_at)
SELECT '3c79aede-12a4-4f65-9307-6d4d1a35b1c8', 'a7573efa-3053-4b86-9ece-db7cb18d0d39', 2500.00, NOW()
WHERE NOT EXISTS (SELECT 1 FROM accounts WHERE id = '3c79aede-12a4-4f65-9307-6d4d1a35b1c8');
--