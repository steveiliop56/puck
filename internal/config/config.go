package config

type Server struct {
	Name string `validate:"required"`
	Hostname string `validate:"required"`
	Username string `validate:"required"`
	Password string
	PrivateKey string
}

type Config struct {
	Servers []Server
}