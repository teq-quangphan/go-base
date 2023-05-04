
CREATE TABLE if not exists `refresh_token`
(
    `id`          varchar(64) PRIMARY KEY  COMMENT 'uuid' ,
    `token`       text,
    `user_id`     varchar(255),
    `expired_at`  integer NOT NULL
    );

CREATE INDEX `idx_user_id` ON `refresh_token` (`user_id`);

ALTER TABLE `refresh_token` ADD CONSTRAINT uq_token UNIQUE(`user_id`);