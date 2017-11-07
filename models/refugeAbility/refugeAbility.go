package refugeAbility

import (
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(refugeAbility structs.RefugeAbilityStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO RefugeAbility (name) VALUES(?)`, refugeAbility.Name)
	return err
}

func GetList() ([]structs.RefugeAbilityStruct, error) {
	var refugeAbility structs.RefugeAbilityStruct
	refugeAbilityList := make([]structs.RefugeAbilityStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, name FROM RefugeAbility`)
	defer rows.Close()

	if err != nil {
		return []structs.RefugeAbilityStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&refugeAbility.Id, &refugeAbility.Name)
		if err != nil {
			return []structs.RefugeAbilityStruct{}, err
		}

		refugeAbilityList = append(refugeAbilityList, structs.RefugeAbilityStruct{Id: refugeAbility.Id, Name: refugeAbility.Name})
	}

	return refugeAbilityList, nil
}

func GetById(id int) (structs.RefugeAbilityStruct, error) {
	var refugeAbility structs.RefugeAbilityStruct
	refugeAbility.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT name FROM RefugeAbility WHERE id = ?`, id).Scan(&refugeAbility.Name)
	if err != nil {
		return structs.RefugeAbilityStruct{}, err
	}

	return refugeAbility, nil
}

func UpdateById(refugeAbility structs.RefugeAbilityStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE RefugeAbility SET name = ? WHERE id = ?`, refugeAbility.Name, refugeAbility.Id)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM RefugeAbility WHERE id = ?`, id)
	return err
}
