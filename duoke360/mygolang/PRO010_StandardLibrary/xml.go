package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person1 struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func testMarshal_xml() {
	p := Person1{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	// b, _ := xml.Marshal(p)
	// 有缩进格式
	b, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Printf("%v\n", string(b))
}

func testMarshal_xml2() {
	s := `
	<person>
	<name>zhu</name>
	<age>3</age>
	<email>zhu@gmail.com</email>
	</person>
	`

	b := []byte(s)
	var per Person1
	err := xml.Unmarshal(b, &per)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("per:%v\n", per)
}

func read() {
	/*
	   <?xml version="1.0" encoding="UTF-8"?>
	   <person>
	      <name>tom</name>
	      <age>20</age>
	      <email>tom@gmail.com</email>
	   </person>
	*/
	b, _ := ioutil.ReadFile("a.xml")
	var p Person
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func write() {
	type Person struct {
		XMLName xml.Name `xml:"person"`
		Name    string   `xml:"name"`
		Age     int      `xml:"age"`
		Email   string   `xml:"email"`
	}

	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	f, _ := os.OpenFile("a.xml", os.O_WRONLY, 0777)
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Encode(p)
}


func main() {
	testMarshal_xml()

	testMarshal_xml2()
}
