package main

import (
	"fmt"

	"github.com/binadel/vali"
	"github.com/mailru/easyjson"
)

func main() {
	result := vali.Int("product", "price").Min(0).Validate(-100)
	json, _ := easyjson.Marshal(result)
	fmt.Println(string(json))
}
