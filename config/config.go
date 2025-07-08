package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Conf struct {
    Server struct {
        Port int `env:"SERVER_PORT,default=8080"`
    }
    DB struct {
        Host string `env:"DB_HOST"`
        Port int `env:"DB_PORT,default=5432"`
        Name string `env:"DB_NAME,required"`
        User string `env:"DB_USER,required"`
        Pass string `env:"DB_PASS,required"`
    }
}

func New() *Conf {

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Cannot load .env", err)
    }

    var c Conf
    err := envdecode.Decode(&c)
    if err != nil {
        log.Fatalf("Could not pass env: %s", err)
    }
    return &c
}

