-- for backend service
-- CREATE USER webque_backend;
-- CREATE DATABASE webque_backend OWNER webque_backend;

CREATE TABLE que_jobs (
  priority    smallint    NOT NULL DEFAULT 100,
  run_at      timestamptz NOT NULL DEFAULT now(),
  job_id      bigserial   NOT NULL,
  job_class   text        NOT NULL,
  args        json        NOT NULL DEFAULT '[]'::json,
  error_count integer     NOT NULL DEFAULT 0,
  last_error  text,
  queue       text        NOT NULL DEFAULT '',

  CONSTRAINT que_jobs_pkey PRIMARY KEY (queue, priority, run_at, job_id)
);

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
