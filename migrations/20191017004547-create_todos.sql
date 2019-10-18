
-- +migrate Up
CREATE TABLE IF NOT EXISTS todos
(
  id              BINARY(16) NOT NULL PRIMARY KEY,
  user_id         BINARY(16) NOT NULL,
  title           VARCHAR(255) NOT NULL,
  vital           BOOLEAN NOT NULL DEFAULT FALSE,
  done            BOOLEAN NOT NULL DEFAULT FALSE,
  due             DATETIME NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- +migrate Down
DROP TABLE IF EXISTS todos;
