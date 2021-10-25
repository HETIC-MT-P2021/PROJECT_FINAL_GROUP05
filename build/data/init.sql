CREATE TABLE IF NOT EXISTS commands (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    command VARCHAR(255) NOT NULL,
    is_checked BOOLEAN NOT NULL
);