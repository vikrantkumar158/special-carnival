package main

import (
	"net/http"
	"project/models"
	"project/routes"
	"project/utils"
	//"context"
	//"fmt"
	//"github.com/go-redis/redis"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//"golang.org/x/crypto/bcrypt"
	//"html/template"
)

//var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
//var templates *template.Template

func main() {
	models.Init()
	utils.LoadTemplates("src/project/templates/*.html")
	r := routes.NewRouter()
	//templates = template.Must(template.ParseGlob("src/project/templates/*.html"))
	//r := mux.NewRouter()
	//r.HandleFunc("/", handler).Methods("GET")
	//r.HandleFunc("/hello", helloHandler).Methods("GET")
	//r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")
	//r.HandleFunc("/", AuthRequired(indexGetHandler)).Methods("GET")
	//r.HandleFunc("/", AuthRequired(indexPostHandler)).Methods("POST")
	//r.HandleFunc("/login", loginGetHandler).Methods("GET")
	//r.HandleFunc("/login", loginPostHandler).Methods("POST")
	//r.HandleFunc("/register", registerGetHandler).Methods("GET")
	//r.HandleFunc("/register", registerPostHandler).Methods("POST")
	//r.HandleFunc("/test", testGetHandler).Methods("GET")
	//fs := http.FileServer(http.Dir("./src/project/static/"))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	//http.HandleFunc("/", handler) //Before using mux
	_ = http.ListenAndServe(":8000", nil)
}

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	_, _ = fmt.Fprint(w, "Hello world!")
//}
//
//func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
//	_, _ = fmt.Fprint(w, "Goodbye world!")
//}

//func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//session, _ := store.Get(r, "session")
//		session, _ := sessions.Store.Get(r, "session")
//		_, ok := session.Values["username"]
//		if !ok {
//			http.Redirect(w, r, "/login", 302)
//			return
//		}
//		handler.ServeHTTP(w, r)
//	}
//}

//func indexGetHandler(w http.ResponseWriter, r *http.Request) {
//	//comments, err := client.LRange(context.TODO(),"comments", 0, 10).Result()
//	comments, err := models.GetComments()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		_, _ = w.Write([]byte("Internal server error"))
//		return
//	}
//	//_ = templates.ExecuteTemplate(w, "index.html", comments)
//	utils.ExecuteTemplate(w, "index.html", comments)
//}
//
//func indexPostHandler(w http.ResponseWriter, r *http.Request) {
//	_ = r.ParseForm()
//	comment := r.PostForm.Get("comment")
//	//err := client.LPush(context.TODO(), "comments", comment).Err()
//	err := models.PostComment(comment)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		_, _ = w.Write([]byte("Internal server error"))
//		return
//	}
//	http.Redirect(w, r, "/", 302)
//}
//
//func loginGetHandler(w http.ResponseWriter, r *http.Request) {
//	//_ = templates.ExecuteTemplate(w, "login.html", nil)
//	utils.ExecuteTemplate(w, "login.html", nil)
//}
//
//func loginPostHandler(w http.ResponseWriter, r *http.Request) {
//	_ = r.ParseForm()
//	username := r.PostForm.Get("username")
//	password := r.PostForm.Get("password")
//	err := models.AuthenticateUser(username, password)
//	if err != nil {
//		switch err {
//			case models.ErrUserNotFound:
//				//_ = templates.ExecuteTemplate(w, "login.html", "unknown user")
//				utils.ExecuteTemplate(w, "login.html", "unknown user")
//			case models.ErrInvalidLogin:
//				//_ = templates.ExecuteTemplate(w, "login.html", "invalid login")
//				utils.ExecuteTemplate(w, "login.html", "invalid login")
//			default:
//				w.WriteHeader(http.StatusInternalServerError)
//				_, _ = w.Write([]byte("Internal server error"))
//		}
//		return
//	}
//
//	//hash, err := client.Get(context.TODO(), "user:" + username).Bytes()
//	//if err == redis.Nil {
//	//	_ = templates.ExecuteTemplate(w, "login.html", "unknown user")
//	//	return
//	//} else if err != nil {
//	//	w.WriteHeader(http.StatusInternalServerError)
//	//	_, _ = w.Write([]byte("Internal server error"))
//	//	return
//	//}
//	//err = bcrypt.CompareHashAndPassword(hash, []byte(password))
//	//if err != nil {
//	//	_ = templates.ExecuteTemplate(w, "login.html", "invalid login")
//	//	return
//	//}
//	//session, _ := store.Get(r, "session")
//	session, _ := sessions.Store.Get(r, "session")
//	session.Values["username"] = username
//	_ = session.Save(r, w)
//	http.Redirect(w, r, "/", 302)
//}

//func testGetHandler(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "session")
//	untyped, ok := session.Values["username"]
//	if !ok {
//		return
//	}
//	username, ok := untyped.(string)
//	if !ok {
//		return
//	}
//	_, _ = w.Write([]byte(username))
//}

//func registerGetHandler(w http.ResponseWriter, r *http.Request) {
//	//_ = templates.ExecuteTemplate(w, "register.html", nil)
//	utils.ExecuteTemplate(w, "register.html", nil)
//}
//
//func registerPostHandler(w http.ResponseWriter, r *http.Request) {
//	_ = r.ParseForm()
//	username := r.PostForm.Get("username")
//	password := r.PostForm.Get("password")
//	err := models.RegisterUser(username, password)
//	//cost := bcrypt.DefaultCost
//	//hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
//	//if err != nil {
//	//	w.WriteHeader(http.StatusInternalServerError)
//	//	_, _ = w.Write([]byte("Internal server error"))
//	//	return
//	//}
//	//err = client.Set(context.TODO(), "user:" + username, hash, 0).Err()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		_, _ = w.Write([]byte("Internal server error"))
//		return
//	}
//	http.Redirect(w, r, "/login", 302)
//}