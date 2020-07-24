package main

import (
	"log"
	"tour/cmd"
	_ "tour/internal/word"
)

func main()  {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v",err)
	}

}
