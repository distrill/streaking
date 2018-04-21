use streaking;


-- clean up
DELETE FROM streaks;
DELETE FROM users_goals;
DELETE FROM goals;
DELETE FROM users;


-- let there be insertions
INSERT INTO users VALUES
    (1, 'brent 01', 'bh.01@hhindustries.ca'),
    (2, 'brent 02', 'bh.02@hhindustries.ca'),
    (3, 'brent 03', 'bh.03@hhindustries.ca');

INSERT INTO goals VALUES
    (1, 'first goal', 'the first thing I want to get done'),
    (2, 'second goal', 'the second thing I want to get done'),
    (3, 'third goal', 'the third thing i want to get done');

INSERT INTO users_goals VALUES
    (1, 1, 1), -- user 01, goal 01
    (2, 1, 2), -- user 01, goal 02
    (3, 2, 1), -- user 02, goal 01
    (4, 2, 2), -- user 02, goal 02
    (5, 2, 3), -- user 02, goal 03
    (6, 3, 1); -- user 03, goal 01


INSERT INTO streaks VALUES
    (1, 'user 01 goal 01', '200', 'this is how much i would have spent on cigarettes', '2018-04-01', '2018-04-13', 1, 1),
    (2, 'user 02 goal 01', '300', 'this is how much i would have spent on cigarettes', '2018-03-01', '2018-04-20', 2, 1),
    (3, 'user 02 goal 02', '300', 'this is how much i would have spent on booze', '2018-02-01', '2018-03-20', 2, 2),
    (4, 'user 02 goal 03', '500', 'this is how much weight i have lost', '2018-01-01', '2018-04-11', 2, 3);