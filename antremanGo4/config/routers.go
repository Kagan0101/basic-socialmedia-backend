package config

import (
	admin"antremanGo/function"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routers() *httprouter.Router{
	r := httprouter.New()
	r.GET("/",admin.Add{}.Home)
	r.GET("/index",admin.Add{}.Index)
	r.POST("/send",admin.Add{}.Send)
	r.POST("/save",admin.Add{}.Saved)
	r.GET("/saves",admin.Add{}.Saves)
	r.POST("/likes/:id",admin.Add{}.Likes)

	r.ServeFiles("/views/*filepath",http.Dir("views"))
	r.ServeFiles("/img/*filepath",http.Dir("img"))
	return r
}
