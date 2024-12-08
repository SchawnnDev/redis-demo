package demo;

import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import redis.clients.jedis.Jedis;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

public class Main {

    public static void main(String[] args) throws IOException {
        // Create an HTTP server on port 8080
        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);

        // Add a context for the /calculation endpoint
        server.createContext("/calculation/", new FibonacciHandler());

        // Start the server
        server.setExecutor(null); // Use the default executor
        System.out.println("Server is running on http://localhost:8080");
        server.start();
    }

    static class FibonacciHandler implements HttpHandler {
        private final Jedis jedis;

        public FibonacciHandler() {
            // Connect to Redis
            this.jedis = new Jedis("localhost", 6379);
        }

        @Override
        public void handle(HttpExchange exchange) throws IOException {
            String response;
            try {
                // Parse query parameters and path
                String query = exchange.getRequestURI().getQuery();
                boolean useRedis = query != null && query.contains("use_redis=true");
                boolean useRedisHash = query != null && query.contains("use_redis_hash=true");
                boolean withTTL = query != null && query.contains("with_ttl=true");
                String path = exchange.getRequestURI().getPath();
                String nbStr = path.substring("/calculation/".length());
                long nb;

                try {
                    nb = Long.parseLong(nbStr);
                    if (nb < 0) {
                        throw new NumberFormatException("Negative number not allowed");
                    }
                } catch (NumberFormatException e) {
                    response = "Invalid number";
                    exchange.sendResponseHeaders(400, response.length());
                    try (OutputStream os = exchange.getResponseBody()) {
                        os.write(response.getBytes());
                    }
                    return;
                }

                long result;
                if (useRedis) {
                    if (useRedisHash) {
                        result = redisGetHash(nb, withTTL);
                    } else {
                        result = redisGetSimple(nb, withTTL);
                    }
                } else {
                    result = fibonacci(nb);
                }

                // Prepare the response
                response = "Fibonacci(" + nb + ") = " + result;
                exchange.sendResponseHeaders(200, response.length());
            } catch (Exception e) {
                // Handle unexpected errors
                response = "Internal Server Error: " + e.getMessage();
                exchange.sendResponseHeaders(500, response.length());
            }

            try (OutputStream os = exchange.getResponseBody()) {
                os.write(response.getBytes());
            }
        }

        private long fibonacci(long n) {
            if (n <= 1) {
                return n;
            }
            return fibonacci(n - 1) + fibonacci(n - 2);
        }

        private long redisGetSimple(long nb, boolean withTTL) {
            String key = "fibonacci:" + nb;

            // Check if the value exists in Redis
            String value = jedis.get(key);
            if (value != null) {
                return Long.parseLong(value);
            }

            // Calculate and store the value
            long result = fibonacci(nb);

            if (withTTL) {
                jedis.setex(key, 15, String.valueOf(result));
            } else {
                jedis.set(key, String.valueOf(result));
            }

            return result;
        }

        private long redisGetHash(long nb, boolean withTTL) {
            String key = "fibonacci";
            String field = String.valueOf(nb);

            // Check if the value exists in Redis
            String value = jedis.hget(key, field);
            if (value != null) {
                return Long.parseLong(value);
            }

            // Calculate and store the value
            long result = fibonacci(nb);
            String resStr = String.valueOf(result);
            jedis.hset(key, field, resStr);

            if (withTTL) {
                jedis.hexpire(key, 15, field);
            }

            return result;
        }
    }
}
