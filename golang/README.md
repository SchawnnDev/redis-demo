# redis-demo-golang

## Setting up the app

```bash
go mod tidy
```

## Launching the app

```bash
go run main.go
```

## Performances

This benchmark was done on:
```
Host: Windows Subsystem for Linux - Ubuntu-22.04 (2.3.26)
CPU: AMD Ryzen 7 2700X (16) @ 4.15 GHz
```

### Performance Results

| **Scenario**                | **Total Execution Time**   | **Usr Time**              | **Sys Time**              |
|-----------------------------|----------------------------|---------------------------|---------------------------|
| **Http without Redis**       | 5.99 secs                 | 4.56 millis / 199.00 µs   | 0.76 millis / 758.00 µs   |
| **Http with Redis (simple)** |                            |                           |                           |
| - First calculation          | 6.04 secs                 | 5.63 millis / 642.00 µs   | 0.00 millis / 0.00 µs     |
| - Cached result              | 15.74 millis              | 0.15 millis / 150.00 µs   | 6.07 millis / 549.00 µs   |
| **Http with Redis (hash)**   |                            |                           |                           |
| - First calculation          | 6.03 secs                 | 6.87 millis / 670.00 µs   | 2.86 millis / 2.86 µs     |
| - Cached result              | 7.55 millis               | 5.68 millis / 573.00 µs   | 0.21 millis / 215.00 µs   |

### Http without redis

```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    5.99 secs      fish           external
   usr time    4.56 millis  199.00 micros    4.36 millis
   sys time    0.76 millis  758.00 micros    0.00 millis
```

### Http with redis (simple)

First calculation:
```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45?use_redis=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    6.04 secs      fish           external
   usr time    5.63 millis  642.00 micros    4.99 millis
   sys time    0.00 millis    0.00 micros    0.00 millis
```

Second calculation, the result is cached:
```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45?use_redis=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in   15.74 millis    fish           external
   usr time    0.15 millis  150.00 micros    0.00 millis
   sys time    6.07 millis  549.00 micros    5.52 millis
```

### Http with redis (hash)

First calculation:
```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45?use_redis=true&use_redis_hash=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    6.03 secs      fish           external
   usr time    6.87 millis    0.67 millis    6.20 millis
   sys time    2.86 millis    2.86 millis    0.00 millis
```

Second calculation, the result is cached:
```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45?use_redis=true&use_redis_hash=true"
Fibonacci(45) = 1134903170
________________________________________________________
Executed in    7.55 millis    fish           external
   usr time    5.68 millis  573.00 micros    5.10 millis
   sys time    0.21 millis  215.00 micros    0.00 millis
```
