package user

import (
	"errors"
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func CheckLogin(loginStruct structs.LoginStruct) error {
	var exists bool

	err := dbhandler.GetDatabase().QueryRow(`SELECT EXISTS (SELECT 1 FROM User WHERE username = ? AND password = ? LIMIT 1)`, loginStruct.Username, loginStruct.Password).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New(`Login failed`)
	}

	return nil
}

func Create(user structs.UserStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO User (username, password, refugeName, email, roleID, lat, lon) VALUES(?, ?, ?, ?, ?, ?, ?)`, user.Username, user.Password, user.RefugeName, user.Email, user.RoleID, user.Lat, user.Lon)
	return err
}

func GetList() ([]structs.UserStruct, error) {
	var user structs.UserStruct
	userList := make([]structs.UserStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, username, password, refugeName, email, roleID, lat, lon FROM User`)
	defer rows.Close()

	if err != nil {
		return []structs.UserStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.RefugeName, &user.Email, &user.RoleID, &user.Lat, &user.Lon)
		if err != nil {
			return []structs.UserStruct{}, err
		}

		userList = append(userList, structs.UserStruct{Id: user.Id, Username: user.Username, Password: user.Password, RefugeName: user.RefugeName, Email: user.Email, RoleID: user.RoleID, Lat: user.Lat, Lon: user.Lon})
	}

	return userList, nil
}

func GetById(id int) (structs.UserStruct, error) {
	var user structs.UserStruct
	user.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT username, password, refugeName, email, roleID, lat, lon FROM User WHERE id = ?`, id).Scan(&user.Username, &user.Password, &user.RefugeName, &user.Email, &user.RoleID, &user.Lat, &user.Lon)
	if err != nil {
		return structs.UserStruct{}, err
	}

	return user, nil
}

func GetByUsername(username string) (structs.UserStruct, error) {
	var user structs.UserStruct
	user.Username = username

	err := dbhandler.GetDatabase().QueryRow(`SELECT id, password, refugeName, email, roleID, lat, lon FROM User WHERE username = ?`, username).Scan(&user.Id, &user.Password, &user.RefugeName, &user.Email, &user.RoleID, &user.Lat, &user.Lon)
	if err != nil {
		return structs.UserStruct{}, err
	}

	return user, nil
}

func GetByEmail(email string) (structs.UserStruct, error) {
	var user structs.UserStruct
	user.Email = email

	err := dbhandler.GetDatabase().QueryRow(`SELECT id, username, password, refugeName, roleID, lat, lon FROM User WHERE email = ?`, email).Scan(&user.Id, &user.Username, &user.Password, &user.RefugeName, &user.RoleID, &user.Lat, &user.Lon)
	if err != nil {
		return structs.UserStruct{}, err
	}

	return user, nil
}

func UpdateById(user structs.UserStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE User SET username = ?, password = ?, refugeName = ?, email = ?, roleID = ?, lat = ?, lon = ? WHERE id = ?`, user.Username, user.Password, user.RefugeName, user.Email, user.RoleID, user.Lat, user.Lon, user.Id)
	return err
}

func UpdateByUsername(user structs.UserStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE User SET password = ?, refugeName = ?, email = ?, roleID = ?, lat = ?, lon = ? WHERE username = ?`, user.Password, user.RefugeName, user.Email, user.RoleID, user.Lat, user.Lon, user.Username)
	return err
}

func UpdateByEmail(user structs.UserStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE User SET username = ?, password = ?, refugeName = ?, roleID = ?, lat = ?, lon = ? WHERE email = ?`, user.Username, user.Password, user.RefugeName, user.RoleID, user.Lat, user.Lon, user.Email)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM User WHERE id = ?`, id)
	return err
}

func DeleteByUsername(username string) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM User WHERE username = ?`, username)
	return err
}

func DeleteByEmail(email string) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM User WHERE email = ?`, email)
	return err
}
