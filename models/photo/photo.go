package photo

import (
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(photo structs.PhotoStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO Photo (url, lat, lon) VALUES(?, ?, ?)`, photo.Url, photo.Lat, photo.Lon)
	return err
}

func GetList() ([]structs.PhotoStruct, error) {
	var photo structs.PhotoStruct
	photoList := make([]structs.PhotoStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, url, lat, lon FROM Photo`)
	defer rows.Close()

	if err != nil {
		return []structs.PhotoStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&photo.Id, &photo.Url, &photo.Lat, &photo.Lon)
		if err != nil {
			return []structs.PhotoStruct{}, err
		}

		photoList = append(photoList, structs.PhotoStruct{Id: photo.Id, Url: photo.Url, Lat: photo.Lat, Lon: photo.Lon})
	}

	return photoList, nil
}

func GetById(id int) (structs.PhotoStruct, error) {
	var photo structs.PhotoStruct
	photo.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT url, lat, lon FROM Photo WHERE id = ?`, id).Scan(&photo.Url, &photo.Lat, &photo.Lon)
	if err != nil {
		return structs.PhotoStruct{}, err
	}

	return photo, nil
}

func UpdateById(photo structs.PhotoStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE Photo SET url = ?, lat = ?, lon = ? WHERE id = ?`, photo.Url, photo.Lat, photo.Lon, photo.Id)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM Photo WHERE id = ?`, id)
	return err
}
