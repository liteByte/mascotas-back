package dbhandler

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liteByte/frango"
	"github.com/liteByte/mascotas-backend/config"
)

var db *sql.DB

func ConnectToDatabase() {
	var err error

	db, err = sql.Open(config.GetConfig().DB_TYPE, config.GetConfig().DB_USERNAME+":"+config.GetConfig().DB_PASSWORD+"@tcp("+config.GetConfig().DB_HOST+":"+config.GetConfig().DB_PORT+")/"+config.GetConfig().DB_NAME)
	frango.PrintErr(err)

	err = db.Ping()
	frango.PrintErr(err)

	createSchema()
}

func GetDatabase() *sql.DB {
	return db
}

func createSchema() {
	createUserTable()
	createRoleTable()
	createPhotoTable()
	createRefugeAbilityTable()
	createAuditorReviewTable()
	createPetTable()
	createPetPhotoTable()

}
