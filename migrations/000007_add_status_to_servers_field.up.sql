ALTER TABLE `servers` ADD COLUMN status ENUM('not_serving', 'ok', 'unknown') DEFAULT 'unknown';
