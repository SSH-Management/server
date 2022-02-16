ALTER TABLE groups ADD COLUMN is_system_group boolean DEFAULT FALSE;

CREATE INDEX idx_is_system_group ON groups (is_system_group);
