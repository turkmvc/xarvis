package infrastructure

import (
	"log"
	"net/http"

	"github.com/alamin-mahamud/go-bootstrap/v1/infrastructure/routers"
)

func Run() {
	// settings.Init()
	r := routers.Init()
	log.Fatal(http.ListenAndServe(":8080", r))
}