package main

import (
	"fmt"

	"github.com/binadel/vali"
	"github.com/mailru/easyjson/jwriter"
)

func main() {
	result := vali.Int("product", "price").Validate(100)

	writer := &jwriter.Writer{}
	result.MarshalEasyJSON(writer)
	json, _ := writer.BuildBytes()
	fmt.Println(string(json))
}
