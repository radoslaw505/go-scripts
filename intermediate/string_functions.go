package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// 🔍 Contains
	fmt.Println("Contains 'world':", strings.Contains("hello world", "world"))

	// 🚪 HasPrefix / HasSuffix
	fmt.Println("HasPrefix 'go':", strings.HasPrefix("golang", "go"))
	fmt.Println("HasSuffix 'lang':", strings.HasSuffix("golang", "lang"))

	// 📍 Index / LastIndex
	fmt.Println("Index of 'l':", strings.Index("hello", "l"))
	fmt.Println("LastIndex of 'l':", strings.LastIndex("hello", "l"))

	// 🔄 Replace
	fmt.Println("Replace 'foo' with 'bar':", strings.Replace("foo foo foo", "foo", "bar", -1))

	// 🍴 Split
	parts := strings.Split("a,b,c", ",")
	fmt.Println("Split:", parts)

	// 🔗 Join
	joined := strings.Join([]string{"a", "b", "c"}, "-")
	fmt.Println("Join:", joined)

	// 🔠 ToUpper / ToLower
	fmt.Println("ToUpper:", strings.ToUpper("golang"))
	fmt.Println("ToLower:", strings.ToLower("GOLANG"))

	// ✂️ Trim / TrimSpace
	fmt.Println("Trim '.':", strings.Trim("...hello...", "."))
	fmt.Println("TrimSpace:", strings.TrimSpace("   hello   "))

	// 🔁 Repeat
	fmt.Println("Repeat 'na' 3 times:", strings.Repeat("na", 3))

	// 🧪 Custom sanitizer example
	input := "   BADword is bad   "
	sanitized := sanitize(input)
	fmt.Println("Sanitized:", sanitized)

	// Builder pattern for complex strings
	user := User{Name: "Radosław", Age: 30}
	var builder strings.Builder
	builder.WriteString("This is a complex message:\n")
	builder.WriteString(fmt.Sprintf("- User: %s\n", user.Name))
	builder.WriteString(fmt.Sprintf("- Age: %d\n", user.Age))
	builder.WriteString("- Status: Learning Go!\n")
	fmt.Println(builder.String())

	// Resetting the builder
	builder.Reset()
	builder.WriteString("Builder reset and reused.")
	fmt.Println(builder.String())
}

// sanitize removes spaces, lowers case, and censors "badword"
func sanitize(input string) string {
	trimmed := strings.TrimSpace(input)
	lowered := strings.ToLower(trimmed)
	return strings.ReplaceAll(lowered, "badword", "***")
}
