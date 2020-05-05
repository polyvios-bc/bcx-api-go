package configuration

import (
	"time"
)

type Env string

var (
	PROD    Env = "PROD"
	STAGING Env = "STAGING"
)

type Configuration struct {
	Host    string
	ApiKey  string
	Timeout time.Duration
	Env     Env
}
