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

## Storing data with expiration

```bash
set mykey "Hello" EX 10
```

## Retrieving data expiration time

```bash
ttl mykey
```

## Storing data only if the key already exists

```bash
set mykey "Hello" XX
```

## Storing data only if the key does not exist

```bash
set mykey "Hello" NX
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

## Setting an expire time for a hash field

```bash
hexpire myhash 10 fields 1 field1
```

## Getting the expire time for a hash field

```bash
httl myhash fields 1 field1
```