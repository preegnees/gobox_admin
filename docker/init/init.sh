#!/bin/bash

# set -e
psql -U "$POSTGRES_USER" <<-EOSQL
	CREATE USER docker SUPERUSER;
	CREATE DATABASE docker;
  ALTER ROLE docker WITH PASSWORD 'docker';
	GRANT ALL ON schema public TO docker;
EOSQL

psql -U "docker" <<-EOSQL
create table users (
    user_id int generated always as identity not null,
    username text unique not null,
    password_hash text unique not null,
    user_role varchar(1) default 'u',
    email text unique not null constraint CHK_users_user_role check (user_role in ('r', 'a', 'u')),
    constraint PK_users_user_id primary key(user_id)
);

create table tokens (
  token_id int generated always as identity not null,
  real_token int generated always as identity not null,
  temp_token text unique not null,
  user_id int not null,
  # deleted boolean default 'false',

  constraint FK_tokens_users foreign key (user_id) references users(user_id),
  constraint PK_tokens_token_id primary key (token_id)
);

insert into users (username, password_hash, user_role, email)
values ('roman', 'meiogr34555', 'a', 'roman@mail.ru');
insert into users (username, password_hash, user_role, email)
values ('ivan', 'fmrthoewpr45', 'u', 'ivan@gmail.com');
insert into users (username, password_hash, user_role, email)
values ('radmir', 'etun3v8theo_', 'r', 'radmir@gmail.com');
insert into users (username, password_hash, user_role, email)
values ('alesya', 'w43435435', 'u', 'alesya@mail.nl');

insert into tokens (temp_token, user_id)
values ('token1hello', 1);
insert into tokens (temp_token, user_id)
values ('123token2geeeo', 1);
insert into tokens (temp_token, user_id)
values ('ekrjt3556', 1);

insert into tokens (temp_token, user_id)
values ('ny8c3g4yt348t435', 2);

insert into tokens (temp_token, user_id)
values ('ociun4t8348ty34t4', 3);
insert into tokens (temp_token, user_id)
values ('v;inuy4by59mmxi4tn43e5yb6', 3);
EOSQL



