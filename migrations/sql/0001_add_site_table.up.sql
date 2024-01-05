CREATE TABLE sites (
    id  INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    access_time INT,
    state VARCHAR(255),
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at TIMESTAMP
);
