package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 🔍 MatchString: check if string matches pattern
	re := regexp.MustCompile(`^go\d+`)
	fmt.Println("Match 'go123':", re.MatchString("go123"))   // true
	fmt.Println("Match 'golang':", re.MatchString("golang")) // false

	// 🔎 FindString: extract first match
	reDigits := regexp.MustCompile(`\d+`)
	fmt.Println("First number in 'Age: 42':", reDigits.FindString("Age: 42")) // "42"

	// 🔎 FindAllString: extract all matches
	text := "There are 3 cats, 7 dogs, and 12 birds."
	allNumbers := reDigits.FindAllString(text, -1)
	fmt.Println("All numbers:", allNumbers) // ["3", "7", "12"]

	// ✂️ ReplaceAllString: replace matches
	reSpaces := regexp.MustCompile(`\s+`)
	cleaned := reSpaces.ReplaceAllString("Hello     world!", " ")
	fmt.Println("Cleaned:", cleaned) // "Hello world!"

	// 🧩 FindStringSubmatch: capture groups
	reEmail := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
	email := "Contact: rado@example.com"
	matches := reEmail.FindStringSubmatch(email)
	if len(matches) > 0 {
		fmt.Println("Full match:", matches[0]) // rado@example.com
		fmt.Println("Username:", matches[1])   // rado
		fmt.Println("Domain:", matches[2])     // example.com
	}

	// 🧪 Custom validator
	fmt.Println("Valid email:", isValidEmail("test@go.dev"))    // true
	fmt.Println("Invalid email:", isValidEmail("not-an-email")) // false
}

// isValidEmail checks if a string is a valid email format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
