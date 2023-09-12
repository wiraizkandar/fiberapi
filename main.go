package main

import (
	"fiberapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	godotenv.Load()
	// Register Routes
	routes.Register(app)

	// Start the Fiber app with TLS

	if os.Getenv("SECURE") == "true" {
		log.Println("Secure")
		// Path to your SSL certificate and private key files
		certFile := "certs/fiberapi.test.pem"
		keyFile := "certs/fiberapi.test-key.pem"
		err := app.ListenTLS(":9843", certFile, keyFile)
		if err != nil {
			log.Fatalf("Error starting Fiber with TLS: %v", err)
		}

	} else {
		app.Listen(":3000")
		log.Fatalf("Server Localhost with port 3000")
	}

}
