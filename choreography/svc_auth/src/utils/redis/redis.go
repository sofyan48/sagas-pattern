package redis

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

type RedisLib struct {
	Store redis.Conn
}

func RedisLibHandler() *RedisLib {
	return &RedisLib{
		Store: getConnection(),
	}
}

type RedisLibInterface interface {
	RowsCached(keys string, data []byte, ttl int) ([]byte, error)
	GetRowsCached(keys string) (string, error)
	GetStore() redis.Conn
}

// InitRedis func
// return: redis.Conn
func initRedis() (redis.Conn, error) {
	rdhost := os.Getenv("REDIS_HOST")
	rdport := os.Getenv("REDIS_PORT")
	rdpass := os.Getenv("REDIS_PASSWORD")

	connRedis, err := redis.Dial("tcp", fmt.Sprintf(
		"%s:%s", rdhost, rdport))
	if rdpass != "" {
		if _, err := connRedis.Do("AUTH", rdpass); err != nil {
			connRedis.Close()
			fmt.Println(fmt.Sprintf("failed authorization to redis : %v", err))
			return nil, err
		}
	}
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to connect to redis from environment: %v", err))
		fmt.Println("Trying Local Connection")
		connRedis, err := redis.Dial("tcp", fmt.Sprintf(
			rdhost+":"+rdport))
		return connRedis, err
	}
	return connRedis, nil
}

// GetConnection function
// return store
func getConnection() redis.Conn {
	var Store redis.Conn
	if Store == nil {
		Store, _ = initRedis()
	}
	return Store
}

// GetStore function
// return store
func (rds *RedisLib) GetStore() redis.Conn {
	return rds.Store
}

// RowsCached params
// @keys: string
// @data: []byte
// @ttl: int
// return []byte, error
func (rds *RedisLib) RowsCached(keys string, data []byte, ttl int) ([]byte, error) {
	_, err := redis.String(rds.Store.Do("SET", keys, data))
	if err != nil {
		return nil, err
	}
	redis.String(rds.Store.Do("EXPIRE", keys, ttl))
	return data, nil
}

// GetRowsCached params
// @keys: string
// return string, error
func (rds *RedisLib) GetRowsCached(keys string) (string, error) {
	value, err := redis.String(rds.Store.Do("GET", keys))
	if err != nil {
		return "", err
	}
	return value, nil
}
