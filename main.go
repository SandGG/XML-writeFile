package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type teacher struct {
	Name     string     `xml:"name"`
	Age      int        `xml:"age"`
	Salary   int        `xml:"salary"`
	Subjects []subjects `xml:"subjects"`
}

type subjects struct {
	Name  string `xml:"name"`
	Group string `xml:"group"`
}

func main() {
	writeFile()
	readFile()
}

func writeFile() {
	var file, err = os.OpenFile("./files/teachers.xml", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var teachers = []teacher{
		{
			Name:   "Maria Cuevas",
			Age:    35,
			Salary: 190,
			Subjects: []subjects{
				{
					Name:  "Spanish",
					Group: "1A",
				},
			},
		},
		{
			Name:   "Nicolas Chaves",
			Age:    40,
			Salary: 150,
			Subjects: []subjects{
				{
					Name:  "Physical",
					Group: "4B",
				},
			},
		},
		{
			Name:   "Fernanda Suarez",
			Age:    39,
			Salary: 140,
			Subjects: []subjects{
				{
					Name:  "Math",
					Group: "3A",
				},
			},
		},
	}

	var b, errC = xml.MarshalIndent(teachers, " ", " ")
	if errC != nil {
		log.Fatal(errC)
	}
	file.Write(b)

	fmt.Println("File Updated Successfully")
}

func readFile() {
	var file, err = os.OpenFile("./files/teachers.xml", os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var b = make([]byte, 10)
	for {
		n, err := file.Read(b)
		fmt.Print(string(b[:n]))
		if err == io.EOF {
			break
		}
	}
}
