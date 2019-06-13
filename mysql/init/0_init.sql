CREATE DATABASE test_db;

CREATE TABLE test_db.meeting(
id         int NOT NULL AUTO_INCREMENT PRIMARY KEY,
title      varchar(256),
created_at datetime,
updated_at datetime
)
;

INSERT INTO test_db.meeting (title, created_at) VALUES
('テスト１', NOW()),
('test2'  , NOW()),
('test3'  , NOW()),
('テスト４', NOW())
;

