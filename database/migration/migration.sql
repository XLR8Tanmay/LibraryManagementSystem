CREATE TABLE IF NOT EXISTS users (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(50),
    `password` VARCHAR(100),
    `email` VARCHAR(100) UNIQUE,
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
    `isbn` VARCHAR(13),
    `genre` VARCHAR(20),
    `publication_year` INT,
    `total_copies` INT,
    `available_copies` INT,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

CREATE TABLE IF NOT EXISTS book_pricing (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `book_id` INT,
    `price` FLOAT,
    `discount_percentage` FLOAT,
    `status` TINYINT,
    FOREIGN KEY (book_id) REFERENCES books(id)
);

CREATE TABLE memberships (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT,
	`membership_id` VARCHAR(30) UNIQUE,
    `membership_type` INT,
    `membership_status` INT,
    `membership_expiry_date` DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);