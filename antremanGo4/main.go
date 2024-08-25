package main

import (
	"net/http"
	admin_database"antremanGo/database"
	config"antremanGo/config"
	admin_database2 "antremanGo/savedatabase"
)




func main() {
	admin_database.Users{}.Migrate()
	admin_database2.Saves{}.Migrate()
	http.ListenAndServe(":8080",config.Routers())
}









