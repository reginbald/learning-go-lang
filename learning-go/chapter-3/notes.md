# Chapter 3 - Composite Types

## Array
* Rarely used in Go directly
* Only one-dimensional arrays but can simulate multidimensional arrays
* Can't read/write past end of array or use negative index* Go considers the size of the array to be part of the type of the array
    * ```[2]int``` and ```[3]int``` are not the same
    * Can't use variable for size of array -> types are resolved at compile time
    * Can't use type conversion to convert arrays of different sizes to identical types
        * Can't write functions for arrays of any sizes
        * Can't assign arrays of different sizes to the same variable
* Don't use arrays unless you know the exact length you need.
* Mainly exist to provide baking store for **slices**.
* ```var x [10]int``` &rarr; creates an array of 10 ints with all values as 0
* ```var x = [2]int{1,2}``` &rarr; use array literal to initialize values
* ```var x = [10]int{1, 2:2, 8:10}``` &rarr; sparse array, can use index with values rest is set to 0 (index:value)
* ```var x = [...]int{1,2,3}```&rarr; can leave size when using array literal
* == and != can be used to compare arrays
* ```var x [2][4]int``` &rarr; simulate multidimensional arrays, array of length 2 with type array of ints of length 4
* ```len(x)``` &rarr; returns length of array x
* ```x[0] = 10``` &rarr; bracket syntax to write
* ```x[0]``` &rarr; bracket syntax to read

## Slices
* Most of the time use slices to hold sequence of values
* Length is not a part of the type of a slice
* Can write a function that processes slices of any size, unlike arrays
* Can't read/write past end oof slice or use negative index
* Slice type is not comparable, == results in compile-time error
    * can only compare a slice with **nil**
* ``` var x = []int{1,2,3}``` &rarr; slice literal creates a slice of 3 ints, (no size specified)
    * [...] makes an array - [] makes a slice
* ```var x = []int{1, 2:2, 8:10}``` &rarr; creates a sparse slice (index:value)
* ```var x [][]int``` &rarr; simulates multidimensional slice of slices
* ```x[0] = 10``` &rarr; bracket syntax to write
* ```x[0]``` &rarr; bracket syntax to read
* ```var x []int``` &rarr; slice declaration with literal, creates a slice of ints with zero value which is **nil** (**nil** slice contains nothing)
    * **nil** is different from null in other languages
    * **nil** is an identifier that represent the lack of value for some types
    * **nil** has no type so it can be compared against values of different types
* ```var x = []int{}``` &rarr; creates a non-nil zero-length slice, works identically to nil slice. Only scenario where a zero-length slice is useful is when converting a slice to JSON. 

## len
* len() built in function works for both arrays and slices.
* len() works as well for maps and strings
* len() for a nil slice returns 0
* len() is a built-in function since you can't develop a function that can do what built-in functions can

## append
* built-in function to grow slices.
* Takes two parameters, slice of any type and value of that type.
* returns a slice of same type
```go 
var x = []int{1,2}
x = append(x, 3) 
```
* slice can be append to another with the **...** operator to expand source slice into individual values.
```go
var y = []int{4,5,6}
x = append(x, y...)
```
* You'll get an compile-time error if result of append is not assigned
    * Go is a call by value language
    * Go makes a copy of parameter values every time you pass it to a function
    * append adds values to the copy and returns the copy
* when appending to a slice causes the length to reach capacity size the go runtime will allocate a new slice with larger capacity and copy the original slice to the new slice and the new values are added to the end and the new slice returned.
* It's time consuming to increase capacity so the Go runtime will double capacity while less than 1024 and the 25% after that.

## Capacity
* Each item in a slice is assigned to consecutive memory locations
* Capacity is the number of consecutive memory locations reserved by a slice
* Capacity can be larger than the length of a slice
* cap() returns the current capacity of a slice
* Infrequently used, mainly to check if slice has capacity to hold new data or if call to make is needed to create a new slice
* cap() will alway return the length for an array

## Go runtime
* Set of libraries to enable programs written in Go to run 
* Provides memory allocation and garbage collection, concurrency suppot, networking and built-in types and functions
* The runtime is compiled into every binary
    * As appose to Java and others that use virtual machines that need to be installed separately
    * makes distribution easier
    * Avoids worries about compatibility between runtime and program

## make
* allows us to create a slice that is empty yet has a length or capacity specified.
    * compared to creating a slice with a slice literal or nil. 
* **make** allows us to specify type, length and optionally capacity.
* ```x := make([]int, 5)``` &rarr; int slice with len and cap of 5 with all values set to 0.
* ``` x := make([]int, 0, 10)``` &rarr; int slice with length 0 and capacity 10
    * We can't directly index into the slice since it a non-nil slice with length 0
    * But we can append values to it &rarr;```x = append(x, 5,6,7,8)```
* You'll get a compile-time error if you specify a length that is greater then the capacity with a constant or numeric literal, you'll get a runtime panic if the length is stored in a variable and is greater then the capacity

## Declaring a slice
* Goal is to minimize the amount of times the slice needs to grow
1. If there is a possibility that the slice doesn't need to grow, use a **var** declaration with no assignment (creates a **nil** slice)
    * ```var x []int```
2. If you starting values &rarr; use slice literal 
    * ```x := []int{1,2,3}```
3. If a slices values aren't going to change &rarr; use slice literal
    * ```x := []int{1,2,3}```
4. If you know how large slice should be but not the values &rarr; use **make**
    1. If using a slice as a buffer &rarr; specify a nonzero length
        * ```x := make([]int, 10, 10)```
    2. If you know the exact size, you can specify the length and index into it to set the values. (Might result in a bug when accessing values that don't exist)
        * ```x := make([]int, 5, 10)```
    3. Otherwise use **make** with zero length and specified capacity

## Slicing Slices
* Slice expression creates a slice from a slice
* Consists of a starting offset and ending offset separated by a colon 
* leaving starting offset results in a 0 offset
* leaving ending offset results in a end of slice offset
```go
x := []int{1,2,3,4}
y := x[:2] // result: 1,2
z := x[:] // results 1,2,3,4
```
* <span style="color:#e74c3c">When taking a slice from a slice you are not making a copy</span>
    * You have two variables sharing memory
    * changes in a slice affect all slices that share the element
* <span style="color:#e74c3c">Don't use append with sub slices to avoid complicated slice scenarios</span>
    * Subslice's capacity is set to the capacity of the original slice minus the start offset &rarr; any unused capacity is shared 
    * multiple slices appending causes overwrites of each others data
    * Use a **full slice expression** to avoid overwrites &rarr; makes it clear how much memory is shared between parent and subslice ```y := [2:2:2] // cap = 2```
        * last or third part is the position in the parents capacity that is available to the subslice 

## Converting Arrys to Slices
* You can take a subslice of an array
* Useful way to bridge an array to a function that only takes slices
* Has same memory properties as slice of slice

## copy
* To create a slice that is independent from the original
* Takes destination slice and source slice input
* Copies as many values as possible limited by the smaller sized slice
* returns number of elements copied
* can be used with array by taking a slice of array ```a[:]```

```go
x := []int{1,2,3,4}
y := make([]int, 2)
copy(y, x[2:]) // result [1,2] into y
```
* Can copy slices with in a slice
```go
x := []int{1,2,3,4}
copy(x[:3], x[1:]) // result [1,3,4,4]
```

## Strings and Runes and Bytes
* Strings are not made out of runes
* Strings are sequences of bytes
* Bytes don't have to be in any particular encoding
* Go libraries assume that string is composed of a sequence of UTF-8 encoded code points
* string literals are written in UTF-8 (unless hexadecimal escapes are used)
* **Index expression** can be used to extract a single value 
    * indexes are zero-based
```go
s := "Hello"
var b byte = s[0]
```
* **Slice expression** also works with strings
```go
var s string = "hello"
s2 := s[:2] // results in "he"
```
* <span style="color:#e74c3c">String is a sequence of bytes and UTF-8 code points can be 1 to 4 bytes</span>
* **len()** will return the length of the string in bytes not code points
* <span style="color:#e74c3c">Only use slicing or indexing when you know that the string only contains ony byte characters</span>
* Better to use the strings and unicode/utf8 packages to slice and index strings

## Maps
* Useful when data is organized based on a value that is not strictly increasing order.
    * Map when order doesn't matter, slice when order of elements does matter
* Associate one value with another written as *map[keyType]valueType*
* Map in Go is a hash map &rarr; fast lookup, array internally key is turned into number with hash algorithm. Each element in the array is a bucket that can store multiple key-value pairs (collision, two keys map to same bucket). 
* Automatically grow with added key-value pairs
* **len()** outputs number of key-value pairs in a map
* zero value is *nil*
* Not comparable but can compared to *nil*
* keys in a map are comparable type
* ```var nilM map[string]int``` &rarr; var declaration of map set to nil
    * nil map has the length of 0
    * reading a nil map always returns 0
    * writing to a nil map causes panic
* ```m := map[string]int{}``` &rarr; := declaration with map literal
    * has length 0
    * can read/write to map
* ```score := map[string]int{"a": 2, "b": 3}``` &rarr; nonempty map literal
    * written has key followed by colon and then value
    * key-value pair is separated by a comma 
* ```grades := make(map[int]int, 10)``` &rarr; can use make if amount of key-value pairs is known
    * has length 0 and can grow past initial size

### Comma ok idiom
* To tell difference between a key that's associated with a zero value and a key that is not in the map
* true key is present in the map and false if it is not
```go
v, ok := valueMap["key"]
```

### Deleting from Maps
* build-in **delete()** function removes key-value pairs from maps
* takes a map and a key as input
* removes key-value pair if it exists otherwise it does nothing
* Does nothing if the map is nil
* returns nothing

### Simulate Sets with Maps
* Go doesn't include set.
* Use the key of the map for the type that you want to put into the set and use a bool for the value.
* You need to write your own union, intersection and subtraction functions
* You can use a struct{} for the value as that uses zero bytes when it is empty while boolean uses one byte &rarr; downside is that it makes the code less readable.

## Structs
* When you have related data that you want to group together
* Go doesn't have classes because it doesn't have inheritance
```go
type car struct {
    name string
    year int
    manufacturer string
}
```
* Defined with keyword type, name of the struct type and keyword struct and braces
    * Within the braces you define the fields of the struct
    * struct field name follow by struct field type
* Can be defined within or outside functions
    * Struct types defines within a function can only be used within the function
    * Technically it can be defined at any block level
* Once defined we can declare variables of that struct type
* structs zero value is the fields set to its zero values
    * There is no difference in assigning an empty struct and not assigning a value at all
```go
ford := car{}
```
* First struct literal format needs every field to be specified and values assigned to the fields in the order they are declared in the struct definition.
    * All values most be declared
    * Can be used for small structs
```go
volvo := car {
    "xc90",
    2020,
    "volvo"
}
```
* Second struct literal format looks like a map literal:
    * names of fields are used
    * fields can be in any order
    * any field not specified gets its zero value
    * This is more verbose, makes clear what is being assigned
    * More maintainable, future version might add more fields causing the above type not to compile
```go
honda := car {
    "year": 2016,
    "name": "CR-V"
}
```
* ```honda.name```Fields are accessed with dotted notation, both reads and writes

### Anonymous structs
* Variable can be declared to implement a struct type with out giving a struct type name.
```go
var animal struct {
    name string
    age int
}

animal.name = "Meowser"
animal.age = 4

pet := struct {
    name string
    kind string
}{
    "Tinni",
    "dog"
}
```
* Two scenarios when this is good
    * When translating external data into struct or vice versa.
        * Unmarshaling and marshaling data 
    * Writing tests 

### Comparing and converting structs
* Structs that are entirely composed of comparable types are comparable 
    * Does not contain: slices, maps, function and channel
    * No method can be overridden to redefine equality to make == and != work for incomparable structs
    * Cannot compare between variables of different struct types
* Type conversion can be done between two struct types if fields of both structs have the same names, order and types
    * same with anonymous struct types