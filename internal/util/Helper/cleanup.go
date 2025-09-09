package util

import storage "github.com/SLANGERES/Service-Discovery/internal/Storage"

func Cleanup(id string) {
	storage.RemoveFromStorage(id)
}
