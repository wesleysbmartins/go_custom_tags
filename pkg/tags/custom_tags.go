package tags

import (
	"fmt"
	"reflect"
)

type CustomTags struct{}

type ICustomTags interface {
	Validate(s interface{}) error
}

func (t *CustomTags) Validate(s interface{}) error {
	typeStruct := reflect.TypeOf(s)
	valueStruct := reflect.ValueOf(s)

	fmt.Println("TYPE: ", typeStruct)
	fmt.Println("VALUE: ", valueStruct)

	for i := 0; i < typeStruct.NumField(); i++ {
		field := typeStruct.Field(i)
		fmt.Println("FIELD: ", field)

		required := field.Tag.Get("required")
		fmt.Println("REQUIRED: ", required)

		if required == "" || required == "false" {
			continue
		}

		value := valueStruct.Field(i)
		fmt.Println("FIELD VALUE: ", value)

		switch value.Kind() {
		case reflect.String:
			if value.String() == "" {
				return fmt.Errorf("O campo %s é obrigatório!", field.Name)
			}

		case reflect.Int:
			if value.Int() == 0 {
				return fmt.Errorf("O campo %s é obrigatório!", field.Name)
			}
		}
	}

	return nil
}
