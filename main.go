package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateParams struct {
	Username     string `json:"username" validate:"required"`
	Guests       Guests `json:"guests"`
	RoomType     string `json:"roomType" validate:"required"`
	CheckinDate  string `json:"checkinDate"`
	CheckoutDate string `json:"checkoutDate"`
}

type Guests struct {
	Person []Person `json:"person"`
}

type Person struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
}

func main() {
	v := CreateParams{
		Guests: Guests{
			Person: []Person{
				{Firstname: "foo", Lastname: "bar"},
				{Lastname: "bar"},
			},
		},
	}
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		fmt.Printf("validation1:\n%+v\n\n", err)
	}

	for _, p := range v.Guests.Person {
		if err := validate.Struct(p); err != nil {
			fmt.Printf("validation2:\n%+v\n\n", err)
		}
	}
}
