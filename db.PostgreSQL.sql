--CREATE DATABASE vrach;

--USE vrach;

/*CREATE TABLE customer (
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(100) NOT NULL,
  lastname VARCHAR(100) NOT NULL,
  data_of_birth VARCHAR(100) NOT NULL,
  email	VARCHAR(100) NOT NULL,
  gender INTEGER NOT NULL
);*/

--ALTER TABLE customer RENAME COLUMN data_of_birth TO date_of_birth;
--ALTER TABLE customer ALTER COLUMN token_created_at DROP DEFAULT;
-- ALTER TABLE customer add COLUMN nickname VARCHAR NULL DEFAULT "Customer";
/*ALTER TABLE customer
	ADD COLUMN username VARCHAR NOT NULL,
	ADD COLUMN password VARCHAR NOT NULL,
	ADD COLUMN account_create_at TIMESTAMP DEFAULT NOW(),
	ADD COLUMN account_update_at TIMESTAMP DEFAULT NOW(),
	ADD COLUMN token VARCHAR,
	ADD COLUMN token_created_at TIMESTAMP DEFAULT NOW(),
	ADD COLUMN token_expired_at TIMESTAMP,
	ADD COLUMN is_token_active INT;
	*/
SELECT * FROM customer ORDER BY id DESC;