package service

import (
    "net/http"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/{id}", testHandler(formatter)).Methods("GET")
    mx.HandleFunc("/repos/{owner}/{repo}",reposHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]
        formatter.JSON(w, http.StatusOK, struct{Current_user_url  string;Emails_url string;Repository_url string;}{"https://localhost:8080/"+id,"https://localhost:8080/"+id+"/emails","https://localhost:8080/repos/{owner}/{repo}"})
    }
}

func reposHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        owner := vars["owner"]
        name  := vars["repo"]
        formatter.JSON(w, http.StatusOK, struct{Owener  string;Name string;}{owner,name})
    }
}