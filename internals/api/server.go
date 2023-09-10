package api

import (
	"os"
)

func Setup() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8888"
	}
}
