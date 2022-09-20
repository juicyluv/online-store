create database online_shop;

\c online_shop

create schema web;

create table if not exists web.users
(
  id          bigserial primary key,
  tel         text        not null,
  password    text        not null,
  email       text,
  created_at  timestamptz not null default now(),
  modified_at timestamptz
);