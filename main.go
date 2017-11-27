package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/HinanawiTenshi/cloudgo_data/entities"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

const (
	dbPath string = "root:root@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=true"
)

func main() {
	port := ":8080"

	router := mux.NewRouter()
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	mySQLEngine, err := xorm.NewEngine("mysql", dbPath)
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/service/user",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			user := entities.NewUser(r.Form["username"][0],
				r.Form["department"][0])

			// database/sql VS xorm
			// entities.UserService.Insert(user)
			_, err := mySQLEngine.Table("users").Insert(user)
			if err != nil {
				panic(err)
			}

			formatter.JSON(w, http.StatusOK, user)
		}).Methods("POST")
	router.HandleFunc("/service/user",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			if len(r.Form["uid"][0]) != 0 {
				id, _ := strconv.ParseInt(r.Form["uid"][0], 10, 32)

				// database/sql VS xorm
				// user := entities.UserService.FindByID(int(id))
				user := entities.User{}
				mySQLEngine.Table("users").Where("uid = ?", int(id)).Get(&user)

				formatter.JSON(w, http.StatusOK, user)
			} else {
				// userList := entities.UserService.FindAll()
				userList := make([]entities.User, 0)
				mySQLEngine.Table("users").Find(&userList)

				formatter.JSON(w, http.StatusOK, userList)
			}
		}).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)

	fmt.Printf("Listening to port %v\n", port)
	http.ListenAndServe(port, n)
}
