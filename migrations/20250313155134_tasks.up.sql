-- Удаляем старую таблицу если есть
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;

-- Создаем таблицу users с корректными ограничениями
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(191) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    role VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Явно создаем пользователя по умолчанию
INSERT INTO users (id, email, password, name, role) 
VALUES (1, 'default@example.com', '$2a$10$xJwL5vY3Jg5H3oYw7UvX.e5L9z9rF8S5J8mzYc6W3bX1VZ2K3N4D', 'Default User', 'user')
ON CONFLICT (id) DO NOTHING;

-- Создаем таблицу tasks с явным именем ограничения
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    text TEXT,
    is_done BOOLEAN DEFAULT FALSE,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_users_tasks FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);