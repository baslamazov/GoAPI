-- Создаем базу данных users  
CREATE DATABASE users;  

-- Подключаемся к базе данных users  
\c users;  

SET NAMES 'utf8';  

-- Удаляем таблицы, если они существуют  
DROP TABLE IF EXISTS items;  
DROP TABLE IF EXISTS users;  

-- Создаем таблицу items  
CREATE TABLE items (  
    id SERIAL PRIMARY KEY,  
    title VARCHAR(255) NOT NULL,  
    description TEXT NOT NULL,  
    updated VARCHAR(255) DEFAULT NULL  
);  

-- Вставляем данные в таблицу items  
INSERT INTO items (id, title, description, updated) VALUES  
    (1, 'database/sql', 'Рассказать про базы данных', 'rvasily'),  
    (2, 'memcache', 'Рассказать про мемкеш с примером использования', NULL);  

-- Создаем таблицу users  
CREATE TABLE users (  
    user_id SERIAL PRIMARY KEY,  
    login VARCHAR(255) NOT NULL,  
    password VARCHAR(255) NOT NULL,  
    email VARCHAR(255) NOT NULL,  
    info TEXT NOT NULL,  
    updated VARCHAR(255) DEFAULT NULL  
);  

-- Вставляем данные в таблицу users  
INSERT INTO users (user_id, login, password, email, info, updated) VALUES  
    (1, 'rvasily', 'love', 'rvasily@example.com', 'none', NULL);  