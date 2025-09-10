package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

// Struct with multiple tags
type User struct {
	ID       int    `json:"id" xml:"id" validate:"required"`
	Name     string `json:"name" xml:"name" validate:"required,min=3"`
	Email    string `json:"email,omitempty" xml:"email,omitempty"`
	Password string `json:"-" xml:"-"` // excluded from JSON and XML
	Note     string `custom:"important"`
}

func main() {
	// === JSON Encoding ===
	u := User{ID: 1, Name: "Radosław", Email: "r@example.com", Password: "secret", Note: "VIP"}
	jsonBytes, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("JSON Output:\n", string(jsonBytes))

	// === XML Encoding ===
	xmlBytes, _ := xml.MarshalIndent(u, "", "  ")
	fmt.Println("\nXML Output:\n", string(xmlBytes))

	// === Reflection: Accessing Tags ===
	t := reflect.TypeOf(u)
	fmt.Println("\nStruct Tags:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s\n", field.Name)
		fmt.Printf("  json:   %s\n", field.Tag.Get("json"))
		fmt.Printf("  xml:    %s\n", field.Tag.Get("xml"))
		fmt.Printf("  validate: %s\n", field.Tag.Get("validate"))
		fmt.Printf("  custom: %s\n", field.Tag.Get("custom"))
	}

}
