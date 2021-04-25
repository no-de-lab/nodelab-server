CREATE TABLE IF NOT EXISTS nodelab.user (
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '이메일',
    `username` varchar(20) COLLATE utf8mb4_unicode_ci COMMENT '유저닉네임',
    `profile_image_id` int DEFAULT NULL COMMENT '프로필 이미지',
    `intro` text COLLATE utf8mb4_unicode_ci COMMENT '소개글',
    `github_url` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Github URL',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email_UNIQUE` (`email`),
    UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='회원';

CREATE TABLE IF NOT EXISTS nodelab.account
(
    `id` int NOT NULL AUTO_INCREMENT,
    `email`         varchar(40) NOT NULL COMMENT '유저이메일',
    `password` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '패스워드',
    `provider`      varchar(20) COMMENT '프로바이더',
    `provider_id`   varchar(50) COMMENT '프로바이더 id',
    `access_token`  varchar(50) COMMENT '액세스 토큰',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='로그인 정보';
