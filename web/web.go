package web

import (
	"fmt"
	"inv-id-oice/conf"
	"inv-id-oice/util"
	"inv-id-oice/web/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

type ServiceRunner struct {
	configfile string
}

func NewServiceRunner(confile string) *ServiceRunner {
	app := ServiceRunner{configfile: confile}
	return &app
}

func (rs *ServiceRunner) RunService() error {

	if _, err := conf.ReadConfig(rs.configfile); err != nil {
		return err
	}

	serverurl := conf.Current.ServiceURL
	serverurl = strings.Replace(serverurl, "0.0.0.0", "localhost", 1)
	serverurl = strings.Replace(serverurl, "127.0.0.1", "localhost", 1)
	log.Println("Server started with URL ", serverurl)

	// Site should be /
	log.Println("Try this url for: ", fmt.Sprintf("http://%s", serverurl))
	staticDirSrv := http.Dir(util.GetFullPath(fmt.Sprintf("static/%s", conf.Current.StaticAppDir)))
	log.Println("static dir", staticDirSrv)
	http.Handle("/", http.StripPrefix("/", http.FileServer(staticDirSrv)))

	myApp, err := app.NewApp()
	if err != nil {
		return err
	}
	http.HandleFunc(conf.Current.RootURLPattern, myApp.APiHandler)

	srv := &http.Server{
		Addr: serverurl,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      nil,
	}
	go func() {
		log.Println("start listening web with http")
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Server is not listening anymore: ", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	log.Println("Enter in server blocking loop")

loop:
	for {
		select {
		case <-sig:
			log.Println("stop because interrupt")
			break loop
		}
	}

	log.Println("Bye, service")
	return nil
}
