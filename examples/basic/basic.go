package main

import (
	"fmt"

	"github.com/binadel/vali"
	"github.com/mailru/easyjson"
)

func main() {
	result := vali.String("product", "title").MaxLength(10).Validate("Some product name")
	json, _ := easyjson.Marshal(result)
	fmt.Println(string(json))

	result = vali.String("user", "email").Pattern(vali.PatternEmail).Validate("somone@somewhere.com")
	json, _ = easyjson.Marshal(result)
	fmt.Println(string(json))
}
