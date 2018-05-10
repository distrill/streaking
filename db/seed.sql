/*
  Streaking - productivity/etc streak tracking
  Brent Hamilton <hamilton.bh9@gmail.com>
*/


use streaking;


-- clean up
DELETE FROM streaks;
DELETE FROM goals;
DELETE FROM users;


-- let there be insertions
INSERT INTO users VALUES
    (1, 'brent 01', 'bh.01@hhindustries.ca', 'STREAKING', '12345'),
    (2, 'brent 02', 'bh.02@hhindustries.ca', 'STREAKING', '23456'),
    (3, 'brent 03', 'bh.03@hhindustries.ca', 'STREAKING', '34567');

INSERT INTO goals VALUES
    (1, 1, '01 first goal', 'the first thing 01 want to get done', 'teal', 'day', 'cigarette money', '100', 'how much i would have spent on cigarettes'),
    (2, 2, '02 first goal', 'the first thing 02 want to get done', 'indigo', 'day', 'cigarette money', '120', 'how much i would have spent on cigarettes'),
    (3, 2, '02 second goal', 'the second thing 02 want to get done', 'light-blue', 'day', 'booze money', '230', 'how much i would have spent on booze'),
    (4, 2, '02 third goal', 'the third thing 02 want to get done', 'teal', 'week', 'miles run', '5', "miles i've run"),
    (5, 3, '03 first goal', 'the first thing 03 want to get done', 'indigo', 'week', 'booze money', '150', 'how much i would have spent on booze'),
    (6, 3, '03 second goal', 'the second thing 02 want to get done', 'light-blue', 'week', 'miles run', '3', "miles i've run");

INSERT INTO streaks VALUES
    (1, 1, '2018-04-01', '2018-04-13'),
    (2, 2, '2018-03-01', '2018-04-28'),
    (3, 3, '2018-02-01', '2018-03-20'),
    (4, 4, '2017-12-01', '2017-12-20'),
    (5, 4, '2018-01-10', '2018-03-20'),
    (6, 5, '2018-01-01', '2018-04-11');
