CREATE DATABASE vrach;

USE vrach;

CREATE TABLE customer (
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(100) NOT NULL,
  lastname VARCHAR(100) NOT NULL,
  data_of_birth VARCHAR(100) NOT NULL,
  email	VARCHAR(100) NOT NULL,
  gender INTEGER NOT NULL
);

ALTER TABLE customer RENAME COLUMN data_of_birth TO date_of_birth;