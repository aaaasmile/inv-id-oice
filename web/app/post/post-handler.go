package post

import (
	"inv-id-oice/db"
	"inv-id-oice/util"
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

func NewPostHandler(dbg bool, litedb *db.LiteDB) *PostHandler {
	ret := PostHandler{debug: dbg, liteDB: litedb}
	return &ret
}

func (ph *PostHandler) HandlePost(w http.ResponseWriter, req *http.Request) error {
	ph.start = time.Now()
	remPath := ""
	ph.lastPath, remPath = util.GetLastPathInUri(req.RequestURI)
	if ph.debug {
		log.Println("[handlePost] uri requested is: ", ph.lastPath, remPath)
	}

	elapsed := time.Since(ph.start)
	log.Printf("[WARN] ignored request. Total call duration: %v\n", elapsed)
	return nil
}
