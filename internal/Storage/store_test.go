package storage

import (
	"fmt"
	"testing"

	"github.com/SLANGERES/Service-Discovery/internal/models"
)

// helper function to reset storages before each test
func resetStorage() {
	Mu.Lock()
	defer Mu.Unlock()
	Storages = make(map[string]models.Service)
}

func TestAddIntoStorage(t *testing.T) {
	resetStorage()

	service := models.Service{Name: "test-service", Host: "0.0.0.0", Port: 8080, TTl: 30}
	AddIntoStorage(service)

	key := fmt.Sprintf("%s:%d", service.Name, service.Port)

	Mu.RLock()
	got, exists := Storages[key]
	Mu.RUnlock()

	if !exists {
		t.Fatalf("expected service to exist in storage, but it doesn't")
	}

	if got != service {
		t.Errorf("expected %+v, got %+v", service, got)
	}
}

func TestRemoveFromStorage(t *testing.T) {
	resetStorage()

	service := models.Service{Name: "removable", Port: 9090, TTl: 30}
	AddIntoStorage(service)
	RemoveFromStorage(service.Host)

	key := fmt.Sprintf("%s:%d", service.Name, service.Port)

	Mu.RLock()
	_, exists := Storages[key]
	Mu.RUnlock()

	if exists {
		t.Errorf("expected service to be removed, but found in storage")
	}
}

func TestGetAll(t *testing.T) {
	resetStorage()

	s1 := models.Service{Name: "s1", Port: 1111, TTl: 30}
	s2 := models.Service{Name: "s2", Port: 2222, TTl: 30}
	AddIntoStorage(s1)
	AddIntoStorage(s2)

	all := GetAll()

	if len(all) != 2 {
		t.Errorf("expected 2 services, got %d", len(all))
	}

	// check copy safety (mutating the copy shouldn’t affect storage)
	all["s1:1111"] = models.Service{}
	real := GetAll()
	if real["s1:1111"] != s1 {
		t.Errorf("storage was mutated from outside!")
	}
}

func TestGetServicesByName(t *testing.T) {
	resetStorage()

	s1 := models.Service{Name: "my-service", Port: 8080, TTl: 30}
	s2 := models.Service{Name: "my-service", Port: 9090, TTl: 30}
	s3 := models.Service{Name: "other-service", Port: 7070, TTl: 30}

	AddIntoStorage(s1)
	AddIntoStorage(s2)
	AddIntoStorage(s3)

	services := GetServicesByName("my-service")

	if len(services) != 2 {
		t.Errorf("expected 2 services, got %d", len(services))
	}

	for k, v := range services {
		if v.Name != "my-service" {
			t.Errorf("unexpected service in result: key=%s val=%+v", k, v)
		}
	}
}

func TestUpdateTTL(t *testing.T) {
	resetStorage()

	svc := models.Service{Name: "ttl-service", Port: 3000, TTl: 10}
	AddIntoStorage(svc)

	UpdateTTL(svc)

	key := fmt.Sprintf("%s:%d", svc.Name, svc.Port)

	Mu.RLock()
	updated := Storages[key]
	Mu.RUnlock()

	if updated.TTl != 60 {
		t.Errorf("expected TTL=60, got %d", updated.TTl)
	}
}
