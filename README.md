# Golang Custom Tags
Criação e validação de tags customizadas para structs em golang, usando o pacote nativo reflect.


## Entidade
A entidade users contém a tag customizada **required**, onde iremos validar posteriormente.
```go
package entities

type User struct {
	Name     string `required:"true"`
	Age      int
	Email    string `required:"true"`
	Password string `required:"true"`
}
```

## Validação
Para a validação utilizaremos o pacote **reflect** do Go, onde é possível identificar os tipos, nomes e valores da struct e seus campos.
```go
package tags

import (
	"fmt"
	"reflect"
)

type CustomTags struct{}

type ICustomTags interface {
	Validate(s interface{}) error
}

// utilizando generics para a struct que validaremos
func (t *CustomTags) Validate(s interface{}) error {
    // tipo da struct (entities.User)
	typeStruct := reflect.TypeOf(s)

    //valor da struct
	valueStruct := reflect.ValueOf(s)

    // iterando struct para passar por todos os campos
	for i := 0; i < typeStruct.NumField(); i++ {
        // nome do campo
		field := typeStruct.Field(i)

        // buscando tag required
		required := field.Tag.Get("required")

        // caso não tenha a flag ou seu valor seja false
		if required == "" || required == "false" {
			continue
		}

        // valor do campo
		value := valueStruct.Field(i)

        // validação
		switch value.Kind() {
		case reflect.String:
            // se for do tipo string, required true, mas não populado corretamente
			if value.String() == "" {
				return fmt.Errorf("O campo %s é obrigatório!", field.Name)
			}

        // se for do tipo int, required true, mas não populado corretamente
		case reflect.Int:
			if value.Int() == 0 {
				return fmt.Errorf("O campo %s é obrigatório!", field.Name)
			}
		}
	}

	return nil
}
```

## Main
```go
package main

import (
	"go_custom_tags/entities"
	"go_custom_tags/pkg/tags"
)

func main() {
	user := entities.User{
		Name: "Wesley Martins",
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
```

Coloque prints ou debugue o método de validação e entenda melhor o que cada operação do reflect retorna:
```
STRUCT:
    TYPE:  entities.User
    VALUE:  {Wesley Martins 0 email_example@gmail.com pass1234}

CAMPOS:
    FIELD:  {Name  string required:"true" 0 [0] false}
    REQUIRED:  true
    FIELD VALUE:  Wesley Martins

    FIELD:  {Age  int  16 [1] false}
    REQUIRED:

    FIELD:  {Email  string required:"true" 24 [2] false}
    REQUIRED:  true
    FIELD VALUE:  email_example@gmail.com

    FIELD:  {Password  string required:"true" 40 [3] false}
    REQUIRED:  true
    FIELD VALUE:  pass1234
```
