DROP INDEX idx_is_system_group ON `groups`;

ALTER TABLE `groups` DROP COLUMN is_system_group;

