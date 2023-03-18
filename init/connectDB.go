package init

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func connectDB() {
	var err error

	db, err = gorm.Open("postgres", "host=localhost user=myuser dbname=mydb sslmode=disable password=mypassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

}
