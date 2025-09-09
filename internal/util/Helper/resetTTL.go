package util

import (
	"time"

	storage "github.com/SLANGERES/Service-Discovery/internal/Storage"
)

func ResetTTL(id string) {
	
	s := storage.Storages[id]
	s.TTl = 180
	s.Expires = time.Now().Add(time.Duration(180) * time.Second)
	storage.Storages[id] = s
}
