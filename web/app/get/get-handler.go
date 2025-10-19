package get

import (
	"fmt"
	"inv-id-oice/db"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type PageCtx struct {
	RootUrl string
}

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
	gh.lastPath, gh.remPath = getLastPathInUri(req.RequestURI)
	if gh.debug {
		log.Println("Check the last path ", gh.lastPath, remPath)
	}
	switch gh.lastPath {
	case "users":
		gh.handleUsers(w)
	}

	*status = http.StatusNotFound
	return fmt.Errorf("[WARN] invalid GET request for %s", gh.lastPath)

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
