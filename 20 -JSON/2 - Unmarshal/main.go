package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJason := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	fmt.Println(c)

	if erro := json.Unmarshal([]byte(cachorroEmJason), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJson := `{"nome":"Toby", "raca":"Poodle"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
