package util

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	storage "github.com/SLANGERES/Service-Discovery/internal/Storage"
)

// check service health
func check(url string) bool {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func HeartBeat() {
	s := storage.Storages 

// run in background
	go func() {
		for {
			for id, serv := range s {
				// Build URL properly
				url := "http://" + serv.Host + ":" + strconv.Itoa(serv.Port) + "/"

				healthy := check(url)
				if healthy {
					key := serv.Name + fmt.Sprintf(":%d", serv.Port)
					// ⚠️ if ResetTTL is in util package, just call ResetTTL(key)
					// util.ResetTTL(key) will cause import cycle
					ResetTTL(key)
					fmt.Printf("✅ Service %s is healthy (%s)\n", id, url)
				} else {
					fmt.Printf("❌ Service %s at %s is DOWN\n", id, url)
					// optional: remove from storage or mark unhealthy
				}
			}
			time.Sleep(10 * time.Second) // heartbeat interval
		}
	}()
}
