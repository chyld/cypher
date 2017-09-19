# psql -h 0.0.0.0 -U postgres -d temp -f schema.sql

CREATE TABLE logins (
  id          BIGSERIAL PRIMARY KEY NOT NULL,
  email       VARCHAR(255),
  username    VARCHAR(255),
  password    VARCHAR(255),
  pin         VARCHAR(255),
  site        VARCHAR(255),
  meta        TEXT,
  created_at  TIMESTAMP
);
