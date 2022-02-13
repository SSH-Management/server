ALTER TABLE groups ADD COLUMN is_system_group TINYINT(1) DEFAULT 0;

CREATE INDEX idx_is_system_group ON groups (is_system_group);
