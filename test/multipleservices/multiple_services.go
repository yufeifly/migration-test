package multipleservices

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/validator/cuserr"
	"strconv"
	"sync"
	"time"
)

const (
	defaultProxyPort    = "6788"
	defaultMigratorPort = "6789"
)

var (
	src       = "192.168.227.144"
	dst       = "192.168.227.147"
	serverSrc = "192.168.1.207"
	serverDst = "192.168.1.207"
)

type TestOptions struct {
	Platform string
	Number   int
}

// migOpts ...
type migOpts struct {
	service       string
	checkpointID  string
	checkpointDir string
	srcAddr       string
	dstAddr       string
}

// AccessRedis ...
func accessRedis(service string, wg *sync.WaitGroup) {
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
func triggerMigration(opts migOpts, wg *sync.WaitGroup) {
	data := make(map[string]string, 5)
	data["Service"] = opts.service
	data["CheckpointID"] = opts.checkpointID
	data["CheckpointDir"] = opts.checkpointDir
	data["Src"] = opts.srcAddr
	data["Dst"] = opts.dstAddr
	ro := grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + buildAddress(src, defaultProxyPort) + "/service/migrate"
	resp, err := grequests.Post(url, &ro)
	if err != nil {
		logrus.Errorf("service: %v, TriggerMigration err: %v", opts.service, err)
	}
	logrus.Infof("service: %v, resp: %v", opts.service, resp.RawResponse.Body)
	wg.Done()
}

// Multiple test migrating multiple services
func TestMultipleService(opts TestOptions) error {
	switch opts.Platform {
	case "server":
		src = serverSrc
		dst = serverDst
	case "pc":
	default:
		return cuserr.ErrBadParams
	}
	serviceCount := opts.Number // number of serviceCount, or services
	// todo check if these services do exist
	// services
	services := make([]string, serviceCount)
	for i := 0; i < serviceCount; i++ {
		services[i] = "service" + strconv.Itoa(i+1) // add service name to services
	}
	logrus.Infof("services: %v", services)

	wg := sync.WaitGroup{}
	// start accessing the redis service, imitate the real-world accesses
	for i := 0; i < serviceCount; i++ {
		wg.Add(1)
		go accessRedis(services[i], &wg)
	}

	// sleep for a while, then migrate it
	time.Sleep(100 * time.Microsecond)
	// start migration, concurrently
	for i := 0; i < serviceCount; i++ {
		wg.Add(1)
		opts := migOpts{
			service:       services[i],
			checkpointID:  "chkp-" + services[i],
			checkpointDir: "/tmp",
			srcAddr:       buildAddress(src, defaultMigratorPort),
			dstAddr:       buildAddress(dst, defaultMigratorPort),
		}
		go triggerMigration(opts, &wg)
	}
	// wait serviceCount to finish
	wg.Wait()
	return nil
}

func buildAddress(ip, port string) string {
	return fmt.Sprintf("%s:%s", ip, port)
}
