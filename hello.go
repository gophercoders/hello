// This is a comment
// Go ignores these, the godoc tool uses then and humans read then! :)
// Comments are important.

// This program is called 'hello' it prints a message to the terminal
// and draws the go gopher.

// This is the main package.
// Go looks in this file for the main function.
package main

// Import the fmt package.
// The fmt package contains functions to print to the terminal.
// "fmt" is short for format.
import "fmt"

// The main fuction.
// The computer starts to run the program at the main function.
// The program runs from the opening brace, the {, until closing brace,
// the }, in the last line.
// func is short for function.
// () means there are no parameters. We'll explain these later!
func main() {
	// Now we want to print something to the terminal.
	// To do this we need to use the function called Println from the fmt package.
	fmt.Println("Hello Gophers!")
	// You read fmt.Println bit as
	// "In the package called fmt use the function called Println".
	fmt.Println("and now for the clever bit...")
	// Anything in-between the "'s is printed to the screen.
	// The bit between the quote marks is called a "string literal".
	// The characters \ and " are both special. We have to put a \
	// in front of them in order to print them.
	// If we did not do this we would have a syntax error
	fmt.Println("         ,_---~~~~~----._")
	fmt.Println("  _,,_,*^____       ____``*g*\\\"*,")
	fmt.Println(" / __/  /'    ^.   /     \\ ^@q   f")
	fmt.Println("[  @f  | @))   |   | @))  l   0 _/")
	fmt.Println(" \\`/    \\~____ / __ \\_____/   \\")
	fmt.Println(" |             _l__l_         I")
	fmt.Println(" }            [______]         I")
	fmt.Println(" ]              | | |          |")
	fmt.Println(" ]               ~ ~           |")
	fmt.Println(" |                            |")
	fmt.Println("  |                           |")
} // The program stops, or exits here.
