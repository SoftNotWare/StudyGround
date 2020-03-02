CREATE SCHEMA `study_ground` DEFAULT CHARACTER SET utf8mb4 ;

ALTER TABLE `study_ground`.`users` AUTO_INCREMENT = 100001;

CREATE TABLE `study_ground`.`users` (
  `id` INT NOT NULL,
  `email` VARCHAR(45) NOT NULL COMMENT '邮箱地址',
  `password` VARCHAR(64) NOT NULL COMMENT '密码',
  `created_at` DATETIME NOT NULL COMMENT '注册时间',
  `login_at` DATETIME NOT NULL COMMENT '最近登录时间\n',
  `score` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '积分',
  `avatar` VARCHAR(128) NULL COMMENT '头像',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `unique` (`email` ASC)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;