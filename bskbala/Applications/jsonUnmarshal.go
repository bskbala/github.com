package main

import (
	"encoding/json"
	"fmt"
)

type person []struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main ()  {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Println("%T\n",s)
	fmt.Println("%T\n",bs)

	var people [] person

	err := json.Unmarshal(bs, &people)
	if err !=nil {
		  fmt.Println(err)
	}

	fmt.Println("\nall of data",people)

	for i,v :=range people {
		fmt.Println("\n Person Number" ,i)
		fmt.Println(v.First, v.last, v.age)
	}
}