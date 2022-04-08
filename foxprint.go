package main

import (
	"fmt"
	"io/ioutil"
)

func ShowFox(file string) {
	foxes_text, err := ioutil.ReadFile("./" + file)

	if err != nil {
		return
	}

	fmt.Println(string(foxes_text))
}
