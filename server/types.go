package server

import (
	"github.com/gorilla/mux"
	"github.com/srodrmendz/api-auth/service"
)

const bearerTokenHeaderLength = 2

type services struct {
	AuthService service.Service
}

type config struct {
	Version   string
	BuildDate string
	SecretKey string
}

type App struct {
	Services services
	Config   config
	Router   *mux.Router
}
