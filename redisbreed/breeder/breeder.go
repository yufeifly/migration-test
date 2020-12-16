package breeder

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/validator/redisbreed/generator"
	"github.com/yufeifly/validator/utils"

	"strconv"
)

const (
	valueLength = 200
)

type BreedOpts struct {
	RedisServer string
	KeyName     string
	Range       string
}

// BreedRedis feed redis services on kv pairs
func BreedRedis(opts BreedOpts) error {
	// parse address
	addr, err := utils.ParseAddress(opts.RedisServer)
	if err != nil {
		return err
	}
	// parse range
	r, err := utils.ParseRange(opts.Range)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// get redis connection
	ctx := context.Background()
	redisCli := redisConn(utils.BuildAddress(addr.IP, addr.Port))
	_, err = redisCli.Ping(ctx).Result()
	if err != nil {
		logrus.Error(err)
		return err
	}

	// set kv pairs
	for i := r.Start; i < r.End; i++ {
		key := r.Name + strconv.Itoa(i)
		valStr := generator.RandStringBytesMaskImprSrc(valueLength)
		err := redisCli.Set(ctx, key, valStr, 0).Err()
		if err != nil {
			logrus.Errorf("BreedRedis set kv err: %v", err)
			return err
		}
	}
	return nil
}

// redisConn get redis connection
func redisConn(redisServer string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisServer,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
