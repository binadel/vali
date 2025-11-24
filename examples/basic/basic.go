package main

import (
	"fmt"

	"github.com/binadel/vali"
	"github.com/google/uuid"
	"github.com/mailru/easyjson"
)

func main() {
	result := vali.String("product", "title").MaxLength(10).Validate("Some product name")
	json, _ := easyjson.Marshal(result)
	fmt.Println(string(json))

	result2 := vali.String("user", "email").Pattern(vali.PatternEmail).Email().Validate("somone@somewhere.com")
	json, _ = easyjson.Marshal(result2)
	fmt.Println(string(json))

	id := uuid.New().String()
	result3 := vali.String("user", "id").UUID().Version(7).Validate(id)
	json, _ = easyjson.Marshal(result3)
	fmt.Println(string(json))

	result4 := vali.String("user", "ip").IP().Version4().Validate("2f::1")
	json, _ = easyjson.Marshal(result4)
	fmt.Println(string(json))
}
