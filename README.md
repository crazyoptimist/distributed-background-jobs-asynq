# Distributed Background Jobs with Asynq

This applicaton demonstrates distributed background jobs handling with [asynq](https://github.com/hibiken/asynq).

Asynq is very impressive and easy to use. It ensures a task to be run at least once and only once.

## Instructions for Demo

```bash
make docker
make upd
make log
```

```bash
for i in {1..11}; do; echo $i:; http http://localhost:8080/sleep seconds:=5; done;
```

And watch logs to see what's happening. Individual worker containers run tasks concurrently. Concurrency in inside a worker has been disabled for demonstration purpose, which is configurable in the worker configuration code.

## MakeFile

build docker images
```bash
make docker
```

run with docker compose
```bash
make up
```

run in detached mode
```bash
make upd
```

docker compose log
```bash
make log
```

docker compose down
```bash
make down
```

run the test suite
```bash
make test
```