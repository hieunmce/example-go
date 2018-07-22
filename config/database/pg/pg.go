package pg

import (
	"log"

	"github.com/jinzhu/gorm"
)

// New create new postgres database connection
// it return postgres connection and cleanup postgres connection
func New(ds string) (*gorm.DB, func()) {
	db, err := gorm.Open("postgres", ds)
	if err != nil {
		panic(err)
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			log.Println("Failed to close DB by error", err)
		}
	}
}
