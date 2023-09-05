# drop database todo_app;
#
# CREATE DATABASE todo_app;

#USE todo_app;

CREATE TABLE IF NOT EXISTS todo_app.todo_task (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    due_date    DATETIME NULL,
    created_at  DATETIME NULL ON UPDATE CURRENT_TIMESTAMP,
    completed   TINYINT(1) NOT NULL
    );
