 //---------------------------Blocks---------------------------
 
/** 
Each place where a declaration occurs is called a block. Variables, constants, types,
 and functions declared outside of any functions are placed in the package block.

 Within a function, every set of braces ({}) defines another block

   
*/

//Shadowing Variables
// Example : Shadowing variables
func main() {
x := 10
if x > 5 {
fmt.Println(x) //10
x := 5 //shadowing variable ,this shadowing only exists within this block
fmt.Println(x) //5
}
fmt.Println(x) //10
}

//Example : Shadowing with multiple assignment
func main() {
	x := 10
	if x > 5 {
	x, y := 5, 20 //:= only reuses variables that are declared in the current block
	fmt.Println(x, y) //When using :=, make sure that you don’t have any variables from
	//                 an outer scope on the lefthand side, unless you intend to shadow them.
	}
	fmt.Println(x)
	}
	Running this code gives you:
	5 20
	10
//You also need to be careful to ensure that you don’t shadow a package import. As in you declare
//  and assign fmt a new value which will make fmt package to be unusable for the rest of the main function


//Detecting Shadowed Variables
/*
Install shadow linter 

		If you are building with a Makefile, consider including shadow in the vet task:
		vet:
		go vet ./...
		shadow ./...
		.PHONY:vet


*/

//Universal block
//GO only has 25 keywords because it considers the built-in types (like int and string), 
// constants (like true and false), and functions (like make
 // or close) ,nil to be predeclared identifiers rather than a keyword. predeclared identifiers are placed in universal block
//  which is the block that contains all other blocks

//*****************Not even shadow detects shadowing of universe block identifiers***************


//---------if---------
// you don’t put parenthesis around the condition
// What Go adds is the ability to declare vari‐
// ables that are scoped to the condition and to both the if and else blocks

//Example  Scoping a variable to an if statement
if n := rand.Intn(10); n == 0 {
	fmt.Println("That's too low")
	} else if n > 5 {
	fmt.Println("That's too big:", n)
	} else {
	fmt.Println("That's a good number:", n)
	}


//for, Four Ways
//What makes
// Go different from other languages is that for is the only looping keyword in the lan‐
// guage. Go accomplishes this by using the for keyword in four different formats:

/*
• A complete, C-style for
• A condition-only for
• An infinite for
• for-range
*/

//• A complete, C-style for

	// Example 4-8. A complete for statement
	for i := 0; i < 10; i++ {
	fmt.Println(i)
	}
	//Rules to remember
	 // you must use := to initialize the variables
	 // just like variable declarations in if statements, you can shadow a variable here.

// • A condition-only for
// Example 4-9. A condition-only for statement  (like while statement)
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}


//• An infinite for
// Example 4
package main
import "fmt"
func main() {
	for {
	fmt.Println("Hello")
	}
}

//go version of do-while to run at least once 
	do {
	// things to do in the loop
	} while (CONDITION);
	// The Go version looks like this:
	for {
	// things to do in the loop
		if !CONDITION {
			break
			}
		}

//continue 
// Example  Using continue to make code clearer
for i := 1; i <= 100; i++ {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("FizzBuzz")
		continue
	}
	if i%3 == 0 {
		fmt.Println("Fizz")
		continue
	}
	if i%5 == 0 {
		fmt.Println("Buzz")
		continue
	}
}



// • for-range
// to iterate over elements in some of Go’s built-in types. like strings,
// arrays, slices, and maps.

// Example The for-range loop
evenVals := []int{2, 4, 6, 8, 10, 12}

for i, v := range evenVals {
		fmt.Println(i, v)
	}

// Running this code produces the following output:
0 2
1 4
2 6
3 8
4 10
5 12
