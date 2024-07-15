package main

import (
	"log"

	"github.com/ssr0016/personal-finance/internal/config"
	"github.com/ssr0016/personal-finance/internal/database"
	"github.com/ssr0016/personal-finance/internal/model"
	"github.com/ssr0016/personal-finance/pkg/util"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseUrl)
	password, err := util.HashPassword("password")
	if err != nil {
		log.Fatalf("Error generating password: %v\n", err)
	}
	users := []model.User{
		{
			Username: "admin",
			Password: password,
		},
	}
	_, err = db.NamedExec(
		`INSERT INTO users(username, password)
		 VALUES (:username, :password)`,
		users,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v\n", err)
	}

	log.Printf("Successfully inserted users: %v\n", users)
}
