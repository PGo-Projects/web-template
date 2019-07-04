package security

import (
	"github.com/PGo-Projects/web-template/internal/config"
	"github.com/PGo-Projects/web-template/internal/database"
	"github.com/PGo-Projects/webauth"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func Setup(mux *chi.Mux, allowedOrigins []string) {
	// Setup security middleware
	mux.Use(webauth.ExpirationMiddleware)

	authenticationKey := viper.GetString(config.AuthenticationKey)
	encryptionKey := viper.GetString(config.EncryptionKey)
	csrfAuthenticationKey := viper.GetString(config.CSRFAuthenticationKey)

	webauth.RegisterDatabase(database.Client())

	webauth.SetupSessions([]byte(authenticationKey), []byte(encryptionKey))
	webauth.SessionOptions.Secure = config.ProdRun

	webauth.SetupCORS(mux, false, allowedOrigins)
	webauth.SetupCSRF(mux, []byte(csrfAuthenticationKey), config.ProdRun)
	webauth.RegisterEndPoints(mux)
}
