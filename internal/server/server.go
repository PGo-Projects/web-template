package server

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/web-template/internal/config"
	"github.com/PGo-Projects/web-template/internal/security"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/lpar/gzipped"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const SOCK = "/tmp/<INSERT SOCK NAME>.sock"

func isValidStaticAssetPath(path string) bool {
	webAssetsDirectory := viper.GetString(config.WebAssetsPathKey)
	resourcePath := filepath.Join(webAssetsDirectory, path)
	_, err := os.Stat(resourcePath)
	return !os.IsNotExist(err)
}

func serveStaticOrIndex(w http.ResponseWriter, r *http.Request) {
	potentialStaticAssetPath := strings.TrimLeft(r.URL.Path, "/")
	if r.URL.Path == "/" || !isValidStaticAssetPath(potentialStaticAssetPath) {
		r.URL.Path = "/index.html"
		w.Header().Set("X-CSRF-TOKEN", csrf.Token(r))
	}
	webAssetsDirectory := http.Dir(viper.GetString(config.WebAssetsPathKey))
	gzipped.FileServer(webAssetsDirectory).ServeHTTP(w, r)
}

func Run(cmd *cobra.Command, arg []string) {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	// TODO: Fill in all allowed origins
	allowedOrigins := []string{"http://localhost:8080"}
	security.Setup(mux, allowedOrigins)
	mux.MethodFunc(http.MethodGet, "/*", serveStaticOrIndex)

	if config.DevRun {
		output.Println("Attempting to run on localhost:8080...", output.BLUE)
		log.Fatal(http.ListenAndServe(":8080", mux))
	} else {
		os.Remove(SOCK)
		unixListener, err := net.Listen("unix", SOCK)
		if err != nil {
			output.Error(err)
		}
		os.Chmod(SOCK, 0666)
		defer unixListener.Close()

		output.Println("Attempting to run on unix socket...", output.BLUE)
		log.Fatal(http.Serve(unixListener, mux))
	}
}
