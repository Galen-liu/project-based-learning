package redis_service

import "github.com/gomodule/redigo/redis"

const network = "tcp"
const address = ":6379"

func AddShortenUrlMap(key, value string) error {
	c, err := redis.Dial(network, address)
	if err != nil {
		return err
	}
	defer c.Close()

	_, redisErr := c.Do("setnx", key, value)
	return redisErr
}

func GetShortenUrlMap(key string) (value string, err error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return "", err
	}
	defer c.Close()

	value, err = redis.String(c.Do("get", key))
	if err != nil {
		return "", err
	}

	return value, nil
}


