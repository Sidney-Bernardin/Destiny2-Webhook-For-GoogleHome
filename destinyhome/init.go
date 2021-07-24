package destinyhome

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"root/destinyhome/bungie"
	"root/destinyhome/repository"
)

var (
	repo      repository.Repository
	bungieSrv bungie.Service
)

func init() {

	// Get the projectID.
	projectID, ok := os.LookupEnv("PROJECT_ID")
	if !ok {
		log.Fatalf("missing env-var: PROJECT_ID")
	}

	// Get the API key.
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalf("missing env-var: API_KEY")
	}

	// Init dependencies.
	initRepo(projectID)
	initBungie(apiKey)
}

func initRepo(projectID string) {

	// If running with go test, don't init repo.
	if !strings.HasSuffix(os.Args[0], ".test") {
		var err error
		repo, err = repository.NewRepository(projectID, time.Duration(9)*time.Second)
		if err != nil {
			log.Fatalf("cannot create repository: %v", err)
		}
	}
}

func initBungie(apiKey string) {
	var err error
	bungieSrv, err = bungie.NewService(&http.Client{}, apiKey)
	if err != nil {
		log.Fatalf("cannot create bungo service: %v", err)
	}
}
