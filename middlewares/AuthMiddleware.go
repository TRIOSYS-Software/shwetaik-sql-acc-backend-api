package middlewares

import (
	"log"
	"net/http"
	"shwetaik-sql-acc-backend-api/config"
	"shwetaik-sql-acc-backend-api/utilities"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Printf("Error getting config: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		presharedKey := c.Request().Header.Get("ShweTaik")
		if presharedKey == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		decryptedKey, err := utilities.Decrypt(presharedKey, cfg.DefinedKey)
		if err != nil {
			log.Printf("Error decrypting preshared key: %v", err)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		if decryptedKey != cfg.DefinedPreShareKey {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}
