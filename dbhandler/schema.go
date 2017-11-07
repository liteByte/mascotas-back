package dbhandler

import (
	"github.com/liteByte/frango"
)

func createUserTable() {
	query := `CREATE TABLE IF NOT EXISTS User (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		username varchar(255) UNIQUE NOT NULL,
		password varchar(255) NOT NULL,
		refugeName varchar(255),
		email varchar(255) UNIQUE NOT NULL,
		roleID int NOT NULL,
		lat float,
		lon float
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func createRoleTable() {
	query := `CREATE TABLE IF NOT EXISTS Role (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		name varchar(255)
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func createPhotoTable() {
	query := `CREATE TABLE IF NOT EXISTS Photo (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		url varchar(255),
		lat float NOT NULL,
		lon float NOT NULL
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func createRefugeAbilityTable() {
	query := `CREATE TABLE IF NOT EXISTS RefugeAbility (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		name varchar(255)
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func createAuditorReviewTable() {
	query := `CREATE TABLE IF NOT EXISTS AuditorReview (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		idUser int NOT NULL,
		idRefuge int NOT NULL,
		idRefugeAbility int NOT NULL,
		score int NOT NULL
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func createPetTable() {
	query := `CREATE TABLE IF NOT EXISTS Pet (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		refugeId int NOT NULL,
		familyId int,
		breed varchar(255) NOT NULL,
		color varchar(255) NOT NULL,
		age int NOT NULL
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func createPetPhotoTable() {
	query := `CREATE TABLE IF NOT EXISTS PetPhoto (
		id int UNIQUE NOT NULL AUTO_INCREMENT,
		idPet int NOT NULL,
		idPhoto int NOT NULL
	);`
	_, err := db.Exec(query)
	frango.PrintErr(err)
}
