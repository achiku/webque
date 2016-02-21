-- for backend service
-- CREATE USER webque_backend;
-- CREATE DATABASE webque_backend OWNER webque_backend;

CREATE TABLE account (
  id          serial    PRIMARY KEY
  ,name       text      NOT NULL
  ,created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
  ,updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
);


CREATE TABLE current_deposit (
  account_id integer   NOT NULL UNIQUE
  ,amount     integer   NOT NULL default 0
  ,created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
  ,updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
);
