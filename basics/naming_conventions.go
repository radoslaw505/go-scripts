package main

import "fmt"

func main() {

	// Variable Naming Conventions in Go
	var userName string = "john_doe" // camelCase for multi-word variable names
	var IsActive bool = true         // Start with uppercase for exported variables

	// Constants
	const MAX_RETRIES int = 5 // ALL_CAPS with underscores for constants

	// Struct Naming
	type UserProfile struct { // CamelCase for struct names
		Name  string
		Email string
	}

	// Acronyms
	type HTTPServer struct { // Treat acronyms as words
		Address string
		Port    int
	}

	// Avoiding Stuttering
	// If the package is "user", use "Profile" instead of "UserProfile"
	var Profile UserProfile

	// Print the variables to avoid unused variable errors
	fmt.Println(userName, IsActive, MAX_RETRIES, Profile)

}

// ------------------------

// Naming Conventions in Go
// ------------------------

// 1. Package Names
// - Should be short, lowercase, and meaningful.
// - Avoid underscores and mixed caps.
// - Example: "fmt", "http", "json"

// 2. Variable Names
// - Use camelCase for multi-word variable names.
// - Start with a lowercase letter for package-level variables (unexported).
// - Start with an uppercase letter for exported variables.
// - Example: userName, userAge, IsActive

// 3. Function Names
// - Follow the same conventions as variable names.
// - Example: calculateSum, GetUserInfo

// 4. Constants
// - Use ALL_CAPS with underscores for multi-word constants.
// - Example: MAX_RETRIES, DEFAULT_TIMEOUT

// 5. Structs and Interfaces
// - Use CamelCase for struct and interface names.
// - Example: UserProfile, DataFetcher

// 6. Acronyms
// - Treat acronyms as words.
// - Example: HTTPServer, URLParser

// 7. Avoiding Stuttering
// - Do not repeat the package name in the variable or function name.
// - Example: If the package is "user", use "Profile" instead of "UserProfile".

// 8. Comments
// - Use comments to explain the purpose of variables, functions, and packages.
// - Follow the format: // FunctionName does something...
