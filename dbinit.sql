SET sql_safe_updates = FALSE;

USE defaultdb;
DROP DATABASE IF EXISTS bank CASCADE;
CREATE DATABASE IF NOT EXISTS bank;

USE bank;

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    balance INT8
);