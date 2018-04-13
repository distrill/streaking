use streaking;


-- clean up
DELETE FROM streaks;
DELETE FROM goals;
DELETE FROM users;


-- let there be insertions
INSERT INTO users VALUES (1, 'brent 01', 'bh@hhindustries.ca');

INSERT INTO goals VALUES (1, 'first goal', 'the first thing i want to get done');

INSERT INTO streaks VALUES (1, 'money saved', '200', 'this is how much i would have spent on cigarretes', '2018-04-01', '2018-04-13', 1, 1);