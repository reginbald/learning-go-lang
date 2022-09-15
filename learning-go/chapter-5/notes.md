# Chapter 5 - Functions

## Declaring and calling functions
* Every program has a main function which is the starting point for every Go program.
    * has no input parameters and no return values
```go
func div(numerator int, denominator int) int {
    if denominator == 0 {
        return 0
    }
    return numerator / denominator
}
```
* function declaration has 4 parts
    * keyword **func**
    * function name
    * input parameters &rarr; in parenthesis, separated with commas, parameter name and followed by type
    * return type &rarr; written after input parenthesis and before function body braces
* **return** keyword is used to return values from functions
    * **return** is not needed if function does not return a value
* Use empty parenthesis when there are no input parameters
* Leave blank when function doesn't return a value
* ```func test(a, b int) int {...}``` &rarr; input parameters can be written like this when the are of the same type

## Simulating named and optional parameters
* Go does not have named and optional input parameters
* You must supply all parameters to a function
* Can be emulated with a struct that has fields that match the desired parameters and pass the struct to the function
* A function shouldn't have many parameters &rarr; should refactor as it might be to complicated

## Variadic Input Parameters and Slices
* variadic parameter must be the last or only parameter
* indicated with three dots ... before type
* The variable that is created is a slice of that type
* ```addTo(3, []int{1, 2, 3}...)``` &rarr; You can supply a slice to the variadic parameter but it must be followed by ...

## Multiple return values
* Go allows multiple return values
* ```func test(a int)(int, int, error) { return a, a, nil}```
* return types are listed in parenthesis separated by commas
* You must return all values listed separated by commas
* You use multiple return value support to return an error from a function, we return nil if everything was successful
    * Convention states that error is always the last or only value returned
* Multiple return values are multiple values &rarr; it's not a tuple you must assign each value returned or you'll get a compile-time error.
* **_** is used if you don't need some of the returned values
    * ```result, _, err := test(10)```
* ```test(10)``` &rarr; implicitly ignores all returned values

## Named Return Values
```go
func test(a, b int) (result int) {
    result = a + b
    return result
}
```
* Supplying names to return values pre-declares variables with those names to be used in the function
* Named return values must be surrounded with parenthesis
* Named return values are local to the function, no enforcement of names outside of the function
* ```func test(a, b int) (result int, _ error)``` &rarr; **_** can be used to only name some return values
* Problems
    * It's possible to shadow a named return value, make sure you are assigning to the return value and not a shadow
    * You don't need to return named return values
        * Go compiler assigns whatever is returned to the return variables
    
## Blank Return - Don't use
* When you have named return values, you can return without any values. This returns last values assigned to the named return values.
* Makes it harder to understand the data flow &rarr; good software is clear and readable 
* If function returns values never use blank returns 

## Functions are values
* The type of the function is built from the keyword **func** and the input parameter types and the return values.
    * Called a *signature* of a function
    * Any function that has the same number and types of input parameters and return values has the same *signature*
* Example: You can have functions in maps
```go
func add(a, b int) int {return a + b}
var opMap = map[string]func(int,int) int { "+", add}
```

## Function Type Declarations
* You can use the **type** keyword to define a function type
```go
type opFuncType func(int, int) int
func add(a, b int) int {return a + b}
var opMap = map[string]opFuncType { "+", add}
```
* Any function that has two int input parameters and int return value type automatically meets the type and can be assigned to the map above
* This is useful for documentation, especially when you are going to refer to it multiple times

## Anonymous functions
* You can assign functions to variables
* You can define new functions within a function and assign them to variables
* Inner functions are anonymous functions
* Don't have a name and you don't have to assign them to variables
* You can write them inline and call them immediately
```go
func main() {
    func(a, b int) int {return a + b}(1, 2)
}
```
* Anonymous functions are declared with the **func** keyword followed by input parameters, return values and the execution body
* Calling Anonymous functions immediately are useful when using the **defer** statement and launching **goroutines**

## Closures
* Functions declared inside of functions are *closures*. 
    * Means that these function can access and modify variables declared in the outer function
* Benefits
    * Limit function's scope, if function is going to be called multiple times by a single function &rarr; inner functions can be used to hide the function
    * Reduces number of declarations at package level &rarr; can make it easier to find unused names
    * Can be passed to functions or returned by functions allows you to use variables within a function outside of the function

## Passing functions as parameters
* Passing a function as input parameter will include it's references to local variables. &rarr; call captured variables
* Useful when performing different operations on same kind of data
```go
numbers := []int{4,3,2,1}
sort.Slice(numbers, func(a, b int) bool {return a > b})
fmt.Println(numbers)
```

## Returning functions from functions
* returned functions have access to local variables &rarr; captured variables
```go
func makeMult(base int) func(int) int {
    return func(factor int) int { return base * factor}
}

twoBase := makeMult(2)
twoBase(2)
```
* Higher order functions &rarr; function has a function for an input or return value

## Defer 
```go
func main() {
    f, err := os.Open("notes.md")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
}
```

* Clean up code is attached to the function with the **defer** keyword
    * Something that must always happen after successful or failed execution of a function
    * Release resources
    * temporary resources &rarr; files, network connections
* **defer** delays the invocation until the surrounding function exits.
* You can defer multiple closures in a function
    * Last in first out order of execution
* Code within defer closures runs after return statement
    * You can supply a function that returns values to defer but there is no way of reading the returned values as it is the last code to run (after return statement)
* function can be supplied to a defer, variables passed into it are evaluated until the closure runs
* **defer** can examine and modify named return values
    * Allows actions to be taken based on an error, such as adding context or rolling back database changes

```go
func InsertIntoDB(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer func(){
        if err == nil {
            err = tx.Commit()
        }
        if err != nil {
            tx.Rollback()
        }
    }()
    _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
    if err != nil {
        return err
    }
    return nil
}
```

* It is a common pattern to return a closure from a function that allocates a resource for clean up. 
```go
func getFile(name string)(*os.File, func(), err) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, err
    }
    return file, func() {file.Close()}, nil
}

main() {
    f, closer, err := getFiles(os.Args[1])
    if err != ni {
        log.Fatal(err)
    }
    defer closer()
}
```

## Go is Call By Value
* When supplying a variable for a parameter to a function, Go will always make a copy of the value in the variable
* This is true for primitive types and structs
* Maps and slices are implemented with pointers so they can be modified
    * Any changes inside a function to a map passed into it will be reflected in the variable passed in
    * You can modify an element in a slice but not lengthen it
* Every type in Go is a value type. It's just that sometimes the value is a pointer
* This is why limited support for constants is a minor issue, you can be sure that calling a function with a variable will not changes that variable (unless it is a map or slice)
* This makes it easier to understand the flow of data in a Go program