CREATE TABLE IF NOT EXISTS directory (
    id INT NOT NULL AUTO_INCREMENT,
    dname VARCHAR(255) NOT NULL,
    is_hide TINYINT NOT NULL,
    PRIMARY KEY (id)
);