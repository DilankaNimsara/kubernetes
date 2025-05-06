package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Start() {

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/api/v1/readEnv/configmap", readEnvFromConfigMap).Methods("GET")
	r.HandleFunc("/api/v1/readEnv/secret", readEnvFromSecret).Methods("GET")

	fmt.Println("Server starting .....")
	http.ListenAndServe(":8080", r)

}

func readEnvFromSecret(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Read Env")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"DB_PASSWORD": "%s"} {"DB_USERNAME": "%s"}`, os.Getenv("DB_PASSWORD"), os.Getenv("DB_USERNAME"))))
	fmt.Println("Env read successfully")
}

func readEnvFromConfigMap(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Read Env")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"DB_URL": "%s"}`, os.Getenv("DB_URL"))))
	fmt.Println("Env read successfully")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Write([]byte("Welcome to K8s REST API"))
}
