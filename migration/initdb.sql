CREATE TABLE IF NOT EXISTS `expenses` (
  `id` INT(10) unsigned NOT NULL,
  `description` VARCHAR(20) NOT NULL,
  `amount` DECIMAL(13,2) NOT NULL,
  `date` DATETIME DEFAULT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY `id` (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `incomes` (
  `id` INT(10) unsigned NOT NULL,
  `description` VARCHAR(20) NOT NULL,
  `amount` DECIMAL(13,2) NOT NULL,
  `date` DATETIME DEFAULT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY `id` (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci AUTO_INCREMENT=1;
