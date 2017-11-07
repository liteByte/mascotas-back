package pet

import (
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(pet structs.PetStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO Pet (refugeId, familyId, breed, color, age) VALUES(?, ?, ?, ?, ?)`, pet.RefugeId, pet.FamilyId, pet.Breed, pet.Color, pet.Age)
	return err
}

func GetList() ([]structs.PetStruct, error) {
	var pet structs.PetStruct
	petList := make([]structs.PetStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, refugeId, familyId, breed, color, age FROM Pet`)
	defer rows.Close()

	if err != nil {
		return []structs.PetStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&pet.Id, &pet.RefugeId, &pet.FamilyId, &pet.Breed, &pet.Color, &pet.Age)
		if err != nil {
			return []structs.PetStruct{}, err
		}

		petList = append(petList, structs.PetStruct{Id: pet.Id, RefugeId: pet.RefugeId, FamilyId: pet.FamilyId, Breed: pet.Breed, Color: pet.Color, Age: pet.Age})
	}

	return petList, nil
}

func GetById(id int) (structs.PetStruct, error) {
	var pet structs.PetStruct
	pet.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT refugeId, familyId, breed, color, age FROM Pet WHERE id = ?`, id).Scan(&pet.RefugeId, &pet.FamilyId, &pet.Breed, &pet.Color, &pet.Age)
	if err != nil {
		return structs.PetStruct{}, err
	}

	return pet, nil
}

func UpdateById(pet structs.PetStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE Pet SET refugeId = ?, familyId = ?, breed = ?, color = ?, age = ? WHERE id = ?`, pet.RefugeId, pet.FamilyId, pet.Breed, pet.Color, pet.Age, pet.Id)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM Pet WHERE id = ?`, id)
	return err
}
