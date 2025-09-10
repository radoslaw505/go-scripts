package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	raw := "https://www.example.com/search?q=golang"
	parsed, err := url.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", parsed.Scheme)
	fmt.Println("Host:", parsed.Host)
	fmt.Println("Path:", parsed.Path)
	fmt.Println("Query:", parsed.RawQuery)

	// Constructing a URL
	u := &url.URL{
		Scheme: "https",
		Host:   "www.example.com",
		Path:   "/users",
	}
	q := u.Query()
	q.Set("id", "123")
	u.RawQuery = q.Encode()
	fmt.Println("Constructed URL:", u.String())

	// Parsing query string
	values, _ := url.ParseQuery("q=golang&lang=en")
	fmt.Println("Query param 'q':", values.Get("q"))

	// Escaping and unescaping
	escaped := url.QueryEscape("golang & programming")
	fmt.Println("Escaped:", escaped)
	unescaped, _ := url.QueryUnescape(escaped)
	fmt.Println("Unescaped:", unescaped)

	// Resolving reference
	base, _ := url.Parse("https://www.example.com/base/")
	rel, _ := url.Parse("path/to/resource")
	resolved := base.ResolveReference(rel)
	fmt.Println("Resolved URL:", resolved.String())

}
