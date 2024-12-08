import redis
from flask import Flask, request, jsonify

app = Flask(__name__)

# Connect to Redis
def connect_to_redis():
    return redis.StrictRedis(host='localhost', port=6379, decode_responses=True)

redis_client = connect_to_redis()

# Calculate Fibonacci recursively
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

# Retrieve or calculate Fibonacci value using simple Redis keys
def redis_get_simple(client, nb, with_ttl):
    key = f"fibonacci:{nb}"

    # Try to get the value from Redis
    value = client.get(key)
    if value is not None:
        return int(value)

    # Calculate and store the value in Redis
    result = fibonacci(nb)
    client.set(key, result, ex=15 if with_ttl else None)
    return result

# Retrieve or calculate Fibonacci value using Redis hashes
def redis_get_hash(client, nb, with_ttl):
    key = "fibonacci"
    field = str(nb)

    # Try to get the value from Redis
    value = client.hget(key, field)
    if value is not None:
        return int(value)

    # Calculate and store the value in Redis hash
    result = fibonacci(nb)
    client.hset(key, field, result)

    if with_ttl:
        client.hexpire(key, 15, field)

    return result

@app.route('/calculation/<int:nb>', methods=['GET'])
def calculation(nb):
    use_redis = request.args.get('use_redis', 'false').lower() == 'true'
    use_redis_hash = request.args.get('use_redis_hash', 'false').lower() == 'true'
    with_ttl = request.args.get('with_ttl', 'false').lower() == 'true'

    if nb < 0:
        return jsonify(error="Invalid number"), 400

    try:
        if use_redis:
            if use_redis_hash:
                result = redis_get_hash(redis_client, nb, with_ttl)
            else:
                result = redis_get_simple(redis_client, nb, with_ttl)
        else:
            result = fibonacci(nb)

        return jsonify(f"Fibonacci({nb}) = {result}")
    except Exception as e:
        return jsonify(error=str(e)), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
