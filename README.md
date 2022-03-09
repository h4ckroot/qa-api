# ![Simple Question/Answer API](logo.png)

This repository contains a simple Rest API for a stackoverflow like API. It implements the following users stories:
- As a user, I should be able to submit questions.
- As a user, I should be able to answer submitted questions.
- As a user, I should be able to assign multiple tags to questions.

## Stack used:
  - Golang
  - Postgres
  - Gin
  - Gorm
  - Helm
  - Docker


## Deployment:
  - You can run the code localy using docker-compose by :

```
docker-compose up -d
```

  - You can install it in a kuberentes cluster using the helm chart under deploy/helm using the following:

```
cd helm/deploy/qa-api
helm --namespace <namespace> install test .
```


## Testing:
  - Please refer to test_samples for some curl command samples to test the basic functions.

## ToDo:
  - Add PgBouncer for connection pooling.
  - Add Makefile for basic task (build, tests, push image).
  - Add basic unit test.
  - Enhance validations and code structure.