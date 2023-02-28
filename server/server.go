package server

import "os"

func Init() {
	PORT := os.Getenv("SERVER_PORT")
	HOST := "0.0.0.0"

	r := NewRouter()
	r.Run(HOST + ":" + PORT)
}
