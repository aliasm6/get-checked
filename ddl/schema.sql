DROP DATABASE IF EXISTS getchecked;
CREATE DATABASE getchecked;
\c getchecked;

DROP TABLE IF EXISTS testing_centers;
CREATE TABLE testing_centers (
  id serial primary key,
  center_name varchar(255) NOT NULL,
  address text NOT NULL,
  days_open text NOT NULL,
  time_open int,
  time_closed int,
  website text,
  need_appointment boolean
);

\i /Users/alias/work/src/github.com/aliasm6/get-checked/dml/test_center.sql;
