# webque
Go HTTP Queue/Worker Experiment


## What it can do

- Publish events using curl
- Subscribe events and execute jobs
- Jobs access external service by HTTP
- Jobs write data to DB
- Retry jobs if they fail with exp backoff


## Prep

```
-- for proxy service
CREATE USER webque_proxy;
CREATE DATABASE webque_proxy OWNER webque_proxy;

-- for backend service
CREATE USER webque_backend;
CREATE DATABASE webque_backend OWNER webque_backend;
```
