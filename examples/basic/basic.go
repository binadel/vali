package main

import (
	"fmt"

	"github.com/binadel/vali"
	"github.com/mailru/easyjson"
)

func main() {
	result := vali.Int("product", "price").ExclusiveMinimum(0).MultipleOf(50).Validate(1111)
	json, _ := easyjson.Marshal(result)
	fmt.Println(string(json))
}
