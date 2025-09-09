package models

import "time"

type Service struct {
	Name string `json:"name" validate:"required"`
	Host string `json:"host" validate:"required"`
	Port int    `json:"port" validate:"required"`
	TTl  int    `json:"ttl"`
	Expires time.Time `json:"-"`

}
