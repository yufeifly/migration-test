package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

const (
	src                 = "192.168.227.144"
	dst                 = "192.168.227.147"
	defaultProxyPort    = "6788"
	defaultMigratorPort = "6789"
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
func AccessRedis(service string, wg *sync.WaitGroup) {
	for i := 0; i < 300; i++ {
		data := make(map[string]string, 3)
		data["key"] = "name" + strconv.Itoa(i)
		data["value"] = "value" + strconv.Itoa(i)
		data["service"] = service
		ro := grequests.RequestOptions{
			Data: data,
		}
		url := "http://" + buildAddress(src, defaultProxyPort) + "/redis/set"
		resp, err := grequests.Post(url, &ro)
		if err != nil {
			logrus.Errorf("service: %v, AccessRedis.Post err: %v", service, err)
			time.Sleep(1 * time.Microsecond)
			continue
		}

		logrus.Infof("service: %v, i: %d, resp: %v", service, i, resp.String())
		time.Sleep(1 * time.Microsecond)
	}
	wg.Done()
}

// TriggerMigration ...
func TriggerMigration(opts MigOpts, wg *sync.WaitGroup) {
	data := make(map[string]string, 5)
	data["Service"] = opts.Service
	data["CheckpointID"] = opts.CheckpointID
	data["CheckpointDir"] = opts.CheckpointDir
	data["Src"] = opts.Src
	data["Dst"] = opts.Dst
	ro := grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + buildAddress(src, defaultProxyPort) + "/service/migrate"
	resp, err := grequests.Post(url, &ro)
	if err != nil {
		logrus.Errorf("service: %v, TriggerMigration err: %v", opts.Service, err)
	}
	logrus.Infof("service: %v, resp: %v", opts.Service, resp.RawResponse.Body)
	wg.Done()
}

func main() {
	// services
	services := []string{
		"service1",
		"service2",
	}
	serviceCount := 2 // number of serviceCount, or services
	wg := sync.WaitGroup{}
	// start accessing the redis service, imitate the real-world accesses
	for i := 0; i < serviceCount; i++ {
		wg.Add(1)
		go AccessRedis(services[i], &wg)
	}

	// sleep for a while, then migrate it
	time.Sleep(100 * time.Microsecond)
	// start migration, concurrently
	for i := 0; i < serviceCount; i++ {
		wg.Add(1)
		opts := MigOpts{
			Service:       services[i],
			CheckpointID:  "chkp-" + services[i],
			CheckpointDir: "/tmp",
			Src:           buildAddress(src, defaultMigratorPort),
			Dst:           buildAddress(dst, defaultMigratorPort),
		}
		go TriggerMigration(opts, &wg)
	}
	// wait serviceCount to finish
	wg.Wait()
}

func buildAddress(ip, port string) string {
	return fmt.Sprintf("%s:%s", ip, port)
}
