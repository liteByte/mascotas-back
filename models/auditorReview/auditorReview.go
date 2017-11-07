package auditorReview

import (
	"github.com/liteByte/mascotas-backend/dbhandler"
	"github.com/liteByte/mascotas-backend/structs"
)

func Create(auditorReview structs.AuditorReviewStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`INSERT INTO AuditorReview (idUser, idRefuge, idRefugeAbility, score) VALUES(?, ?, ?, ?)`, auditorReview.IdUser, auditorReview.IdRefuge, auditorReview.IdRefugeAbility, auditorReview.Score)
	return err
}

func GetList() ([]structs.AuditorReviewStruct, error) {
	var auditorReview structs.AuditorReviewStruct
	auditorReviewList := make([]structs.AuditorReviewStruct, 0)

	rows, err := dbhandler.GetDatabase().Query(`SELECT id, idUser, idRefuge, idRefugeAbility, score FROM AuditorReview`)
	defer rows.Close()

	if err != nil {
		return []structs.AuditorReviewStruct{}, err
	}

	for rows.Next() {
		err = rows.Scan(&auditorReview.Id, &auditorReview.IdUser, &auditorReview.IdRefuge, &auditorReview.IdRefugeAbility, &auditorReview.Score)
		if err != nil {
			return []structs.AuditorReviewStruct{}, err
		}

		auditorReviewList = append(auditorReviewList, structs.AuditorReviewStruct{Id: auditorReview.Id, IdUser: auditorReview.IdUser, IdRefuge: auditorReview.IdRefuge, IdRefugeAbility: auditorReview.IdRefugeAbility, Score: auditorReview.Score})
	}

	return auditorReviewList, nil
}

func GetById(id int) (structs.AuditorReviewStruct, error) {
	var auditorReview structs.AuditorReviewStruct
	auditorReview.Id = id

	err := dbhandler.GetDatabase().QueryRow(`SELECT idUser, idRefuge, idRefugeAbility, score FROM AuditorReview WHERE id = ?`, id).Scan(&auditorReview.IdUser, &auditorReview.IdRefuge, &auditorReview.IdRefugeAbility, &auditorReview.Score)
	if err != nil {
		return structs.AuditorReviewStruct{}, err
	}

	return auditorReview, nil
}

func UpdateById(auditorReview structs.AuditorReviewStruct) error {
	_, err := dbhandler.GetDatabase().Exec(`UPDATE AuditorReview SET idUser = ?, idRefuge = ?, idRefugeAbility = ?, score = ? WHERE id = ?`, auditorReview.IdUser, auditorReview.IdRefuge, auditorReview.IdRefugeAbility, auditorReview.Score, auditorReview.Id)
	return err
}

func DeleteById(id int) error {
	_, err := dbhandler.GetDatabase().Query(`DELETE FROM AuditorReview WHERE id = ?`, id)
	return err
}
