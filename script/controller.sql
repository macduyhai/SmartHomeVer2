DROP TABLE IF EXISTS `controllers`;

CREATE TABLE `controllers` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT(20),
    `chip_id` VARCHAR(64) UNIQUE,
    `flash_chip_id` VARCHAR(64) UNIQUE,
    `ide_flash_size` VARCHAR(64) ,
    `real_flash_size` VARCHAR(64) ,
    `soft_ap_ip` VARCHAR(64) ,
    `soft_ap_mac` VARCHAR(64) ,
    `station_mac` VARCHAR(64) ,
    `serial` VARCHAR(64) ,
    `name` VARCHAR(64) ,
    `type` VARCHAR(64) ,
    `active` BOOLEAN,
    `number_device` BIGINT(20),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX device_name (`name`)
);
