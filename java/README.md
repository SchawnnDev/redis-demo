# redis-demo-java

## Setting up the project

```bash
mvn clean package
```

## Launching the app

```bash
java -jar target/redis-demo-1.0-SNAPSHOT-jar-with-dependencies.jar
```

## Performances

This benchmark was done on:
```
Host: Windows Subsystem for Linux - Ubuntu-22.04 (2.3.26)
CPU: AMD Ryzen 7 2700X (16) @ 4.15 GHz
```

### Http without redis

```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45"
"Fibonacci(45) = 1134903170"

________________________________________________________
Executed in  418.78 secs      fish           external
   usr time    0.28 millis  285.00 micros    0.00 millis
   sys time   23.80 millis  474.00 micros   23.33 millis
```

### Http with redis (simple)

Cached result:
```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45?use_redis=true"
"Fibonacci(45) = 1134903170"

________________________________________________________
Executed in   22.69 millis    fish           external
   usr time    9.03 millis  216.00 micros    8.82 millis
   sys time    0.32 millis  324.00 micros    0.00 millis
```

### Http with redis (hash)

Cached result:
```bash
redis-demo (main)> time curl "http://localhost:8080/calculation/45?use_redis=true&use_redis_hash=true"
"Fibonacci(45) = 1134903170"

________________________________________________________
Executed in   13.03 millis    fish           external
   usr time    0.31 millis  311.00 micros    0.00 millis
   sys time    9.67 millis  466.00 micros    9.21 millis
```