package function

import (
	database "antremanGo/database"
	savee "antremanGo/savedatabase"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Post struct{}
type Add struct{}
type Data struct {
	Ideas string `json:"idea"`
    Likeclass1 string `json:"classname"`
	Likenumberr int `json:"intval"`
}


func (add Add) Home(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	view, err := template.ParseFiles("views/idea.html")
    if err != nil {
        http.Error(w, "Template yüklenemedi: "+err.Error(), http.StatusInternalServerError)
        return
    }

    posts := database.Users{}.GetAll()
    if posts == nil {
        http.Error(w, "Postlar alınamadı", http.StatusInternalServerError)
        return
    }

    reversedPosts := make([]database.Users, len(posts))
    for i, p := range posts {
        reversedPosts[len(posts)-1-i] = p
    }

    data := make(map[string]interface{})
    data["Posts"] = reversedPosts

    err = view.Execute(w, data)
    if err != nil {
        http.Error(w, "Template render edilemedi: "+err.Error(), http.StatusInternalServerError)
    }
}

func(add Add) Index(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	views,_ := template.ParseFiles("views/index.html")
	views.Execute(w,nil)
}

func (add Add) Send(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	idea := r.FormValue("Idea")
	database.Users{
		Idea: idea,
	}.Add()
	http.Redirect(w,r,"/",http.StatusSeeOther)
}

func (add Add) Saved(w http.ResponseWriter,r *http.Request, params httprouter.Params){
			var data Data
            err := json.NewDecoder(r.Body).Decode(&data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            savee.Saves{
				İdea: data.Ideas,
			}.Add()
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Data saved successfully"))
}


func (add Add) Saves(w http.ResponseWriter,r *http.Request, params httprouter.Params){
	views,_ := template.ParseFiles("views/saves.html")
	data := make(map[string]interface{})
	data ["Saves"] = savee.Saves{}.GetAll()
	views.Execute(w,data)
}

func(add Add) Likes(w http.ResponseWriter,r *http.Request, params httprouter.Params){
			post := database.Users{}.Get(params.ByName("id"))
			var data Data
            err := json.NewDecoder(r.Body).Decode(&data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
			post.Updates(database.Users{
                Likenumber: data.Likenumberr,
                Likeclass: data.Likeclass1,
            })
           	
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Data saved successfully"))
}

