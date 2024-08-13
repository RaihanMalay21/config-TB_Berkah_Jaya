package main 

import (
	"log"
	"net/http"
	"gorm.io/gorm"
	config "github.com/RaihanMalay21/config-TB_Berkah_Jaya"
)

var (
	db *gorm.DB
	err error
)

func main() {
	db, err = config.DB_Connection()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Raw("SELECT 1").Error; err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while trying to connect to database"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ping succes to database"))
		log.Println("Ping succes")
	}
}