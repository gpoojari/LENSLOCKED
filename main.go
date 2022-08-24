package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my Awesome site!</h1>")
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "hi %v", id)

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch , email me at <a href=\"mailto:poojari.girish@gmail.com\">Girish Poojari</a>.")
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "hi %v", id)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//fmt.Fprint(w, "<h1>Faq Page</h1> <p> Q: Is there  a free Version?<br> A: Yes! We offer a free trial  for 30 dayson any paid Plans<br><br> Q: What are your Support hours? <br> A: Yes We have support staff</p>")

	fmt.Fprint(w, `<h1>Faq Page</h1> 
	<ul>
	  <li> <b>Q: Is there  a free Version? </b> <br> A: Yes! We offer a free trial  for 30 dayson any paid Plans<br><br> 
	 <li> <b> Q: What are your Support hours?</b> <br> A: Yes We have support staff</li>
	<ul>`)
}

/*
func pathHandler(w http.ResponseWriter, r *http.Request) {
*/
/*if r.URL.Path == "/" {
	homeHandler(w,r)
}
if r.URL.Path == "/contact" {
	contactHandler(w,r)
}*/

/* switch r.URL.Path {
case "/":
	homeHandler(w, r)
case "/contact":
	contactHandler(w, r)
default: */
//TODO: Handle the page not found error
/*w.WriteHeader(http.StatusNotFound)
//http.NotFound(w, r)
fmt.Fprint(w, "Page not found")
*/
//http.Error(w, "Page not found", http.StatusNotFound)
/*http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	//fmt.Fprintln(w, r.URL.Path)
	//fmt.Fprintln(w, r.URL.RawPath)
}
*/

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
func main() {
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)
	//http.HandleFunc("/", pathHandler)
	//var router Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000....")
	http.ListenAndServe(":3000", r)
}
