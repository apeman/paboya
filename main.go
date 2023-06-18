package main
import (
	"html/template"
	"net/http"
  	"log"
	"fmt"
	"strings"
	"os"
	 
    "github.com/julienschmidt/httprouter"
)

const PORT = os.Getenv("PORT")
var tmpl = template.Must(template.ParseGlob("templates/*.html"))


func main() {
	HandleRoutes()
}


func HandleRoutes() {

	router := httprouter.New()
//--------Blog--------//
	router.GET("/", Index)
	router.GET("/Index", Blog)
	router.GET("/new", CreateBlogPost)
	router.POST("/new", CreateBlogPost)
	router.GET("/edit/:postid", EditBlogPost)
	router.POST("/edit/:postid", EditBlogPost)
	router.GET("/view/:postid", ViewBlogPost)
	router.POST("/delete/:postid", DeleteBlogPost)

//--------Login--------//
	router.GET("/register", Register)
	router.POST("/register", Register)
	router.GET("/login", loginHandler)
	router.POST("/login", loginHandler)
	router.POST("/logout", logoutHandler)

//--------FileServer--------//
	router.NotFound = http.FileServer(http.Dir(""))
	router.ServeFiles("/img/*filepath", http.Dir("uploads"))
	router.ServeFiles("/static/*filepath", http.Dir("static"))

//--------Server--------//
	log.Println("Starting erver on ", PORT)
	err := http.ListenAndServe(PORT, router)
//err := http.ListenAndServe(GetPort(), router)
 	if err != nil {
		log.Fatal("error starting http server : ", router)
 	}

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("path", r.URL.Path)
	if r.Method == "GET" {
		m := rdxHgetall("newpost")
			for k, v := range m {
				fmt.Println(">>", k)
				parts := strings.Split(v, ":::")
				fmt.Println(parts[0])
				fmt.Println(parts[1])
			}
			fmt.Println(m)
		if isLoggedIn(r) {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nav.html", LoginStatus{LoggedIn: "true"})
			tmpl.ExecuteTemplate(w, "index.html", m)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		} else {
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "nonloginhome.html", nil)
			tmpl.ExecuteTemplate(w, "footer.html", nil)
		}
	}
}
