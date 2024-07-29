package main

import (
	"go_custom_tags/entities"
	"go_custom_tags/pkg/tags"
)

func main() {
	user := entities.User{
		Name:     "Wesley Martins",
		Age:      25,
		Email:    "email_example@gmail.com",
		Password: "pass1234",
	}

	tags := tags.CustomTags{}

	err := tags.Validate(user)
	if err != nil {
		panic(err)
	}
}
