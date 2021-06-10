CREATE TABLE IF NOT EXISTS nodelab.account
(
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '이메일',
    `password` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '패스워드',
    `provider`      varchar(20) COMMENT '프로바이더',
    `provider_id`   varchar(50) COMMENT '프로바이더 id',
    `access_token`  varchar(200) COMMENT '액세스 토큰',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='로그인 정보';

CREATE TABLE IF NOT EXISTS nodelab.user (
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '이메일',
    `username` varchar(20) COLLATE utf8mb4_unicode_ci COMMENT '유저닉네임',
    `profile_image_id` int DEFAULT NULL COMMENT '프로필 이미지',
    `intro` text COLLATE utf8mb4_unicode_ci COMMENT '소개글',
    `github_url` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Github URL',
    `position` ENUM('FRONTEND', 'BACKEND', 'FULLSTACK', 'DEVOPS', 'MOBILE', 'MANAGER', 'AI', 'DATA', 'DATABASE', 'DESIGNER') COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '포지션',
    `interest` ENUM('FRONTEND', 'BACKEND', 'FULLSTACK', 'DEVOPS', 'MOBILE', 'MANAGER', 'AI', 'DATA', 'DATABASE', 'DESIGNER') COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '관심사',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`email`) REFERENCES account(`email`) ON DELETE CASCADE,
    UNIQUE KEY `email_UNIQUE` (`email`),
    UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='회원';

CREATE TABLE IF NOT EXISTS nodelab.study (
   `id` int NOT NULL AUTO_INCREMENT,
   `created_at` datetime DEFAULT NULL,
   `updated_at` datetime DEFAULT NULL,
   `deleted_at` datetime DEFAULT NULL,
   `studyname` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '스터디 모임 명',
   `title` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '스터디 제목',
   `content` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '스터디 설명',
   `summary` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '스터디 요약',
   `notice` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '스터디 공지사항',
   `start_date` datetime NOT NULL ,
   `finish_date` datetime NOT NULL,
   `limit` int NOT NULL,
   `thumbnail` varchar(500) NOT NULL,
   `status` ENUM('OPEN', 'PROGRESS', 'CLOSED') COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '포지션',
   `leader_id` int NOT NULL,
   PRIMARY KEY (`id`),
   FOREIGN KEY (`leader_id`) REFERENCES user(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='스터디';
