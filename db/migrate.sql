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
  +-----------------------------------+
  | goals                             |
  +-----------------------------------+
  | id                      (int)     |
  | user_id                 (int)     |
  | name                    (varchar) |
  | description             (text)    |
  | color                   (varchar) |
  | update_interval         (string)  |
  | accumulator_key         (varchar) | *
  | accumulator_increment   (text)    | *
  | accumulator_description (text)    | *
  +-----------------------------------+
  +-----------------------------------+
  | streaks                           |
  +-----------------------------------+
  | id                      (int)     |
  | date_start              (date)    |
  | date_end                (date)    |
  | goal_id                 (int)     |
  +-----------------------------------+
  * think money saved not buying cigarettes  
*/


use streaking;


DROP TABLE IF EXISTS streaks;
DROP TABLE IF EXISTS users_goals;
DROP TABLE IF EXISTS goals;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(191),
  email VARCHAR(191),

  PRIMARY KEY (id),

  UNIQUE KEY (email)
);


CREATE TABLE goals (
  id BIGINT NOT NULL AUTO_INCREMENT,
  user_id BIGINT,
  name VARCHAR(191),
  description text,
  color VARCHAR(191),
  update_interval VARCHAR(191),
  accumulator_key VARCHAR(191),
  accumulator_increment text,
  accumulator_description text,

  PRIMARY KEY (id),

  UNIQUE KEY (user_id, name),

  FOREIGN KEY (user_id)
  REFERENCES users(id)
);


CREATE TABLE streaks (
  id BIGINT NOT NULL AUTO_INCREMENT,
  goal_id BIGINT,
  date_start DATE,
  date_end DATE,
  
  PRIMARY KEY (id),

  UNIQUE KEY (goal_id, date_start),

  FOREIGN KEY (goal_id)
  REFERENCES goals(id)
);