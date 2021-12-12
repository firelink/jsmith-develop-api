package argModel

import (
	"encoding/json"
)

type password string

type User struct {
	//Model
	Id        string   `json:"userID" bson:"_id,omitempty"`
	UserName  string   `json:"userName" bson:"userName,omitempty"`
	Password  password `json:"password" bson:"password,omitempty"`
	FirstName string   `json:"firstName" bson:"firstName,omitempty"`
	LastName  string   `json:"lastName" bson:"lastName,omitempty"`
	Email     string   `json:"email" bson:"email,omitempty"`
	Title     string   `json:"title" bson:"title,omitempty"`
	Picture   string   `json:"picture" bson:"picture,omitempty"`
	About     string   `json:"about" bson:"about,omitempty"`
}

func (user *User) New() Entity {
	return &User{}
}

func (user *User) GetID() string {
	return user.Id
}

func (user *User) SetID(Id string) {
	user.Id = Id
}

func (user *User) ToString() string {
	json, err := json.Marshal(user)

	if nil != err {
		return "Unable to marshal User"
	}

	return string(json)
}

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
