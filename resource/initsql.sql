CREATE DATABASE IF NOT EXISTS `Direwolf`
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_unicode_ci;
use Direwolf;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20)   NOT NULL AUTO_INCREMENT,
    `email`       varchar(255) NOT NULL,
    `username`    varchar(255) NOT NULL,
    `password`    varchar(255) NOT NULL,
    `admin`       tinyint(1)   NOT NULL,
    `active`      tinyint(1)            DEFAULT NULL,
    `name`        varchar(255)          DEFAULT NULL,
    `description` varchar(255)          DEFAULT NULL,
    `avatar`      varchar(255)          DEFAULT NULL,
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by`   bigint(20)   NOT NULL DEFAULT '0',
    `update_time` timestamp    NULL ON UPDATE CURRENT_TIMESTAMP,
    `update_by`   bigint(20)            DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_email` (`email`),
    KEY `idx_username` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 模型表
DROP TABLE IF EXISTS `maas`;
CREATE TABLE `maas` (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) NOT NULL,
                         `description` varchar(255) DEFAULT NULL,
                         `status` ENUM('active', 'inactive', 'deprecated') DEFAULT 'active',
                         `model` varchar(255) NOT NULL,
                         `model_type` varchar(255) NOT NULL,
                         `base_url` varchar(255) NOT NULL,
                         `api_key` varchar(255) NOT NULL,
                         `avatar` varchar(255) DEFAULT NULL,
                         `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `create_by` bigint(20) NOT NULL DEFAULT '0',
                         `update_time` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
                         `update_by` bigint(20) DEFAULT NULL,
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE KEY `uk_model_base_url` (`model`, `base_url`),
                         KEY `idx_model` (`model`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;


-- 会话表
DROP TABLE IF EXISTS `conversations`;
CREATE TABLE conversations (
                               id BIGINT PRIMARY KEY AUTO_INCREMENT,
                               session_id VARCHAR(255) NOT NULL,  -- 【建议添加】
                               user_id BIGINT NOT NULL,
                               title VARCHAR(255) NOT NULL ,
                               status ENUM('active', 'archived', 'deleted') DEFAULT 'active', -- 【建议添加】
                               last_message_at TIMESTAMP NULL DEFAULT NULL,  -- 【建议添加】
                               total_messages INT DEFAULT 0,  -- 【建议添加】
                               create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               update_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              
                               UNIQUE KEY uk_session_id (session_id),  -- 【建议添加】
                               KEY `idx_user_id` (`user_id`),
                               KEY `idx_last_message` (`last_message_at`),
                               KEY `idx_status` (`status`),
                               CONSTRAINT `fk_conversations_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 会话-模型关联表
DROP TABLE IF EXISTS `conversation_maas`;
CREATE TABLE conversation_maas (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                    `conversation_session_id` varchar(64) NOT NULL,
                                    `maas_id` bigint(20) NOT NULL,
                                     status ENUM('active', 'inactive') DEFAULT 'active', -- 【建议添加】
                                     create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `uk_conv_maas` (`conversation_session_id`, `maas_id`),
                                    KEY `idx_conversation_session_id` (`conversation_session_id`),
                                    CONSTRAINT `fk_conversation_maas_session_id` FOREIGN KEY (`conversation_session_id`) REFERENCES `conversations` (`session_id`),
                                    CONSTRAINT `fk_conversation_maas_id` FOREIGN KEY (`maas_id`) REFERENCES `maas` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 消息表
DROP TABLE IF EXISTS `messages`;
CREATE TABLE messages (
                          `id` bigint(20) NOT NULL AUTO_INCREMENT,
                          `conversation_session_id` varchar(64) NOT NULL,
                          `message_type` ENUM('user', 'assistant') NOT NULL,
                          `content` text NOT NULL,
                          `context` longtext DEFAULT NULL,
                          `expand` longtext DEFAULT NULL,
                          `maas_id` bigint(20) DEFAULT NULL,
                          `sequence_number` int(11) NOT NULL,
                          `parent_question_id` bigint(20) DEFAULT NULL,
                          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,


                          PRIMARY KEY (`id`),
                          KEY `idx_conv_seq` (`conversation_session_id`, `sequence_number`),
                          KEY `idx_parent_question` (`parent_question_id`),
                          KEY `idx_conversation_session_id` (`conversation_session_id`),
                          CONSTRAINT `fk_messages_session_id` FOREIGN KEY (`conversation_session_id`) REFERENCES `conversations` (`session_id`),
                          CONSTRAINT `fk_messages_maas_id` FOREIGN KEY (`maas_id`) REFERENCES `maas` (`id`),
                          CONSTRAINT `fk_messages_parent_question_id` FOREIGN KEY (`parent_question_id`) REFERENCES `messages` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;