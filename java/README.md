# redis-demo-java

## Launching Redis

```bash
docker run --name redis-demo --rm -p 6379:6379 -d redis
```

## Accessing Redis

```bash
docker exec -it redis-demo redis-cli
```

## Performances

This benchmark was done on:
```
Host: Windows Subsystem for Linux - Ubuntu-22.04 (2.3.26)
CPU: AMD Ryzen 7 2700X (16) @ 4.15 GHz
```

### Http without redis

```bash
redis-demo (main)> time curl "http://192.168.1.169:8080/calculation/45"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    6.79 secs      fish           external
   usr time    7.56 millis  146.00 micros    7.41 millis
   sys time    0.49 millis  495.00 micros    0.00 millis
```

### Http with redis (simple)

First calculation:
```bash
redis-demo (main)> time curl "http://192.168.1.169:8080/calculation/45?use_redis=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    6.73 secs      fish           external
   usr time    0.00 millis    0.00 micros    0.00 millis
   sys time    6.83 millis  548.00 micros    6.29 millis
```

Second calculation, the result is cached:
```bash
redis-demo (main)> time curl "http://192.168.1.169:8080/calculation/45?use_redis=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in   23.23 millis    fish           external
   usr time    0.10 millis  104.00 micros    0.00 millis
   sys time    5.98 millis  331.00 micros    5.65 millis
```

### Http with redis (hash)

First calculation:
```bash
redis-demo (main)> time curl "http://192.168.1.169:8080/calculation/45?use_redis=true&use_redis_hash=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    6.74 secs      fish           external
   usr time    5.32 millis  127.00 micros    5.20 millis
   sys time    0.39 millis  389.00 micros    0.00 millis
```

Second calculation, the result is cached:
```bash
redis-demo (main)> time curl "http://192.168.1.169:8080/calculation/45?use_redis=true&use_redis_hash=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    8.01 millis    fish           external
   usr time    5.22 millis    0.00 micros    5.22 millis
   sys time    0.55 millis  553.00 micros    0.00 millis
```