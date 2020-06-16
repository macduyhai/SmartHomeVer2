DROP TABLE IF EXISTS `devices`;

CREATE TABLE `devices` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `chip_id` VARCHAR(64) UNIQUE,
    `flash_chip_id` VARCHAR(64) UNIQUE,
    `ide_flash_size` VARCHAR(64) UNIQUE,
    `real_flash_size` VARCHAR(64) UNIQUE,
    `soft_ap_ip` VARCHAR(64) UNIQUE,
    `soft_ap_mac` VARCHAR(64) UNIQUE,
    `Station_mac` VARCHAR(64) UNIQUE,
    `serial` VARCHAR(64) UNIQUE,
    `name` VARCHAR(64) UNIQUE,
    `type` VARCHAR(64) UNIQUE,
    `laststate` BOOLEAN,
    `newstate` BOOLEAN,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX device_name (`name`)
);
