CREATE DATABASE IF NOT EXISTS streaking;
DROP USER IF EXISTS streaking;
CREATE USER 'streaking'@'%' IDENTIFIED BY 'streaking';
GRANT ALL ON `streaking`.* TO 'streaking'@'%' IDENTIFIED BY 'streaking';
