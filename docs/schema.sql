CREATE DATABASE IF NOT EXISTS `htz_blog` DEFAULT CHARACTER SET utf8;

USE `htz_blog`;

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(100) NOT NULL DEFAULT '',
	`content_abstract` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
	`content` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
	`nav_name` VARCHAR(50) NOT NULL,
	`read_count` INT UNSIGNED NOT NULL DEFAULT 0,
	`created_at` DATE NOT NULL,
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `mime` VARCHAR(50) NOT NULL,
  `size` INT UNSIGNED NOT NULL,
  `created_at` DATE NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `navs`;
CREATE TABLE `navs` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


INSERT INTO `navs`(`name`) VALUES(`首页`）,（`生理健康`）,（`心理自在`）,（`事理圆满`）,（`国学园地`）,（`修行原理`）,（`讲师影片`）,（`书院活动`）,（`学院心得`）,（`其他`);
