CREATE DATABASE GOFEEDS;
GO

USE GOFEEDS;
GO

CREATE TABLE USERS (
    ID VARCHAR(36) PRIMARY KEY,
    hashed_password VARCHAR(100),
    salt VARCHAR(16),
    first_name VARCHAR(32),
    last_name VARCHAR(32),
    user_name VARCHAR(64) UNIQUE
);

CREATE TABLE USER_USER (
    fk_user_id VARCHAR(36),
    fk_follower_id VARCHAR(36),
    PRIMARY KEY (fk_user_id, fk_follower_id),
    CONSTRAINT fk_user FOREIGN KEY (fk_user_id) REFERENCES USERS(ID),
    CONSTRAINT fk_follower FOREIGN KEY (fk_follower_id) REFERENCES USERS(ID)
);

CREATE TABLE POSTS (
    ID VARCHAR(36) PRIMARY KEY,
    fk_user_id VARCHAR(36),
    content_text VARCHAR(512),
    content_image_path VARCHAR(100),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_post_user FOREIGN KEY (fk_user_id) REFERENCES USERS(ID)
);

CREATE TABLE COMMENTS (
    ID VARCHAR(36) PRIMARY KEY,
    fk_post_id VARCHAR(36),
    fk_user_id VARCHAR(36),
    content VARCHAR(150),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_comment_post FOREIGN KEY (fk_post_id) REFERENCES POSTS(ID),
    CONSTRAINT fk_comment_user FOREIGN KEY (fk_user_id) REFERENCES USERS(ID)
);

CREATE TABLE LIKES (
    fk_post_id VARCHAR(36),
    fk_user_id VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (fk_post_id, fk_user_id),
    CONSTRAINT fk_like_post FOREIGN KEY (fk_post_id) REFERENCES POSTS(ID),
    CONSTRAINT fk_like_user FOREIGN KEY (fk_user_id) REFERENCES USERS(ID)
);
