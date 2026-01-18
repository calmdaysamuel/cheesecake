package main

import (
	"fmt"

	"github.com/calmdaysamuel/cheesecake/application"
	"github.com/calmdaysamuel/cheesecake/widgets/textfield"
)

func main() {
	fmt.Println("yup")
	_ = application.Start(textfield.New())
}
