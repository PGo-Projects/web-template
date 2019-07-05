package security

import (
	"context"

	"github.com/PGo-Projects/web-template/internal/config"
	"github.com/PGo-Projects/web-template/internal/securitydb"
	"github.com/PGo-Projects/webauth"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func MustSetup(mux *chi.Mux) {
	// Setup security middleware
	mux.Use(webauth.ExpirationMiddleware)

	authenticationKey := viper.GetString(config.AuthenticationKey)
	encryptionKey := viper.GetString(config.EncryptionKey)
	csrfAuthenticationKey := viper.GetString(config.CSRFAuthenticationKey)

	mongoClient := securitydb.MustMongoClient(context.TODO(), "mongodb://localhost:27017")
	webauth.RegisterDatabase(mongoClient)

	webauth.SetupSessions([]byte(authenticationKey), []byte(encryptionKey))
	webauth.SessionOptions.Secure = config.ProdRun

	allowedOrigins := viper.GetStringSlice(config.AllowedOrigins)
	webauth.SetupCORS(mux, false, allowedOrigins)
	webauth.SetupCSRF(mux, []byte(csrfAuthenticationKey), config.ProdRun)
	webauth.RegisterEndPoints(mux)
}
