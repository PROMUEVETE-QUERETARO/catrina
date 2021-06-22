package internal

import (
	"log"
	"net/http"
)

// StartServer run a proof server as defined in the configuration file.
func StartServer(config Config) {
	log.Printf("Listen server in http://localhost%v...", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, http.FileServer(http.Dir(config.BuildPath))))
}
