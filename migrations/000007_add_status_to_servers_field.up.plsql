CREATE TYPE server_status AS ENUM ('ok', 'not_serving', 'unknown');

ALTER TABLE servers ADD COLUMN status server_status DEFAULT 'unknown';
