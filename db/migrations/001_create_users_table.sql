-- Create Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);