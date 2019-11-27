package storages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

// RedisClient redis client type
type RedisClient struct{ *redis.Client }

var once sync.Once
var redisClient *RedisClient

type data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func init() {

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Fatal(err)
	}

	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_DB_URL"),
			Password: os.Getenv("REDIS_DB_PASSWORD"),
			DB:       db,
		})

		redisClient = &RedisClient{client}
	})

	pong, err := redisClient.Ping().Result()

	fmt.Println(pong)

	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}
}

// SetKey set the key and value in the redis DB
func SetKey(w http.ResponseWriter, r *http.Request) {
	var params data

	_ = json.NewDecoder(r.Body).Decode(&params)

	// key, value, expiry time
	err := redisClient.Set(params.Key, params.Value, 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode("Successfully Saved!")
}

// GetKey return the value of the key from the redis DB
func GetKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	val, err := redisClient.Get(params["key"]).Result()

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(val)
}

// Get return all the keys in the redis DB
func Get(w http.ResponseWriter, r *http.Request) {
	val, err := redisClient.Keys("*").Result()

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(val)
}
