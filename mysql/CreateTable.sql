CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL,
	`nickname` VARCHAR(64) NULL DEFAULT NULL,
    `password` VARCHAR(64) NULL DEFAULT NULL,
	`coins` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`uid`),
	UNIQUE KEY (`username`)
);