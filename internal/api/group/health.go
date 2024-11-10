package group

import (
	"aie/internal/providers"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func checkHealth() string {
	// Test the DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	err := providers.DB.Client().Ping(ctx, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("MongoDB ping failed")
	}

	cancel()

	log.Info().Msg("Group service health: LIVE")
	return "LIVE"
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": checkHealth(),
	})
}
