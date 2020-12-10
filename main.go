package main

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

// MigOpts ...
type MigOpts struct {
	Service       string
	CheckpointID  string
	CheckpointDir string
	Src           string
	Dst           string
}

// AccessRedis ...
func AccessRedis(wg *sync.WaitGroup) {
	for i := 0; i < 300; i++ {
		data := make(map[string]string, 3)
		data["key"] = "name" + strconv.Itoa(i)
		data["value"] = "value" + strconv.Itoa(i)
		data["service"] = "service1"
		ro := grequests.RequestOptions{
			Data: data,
		}
		url := "http://127.0.0.1:6788/redis/set"
		resp, err := grequests.Post(url, &ro)
		if err != nil {
			logrus.Errorf("AccessRedis.Post err: %v", err)
			continue
		}

		logrus.Infof("i: %d, resp: %v", i, resp.String())
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

// TriggerMigration ...
func TriggerMigration(opts MigOpts) {
	data := make(map[string]string, 5)
	data["Service"] = opts.Service
	data["CheckpointID"] = opts.CheckpointID
	data["CheckpointDir"] = opts.CheckpointDir
	data["Src"] = opts.Src
	data["Dst"] = opts.Dst
	ro := grequests.RequestOptions{
		Data: data,
	}
	url := "http://127.0.0.1:6788/service/migrate"
	resp, err := grequests.Post(url, &ro)
	if err != nil {
		logrus.Errorf("TriggerMigration err: %v", err)
	}
	logrus.Infof("resp: %v", resp.RawResponse.Body)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// start accessing the redis service, imitate the real-world accesses
	go AccessRedis(&wg)

	// sleep for a while, then migrate it
	time.Sleep(2 * time.Second)

	opts := MigOpts{
		Service:       "service1",
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Src:           "192.168.227.144:6789",
		Dst:           "192.168.227.147:6789",
	}
	TriggerMigration(opts)
	wg.Wait()
}
