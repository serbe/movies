package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

func httpLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlStr := r.URL.String()
		if !strings.Contains(urlStr, "/public/") {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		}
		h.ServeHTTP(w, r)
	})
}

func serveFileHandler(w http.ResponseWriter, req *http.Request) {
	fname := path.Base(req.URL.Path)
	http.ServeFile(w, req, "./public"+fname)
}

func (app *application) initServer() {
	var h = http.NewServeMux()

	h.HandleFunc("/", app.index)
	h.HandleFunc("/search", app.search)
	// h.HandleFunc("/b/", app.bindex)

	// h.HandleFunc("/about", about)

	h.HandleFunc("/favicon.ico", serveFileHandler)

	h.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	var serverHandler http.Handler
	if app.cfg.Web.Log {
		serverHandler = httpLogger(h)
	} else {
		serverHandler = h
	}

	s := &http.Server{
		Addr:           ":" + app.cfg.Web.Port, // server port
		Handler:        serverHandler,          // handler to invoke, http.DefaultServeMux if nil
		ReadTimeout:    10 * time.Second,       // maximum duration before timing out read of the request
		WriteTimeout:   10 * time.Second,       // maximum duration before timing out write of the response
		MaxHeaderBytes: 1 << 20,                // maximum size of request headers, 1048576 bytes
	}

	log.Fatal(s.ListenAndServe())

}

// func render(w http.ResponseWriter, cont context, name string) error {
// 	tmpl, ok := templates[name]
// 	if !ok {
// 		return fmt.Errorf("The template %s does not exist./n", name)
// 	}
// 	err := tmpl.ExecuteTemplate(w, "base.html", cont)
// 	if err != nil {
// 		log.Print("template executing error: ", err)
// 	}
// 	return err
// }

func render(w http.ResponseWriter, cont context, tmpl string) error {
	tmplList := []string{"templates/base.html",
		fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmplList...)
	if err != nil {
		log.Print("template parsing error: ", err)
		return err
	}
	err = t.Execute(w, cont)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	return err
}

func renderBulma(w http.ResponseWriter, cont context, tmpl string) error {
	tmplList := []string{"templates/bbase.html",
		fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmplList...)
	if err != nil {
		log.Print("template parsing error: ", err)
		return err
	}
	err = t.Execute(w, cont)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	return err
}
