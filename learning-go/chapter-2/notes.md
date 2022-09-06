# Chapter 2 - Primitive Types an Declarations

## The Zero Value
* Go assigns a default zero value to any declares variable that is not assigned a value

## Literals
* There are 4 literals in Go (plus a fifth rare literal)
* Integer literals
    * Sequence of numbers usually in base 10
    * Different prefixes denote other bases: **0b** -> binary, **0o** -> octal, or **0x** -> hexadecimal
    * Go allows underscores in the middle of long literals to make them more readable. (1_234)
* Floating point literals
    * Have decimal point to indicate the fractional portion
    * The can have an exponent specified with an **e** and a positive or negative number, 1.23e45
    * You can use underscores to increase readability
    * You can use **0x** for hexadecimal and **p** for the exponent
* Rune literals
    * represents characters and are surrounded by single quotes
    * Single and double quotes are not interchangeable
    * Can be written as single unicode characters ('a'), 8-bit octal numbers ('\141'), 8-bit hexadecimal numbers (\x61), 16-bit hexadecimal numbers (\u0061) or 32-bit unicode numbers ('\U00000061')
    * Backslash escaped rune literals: ```'\n', '\t', '\'', '\"' and '\\'```
    * Use base ten for number literals -> unless context makes it clearer
    * Octal mostly used for POSIX per mission flag (0o777 -> rtwxtwxtwx)
    * Hexadecimal and binary sometimes used for bit filters and networking/infrastructure applications.

* String literals
    * Interpreted string literal is indicated with double quotes. Contains zero or more runes.
        * Not allowed: Unescaped backlashes, unescaped newlines, and unescaped double quotes
    * Raw string literal is indicated with backquotes (`)
        * Cannot include backquotes
        * Can include backslashes, double quotes and new lines
    
* Go literals are untyped
    * can use integer literal in floating point expression or assign integer literal to a floating point variable.
    * Can interact with any literal that's compatible with the literal
    * Can't assign string variable with numeric type 
    * Can't assign literal number to string variable
    * Can't assign float literal to an int
    * You'll get an compile error when assigning literal whose value overflows the specified variable (1000 literal to variable of type byte)

## Booleans
* bool type represents Boolean variables.
* two values true or false
* zero value of bool is false

## Numeric types
* 12 different types
* Integer types
    * Both signed and unsigned types, from one to eight bytes
    * int8, int16, int32, int64, uint8, uint16, uint32, uint64
    * zero values is 0
    * Special names
        * **byte** is an alias for uint8
            * legal to assign compare, calculations with uint8
            * byte is usually used instead of uint8
        * **int** on 32bit system int is int32 ont 64bit system int is int64
            * It is a compile time error to compare, assign, etc an int and int32/int64 because it depends on the system (without type conversion)
            * integer literals default to being int type
        * **uint** follows same rules as int except it is unsigned
    * **When to use rules**
        * Working with binary files or networking protocol that has an integer of specific size or sign use corresponding integer type
        * Writing library function that should work with any integer type, write a pair of functions, one with int64 and the other with uint64
        * For performance or integration purposes use explicit size and sign
        * All other cases use int
    * Integer operators
        * Supports usual arithmetic ops: +, -, *, / and %
        * Integer division results in an integer -> use type conversion to float
        * Dividing by 0 results in a panic
        * Integer division follows truncation toward zero.
        * Combine arithmetic op with = to modify variable: +=, -=, *=, /= and %=
        * Compare integers with ==, !=, <, > <= and >=
        * Bit shift left and right with << and >> 
        * Bit mask with & (logical AND), | (logical OR), ^ (logical XOR), &^ (logical AND NOT) 
        * bit shift and bit mask can be combined with = 
* Floating point types
    * Two types: float32 and float64
    * zero value is 0
    * Uses IEEE 754 specification for floating point math
    * **Pick float64 unless you need to be compatible with an existing format**
    * Default type for floating point literals is float64
    * Use profiler to determine if difference in memory size is a problem.
    * Store nearest approximation of value in a range.
    * Can only be used where inexact values are acceptable or when rules are clearly understood
    * That limits to graphics and scientific operations
    * Don't use it to represent money
    * You can use all standard mathematical and comparison operators with float except %
        * Dividing nonzero floating point value with 0 results in +Inf or -Inf
        * Dividing floating point variable set to 0 by 0 return NaN
    * **Go Allow use of == and != but don't use it** because of the inexact nature of floats -> define a max variance and see if difference of two floats is less (sometimes called epsilon)
* Complex types
    * You should consider other languages for these types of calculations
    * complex64 uses float32 values to represent the real and imaginary part
    * complex128 uses float64 values
    * use epsilon technique to compare two complex values
    * zero value is 0 for both parts

## String and Runes
* strings are a built-in type
* strings zero value is an empty string
* Supports unicode
* compared for equality with == and !=
* compared for ordering with <, >, <= and >=
* concatenation with + 
* String are immutable -> you reassign the value of a string variable but not modify the string value that is assigned to it
* **rune** is an alias for int32
    * represents single code point
    * use rune type for characters

## Explicit Type Conversion
* Go doesn't allow for automatic type promotion as in other languages
* Must use type conversion -> makes it clear exactly what resulting type should be
* Can't treat other Go types as boolean
* Can't convert other types to bool
    * Need to use comparison operator to create a bool from other types

## **var** *vs* **:=**
* ```var x int = 10``` (Most verbose way)
* ```var x = 10``` (can leave out type since type on righthand side is the expected type)
* ```var x int``` (declares int value and assigns 0)
* ```var x, y int = 1, 2``` (declare multiple values in a single line) 
* ```var x, y int``` (multiple values all assigned zero)
* ```var x, y = 10, "hello"``` (multiple values different types)
* Declaration list for multiple vars
```go
var (
    x int
    y = 10
    z = "yo"
)
```
* ```x := 10``` Short declaration format when inside a function
* ```x , y := 10, "hello"``` Short declaration format with multiple variables
* := Can assign value to existing variable but not var
```go
x := 10
x, y := 20, "hello"
```
* Must use var on package level as := is only allow inside functions
* **Usage Rules**
    * Use := inside functions
        * Not when initializing variable to its zero value
        * Not when assigning untyped constant or literal to variable when their type is not the intended resulting type
        * Idiomatic to write var x byte = 20 rather than x := byte(20)
        * Sometimes creates new variables when you think you are using an existing variable
    * use declaration lists on rare occasions you are declaring multiple package-level variables
    * Only use one-line multi variable declaration when assigning multiple values returned from a function or the "comma ok idiom"
    * Rarely declare variables outside of functions (package block)
        * Changing package-level variables are a bad idea, as it can be hard to track the changes and understand the data flow
        * Only declare immutable variables in package block if possible

## Const
* Const keyword to declare that a value is immutable
* Const in go is a way to give names to literals
    * numeric literals
    * true and false
    * strings
    * runes
    * built-in functions complex, real, imag, len and cap
    * Expressions consisting of operators and preceding values
* Can't declare a runtime calculated value as immutable
    * There are no immutable arrays, slices, maps or structs
* You should declare a group of constants with a set parenthesis
```go
const (
    instrumentKey = "id"
    apiKey = "key"
)
``` 

## Typed and untyped constants
* Untyped works like a literal
* untyped constant has no type of its own but does have a default type which is used when no other type can be inferred. 
```go
// Untyped constant
const x = 10
// legal assignments
var y int  = x
var z float64 = x
```
* Typed constant can only be directly assigned to a variable of that type
```go
// Typed constant can only assigned to variable of type int 
const typedX int = 10
```

## Unused variables
* Every declared local variable must be read
    * results in compile time error if it is not
    * does not stop package level unread variables
* Unused constants are not included in the compiled binary

## Naming variables and constants
* identifier names must start with a letter or underscore
* can contain numbers, letters and underscores
* Any unicode character that is considered a letter or digit is allowed
* Don't name variables with look-alike or special unicode code points as it makes the code nonidiomatic
* underscore characters are rarely used in variable names because idiomatic Go doesn't use snake case (file_name)
* Go uses camel case (fileName) when identifier name consists of multiple words
* Constants are not named with underscore and UPPERCASE letters
* Use short names inside functions
    * k and v for key and value usually in for-range loops 
    * i and j for indexes in standard for loops
* Convention for single letter names and variable types
    * first letter of a variable type for variable name, i for int, f for float, etc
* Single letter variables purpose:
    * to eliminate repetitive typing and keep the code short
    * Serves as a check on the code complexity
        * if it is hard to keep track then the code block might be doing to much
* Use descriptive names when naming constants and variables in the package-level block