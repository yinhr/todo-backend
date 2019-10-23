
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
  id              BINARY(16) NOT NULL PRIMARY KEY,
  email           VARCHAR(255) NOT NULL UNIQUE,
  password_digest CHAR(60) NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- +migrate Down
DROP TABLE IF EXISTS users;