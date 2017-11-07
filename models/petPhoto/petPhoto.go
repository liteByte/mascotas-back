package petPhoto

import (
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(petPhoto structs.PetPhotoStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO PetPhoto (idPet, idPhoto) VALUES(?, ?)`, petPhoto.IdPet, petPhoto.IdPhoto)
	return err
}

func GetList() ([]structs.PetPhotoStruct, error) {
	var petPhoto structs.PetPhotoStruct
	petPhotoList := make([]structs.PetPhotoStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, idPet, idPhoto FROM PetPhoto`)
	defer rows.Close()

	if err != nil {
		return []structs.PetPhotoStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&petPhoto.Id, &petPhoto.IdPet, &petPhoto.IdPhoto)
		if err != nil {
			return []structs.PetPhotoStruct{}, err
		}

		petPhotoList = append(petPhotoList, structs.PetPhotoStruct{Id: petPhoto.Id, IdPet: petPhoto.IdPet, IdPhoto: petPhoto.IdPhoto})
	}

	return petPhotoList, nil
}

func GetById(id int) (structs.PetPhotoStruct, error) {
	var petPhoto structs.PetPhotoStruct
	petPhoto.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT idPet, idPhoto FROM PetPhoto WHERE id = ?`, id).Scan(&petPhoto.IdPet, &petPhoto.IdPhoto)
	if err != nil {
		return structs.PetPhotoStruct{}, err
	}

	return petPhoto, nil
}

func UpdateById(petPhoto structs.PetPhotoStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE PetPhoto SET idPet = ?, idPhoto = ? WHERE id = ?`, petPhoto.IdPet, petPhoto.IdPhoto, petPhoto.Id)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM PetPhoto WHERE id = ?`, id)
	return err
}
