package service

import (
	"fmt"
	"time"

	storage "github.com/SLANGERES/Service-Discovery/internal/Storage"
	model "github.com/SLANGERES/Service-Discovery/internal/models"
	util "github.com/SLANGERES/Service-Discovery/internal/util/Helper"
	"github.com/robfig/cron/v3"
)

// Clean checks a single service and removes it if expired
func Clean(data model.Service) {
	if time.Since(data.Expires) > 180*time.Second {
		key := data.Name + fmt.Sprintf(":%d", data.Port)
		fmt.Printf("⏰ Expired! Cleaning up service %s\n", key)
		util.Cleanup(key)
	}
}

// Cleanup schedules a cron job to clean expired services
func Cleanup() {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("* * * * * *", func() {
		s := storage.Storages
		for _, data := range s {
			Clean(data)
		}
	})
	if err != nil {
		fmt.Println("❌ Error scheduling cleanup:", err)
		return
	}

	c.Start()
	fmt.Println("✅ Cleanup job started (runs every second)")
	
}
