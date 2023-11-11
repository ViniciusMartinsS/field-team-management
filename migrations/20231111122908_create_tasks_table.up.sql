CREATE TABLE IF NOT EXISTS tasks (
    id         INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    summary    VARCHAR(2500) NOT NULL,
    `date`     DATE,
    user_id    INT NOT NULL,
    deleted    BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);