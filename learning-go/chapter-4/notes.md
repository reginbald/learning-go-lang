# Chapter 4 - Blocks, Shadows and Control Structures

## Blocks
* Variables, constants, types and functions declared outside of any functions are placed in the **package block** .
* **import** statements define names for other packages that are file that contains the **import** statement. These names are in the **file block**
* All variables defined at the top level of a function (including params) are in a block
* Every set of {} defines another block within a function
* Control structures define their own block using {}
* Identifier defined in any outer block can be access within any inner block.
* **Universe block** is a block that contains all other blocks and contains predeclared identifiers &rarr; built-in types (int and string), constants (true and false) and functions (make and close) 
    * They can therefor be shadowed, linter and shadow linter don't catch it

## Shadowing variables
* Declaration of an identifier in an inner block with same name as an identifier in an outer block &rarr; shadows outer block identifier
* A shadowing variable is a variable that has the same name as a variable in a containing block.
* As long as a shadowing variable exists, you cannot access a shadowed variable.
* **:=** can make is easy to accidentally shadow a variable, that is when we are creating and assigning to multiple variables
    * **:=** only reuses variables that are declared in the current block
* Package imports can be shadowed as well.
```go
func main() {
    x := 10
    fmt.Println(x)
    fmt := "oops"
    fmt.Println(fmt) // results in an error
}
```

### Detecting shadow variables
* go vet and golangcli-lint do not detect shadowing
* Shadow linter can detect shadowing during build process

```
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

## if
* Much like if statements in most programming languages
* You don't put parenthesis around the condition
* variables declared within the if block only exists there
* variables declared in the condition exists in the if else block
    * lets you create variables only where they are needed
    * Shadowing can happen here as well
```go
if n := rand.Intn(10); n == 0 {
    fmt.Println("hello")
}
```

## for, Four ways
* Only looping keyword
* Four formats
    * Complete C-style for statement
    * Condition-only for statement
    * Infinite for statement
    * for-range statement

### Complete for statement
* No parenthesis
* 3 parts
    * initialization of variables (must use **:=**)
        * variables can be shadowed
    * comparison &rarr; expression that evaluates to bool, checked before body is run and must be evaluated to true for the body to run
    * increment, runs immediately after each iteration of the loop, before condition is evaluated
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### Condition-only for statement
* can leave out initialization and increment parts
* functions like while statement in C
```go
i := 1
for i < 100 {
    fmt.Println(i)
    i = i * 2
}
```

### Infinite for statement
* Leave out all parts
* loops forever
* **break** keyword exists the loop immediately
* **continue** keyword skip rest of body and proceeds to next iteration
    * go encourages short if statements, as left-aligned as possible &rarr; makes code more readable
```go
for {
    fmt.Println("Hello")
}
```

### for-range statement
* for iterating over elements in some of Go's built-in types
    * Can only use with built-in compound types or user-defined types based on them
* Resembles the iterators in other languages
* Slices &rarr; get index and value from for range loops
    * index usually gets i as the name for slices and arrays while k for maps
    * value usually gets v as the name
    * Go requires all declared variables to be accessed so you can use **_** to ignore index
        * Use **_** anytime you'd like to ignore a returned value
```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
    fmt.Println(i, v)
}
```
* Maps
    * Output of iterating over a map varies &rarr; security feature, old Go versions usually return the same each run
        * People would write code that assumed the order was fixed
        * If maps hash items to exact same values you can slow down a server with a Hash DoS attack.
        * Map implementation generates a random number every time a map is created and iteration over map varies a bit each time it is looped over &rarr; to combat the above issues
```go
m := map[string]int {"a": 1, "c": 3, "b": 2}
for i := 0; i <3; i++ {
    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
* Strings &rarr; for-range iterates over the runes of a string not bytes
    * multibyte runes in a string are converted from UTF-8 representation to a single 32-bit number
    * bytes that are not valid UTF-8 values are replaced with the unicode character 0xfffd
```go
samples := []string{"hello", "apple_π"}
	for _, s := range samples {
		for i, r := range s {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
```
* for-range value is a copy
    * when iterating over compound type it copies the value to the value variable
    * modifying the value variable will not modify the value in the compound type

### Labeling for statements
* **break** and **continue** keywords apply to the for loop that directly contains them. You can use labels to target an outer loop
* rarely used
```go
func main() {
    samples := []string{"hello", "world"}
outer:
    for _, s := range samples {
        for _, r := range s {
            if r == 'l' {
                continue outer
            }
        }
    }
}

```

### for loop statement chose
* Most of the time for-range loops will be used
    * Works well built-in compound types (strings, slices, maps and channels)
* Complete for loop statement is best when you are not iterating from the beginning to the end
    * You need to use for-range loop for strings since for loop doesn't handle multibyte characters
* Conditional for loop is not frequently used but is good for looping based on calculated value
* Infinite for loop is useful in some situations.

## Switch
* More useful version than in other languages
* No parenthesis are used
* Can declare variable scoped to all branches of the switch statement
* case and default clauses are inside the switches braces
    * their bodies are not surrounded by braces
    * Body can have multiple lines and is one block
* Variables declared inside case blocks are only visible there
* There are no **break** keywords used
    * by default cases don't fall through
        * you can use the **fallthrough** keyword to achieve this
            * Try to restructure your algorithm rather than using it
    * You can use **break** inside switch but consider refactoring 
        * Could be useful to break a for loop with a label
* Multiple matches are separated by comma
* Empty case means that nothing happens
* You can switch on any type that can be compared with **==**
    * Everything except slices, maps, channels, functions and structs with the aforementioned types

```go
switch size := len(word); size {
case 1, 2, 3, 4:
    fmt.Println(word, "is a short word!")
case 5:
    wordLen := len(word)
    fmt.Println(word, "is exactly the perfect length:", wordLen)
case 6, 7, 8, 9:
default:
    fmt.Println(word, "is a long word!")
}
```

### Blank Switches
* You can declare a switch statement without a value to compare to
* Allows you to use any boolean comparison for each case
* Can include an optional variable declaration
* Use expression switch if all cases are equality comparisons to same variable
```go
switch wordLen := len(word); {
case wordLen < 5:
    fmt.Println(word, "is a short word!")
case wordLen > 10:
    fmt.Println(word, "is a long word!")
default:
    fmt.Println(word, "is exactly the right length")
}
```

## IF or switch
* Switch indicates that there is a relationship between th values or comparisons in each case
* Use if/else when doing many unrelated comparisons
* Favour blank switch statements over if/else chains when you have multiple related cases &rarr; code becomes more readable.

## goto
* Likely that you will never use this control statement
    * Try to avoid it
* **goto** statement specifies a labeled line of code and execution jumps to it.
* Go forbids jumps that skip over variable declarations, to inner or parallel blocks
* Try to use rather labeled break and continue to jump out of nested loops or skip iterations
* Few cases where goto makes the code clearer, if we don't want to run middle of a function but we do want to run the end
    * could use boolean flags but makes code more verbose  and less readable
    * could use duplicate code but it makes code harder to maintain