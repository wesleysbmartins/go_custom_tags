package entities

type User struct {
	Name     string `required:"true"`
	Age      int
	Email    string `required:"true"`
	Password string `required:"true"`
}
