CREATE TABLE nodelab.user
(
    `email`          varchar(40)     NOT NULL    COMMENT '이메일',
    `username`       varchar(20)     NOT NULL    COMMENT '유저이름',
    `password`       varchar(200)    NULL        COMMENT '비밀번호',
    `profile_image`  integer         NULL        COMMENT '프로필이미지',
    `intro`          text            NULL        COMMENT '소개',
    `github_url`     varchar(40)     NULL        COMMENT 'Github 주소',
    PRIMARY KEY (email)
);

ALTER TABLE nodelab.user COMMENT '회원';
