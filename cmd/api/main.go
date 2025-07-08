package main

import (
	"fmt"
	"log"

	"github.com/tbm5k/tss/api/router"
	"github.com/tbm5k/tss/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbStringf = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
    c := config.New();

    fmt.Println(c)
    dbString := fmt.Sprintf(dbStringf, c.DB.Host, c.DB.User, c.DB.Pass, c.DB.Name, c.DB.Port)

    db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
    if err != nil {
        log.Fatalf("Cannot connect to db: %v", err)
    }

    router.New(c.Server.Port, db)
}

