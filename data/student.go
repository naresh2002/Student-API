package data

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

var StudentsList = map[int]Student{}

// var StudentsList = map[int]Student{
// 	1: {
// 		ID:        1,
// 		Name:      "name1",
// 		Age:       22,
// 		Email:     "name1@gmail.com",
// 		CreatedAt: time.Now().String(),
// 		UpdatedAt: time.Now().String(),
// 	},
// }

func GetStudents() map[int]Student {
	return StudentsList
}
