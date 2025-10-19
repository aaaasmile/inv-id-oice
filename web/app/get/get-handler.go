package get

import (
	"fmt"
	"inv-id-oice/db"
	"inv-id-oice/util"
	"log"
	"net/http"
	"net/url"
)

type TypeGetHandler struct {
	debug         bool
	lastPath      string
	remPath       string
	liteCommentDB *db.LiteDB
}

func NewTypeGetHandler(dbg bool, litedb *db.LiteDB) *TypeGetHandler {
	ret := TypeGetHandler{debug: dbg, liteCommentDB: litedb}
	return &ret
}

func (gh *TypeGetHandler) HandleGet(w http.ResponseWriter, req *http.Request, status *int) error {
	u, _ := url.Parse(req.RequestURI)

	log.Println("GET requested ", u)

	remPath := ""
	gh.lastPath, gh.remPath = util.GetLastPathInUri(req.RequestURI)
	if gh.debug {
		log.Println("Check the last path ", gh.lastPath, remPath)
	}
	w.Header().Set("Cache-Control", "stale-while-revalidate=3600")
	switch gh.lastPath {
	case "users":
		return gh.handleUsers(w)
	}

	*status = http.StatusNotFound
	return fmt.Errorf("[WARN] invalid GET request for %s", gh.lastPath)

}
