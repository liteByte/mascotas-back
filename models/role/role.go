package role

import (
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(role structs.RoleStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO Role (name) VALUES(?)`, role.Name)
	return err
}

func GetList() ([]structs.RoleStruct, error) {
	var role structs.RoleStruct
	roleList := make([]structs.RoleStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, name FROM Role`)
	defer rows.Close()

	if err != nil {
		return []structs.RoleStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return []structs.RoleStruct{}, err
		}

		roleList = append(roleList, structs.RoleStruct{Id: role.Id, Name: role.Name})
	}

	return roleList, nil
}

func GetById(id int) (structs.RoleStruct, error) {
	var role structs.RoleStruct
	role.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT name FROM Role WHERE id = ?`, id).Scan(&role.Name)
	if err != nil {
		return structs.RoleStruct{}, err
	}

	return role, nil
}

func UpdateById(role structs.RoleStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE Role SET name = ? WHERE id = ?`, role.Name, role.Id)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM Role WHERE id = ?`, id)
	return err
}
