package app

import (
	"inv-id-oice/db"
	"log"
	"net/http"
	"time"
)

type PostHandler struct {
	debug    bool
	lastPath string
	start    time.Time
	liteDB   *db.LiteDB
}

func (ph *PostHandler) handlePost(w http.ResponseWriter, req *http.Request) error {
	ph.start = time.Now()
	remPath := ""
	ph.lastPath, remPath = getLastPathInUri(req.RequestURI)
	if ph.debug {
		log.Println("[handlePost] uri requested is: ", ph.lastPath, remPath)
	}

	elapsed := time.Since(ph.start)
	log.Printf("[WARN] ignored request. Total call duration: %v\n", elapsed)
	return nil
}
