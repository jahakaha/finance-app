package v1

import (
	"encoding/json"
	"finance-app/internal/config"
	"net/http"

	"github.com/sirupsen/logrus"
)

// API for returning version
// When server starts, we set version and use it if necessary

// ServerVersion represents the server version
type ServerVersion struct {
	Version string `json:"version"`
}

// Marshalled JSON
var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}

}

func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version")
	}
}
