package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

type api config struct{
	DB *database.Queries
}