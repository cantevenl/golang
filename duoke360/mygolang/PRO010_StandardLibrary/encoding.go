package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func testMarshal() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func testUnmarshal() {
	b1 := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var m Person
	err := json.Unmarshal(b1, &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m的类型%T\n m: %v\n",m, m)
}

func test10() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["big tom", "kite"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("f: %v\n", f)
}

func test11() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test12() {
	// dec := json.NewDecoder(os.Stdin)
	// a.json : {"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["tom", "kite"]}
	f, _ := os.Open("a.json")
	defer f.Close()
	dec := json.NewDecoder(f)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("v: %v\n", v)
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}

	/*
	   输入 {"Name":"tom","Age":20,"Email":"tom@gmail.com"}
	   输出
	   v: map[Age:20 Email:tom@gmail.com Name:tom]
	   {"Age":20,"Email":"tom@gmail.com","Name":"tom"}
	*/
}





func main() {
	//testMarshal()
	//testUnmarshal()
	//
	//test10()
	//test11()

	test12()
}
