-- Создание таблицы humans
CREATE TABLE IF NOT EXISTS humans (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    dob VARCHAR(255) NOT NULL,
    serial_number VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    code_structure VARCHAR(255) NOT NULL
); 