package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// Five important questions

// 1) How do we run the code in our project?
// -> go run main.go
// More explicitly, with the go CLI Tool
// go 	| run 		| compiles and executes
// 		| build 	| compile and generate an executable, but does not execute
// 		| fmt		| format the code
// 		| get		| get packages from other devs
// 		| install	| install packages from other devs
// 		| test		| execute tests

// 2) What does the line "package main" mean
// Package: collection of common source code files.
// 			one app usually makes one package
// first line must declare the package a file belongs to

// types of packages
// 1) executable -> compilation creates an executable
// 					package name "main" is used specifically for executables
// 					MUST have a function named "main"
// 2) reusable -> code dependencies, libraries with code helpers
// 					any name different from "main" defines reusable code
// 					compilation generates NOTHING!

// 3) What does "import fmt" mean?
// "import" provides my package with access to a different package and all of it's functionality

// 4) What's that func thing?
// function declaration with a name, arguments and body

// 5) How is the main.go file organized?
// Package
// import declarations
// functionality
