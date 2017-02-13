package env

import "gopkg.in/redis.v5"

type Env struct {
	DB *redis.Client
}
