package config

type Server struct {
	Name string `validate:"required"`
	Hostname string `validate:"required"`
	Username string `validate:"required"`
	Password string
	PrivateKey string
	NoSudo bool
}

type Config struct {
	Servers []Server
}