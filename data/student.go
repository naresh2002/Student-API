package data

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required,nameValidator"`
	Age       int    `json:"age" validate:"required,gt=0"`
	Email     string `json:"email"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

// var StudentsList = map[int]Student{}

var StudentsList = map[int]Student{
	1: {
		ID:        1,
		Name:      "Naresh",
		Age:       22,
		Email:     "naresh6436@gmail.com",
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	},
}

func GetStudents() map[int]Student {
	return StudentsList
}

func GetNextID() int {
	maxID := 0
	for id := range StudentsList {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

func (s *Student) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("nameValidator", validateName)

	return validate.Struct(s)
}

func validateName(fl validator.FieldLevel) bool {
	// // Name should be of format : {first_name} then "( ith_name)" where i = 2,3,4....
	// re := regexp.MustCompile(`^[a-zA-Z]+([ ][a-zA-Z]+)*$`)

	// Name can only be consist of atleast one of {[A-Z], [a-z] or ' '}
	re := regexp.MustCompile(`^[a-zA-Z ]+$`)
	matches := re.FindAllString(fl.Field().String(), -1)
	fmt.Println(fl.Field().String())

	return len(matches) >= 1
}
