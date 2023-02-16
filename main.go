package main

import (
	"fmt"
	"os"
	"strconv"
	"tcp-monitor/app"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println("port is not valid")
		os.Exit(1)
	}

	app.Bootstrap(port)
}
