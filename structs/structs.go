package structs

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type LoginStruct struct {
	Username string
	Password string
}

func ParseBodyIntoLoginStruct(body io.ReadCloser) (LoginStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	loginStruct := LoginStruct{}
	err := json.Unmarshal(bodyBytes, &loginStruct)
	return loginStruct, err
}

type UserStruct struct {
	Id         int
	Username   string
	Password   string
	RefugeName string
	Email      string
	RoleID     int
	Lat        float64
	Lon        float64
}

func ParseBodyIntoUserStruct(body io.ReadCloser) (UserStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	userStruct := UserStruct{}
	err := json.Unmarshal(bodyBytes, &userStruct)
	return userStruct, err
}

type RoleStruct struct {
	Id   int
	Name string
}

func ParseBodyIntoRoleStruct(body io.ReadCloser) (RoleStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	roleStruct := RoleStruct{}
	err := json.Unmarshal(bodyBytes, &roleStruct)
	return roleStruct, err
}

type PhotoStruct struct {
	Id  int
	Url string
	Lat float64
	Lon float64
}

func ParseBodyIntoPhotoStruct(body io.ReadCloser) (PhotoStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	photoStruct := PhotoStruct{}
	err := json.Unmarshal(bodyBytes, &photoStruct)
	return photoStruct, err
}

type RefugeAbilityStruct struct {
	Id   int
	Name string
}

func ParseBodyIntoRefugeAbilityStruct(body io.ReadCloser) (RefugeAbilityStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	refugeAbilityStruct := RefugeAbilityStruct{}
	err := json.Unmarshal(bodyBytes, &refugeAbilityStruct)
	return refugeAbilityStruct, err
}

type AuditorReviewStruct struct {
	Id              int
	IdUser          int
	IdRefuge        int
	IdRefugeAbility int
	Score           int
}

func ParseBodyIntoAuditorReviewStruct(body io.ReadCloser) (AuditorReviewStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	auditorReviewStruct := AuditorReviewStruct{}
	err := json.Unmarshal(bodyBytes, &auditorReviewStruct)
	return auditorReviewStruct, err
}

type PetStruct struct {
	Id       int
	RefugeId int
	FamilyId int
	Breed    string
	Color    string
	Age      int
}

func ParseBodyIntoPetStruct(body io.ReadCloser) (PetStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	petStruct := PetStruct{}
	err := json.Unmarshal(bodyBytes, &petStruct)
	return petStruct, err
}

type PetPhotoStruct struct {
	Id      int
	IdPet   int
	IdPhoto int
}

func ParseBodyIntoPetPhotoStruct(body io.ReadCloser) (PetPhotoStruct, error) {
	bodyBytes, _ := ioutil.ReadAll(body)
	petPhotoStruct := PetPhotoStruct{}
	err := json.Unmarshal(bodyBytes, &petPhotoStruct)
	return petPhotoStruct, err
}
