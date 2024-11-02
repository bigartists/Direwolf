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
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 模型表
DROP TABLE IF EXISTS `model`;
CREATE TABLE `model` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



-- 会话表
DROP TABLE IF EXISTS `conversations`;
CREATE TABLE conversations (
                               id BIGINT PRIMARY KEY AUTO_INCREMENT,
                               user_id BIGINT NOT NULL,
                               title VARCHAR(255),
                               status ENUM('active', 'archived', 'deleted') DEFAULT 'active', -- 【建议添加】
                               last_message_at TIMESTAMP,  -- 【建议添加】
                               total_messages INT DEFAULT 0,  -- 【建议添加】
                               create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               FOREIGN KEY (user_id) REFERENCES user(id),
                               INDEX idx_user_id (user_id),
                               INDEX idx_last_message (last_message_at),  -- 【建议添加】
                               INDEX idx_status (status)  -- 【建议添加】
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 会话-模型关联表
DROP TABLE IF EXISTS `conversation_models`;
CREATE TABLE conversation_models (
                                     id BIGINT PRIMARY KEY AUTO_INCREMENT,
                                     conversation_id BIGINT NOT NULL,
                                     model_id BIGINT NOT NULL,
                                     status ENUM('active', 'inactive') DEFAULT 'active', -- 【建议添加】
                                     create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 【建议添加】
                                     FOREIGN KEY (conversation_id) REFERENCES conversations(id),
                                     FOREIGN KEY (model_id) REFERENCES model(id),
                                     UNIQUE KEY uk_conv_model (conversation_id, model_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 消息表
DROP TABLE IF EXISTS `messages`;
CREATE TABLE messages (
                          id BIGINT PRIMARY KEY AUTO_INCREMENT,
                          conversation_id BIGINT NOT NULL,
                          message_type ENUM('user', 'assistant') NOT NULL,
                          content TEXT NOT NULL,
                          model_id BIGINT,           -- 现有设计
                          sequence_number INT NOT NULL,
                          create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          parent_question_id BIGINT, -- 【建议添加】直接关联到对应的问题
                          FOREIGN KEY (conversation_id) REFERENCES conversations(id),
                          FOREIGN KEY (model_id) REFERENCES model(id),
                          FOREIGN KEY (parent_question_id) REFERENCES messages(id),
                          INDEX idx_conv_seq (conversation_id, sequence_number),
                          INDEX idx_created_at (create_time),
                          INDEX idx_parent_question (parent_question_id)  -- 【建议添加】
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;