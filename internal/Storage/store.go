package storage

import (
	"fmt"
	"strings"
	"sync"

	"github.com/SLANGERES/Service-Discovery/internal/models"
)

var (
	mu       sync.RWMutex
	storages map[string]models.Service
)	

func init() {
	storages = make(map[string]models.Service) // initialize map
}

// AddIntoStorage adds a service to the registry
func AddIntoStorage(data models.Service) {
	key := data.Name + fmt.Sprintf(":%d", data.Port)
	mu.Lock()
	storages[key] = data
	mu.Unlock()
	fmt.Println("Service Added Sucessfully",storages[key],"key ->",key)
}

// RemoveFromStorage removes a service from the registry
func RemoveFromStorage(data models.Service) {
	key := data.Name + fmt.Sprintf(":%d", data.Port)

	mu.Lock()
	delete(storages, key)
	mu.Unlock()
}

// GetAll returns a snapshot of all registered services
func GetAll() map[string]models.Service {
	mu.RLock()
	defer mu.RUnlock()

	// return a copy so callers don’t mutate the original map
	copy := make(map[string]models.Service, len(storages))
	for k, v := range storages {
		copy[k] = v
	}
	return copy
}
func GetServicesByName(name string) map[string]models.Service {
	mu.RLock()
	defer mu.RUnlock()

	result := make(map[string]models.Service)
	for k, v := range storages {
		// key is "name:port"
		if strings.HasPrefix(k, name+":") {
			result[k] = v
		}
	}
	return result
}
func UpdateTTL(data models.Service) {
    key := data.Name + fmt.Sprintf(":%d", data.Port)

    mu.Lock()
    defer mu.Unlock()

    if svc, ok := storages[key]; ok {
        svc.TTl = 60
        storages[key] = svc
    }
}
