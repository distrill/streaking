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
  +---------------+
  | users_goals   |
  +---------------|
  | id      (int) |
  | user_id (int) |
  | goal_id (int) |
  +---------------+
  +--------------------------------+
  | streaks                        |
  +--------------------------------+
  | id                (int)        |
  | accumulator_key   (varchar)    | *
  | accumulator_value (text)       | *
  | accumulator_description (text) | *
  | date_start        (date)       |
  | date_end          (date)       |
  | user_goal_id      (int)        |
  +--------------------------------+
  * think money saved not buying cigarettes  
*/


use streaking;


DROP TABLE IF EXISTS streaks;
DROP TABLE IF EXISTS users_goals;
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


CREATE TABLE users_goals (
  id BIGINT NOT NULL AUTO_INCREMENT,
  user_id BIGINT,
  goal_id BIGINT,

  PRIMARY KEY (id),

  FOREIGN KEY (user_id)
  REFERENCES users(id),
  FOREIGN KEY (goal_id)
  REFERENCES goals(id)
);


CREATE TABLE streaks (
  id BIGINT NOT NULL AUTO_INCREMENT,
  accumulator_key VARCHAR(255),
  accumulator_value text,
  accumulator_description text,
  date_start DATE,
  date_end DATE,
  user_goal_id BIGINT,
  
  PRIMARY KEY (id),

  FOREIGN KEY (user_goal_id)
  REFERENCES users_goals(id)
)