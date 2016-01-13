create keyspace test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

use test;

CREATE TABLE if not exists users (
  id int PRIMARY KEY,
  name text,
);

insert into users (id, name) values (1, 'Joe');
