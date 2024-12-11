CREATE TABLE IF NOT EXISTS `urls`
(
    `id`            BIGINT                          NOT NULL AUTO_INCREMENT,
    `original_url`  TEXT                            NOT NULL,
    `short_postfix` VARCHAR(10) COLLATE utf8mb4_bin NOT NULL UNIQUE,
    `created_at`    DATETIME                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);