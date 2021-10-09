package students

type StudentsInput struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
	Age  int    `json:"age" binding:"required"`
}
