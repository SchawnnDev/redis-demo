# Using redis-cli to interact with Redis

## Launching Redis

```bash
docker run --name redis-demo --rm -p 6379:6379 -d redis
```

## Accessing Redis

```bash
docker exec -it redis-demo redis-cli
```

## Storing data

```bash
set mykey "Hello"
```

## Retrieving data

```bash
get mykey
```

## Storing a list

```bash
rpush mylist "Hello"
rpush mylist "World"
```

## Retrieving a list

```bash
lrange mylist 0 -1
```

## Storing a hash

```bash
hmset myhash field1 "Hello" field2 "World"
```

## Retrieving a hash

```bash
hgetall myhash
```

## Getting a value from a hash

```bash
hget myhash field1
```