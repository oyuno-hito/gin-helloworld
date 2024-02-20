CREATE TABLE IF NOT EXISTS `gin-helloworld`.`users` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `login_id` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    `role_id` INT NOT NULL,
    CONSTRAINT `fk_users_role_id`
    FOREIGN KEY (`role_id`)
    REFERENCES `gin-helloworld`.`roles` (`id`)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
)
ENGINE = InnoDB;