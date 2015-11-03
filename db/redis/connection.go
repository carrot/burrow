package db

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool          *redis.Pool
	redisServer   = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "", "")
	redisDb       = flag.Int("redisDb", 0, "[0-15]")
)

func newPool(server, password string, dbNum int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if dbNum > 0 {
				if _, err := c.Do("SELECT", dbNum); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Open() {
	flag.Parse()
	pool = newPool(*redisServer, *redisPassword, *redisDb)
}

func Close() {
	pool.Close()
}

func ConnCount() int {
	return pool.ActiveCount()
}

func Get() redis.Conn {
	return pool.Get()
}

func Flush() {
	conn := Get()
	defer conn.Close()
	conn.Do("FLUSHDB")
}
