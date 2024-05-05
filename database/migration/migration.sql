CREATE TABLE IF NOT EXISTS users (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(50),
    `password` VARCHAR(100),
    `email` VARCHAR(100),
    `mobile` VARCHAR(20),
    `age` INT,
    `address` VARCHAR(150),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS authors (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `author_name` VARCHAR(50),
    `nationality` VARCHAR(50),
    `biography` VARCHAR(1000)
);

CREATE TABLE IF NOT EXISTS books (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(100),
    `author_id` INT,
    `book_price_id` INT,
    `isbn` VARCHAR(13),
    `genre` VARCHAR(20),
    `publication_year` INT,
    `total_copies` INT,
    `available_copies` INT
);

CREATE TABLE IF NOT EXISTS book_pricing (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `book_id` INT,
    `price` FLOAT,
    `discount_percentage` FLOAT
);

CREATE TABLE memberships (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT,
	`membership_id` VARCHAR(30),
    `membership_type` INT
);