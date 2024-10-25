package server

import (
	"aie/internal/logger"
	"aie/internal/providers"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func getPort() string {
	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "3001"
}

// configRuntime - This function will set the number of CPUs that can be
// executing simultaneously and returns the previous setting.
func configRuntime() {
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)
}

func Start(logLevel string) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	logger.InitLogger(logLevel)

	// Start the AIE backend service
	log.Info().Msg("Starting AIE backend service")

	router := gin.New()
	registerRoutes(router)
	configRuntime()

	quit := make(chan bool, 1)
	providers.InitProviders(quit)

	server := &http.Server{
		Addr:    ":" + getPort(),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Err(err).Msg("Server start failed.")
		panic(1)
	}

	quit <- true
	err = server.Close()

	if err != nil {
		log.Err(err).Msg("Server shutdown failed.")
	}
}
