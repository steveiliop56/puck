package config

// Schema for each server
type Server struct {
	Name string `validate:"required"`
	Hostname string `validate:"required"`
	Username string `validate:"required"`
	Password string
	PrivateKey string
	NoSudo bool
}

// Main config schema
type Config struct {
	Servers []Server
}