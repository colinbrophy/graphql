package main

import (
	"github.com/neelance/graphql-go"

	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Resolver struct {
}

type UserResolver struct {
	User
}
type User struct {
	Id         graphql.ID
	FirstName string
	LastName  string
	Avatar     string
}

type jsonData struct {
	Data struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
}
type UserInput struct {
	FirstName string
	LastName string
	Avatar string
	}

type Empty struct {}

func (r *Resolver) CreateUser(args struct {User UserInput}) Empty {
	return Empty{}
}

func (r *Resolver) UpdateUser(args struct {User User}) Empty {
	return Empty{}
}
func (r *Resolver) DeleteUser(args struct {UserId graphql.ID}) Empty {
	return Empty{}
}


func (r *Resolver) Users() (usr []User) {
	users, _ := http.Get("reqres.in/api/users")
	jsonStr, _ := ioutil.ReadAll(users.Body)
	fmt.Print(jsonStr)
	json.Unmarshal(jsonStr, &usr)
	return
}

func (u UserResolver) ID() graphql.ID {
	return u.Id
}

func (u UserResolver) First_name() string {
	return u.FirstName
}

func (u UserResolver) Last_name() string {
	return u.LastName
}

func (u UserResolver) Avatar() string {
	return u.User.Avatar
}

func (r *Resolver) User(args struct {Id graphql.ID}) UserResolver {

	users, err := http.Get("https://reqres.in/api/users/" + string(args.Id))
	if err != nil {
		panic(err)
	}
	jsonStr, _ := ioutil.ReadAll(users.Body)
	var usr jsonData
	err = json.Unmarshal(jsonStr, &usr)
	if err != nil {
		panic(err)
	}

	return UserResolver{User{graphql.ID(usr.Data.ID), usr.Data.FirstName,
		usr.Data.LastName, usr.Data.Avatar}}
}
