# redis-demo

This is a simple project to demonstrate how to use Redis.

Four different ways to store data are shown:
- Using redis-cli to interact with Redis
- A simple Go http server that computes Fibonacci numbers and caches the results in Redis
- A simple Java http server that computes Fibonacci numbers and caches the results in Redis
- A simple Python http server that computes Fibonacci numbers and caches the results in Redis

## Setup

### Launching Redis

```bash
docker run --name redis-demo --rm -p 6379:6379 -d redis
```

### Accessing Redis

```bash
docker exec -it redis-demo redis-cli
```
