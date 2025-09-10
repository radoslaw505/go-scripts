package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	// Encode
	p := Person{Name: "Radosław", Age: 30}
	xmlBytes, _ := xml.MarshalIndent(p, "", "  ")
	fmt.Println("Encoded XML:\n", string(xmlBytes))

	// Decode
	xmlStr := `<person><name>Radosław</name><age>30</age></person>`
	var decoded Person
	xml.Unmarshal([]byte(xmlStr), &decoded)
	fmt.Println("Decoded struct:", decoded)

	// Handle XML with attributes
	type Book struct {
		XMLName xml.Name `xml:"book"`
		Title   string   `xml:"title,attr"`
		Author  string   `xml:"author"`
	}
	bookXML := `<book title="Go Programming"><author>Radosław</author></book>`
	var book Book
	xml.Unmarshal([]byte(bookXML), &book)
	fmt.Println("Book struct:", book)

	// Encode Book struct
	bookStruct := Book{Title: "Learning Go", Author: "Radosław"}
	bookBytes, _ := xml.MarshalIndent(bookStruct, "", "  ")
	fmt.Println("Encoded Book XML:\n", string(bookBytes))

	// Handle nested XML
	type Library struct {
		XMLName xml.Name `xml:"library"`
		Books   []Book   `xml:"book"`
	}
	libraryXML := `<library>
		<book title="Go Programming"><author>Radosław</author></book>
		<book title="Advanced Go"><author>Alice</author></book>
	</library>`
	var library Library
	xml.Unmarshal([]byte(libraryXML), &library)
	fmt.Println("Library struct:", library)

	// Encode Library struct
	libraryStruct := Library{
		Books: []Book{
			{Title: "Go in Action", Author: "Bob"},
			{Title: "Concurrency in Go", Author: "Charlie"},
		},
	}
	libBytes, _ := xml.MarshalIndent(libraryStruct, "", "  ")
	fmt.Println("Encoded Library XML:\n", string(libBytes))

	// Write XML to file
	file, _ := os.Create("person.xml")
	defer file.Close()
	file.Write(xmlBytes)
	fmt.Println("XML written to person.xml")

	// Read XML from file
	data, _ := os.ReadFile("person.xml")
	var filePerson Person
	xml.Unmarshal(data, &filePerson)
	fmt.Println("Read from file:", filePerson)
	fmt.Println("XML read from person.xml:\n", string(data))

	// Handle XML with namespaces
	type Note struct {
		XMLName xml.Name `xml:"ns:note"`
		To      string   `xml:"ns:to"`
		From    string   `xml:"ns:from"`
		Body    string   `xml:"ns:body"`
	}
	noteXML := `<ns:note xmlns:ns="http://example.com/ns">
		<ns:to>Alice</ns:to>
		<ns:from>Radosław</ns:from>
		<ns:body>Hello!</ns:body>
	</ns:note>`
	var note Note
	xml.Unmarshal([]byte(noteXML), &note)
	fmt.Println("Note struct with namespace:", note)

	noteStruct := Note{To: "Bob", From: "Charlie", Body: "Hi!"}
	noteBytes, _ := xml.MarshalIndent(noteStruct, "", "  ")
	fmt.Println("Encoded Note XML with namespace:\n", string(noteBytes))

	// Handle XML with comments
	commentedXML := `<!-- This is a comment -->
	<person>
		<name>Radosław</name>
		<age>30</age>
	</person>`
	var commentedPerson Person
	xml.Unmarshal([]byte(commentedXML), &commentedPerson)
	fmt.Println("Person struct from XML with comment:", commentedPerson)

}
