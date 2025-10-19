package post

import (
	"inv-id-oice/db"
	"log"
	"net/http"
	"strings"
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
	ph.lastPath, remPath = getLastPathInUri(req.RequestURI)
	if ph.debug {
		log.Println("[handlePost] uri requested is: ", ph.lastPath, remPath)
	}

	elapsed := time.Since(ph.start)
	log.Printf("[WARN] ignored request. Total call duration: %v\n", elapsed)
	return nil
}

func getLastPathInUri(uri string) (string, string) {
	arr := strings.Split(uri, "/")
	for i := len(arr) - 1; i >= 0; i-- {
		last := arr[i]
		rem_ix := i
		if last != "" {
			if !strings.HasPrefix(last, "?") {
				remPath := strings.Join(arr[0:rem_ix], "/")
				return last, remPath
			}
		}
	}
	return uri, ""
}
