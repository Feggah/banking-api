package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Feggah/banking-api/domain"
	"github.com/Feggah/banking-api/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	checkConfiguration()

	router := mux.NewRouter()

	dbClient := getDbClient()

	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient)),
	}

	ah := AccountHandler{
		service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient)),
	}

	router.HandleFunc("/customers", ch.getCustomersByStatus).Queries("status", "{status}").Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.CreateTransaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func checkConfiguration() {
	neededEnvVars := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASSWORD", "DB_ADDRESS", "DB_PORT", "DB_NAME"}

	for _, envVar := range neededEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatal(fmt.Sprintf("Environment variable '%s' is not defined", envVar))
		}
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
