package main

import (
	"go-ecommerce-backend-api/internal/routers"
	"log"
)

func main() {

	r :=routers.NewRouter()

  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  //For test github personal
  if err := r.Run(); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}