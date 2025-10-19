package app

import (
	"fmt"
	"inv-id-oice/db"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type PageCtx struct {
	RootUrl        string
	Buildnr        string
	ServerName     string
	VuetifyLibName string
	VueLibName     string
}

type GetHandler struct {
	debug         bool
	lastPath      string
	liteCommentDB *db.LiteDB
}

func (gh *GetHandler) handleGet(w http.ResponseWriter, req *http.Request, status *int) error {
	u, _ := url.Parse(req.RequestURI)

	log.Println("GET requested ", u)

	remPath := ""
	gh.lastPath, remPath = getLastPathInUri(req.RequestURI)
	if gh.debug {
		log.Println("Check the last path ", gh.lastPath, remPath)
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
