package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load()

	connString := os.Getenv("PGX_URL")
	m, err := migrate.New("file://internal/database/migrations", connString)
	if err != nil {
		log.Fatalln("ERROR: ", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World!")
}
