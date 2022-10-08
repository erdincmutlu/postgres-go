CREATE DATABASE testdb;
\c testdb;

CREATE TABLE account (
  id serial primary key,
  email varchar(255) unique not null,
  token varchar(255) not null
);

CREATE TABLE team (
  id int unique not null,
  team_name varchar(255) not null,
  country varchar(255) not null,
  team_value int not null default 0,
  team_budget int not null default 0, 
  FOREIGN KEY (id) REFERENCES account(id)
);

INSERT INTO account VALUES (1, 'test@example.com', 'hashPassword');

INSERT INTO team VALUES (1, 'team name', 'UK', '2000', '500');
