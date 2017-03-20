package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/nemesisesq/middleware"
	"os"
	"github.com/nemesisesq/wearevest-api/api"
)


var RM *middleware.RedisAccessor
var RMQ *middleware.RabbitMQAccessor
var MGO *middleware.DatabaseAccessor

func init(){
	//Set up middleware
	//RM, _ = middleware.NewRedisAccessor(os.Getenv("REDIS_URI"), "", 1)
	//RMQ, _ = middleware.NewRabbitMQAccessor(os.Getenv("RABITMQ_URI"))
	MGO, _ = middleware.NewDatabaseAccessor(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_NAME"), os.Getenv("MONGODB_COLL"))

}

func main() {

	port := "8080"
	//	Routes:
	r := mux.NewRouter()

	//	Root 'home' route
	//r.HandleFunc("/", HomeHandler)

	n := negroni.Classic()

	//n.Use(middleware.NewRedisClient(*RM).Middleware())
	//n.Use(middleware.NewRabbitMQConnection(*RMQ).Middleware())
	n.Use(middleware.NewDatabase(*MGO).Middleware())

	r.HandleFunc("/", HomeHandler)
	r.Handle("/graphql", api.H)



	fmt.Println(fmt.Sprintf("listening on port :%s", port))

	//	To login/logout/signup:
	//	/auth/login
	//	/auth/logout
	//	/auth/signup
	auth := r.PathPrefix("/auth").Subrouter()
	auth.Path("/login").HandlerFunc(LoginHandler)
	auth.Path("/logout").HandlerFunc(LogoutHandler)
	auth.Path("/signup").HandlerFunc(SignupHandler)

	// Posts collection
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)

	//	Accounts
	acctBase := mux.NewRouter()
	r.PathPrefix("/account").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.Wrap(acctBase),
	))

	acct := acctBase.PathPrefix("/account").Subrouter()
	acct.Path("/profile").HandlerFunc(ProfileHandler)

	// Posts singular
	post := r.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	fmt.Println("Starting server on :" + port)

	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":" + port, n))
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")

	if r.URL.Query().Get("password") == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back...")
}

func AcctPrefixHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("AcctPrefixHandler...")
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post edit")
}

/*	Auth handlers	*/
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "login")
}

func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "logout")
}

func SignupHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "sign up")
}

func ProfileHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "account profile")
}

func SettingsHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "account settings")
}


