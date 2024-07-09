package main

import (
	"fmt"
	"log"

	"github.com/valitovgaziz/v_observer/internal/app/apiserver"
)

func main() {
	fmt.Println("App starts ...")
	s := apiserver.New()
	if err := s.Start(); err!= nil {
		log.Fatal(err)
	}
}