package main

import (
	"context"
	"fmt"
	"github.com/redis/rueidis"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cli, err := connectToRedis()
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	http.HandleFunc("/calculation/", func(w http.ResponseWriter, r *http.Request) {
		useRedis := r.URL.Query().Get("use_redis") == "true"
		useRedisHash := r.URL.Query().Get("use_redis_hash") == "true"
		nbStr := r.URL.Path[len("/calculation/"):]
		nb, err := strconv.ParseInt(nbStr, 10, 64)
		if err != nil || nb < 0 {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		var result int64

		if useRedis {

			if useRedisHash {
				result, err = redisGetHash(cli, nb)
			} else {
				result, err = redisGetSimple(cli, nb)
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			result = fibonacci(nb)
		}

		fmt.Fprintf(w, "Fibonacci(%d) = %d", nb, result)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// connectToRedis is a function that connects to Redis
func connectToRedis() (rueidis.Client, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// fibonacci is a recursive function that calculates the Fibonacci number of a given number
func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// redisGetSimple is a function that gets a value from Redis, if the value is not found, it calculates it and sets it in Redis
func redisGetSimple(cli rueidis.Client, nb int64) (int64, error) {

	// We build a unique key for the value
	key := fmt.Sprintf("fibonacci:%d", nb)

	// We get the value from Redis
	cmd := cli.B().Get().
		Key(key).
		Build()

	resp, err := cli.Do(context.Background(), cmd).AsInt64()
	if err != nil {

		// We check if the key is not found, if an error occurs we return it
		if !rueidis.IsRedisNil(err) {
			return 0, err
		}

		// Since the key is not found, we proceed to calculate the value
		result := fibonacci(nb)

		// We set the value in Redis, if an error occurs we print it, but we don't return it since we already have the value
		cmd = cli.B().Set().
			Key(key).
			Value(strconv.FormatInt(result, 10)).
			Build()

		setResp := cli.Do(context.Background(), cmd)
		if setResp.Error() != nil {
			fmt.Println(setResp.Error())
		}

		return result, nil
	}

	return resp, nil
}

// redisGetHash is a function that gets a value from a Redis hash
func redisGetHash(cli rueidis.Client, nb int64) (int64, error) {
	// We build a unique key for the value
	key := "fibonacci"
	hKey := strconv.FormatInt(nb, 10)

	// We get the value from Redis
	cmd := cli.B().Hget().
		Key(key).
		Field(hKey).
		Build()

	resp, err := cli.Do(context.Background(), cmd).AsInt64()
	if err != nil {

		// We check if the key is not found, if an error occurs we return it
		if !rueidis.IsRedisNil(err) {
			return 0, err
		}

		// Since the key is not found, we proceed to calculate the value
		result := fibonacci(nb)

		// We set the value in Redis, if an error occurs we print it, but we don't return it since we already have the value
		cmd = cli.B().Hset().
			Key(key).
			FieldValue().
			FieldValue(hKey, strconv.FormatInt(result, 10)).
			Build()

		setResp := cli.Do(context.Background(), cmd)
		if setResp.Error() != nil {
			fmt.Println(setResp.Error())
		}

		return result, nil
	}

	return resp, nil
}
