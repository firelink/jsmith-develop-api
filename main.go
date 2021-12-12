package main

import (
	"log"
	"net/http"

	"github.com/firelink/api.jsmith-develop.com/argSea/repo"
	"github.com/firelink/api.jsmith-develop.com/argSea/service"
	"github.com/firelink/api.jsmith-develop.com/argSea/structure/argStore"
	"github.com/firelink/api.jsmith-develop.com/argSea/usecase"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()

	if nil != err {
		panic(err)
	}

	//Possibly add debugging?
}

func main() {
	router := mux.NewRouter()

	//Cache credentials
	mHost := viper.GetString("mongo.host") + ":" + viper.GetString("mongo.port")
	mUser := viper.GetString("mongo.user")
	mPass := viper.GetString("mongo.pass")
	mDB := viper.GetString("mongo.dbName")

	userTable := "users"
	projectTable := "projects"
	resumeTable := "resume"

	//User
	userRepo := repo.NewUserRepo(argStore.NewMordor(mHost, mUser, mPass, mDB, userTable))
	userCase := usecase.NewUserCase(userRepo)

	//Project
	projRepo := repo.NewProjectRepo(argStore.NewMordor(mHost, mUser, mPass, mDB, projectTable))
	projCase := usecase.NewProjectCase(projRepo)

	//Resume
	resumeRepo := repo.NewResumeRepo(argStore.NewMordor(mHost, mUser, mPass, mDB, resumeTable))
	resumeCase := usecase.NewResumeCase(resumeRepo)

	//user
	userRouter := router.PathPrefix("/api/1/user/").Subrouter()
	service.NewUserService(userRouter, userCase)

	//Project
	projRouter := router.PathPrefix("/api/1/project/").Subrouter()
	service.NewProjectService(projRouter, projCase)

	//Resume
	resumeRouter := router.PathPrefix("/api/1/resume/").Subrouter()
	service.NewResumeService(resumeRouter, resumeCase)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
