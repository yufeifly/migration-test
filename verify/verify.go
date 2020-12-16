package verify

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/yufeifly/validator/utils"
)

// VerifyOptions ...
type CheckOptions struct {
	Addr  string
	Range string
}

type accessOpts struct {
	ip, port   string
	keyName    string
	start, end int
}

// CheckResult verify the result of migration
func CheckResult(opts CheckOptions) error {
	fmt.Printf("addr: %v, range: %v\n", opts.Addr, opts.Range)
	addr, err := utils.ParseAddress(opts.Addr)
	if err != nil {
		return err
	}
	r, err := utils.ParseRange(opts.Range)
	aOpts := accessOpts{
		ip:      addr.IP,
		port:    addr.Port,
		keyName: r.Name,
		start:   r.Start,
		end:     r.End,
	}
	accessRedis(aOpts)
	return nil
}

// accessRedis ...
func accessRedis(opts accessOpts) int {
	failedCnt := 0
	redisCli := redis.NewClient(&redis.Options{
		Addr:     utils.BuildAddress(opts.ip, opts.port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	for i := opts.start; i < opts.end; i++ {
		key := opts.keyName + strconv.Itoa(i)
		value, err := redisCli.Get(ctx, key).Result()
		if err == redis.Nil {
			failedCnt++
			continue
		}
		if err != nil {
			logrus.Errorf("get kv err: %v", err)
			continue
		}
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": value,
		}).Info("the kv pair")
	}
	return failedCnt
}
