package models

import (
	"fmt"
	"inv-id-oice/db"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ModelHandler struct {
	debug         bool
	lastPath      string
	remPath       string
	liteCommentDB *db.LiteDB
	req           *http.Request
}

func NewModelHandler(dbg bool, litedb *db.LiteDB) *ModelHandler {
	ret := ModelHandler{debug: dbg, liteCommentDB: litedb}
	return &ret
}

func (mh *ModelHandler) HandleModel(w http.ResponseWriter, req *http.Request, status *int) error {
	u, _ := url.Parse(req.RequestURI)

	log.Printf("%s requested %s", req.Method, u)

	remPath := ""
	mh.lastPath, mh.remPath = getLastPathInUri(req.RequestURI)
	if mh.debug {
		log.Println("Check the last path ", mh.lastPath, remPath)
	}
	w.Header().Set("Cache-Control", "stale-while-revalidate=3600")
	mh.req = req
	switch mh.lastPath {
	case "users":
		return mh.handleUsers(w)
	case "login":
		return mh.handleLogin(w)
	case "sample_table":
		return mh.viewSampleTable(w)
	}

	*status = http.StatusNotFound
	return fmt.Errorf("[WARN] invalid %s request for %s", mh.req.Method, mh.lastPath)

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
