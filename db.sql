/*
  Streaking - productivity/etc streak tracking
  Brent Hamilton <hamilton.bh9@gmail.com>

  +-----------------+
  | users           |
  +-----------------|
  | id    (int)     |
  | name  (varchar) |
  | email (varchar) |
  +-----------------+
  +-----------------------+
  | goals                 |
  +-----------------------+
  | id          (int)     |
  | name        (varchar) |
  | description (text)    |
  +-----------------------+
  +-----------------------------+
  | streaks                     |
  +-----------------------------+
  | id                (int)     |
  | accumulator_key   (varchar) | *
  | accumulator_value (text)    | *
  | date_start        (date)    |
  | date_end          (date)    |
  | user_id           (int)     |
  | goal_id           (int)     |
  +-----------------------------+
  * think money saved not buying cigarettes  
*/


CREATE DATABASE IF NOT EXISTS streaking;
DROP USER IF EXISTS streaking;
CREATE USER 'streaking'@'%' IDENTIFIED BY 'streaking';
GRANT ALL ON `streaking`.* TO 'streaking'@'%' IDENTIFIED BY 'streaking';


use streaking;


DROP TABLE IF EXISTS streaks;
DROP TABLE IF EXISTS goals;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  email VARCHAR(255),

  PRIMARY KEY (id)
);


CREATE TABLE goals (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  description text,

  PRIMARY KEY (id)
);


CREATE TABLE streaks (
  id BIGINT NOT NULL AUTO_INCREMENT,
  accumulator_key VARCHAR(255),
  accumulator_value text,
  date_start DATE,
  date_end DATE,
  user_id BIGINT,
  goal_id BIGINT,
  
  PRIMARY KEY (id),

  FOREIGN KEY (user_id)
  REFERENCES users(id),
  FOREIGN KEY (goal_id)
  REFERENCES goals(id)
)