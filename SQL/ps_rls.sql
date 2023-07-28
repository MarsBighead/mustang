-- CREATE role jerry login with password 'jerry';

CREATE TABLE customer(
    id BIGSERIAL constraint pk PRIMARY KEY,
    name char(64),
    age smallint,
    gender char(6),
    ts timestamp NOT null
);

INSERT INTO customer VALUES(1, 'jerry',28,'male',now());

GRANT SELECT,INSERT,UPDATE, DELETE on customer to jerry; --time 1

CREATE POLICY rls_gender_read on customer for SELECT to jerry 
using (gender in ('female'));

ALTER TABLE customer enable row level security;

INSERT INTO customer VALUES(2, 'robin',28,'female',now());


CREATE POLICY rls_gender_write on customer for INSERT to jerry 
WITH check (gender in ('male','female'));

-- User Jerry
INSERT INTO customer VALUES(3, 'master',28,'mtf',now());