package storage

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/SLANGERES/Service-Discovery/internal/models"
)

var (
	Mu       sync.RWMutex
	Storages map[string]models.Service
)	

func init() {
	Storages = make(map[string]models.Service) // initialize map
}

// AddIntoStorage adds a service to the registry
func AddIntoStorage(data models.Service) {
	key := data.Name + fmt.Sprintf(":%d", data.Port)
	Mu.Lock()
	Storages[key] = data
	Mu.Unlock()
	fmt.Println("Service Added Sucessfully",Storages[key],"key ->",key)
}

// RemoveFromStorage removes a service from the registry
func RemoveFromStorage(id string) {
	key := id

	Mu.Lock()
	delete(Storages, key)
	Mu.Unlock()
}

// GetAll returns a snapshot of all registered services
func GetAll() map[string]models.Service {
	Mu.RLock()
	defer Mu.RUnlock()

	// return a copy so callers don’t mutate the original map
	copy := make(map[string]models.Service, len(Storages))
	for k, v := range Storages {
		copy[k] = v
	}
	return copy
}
func GetServicesByName(name string) map[string]models.Service {
	Mu.RLock()
	defer Mu.RUnlock()

	result := make(map[string]models.Service)
	for k, v := range Storages {
		// key is "name:port"
		if strings.HasPrefix(k, name+":") {
			result[k] = v
		}
	}
	return result
}
func UpdateTTL(data models.Service) {
    key := data.Name + fmt.Sprintf(":%d", data.Port)

    Mu.Lock()
    defer Mu.Unlock()

    if svc, ok := Storages[key]; ok {
        svc.TTl = 180
		svc.Expires = time.Now().Add(time.Duration(svc.TTl) * time.Second)
        Storages[key] = svc
    }
}
