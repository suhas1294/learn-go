#### how to import ? 
```go
// import "fmt"		// importing in single line

import (			// importing in multi line
	"fmt"
)
```

#### Printing a statement :
```go
const name , age = "John", 30;		// assignment of multiple variables in single line
fmt.Println("Hello, World!");
fmt.printf("Hello, %s! You are %d years old.\n", name, age);
```

#### how to know type of variable
```go
fmt.Printf("Type of name: %T\t and type of age is : %T", name, age);
```
***
#### What is encoding  ?
Computer works on switches : On (1) and Off (0)
It can be used to represent any data. for example : 

Coding scheme : 
00 : bring pizza
01 : bring burger
10 : bring sandwich
11 : bring salad

roughly translating above as coding scheme, way back world was not having standard coding
scheme : 
ex : 00 meant 'bring pizza' in india, but it could mean 'bring burger' in other country.

So there was a need of standard coding scheme : 
ASCII was one one them in earlier days. it used to represent data by using 8 bits. (7 bits for data and 1 bit for parity)
ASCII can represent 256 characters. (2 ^ 8, where 8 is number of bits)
For example, '01000100' is the binary representation of 'D' in ASCII.
Since ASCII can represent only 256 characters, it is not sufficient for all languages.

Then came UNICODE, which can represent 65536 characters. (2 ^ 16, where 16 is number of bits)

Then came UTF-8, which is a variable length encoding scheme. It can represent 1 to 4 bytes of data.
1. It stores Unicode as binary data.
2. If a character needs 1 byte, it will use 1 byte. If a character needs 2 bytes, it will use 2 bytes.
3. This makes it efficient in terms of memory.
4. Rare characters like doubel heart emoji (💕) takes 32 bits or 4 bytes.
5. UTF-8 is backward compatible with ASCII. So, all ASCII characters are also valid UTF-8 characters.
6. UTF-8 is the most widely used encoding scheme in the world. 
***

#### printing raw string literals
```go
fmt.Println(`Hello, 
    Bharath!,
    Jai Jagganath!
`);
```

__Note__ :
1. Go user manual vs Effective Go  : two different documents
2. Go user manual is for beginners and Effective Go is for advanced users.
3. **TODO**: know meaning of "source code is unicode represented in UTF-8"

#### Variables, Zero values, Blank identifier.
```go
// Variable : A VARIABLE holds a VALUE of a specific TYPE.
var sjbf int = 102; // this also assigns the value to the variable.

//  another way is to use short declaration operator ':='
// it allows to declare and assign a variable in one line.
sljfb := 102; // this also assigns the value to the variable.

// multiple assignments ; 
asd, ads, _ , asdjlnas := 1, 2, "blank identifier", 4; // this also assigns the value to the variable.
// here, _ is a blank identifier. It is used to ignore the value of the variable.
// it is useful when you want to ignore something out of many things a function returns.
// for example, if a function returns 3 values and you want to ignore the second value, you can use _.
```

__Note__ : 
1. unused variables are not allowed in Go. Compiler will throw an error if you try to use an unused variable.
2. General guideline : use short declaration operator ':=' when possible.

__When to use var keyword?__ :
1. when you want to declare a variable without assigning a value to it.
2. when you want to declare a variable with a specific type.
3. when you want to declare a variable in a package scope.
4. when you want to declare a variable in a function scope.
5. when you want to declare a variable in a loop.
6. when you want to declare a variable in a conditional statement.
7. when you want to declare a variable in a switch statement.
8. when you want to declare a variable in a select statement.
9. when you want to declare a variable in a defer statement.

__Note :__
1. Unassigned variables are given a zero value. (default value)
2. Default values for pointers, functions, interfaces, slices, channels, maps is nil.

***
#### values, types, conversions, scope
1. Casting in other languages is called conversion in Go.
2. %v is kind of supeset for all types. It is used to print the value of a variable.
3. official syntax for conversion : T(x) where T is the type and x is the value.
4. short declaration can be used only inside functions.
5. There is no variable hoisting in Go. So, you cannot use a variable before it is declared.

#### Built in types, aggregate types, composite types.
- package 'builtin'.
- built in types : int, float, string, bool, complex.
- aggretate types : array, struct, slice, map.
- composite/compound types : struct, interface

#### generating random number in go
```go
import (
    "fmt"
    "math/rand"
)
fmt.Println(rand.Intn(100)) // this will print a random number between 0 and 99. 
// Intn generated number between [0, n) where n is the number passed to it.
```

#### variadic parameter
```go
fmt.Println("name is", name, "and age is", age); // Println takes multiple 'n' number of arguments. variadic parameter.
```

### scope access :
We dont use the words 'public and private' in Go. We use 'exported and unexported'.


#### Functions : 
```go
function add(a int, b int) int { // function with return type
    return a + b;
}

// this can be shortened to
function addAnother(a, b int) int { // function with multiple return type
    return a + b;
}

// function can return multiple values.
function addAndSubtract(a, b int) (int, int) {
    return a + b, a - b;
}

function swap(a, b string) (string, string) {
    return b, a;
}

// Go's return type can be named aka Naked return type.
function addAndSubtract(a, b int) (sum int, diff int) {
    sum = a + b;
    diff = a - b;
    return
}
```

1. a var keyword can be at function level or package level.
2. uniniaialized variables are given a zero value.

__const and multiple declarations__ :
```go
const d int = 10; // constant variable
var i, isPresent bool; // multiple variable declaration, uninitialised
```

__IMPORTANT__: IF VARIABLES ARE INITIALISED, THEN TYPE CAN BE OMITTED.
```go
example : var i, j = 1, 2; // this is valid ; No need to say its int. 
```
***
__About Rune :__ 
1. '__rune__' is another word for '__char__' in other 
languages.
2. rune is an alias for int32. It is used to represent a Unicode code point.
3. uint8 = unsigned int of 8 bits
4. byte is an alias for uint8. It is used to represent a byte.

#### grouping variable declaration : 
variable declarations can be 'factored' into blocks : 
```go
var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 - 1
    holdsNothing string
    z complex128 = cmplx.Sqrt(-5 + 12i) // import "math/cmplx"
)
```
__Note__ : unsigned is positive number.

__Implicit conversion__: 
`var i = 22` // this is valid. value on right determines the type of variable on left.

__IMPORTANT__: __consts__ cannnot be declared with ':='

***

#### Concept of IOTA
```go
const (
    c0 = iota // 0
    c1 = iota // 1
    c2 = iota // 2
)

const (
    c3 = iota // 0
    c4
    c5
    c6
)

const (
    _ = iota // 0
    a
    b
    c
    d
    e
)

fmt.Println(c0, c1, c2, c3, c4, c5, c6); // 0 1 2 0 1 2 3
fmt.Println("%d \t  %b\n", 1, 1); // 1 1
fmt.Println("%d \t  %b\n", 1<<a, 1 << a); // 1 1
fmt.Println("%d \t  %b\n", 1<<b, 1 << b); // 2 10
fmt.Println("%d \t  %b\n", 1<<c, 1 << c); // 4 100
fmt.Println("%d \t  %b\n", 1<<d, 1 << d); // 8 1000
fmt.Println("%d \t  %b\n", 1<<e, 1 << e); // 16 10000
```

__exercise__ : using IOTA concept, print size of bytes, mega bytes, giga bytes and so on.

```go
type ByteSize int;
const (
    _ = iota;
    KB ByteSize = 1 << (10 * iota) // 1 << (10 * 0) = 1
    MB
    GB
    TB
    PB
    EB
)
fmt.Println(KB, MB, GB, TB, PB, EB); // 1024 1048576 1073741824 1099511627776 1125899906842624 1152921504606846976
fmt.Printf("%d \t %b\n", KB, KB); // 1024 10000000000
fmt.Printf("%d \t %b\n", MB, MB); // 1048576 100000000000000000000
fmt.Printf("%d \t %b\n", GB, GB); // 1073741824 100000000000000000000000000
fmt.Printf("%d \t %b\n", TB, TB); // 1099511627776 100
fmt.Printf("%d \t %b\n", PB, PB); // 1125899906842624 100
fmt.Printf("%d \t %b\n", EB, EB); // 1152921504606846976 100
```

__TODO exercise__: Run hash256 (google for 'sha256 checksum mac command') algorithm on a string and get the hash value.
***

#### creating projects in the local : 
`go mod init <module_name>` // this will create a go.mod file in the current directory.

`GOOS` : Go Operating System
`GOARCH` : Go Architecture
`go env GOARCH GOOS`: this will print the current OS and architecture.
`GOOS=darwin GOARCH=amd64 go build -o hello hello.go` : this will build the hello.go file for darwin OS and amd64 architecture.

***

__GOPATH__ VS __GOROOT__ : 
1. __GOPATH__ is the path where your Go projects are stored. It is set to $HOME/go by default.
    - this contains the source code, binaries and packages. (bin, pkg, src)
2. __GOROOT__ is the path where Go is installed. It is set to /usr/local/go by default.
3. __GOPATH__ is no more required in Go 1.11 and later.
    - Go modules are used to manage dependencies. (go.mod file)
    - Go modules are used to manage packages. (go.sum file)

```txt
go.mod is equivalent to package.json in Node.js.
go.sum is equivalent to package-lock.json in Node.js.
```

#### declaring variables in if statement itself:
In go, if statements, you can declare variables in the if statement itself.
```go
if i := 10; i > 0 {
    fmt.Println("i is positive");
} else {
    fmt.Println("i is negative");
}
```

***

1. `go env` : lists all the environment variables used by go, includes GOROOT and GOPATH.
2. `go get` is a tool that lets you download and install packages from the internet.
3. `go mod tidy` is a tool that lets you remove unused packages from the go.mod file and also add missing packages to the go.mod file which is used in the code.

__Package and exports__:

What makes something visible outside the pacakge or not depends on whether its exported or not.
Exported means it starts with a capital letter. Unexported means it starts with a small letter.
Note : package = folder

sample usage : packageName.FunctionName();

download all dependencies without cleaning: `go mod download`.

go mod tidy : clean all dependencies and download only the required ones. (updates those which are updated in go.mod file)

#### versionining : 

Versioning:
`git tag`
`git tag vN.N.N` (this is called semantic versioning : Major::Minor::Patch)
`git push origin --tags`

Major : need not be backward compatible. (breaking changes)
Minor : backward compatible. (new features)
Patch : backward compatible. (bug fixes)

***

#### INIT function : 
Go program has concept of 'init' function.
- this function is called before the main function.
- this function is used to initialize the package.

#### statement, statement idiom 
```go
if a, b, c := 1, 2, 3; a+b+c > 5 {
    fmt.Println("sum > 5")
}
```

practical examples : 
```go
// 1. Error handling during file open
if file, err := os.Open("config.json"); err != nil {
    log.Fatal(err)
} else {
    defer file.Close()
    // use file
}

// 2. Reading a config or environment variable : 
if value, ok := os.LookupEnv("PORT"); ok {
    fmt.Println("Port is", value)
} else {
    fmt.Println("PORT not set")
}

// 3. JSON unmarshalling
if err := json.Unmarshal(data, &config); err != nil {
    log.Println("Invalid JSON:", err)
}

// 4. Map lookup pattern
if val, ok := myMap["key"]; ok {
    fmt.Println("Found value:", val)
}

// 5. Reading data in a loop
for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
    fmt.Println(line)
}
```
***
#### Comma-ok idiom :
```go
if i, ok := m["key"]; ok {
    fmt.Println("key exists");
}
```
In Go, some operations(or functions) return two values:
- the actual result
- a boolean indicating success
`value, ok := something[key_or_input]`

__NOTE__ : It’s not universal to all functions — only functions or operations designed to return (value, ok).

__examples :__ 
```go
// - Map lookups
myMap := map[string]int{
    "one": 1,
    "two": 2,
}

if val, ok := myMap["three"]; ok {
    fmt.Println("Found:", val)
} else {
    fmt.Println("Key not found")
}

// - Type assertions
var i interface{} = "hello"

if s, ok := i.(string); ok {
    fmt.Println("It's a string:", s)
} else {
    fmt.Println("Not a string")
}
// Without this check, an invalid type assertion would panic.


// - Environment variable lookups

// - Channels (value, ok := <-channel)
ch := make(chan int)
close(ch)

if val, ok := <-ch; ok {
    fmt.Println("Received:", val)
} else {
    fmt.Println("Channel closed")
}
// When a channel is closed, you still receive the zero value of its type, 
// so the ok tells you whether the channel was open when you received.
```

#### select statement :
- The select statement is like a switch statement, but for channels.
- It allows you to wait on multiple channel operations.
```go
select { 
    case msg := <-ch1:
        fmt.Println("Received from ch1:", msg)
    case msg := <-ch2:
        fmt.Println("Received from ch2:", msg)
    }
```
***
#### For loops (3 syntaxes) : 

1. for init ; condition; post {}
2. for condition {} // like a c while loop.
3. for { } // infinite loop

__For-range loop :__
- It is used to iterate over arrays, slices, maps, strings, and channels.
- It returns two values: the index/key and the value.
- It is a shorthand for the traditional for loop.
- It is more readable and less error-prone.
- It is used to iterate over collections.

```go
xi = []int{1, 2, 3, 4, 5};
for i, v := range xi {
    fmt.Println(i, v);
}

m := map[string]int{
    "one": 1, 
    "two": 2
};
for k, v := range m {
    fmt.Println(k, v);
}

// NOTE: if you only need the value, you can use the blank identifier to ignore the key.
for _, v := range m {
	fmt.Println(v);
}

// NOTE : Map does not guarantee the order of iteration.
```
***

#### Slice vs arrays : 
- Slice is a dynamic array, while array is a fixed size array.
- Slice is a reference type, while array is a value type.
- Slice has a length and capacity, while array has a fixed length.
	
Array is mostly used internally by go, we dont use it much often in our code unless we know the size of the array.

While declaring array, if we want compiler to figure out the size of the array, we can use '...' instead of size.
`a := [...]int{1, 2, 3, 4, 5};` // this will create an array of size 5.

```go
var c [2]int;
c[0] = 1;
c[1] = 2;
fmt.Println(c); // [1 2]
fmt.Println("Length of array is", len(c)); // 2
```

##### All about slices : 
```go
xi := []int{1, 2, 3, 4, 5}; // this will create a slice of size 5.
xs := []string{"a", "b", "c", "d", "e"}; // this will create a slice of size 5.

// variadic parameter for append function.
xi = append(xi, 6, 4, 234, 546, 567, 234, 456); // this will append the values to the slice.

// slicing a slice : 
xi := []int{0,1,2,3,4,5,6,7,8,9};
xi = xi[1:3]; // this will create a new slice with the values from index 1 to 3.
/* [inclusive: exclusive]

[:exclusive] : this will create a new slice with the values from index 0 to exclusive.

[inclusive:] : this will create a new slice with the values from index inclusive to end of the slice.

[:] : this will create a new slice with the values from index 0 to end of the slice. IN a way cloning the slice.

*/

// deleting in a slice using append function.
xi = append(xi[:2], xi[3:]...); // this will delete the value at index 2.
// here '...' is called unfurling, its similar to spread operator in JS.

// Make function : 
// it is used to create a slice with a specific length and capacity.
xim := make([]int, 5, 10); // this will create a slice of size 5 and capacity 10.
fmt.println(cap(xim)); // 10
fmt.println(len(xim)); // 5

// --------------------------------------
si := make([]int, 2, 10)
fmt.Println("Hello, size, capicity : ", len(si), cap(si))
for i, v := range si {
    fmt.Println("index : ", i, "\tvalue : ", v)
}
fmt.Println()
si = append(si, 3498, 3479, 234)
for i, v := range si {
    fmt.Println("index : ", i, "\tvalue : ", v)
}
/* output : 
Hello, size, capicity :  2 10
index :  0 	value :  0
index :  1 	value :  0

index :  0 	value :  0
index :  1 	value :  0
index :  2 	value :  3498
index :  3 	value :  3479
index :  4 	value :  234

si[5] = 1000 // this will throw an error because the slice is not big enough to hold the value.
U can only append !!
*/
```

##### slice of slice (Multi dimensional array):
```go
radioActiveElements := []string{"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne"};
metalElements := []string{"Li", "Be", "B", "C", "N", "O", "F", "Ne"};
nonMetalElements := []string{"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne"};

allElements := [][]string{radioActiveElements, metalElements, nonMetalElements};
fmt.Println(allElements); // this will print the slice of slices.
```
***

#### All about maps
```go
var m map[string]int; // this will create a map with string as key and int as value.
m = make(map[string]int); // another way to create map using make function.
m["one"] = 1; // this will add the key value pair to the map.
m["two"] = 2; // this will add the key value pair to the map.
fmt.Printf("%#v\n", m); // this will print the map.

// randing over map : 
for k, v := range m {
    fmt.Println("key : ", k, "\tvalue : ", v);
}
// NOTE: if you only need the value, you can use the blank identifier to ignore the key.

// deleting an element from map :
delete(m, "one"); // this will delete the key value pair from the map.
// NOTE : trying to delete a non existing key will NOT throw an error(PANIC).
// NOTE : trying to access a non existing key will return the zero value of the type.
// right way to check if key exists or not is to use comma-ok idiom.
// example :
if val, ok := m["one"]; ok {
    fmt.Println("key exists");
} else {
    fmt.Println("key does not exist");
}
```
***

#### Structs
```go
type person struct {
    first string
    last string
    age int
}

p2 := person{first: "Jane", last: "Doe", age: 25}; // this will create a struct with the values.
p1 := person{"John", "Doe", 30}; // this will create a struct with the values.

// NOTE : diff b/w %v, %+v, %#v
fmt.Printf("%v\n", p1)   // Output: {John Doe 30}
fmt.Printf("%+v\n", p1)  // Output: {first:John last:Doe age:30}
fmt.Printf("%#v\n", p1)  // Output: main.person{first:"John", last:"Doe", age:30}

// Embedded structs :
type address struct {
    city string
    state string
    country string
}

type person struct {
    first string
    last string
    age int
    address address // this will create a struct with the values.
}

type secretAgent struct {
    person // this will create a struct with the values.
    licenseToKill bool
}

sa1 := secretAgent{
    person{"James", "Bond", 35}, 
    true,
}; // this will create a struct with the values.

sa2 := secretAgent{
    person: person{
        "Jane", 
        "Doe", 
        25
    },
    licenseToKill: true,
}

fmt.Println(sa1); // this will print the struct.
// NOTE : if you want to access the Fields of the embedded struct, you can do it directly.
fmt.Println(sa1.first); // this will print the first name of the person.
fmt.Println(sa1.address.city); // this will print the city of the address.
// NOTE : If there is one struct inside another struct, the inner Fields get promoted to the outer struct.

// anonymous struct :
p3 := struct {
    first string
    last string
    age int
}{
    first: "John",
    last: "Doe",
    age: 30,
}; // this will create a struct with the values.
fmt.Printf("%T\n", p3); // this will print out anonymous struct type.

// creating custom types : 
type foo int;

var x foo = 42; // this will create a custom type with the value.

// IMPORTANT :  In go we dont instantiate a type, we create a value of that type.

// NOTE : Oraganise Fields in struct in such a way : large to small. this helps in performance
```
***
### Functions
- [Chatgpt questions](https://chatgpt.com/share/682a0db8-d504-800f-9e72-1c0f50329dbf)

__covered topics :__
1. variadic parameters
2. unfurling a slice
3. methods, interfaces and polymorphism
4. anonymous functions and func expressions
5. callbacks and closures

__Syntaxes__:
```go
func (receiver type) methodName (parameters) (return types) {
    // method body
}

// Note : Everything in Go is a pass by value.

// no params, no return
func functionName() {
    // function body
}

// 1 param, no return
func functionName(param1 type) {
    // function body
}

// 1 param, 1 return
func functionName(param1 type) (return1 type) {
    // function body
    return return1
}

// 2 params, 2 returns
func functionName(param1 type, param2 type) (return1 type, return2 type) {
    // function body
    return return1, return2
}
```

__Defer statement :__
- this is used to delay the execution of a function until the surrounding function returns.
- this is useful for cleaning up resources, closing files, etc.
- Whenever the outerfunction(which holds another function with defer in it) returns, the deferred function will be executed.


__METHODS vs FUNCTIONS :__
1. A function is a standalone block of code that can be called from anywhere in the program.
2. A method is a function that is associated with a specific type (struct) and can be called on an instance of that type.
3. A method has a receiver, which is the instance of the type that the method is called on.
4. A function does not have a receiver.
5. A method can access the Fields and methods of the type it is associated with, while a function cannot.

6. A method can be called on a pointer to the type, which allows it to modify the Fields of the type.
7. A function cannot be called on a pointer to the type.
8. A method can be called on a nil pointer, which will not cause a panic.
9. A function cannot be called on a nil pointer, which will cause a panic.


```go
type person struct {
    name string
    age  int
}

func (p person) greet() {
    fmt.Println("Hello, my name is", p.name, " and my age is ", p.age);
}

p1: = person{name: "John", age: 30}
p1.greet() // Hello, my name is John and my age is 30
```

__INTEERFACE__ :
1. An interface is a collection of method signatures that a type must implement in order to satisfy the interface.
2. An interface does not contain any implementation, only the method signatures.
3. A type can implement multiple interfaces, and an interface can be implemented by multiple types.

__POLYMORPHISM__ :
1. Polymorphism is the ability of a type to take on multiple forms.
2. In Go, values can be be of more then one type.
3. In Go, polymorphism is achieved through interfaces.

```go
type shape interface {
    area() float64
    perimeter() float64
}
// --------- --------- --------- --------- ---------
type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius;
}
// --------- --------- --------- --------- ---------
type rectangle struct {
    length float64
    breadth float64
}
func (r rectangle) area() float64 {
    return r.length * r.breadth
}
func (r rectangle) perimeter() float64 {
    return 2 * (r.length + r.breadth)
}
// --------- --------- --------- --------- ---------
func printShapeInfo(s shape) {
    fmt.Println("Area:", s.area())
    fmt.Println("Perimeter:", s.perimeter())
}
// --------- --------- --------- --------- ---------
c := circle{radius: 5}
r := rectangle{length: 10, breadth: 5}
printShapeInfo(c)
printShapeInfo(r)
```

(_Repeat Example for polymorphism:_)
```go
type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}
func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func printShapeInfo(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}
```

__Note__:
Best way to remeber interface : __Interface says : "Hey man ! if u implement my methods, you are of my type"__

***

__toString( ) overriding | exploring stringer intercace :__
- Its very similar to overriding toString() method in Java.

```go
type Stringer interface {
    String() string
}
type person struct {
    name string
    age  int
}
func (p person) String() string {
    return fmt.Sprintf("My name is %s and my age is %d", p.name, p.age);
}
p1 := person{name: "John", age: 30}
fmt.Println(p1.String()) // My name is John and my age is 30
fmt.Println(p1)          // My name is John and my age is 30
```

__WRAPER FUNCTION USING INTERFACE :__

```go
type book interface {
    title string
}

type count int;

func (c count) String() string {
    return fmt.Sprintf("Count is %d", c);
}

func (b book) String() string {
    return fmt.Sprintf("Book title is %s", b.title);
}

func logCustom(s fmt.Stringer) {
    log.Println("[custom msg from wrapper]: " + s.String());
}
book1 := book{title: "Go Programming"}
count1 := count(10)
logCustom(book1); // output : [custom msg from wrapper]: Book title is Go Programming
logCustom(count1); // output : [custom msg from wrapper]: Count is 10
```
Wrapper function, Example-2:
calculating time taken by a function to execute : 
```go
func add(x, y int) int {
    return x + y
}
func timeTaken(f func()) {
    start := time.Now()
    f()
    elapsed := time.Since(start)
    fmt.Println("Time taken : ", elapsed)
}

timeTaken(func(){
    add(1, 2)
    fmt.Println("Add function executed")
})
```
***
__Anonymous function__ : 

Anonymous function is a function that is defined without a name.

syntax : 
`func(){code}()`
`func (param1 type, param2 type) (return types) {code} ()`

```go
func(s string) {
    fmt.Println("Hello", s)
}("John");
```

__Func expressions__
```go
fe := func(s string) {
    fmt.Println("Hello", s)
}
fe("John");
```

__returning a function from a function__:

```go
func bar() func() int {
    // returning a anonymous function which inturn returns a int
    return func() int {
        return 42
    }
}

barFunc := bar()
fmt.Println(barFunc()) // 42
```

__contept of callbacks :__
```go
func add(a int, b int) int {
    return a + b
}

func subtract(a int, b int) int {
    return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
    return f(a, b)
}

doMath(10, 5, add)      // 15
doMath(10, 5, subtract) // 5
```

__concept of closures :__

```go
func incrementer() func() int{
    i := 0
    return func() int {
        i++
        return i
    }
}

f := incrementer();
fmt.Println(f()) // 1
fmt.Println(f()) // 2
fmt.Println(f()) // 3

g := incrementer();
fmt.Println(g()) // 1
fmt.Println(g()) // 2
fmt.Println(g()) // 3
```

__Misc examples__:
```go
func doNothing() {
	fmt.Println("foo")
}

func printName(s string) {
	fmt.Println("My name is " + s)
}

func getName(s string) string {
	return fmt.Sprintf("They call me %s", s)
}

func getNameWithAlias(fileName int) (string, string) {
	return fmt.Sprintf("File no : %d They call me Bond,", fileName), "James Bond"
}

func greetAllFellows(prefix string, names ...string) {
	for _, name := range names {
		fmt.Println(prefix + " " + name + "!")
	}
}

// function with variadic parameters :
// this prints sum of all the numbers
func funcWithVariadicParams(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {

	doNothing()

	printName("John")

	fmt.Println(getName("John"))

	name, alias := getNameWithAlias(007)
	fmt.Println("The Name as per file is  : ", name, " and the alias is : ", alias)

	funcWithVariadicParams(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	funcWithVariadicParams() // still a valid one !

	ss := []string{"John", "Doe", "Smith"}
	greetAllFellows("Hello", ss...) // this is how you unfurl a slice
	greetAllFellows("Hello")        // still a valid one !

	// ---------
	c := circle{radius: 5}
	r := rectangle{length: 10, breadth: 5}
	printShapeInfo(c)
	printShapeInfo(r)
	// ---------

}
```
***
### Implementing writer interface:
```go
/*
	We have this standard interface in Go from the package "io" that is used to write data to a stream.
	// The interface is called Writer and it has a single method Write that
	// takes a byte slice and returns the number of bytes written and an error.

	type Writer interface {
		Write(p []byte) (n int, err error)
	}
*/

package main

import (
	"bytes"
	"io"
	"os"
)

type person struct {
	name string
}

func (p person) WritePerson(w io.Writer) {
	// The Write method of the Writer interface is implemented here
	// It takes a byte slice and returns the number of bytes written and an error
	w.Write([]byte(p.name))
}

func anotherMain() {
	// Create a new person
	p := person{name: "John"}

	// Create a new writer
	f, _ := os.Create("person.txt")
	defer f.Close()

	var b bytes.Buffer

	p.WritePerson(f)  // writing person details into file
	p.WritePerson(&b) // writing person details into buffer

	/*
		Explanation:

		what are we trying to achieve here?
		We are trying to write the person details into a file and a buffer.

		Here both file interface and buffer interface are implementing the Writer interface.
		Both of them have the Write method that takes a byte slice and returns the number of bytes written and an error.

		If we call writePerson method with file interface, it will write the person details into the file.
		If we call writePerson method with buffer interface, it will write the person details into the buffer.
		So, we can use the same method to write to different types of writers.
	*/
}

```
***
### POINTERS :
[ChatGPT interaction](https://chatgpt.com/share/683aab63-5d6c-800f-9c0e-8f12e3dea843)

Say we have  : `x := 42`
"x" gives us the value of x, however :

We have three different operators :
1. `&x` gives us the address of x : __[address of x in memory]__
2. `*int` : type of `&x` is `*int` (pointer to an int) : __[return type of `&x`]__
3. `*x` gives us the value at the address of x : __[dereference the pointer]__
3.1. `*&x` gives us the value at the address of x, which is 42.

__Note__: Beforehand : Know the difference between "pass by value", "pass by reference" and "reference type"

In go lang, following data structures are refernce types :
1. Slices
2. Maps
3. Channels
4. Functions
5. Interfaces
6. Pointers

Only above data structures are reference types, rest
(primitive data types like int, string, float, boolean ) all are value types.

__Meaning__ : When we pass the reference type to a function, we are passing the address of the value in memory, not the value itself. example :

```go
func mapDelta(m map[string]int, i int) {
    m["a"] = 749
}
func main() {
    m := map[string]int{"a": 1, "b": 2}
    mapDelta(m, 1) // So HERE when we pass 'm' to function, implicitly we are passing the address of 'm' in memory.
    // So when we change the value of 'm' in function, it will change the value of 'm' in main function as well.
    fmt.Println(m) // map[a:749 b:2]
}
```

```go
func intDelta(i int) {
    i = 749;
}
func intDeltaAddress(i *int) {
    *i = 749;
}
func main() {
    i := 1

    // case-1  : AKA pass by value AKA VALUE SEMANTIC IN GOLANG
    intDelta(i) // So HERE when we pass 'i' to function, implicitly we are passing the value of 'i' in memory.
    // So when we change the value of 'i' in function, it will NOT change the value of 'i' in main function as well.
    fmt.Println(i) // 1

    // case-2 : AKA reference type AKA POINTER SEMANTIC IN GOLANG
    intDeltaAddress(&i) // So HERE when we pass '&i' to function, implicitly we are passing the address of 'i' in memory.
    // So when we change the value of 'i' in function, it will change the value of 'i' in main function as well.
    fmt.Println(i) // 749
}
```

When you write:
```go
type User struct {
    Name string
    Age  int
}
```
And then:
```go
u := User{"John", 30}
// u is an actual value.
```
When you pass u to a function:
- A copy of the entire struct is made (including its Fields).
- If you want to avoid copying and mutate the original, you must pass a pointer:

Go prefers __value semantics__ by default (simpler, safer, easier reasoning about ownership, no implicit references flying around).

One more peculiar problem related to slice although slice is a reference type :
```go
func addNumber(s []int) {
    s = append(s, 100) // underlying array might change
}
func main() {
    s := []int{1, 2, 3}
    addNumber(s)
    fmt.Println(s) // still [1, 2, 3]
}
```
Why? Because s inside addNumber is a copy of the slice header (ptr, len, cap).
When capacity overflows, a new array is created — and the copy inside the function now points to the new array, but original in main is unchanged.

__TAGLINE__ : In golang, you choose mutability by using pointers.

__When to use what__ ? (value semantics vs pointer semantics)

1. use value semantics when possible (u dont have to worry about memory management)
2. use pointer semantic for larger data structures (like slices, maps, structs) to avoid copying the entire data structure.
3. use pointer semantic for mutability (when you want to change the value of the data structure in the function)

```go
type dog struct{
    name string
}

func (d dog) bark() {
    fmt.Println("Woof! My name is", d.name)
}

func (d *dog) barkMore() {
    fmt.Println("Woof! My name is", d.name, "and I bark more!")
}
func main() {
    d1 := dog{name: "Buddy"}
    d1.bark()        // Woof! My name is Buddy
    d1.barkMore()   // Woof! My name is Buddy and I bark more! //TODO: why this works ?

    d2 := &dog{name: "Max"}
    d2.bark()        // Woof! My name is Max //TODO: why this works ?
    d2.barkMore()   // Woof! My name is Max and I bark more!
}
```

*** 
#### CONCEPT OF METHOD SET IN GOLANG 
In Go, a 'method set' is the set of methods attached to a type. This concept is key to the
Go's interface mechanism, and it is associated with both the value types and pointer types.

- The method set of a type T consists of all methods with receiver type T.
    - These methods can be called using variables of type T.

- The method set of a type *T consists of all methods with receiver *T or T
    - These methods can be called using variables of type *T.
    - it can call methods of the corresponding non-pointer type as well

The idea of the method set is integral to how interfaces are implemented and used in Go.

An interface in Go defines a method set, and any type whose method set is a superset of the
interface's method set is considered to implement that interface.

A crucial thing to remember is that in Go, if you define a method with a pointer receiver, the method is only in the method set of the pointer type. This is important in the context of interfaces because if an interface requires a method that's defined on the pointer (not the value), then you can only use a pointer to that type to satisfy the interface, not a value of the type.

```go
type dog struct{
    name string
}

func (d dog) bark() {
    fmt.Println("Woof! My name is", d.name)
}

func (d *dog) barkMore() {
    fmt.Println("Woof! My name is", d.name, "and I bark more!")
}

type youngin interface {
    bark()
    barkMore()
}

function youngBark(y youngin) {
    y.bark()
}

func main() {
    d1 := dog{name: "Buddy"}
    d1.bark()        // Woof! My name is Buddy
    d1.barkMore()   // Woof! My name is Buddy and I bark more! //TODO: why this works ?

    youngBark(d1) // output : COMPILATION ERROR : cannot use d1 (type dog) as type youngin in argument to youngBark:
    // because d1 is of type dog, and dog does not implement the interface youngin

    d2 := &dog{name: "Max"}
    d2.bark()        // Woof! My name is Max //TODO: why this works ?
    d2.barkMore()   // Woof! My name is Max and I bark more!

    youngBark(d2) // output : Woof! My name is Max
    // because d2 is of type *dog, and *dog implements the interface youngin
}
```

__SUMMARY OF METHOD SET:__
In simpler terms:

__definition__ of method set :
The collection of methods that a value of that type can call directly, which
determines whether it satisfies a given interface.

If you have a value of type dog, you can only call methods attached to (d dog).

- If you have a pointer of type `*dog`, you can call both:
    ● Methods attached to (d *dog)
    ● Methods attached to (d dog)

👉 Why can a pointer call value methods?
Because Go can automatically dereference the pointer to access the value and call the method. But it won’t automatically create a pointer if you only have a value and try to call a pointer method — that’s a one-way street.


| Variable | Type   | Method Set             | Implements `youngin`?    |
| :------- | :----- | :--------------------- | :----------------------- |
| `d1`     | `dog`  | `bark()`               | ❌ (missing `barkMore()`) |
| `d2`     | `*dog` | `bark()`, `barkMore()` | ✅                        |

| Type   | Method Set           | Can Call                                  | Can Satisfy Interface            |
| :----- | :------------------- | :---------------------------------------- | :------------------------------- |
| `dog`  | `{ bark }`           | `bark()`, `barkMore()` *(if addressable)* | Only interfaces needing `bark()` |
| `*dog` | `{ bark, barkMore }` | `bark()`, `barkMore()`                    | Any interface needing both       |

```
Receivers 				Values
-----------------------------------------------
(t T) 					T and *T
(t *T) 					*T
```

__"The method set of a type determines the INTERFACES that the type implements....."__


✅ Simple way to "read" the table:
- If method has (t T) receiver:
    - You can call it on a T value.
    - You can call it on a *T — because Go will auto-deref the pointer to T for you.

- If method has (t *T) receiver:
    - You can only call it on a *T
    - You cannot call it on a T unless it’s addressable, in which case Go can auto-take its address.

__Illustration :__

```go
type Dog struct {
    name string
}

// Value receiver
func (d Dog) Bark() {
    fmt.Println("Woof from", d.name)
}

// Pointer receiver
func (d *Dog) BarkMore() {
    fmt.Println("Loud woof from", d.name)
}

func main() {
    d := Dog{name: "Buddy"}
    p := &d

    // ------ Method with value receiver (d Dog) Bark --------
    d.Bark() // ✅ allowed: value calling value receiver
    p.Bark() // ✅ allowed: pointer calling value receiver (Go auto-dereferences)

    // ------ Method with pointer receiver (d *Dog) BarkMore --------
    p.BarkMore() // ✅ allowed: pointer calling pointer receiver

    // d.BarkMore() // ⚠️ sometimes allowed if 'd' is addressable — but in pure table rule: not allowed
    // Example of Non-Addressable Value:
    // Dog{"Rocky"}.BarkMore()   // ❌ compile error — cannot take address of composite literal
}
```

#### Method sets full illustration example : 
```go
package main

import "fmt"

type dog struct {
	name string
}

func (d dog) bark() {
	fmt.Println("Woof! My name is", d.name)
}

func (d *dog) barkMore() {
	fmt.Println("Woof! My name is", d.name, "and I bark more!")
}

type brucelee interface {
	bark()
	barkMore()
}

func karate(b brucelee) {
	fmt.Println("\nKarate time!")
	b.bark()
	b.barkMore()
}

func main() {
	d := dog{name: "Buddy"}

	// example-1
	fmt.Println("\n------ example-1 ------")
	d.bark() // this is straight forward, understandable

	// example-2
	fmt.Println("\n------ example-2 ------")
	d.barkMore()
	/*
		here, inside barkMore, we are calling 'd.name', when u call d.barkMore(),
		barkMore() expects a pointer type(address), but we are sending direct value,
		we cant dereference a value, so it will not work. but how it works ?

		answer ;
		-> d is a named variable and hence addressable.
		-> Go automatically converts d.barkMore() to (&d).barkMore() behind the scenes.
		-> So yes — even though barkMore() expects a pointer, Go auto-addresses an addressable value.


	*/

	dpointer := dog{name: "Max"}
	dAddress := &dpointer

	// example-3 :
	fmt.Println("\n------ example-3 ------")
	dAddress.bark() // Go will internally dereference 'dAddress' and call bark on it, so this will work.

	// example-4 :
	fmt.Println("\n------ example-4 ------")
	dAddress.barkMore() // this is straight forward, understandable

	/*
		thing : dAddress
		type : pointer to value
		method sets : bark() and barkMore()
		does this satisy brucelee interface ? : Yes, because it has both methods bark() and barkMore()

		thing : d
		type : value
		method sets : bark() only
		does this satisy brucelee interface ? : No, because it does not have barkMore() method

		Keeping this in mind, lets try to call karate example :
	*/

	// karate(d) // expected compilation error. and its understandable.
	karate(dAddress) // this will work, because dAddress has both methods bark() and barkMore()
}
```

__User Crud example__:
```go
type User struct {
	Name string
	Age  int
}

type UserManager struct {
	Users []User
}

func (um *UserManager) AddUser(user User) {
	um.Users = append(um.Users, user)
}

func (um *UserManager) IncrementAgeOfAllUsers() {
	for i := range um.Users {
		um.Users[i].Age++
	}
}

func (um *UserManager) FindUserByName(name string) *User {
	for i := range um.Users {
		if um.Users[i].Name == name {
			return &um.Users[i]
		}
	}
	return nil
}

func (um UserManager) printUsers() {
	for i := range um.Users {
		fmt.Print(um.Users[i], "\t")
	}
}

func (u User) printUser() {
	fmt.Println(u, "\n[Printed user]")
}

func (um *UserManager) RemoveUser(name string) {
	ul := um.Users
	removeIdx := -1
	for i := range ul {
		if ul[i].Name == name {
			removeIdx = i
			break
		}
	}
	if removeIdx < 0 {
		return
	}

	newUsers := append([]User{}, ul[:removeIdx]...)
	newUsers = append(newUsers, ul[removeIdx+1:]...)
	um.Users = newUsers
}

func main() {
	u1 := User{Name: "john", Age: 34}
	u2 := User{Name: "cena", Age: 56}
	u3 := User{Name: "rock", Age: 72}
	var us []User
	us = append(us, u1, u2, u3)
	um := UserManager{Users: us}
	um.IncrementAgeOfAllUsers()
	um.printUsers()

	uf := um.FindUserByName("cena")
	if nil != uf {
		uf.Age = 66
	}
	uf.printUser()
	um.printUsers() //See if changes reflect in the original UserManager
	um.RemoveUser("rock")
	um.printUsers()
}
```

***

### Generics : 

```go
func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

type myNumber interface {
	int | float64
}

func addM[T myNumber](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	fmt.Println(addI(5, 12))
	fmt.Println(addF(12.34, 345.34))

	// making uses of generic types.
	fmt.Println(addT(5, 12))
	fmt.Println(addT(12.34, 345.34))

	fmt.Println(addM(5, 12))
	fmt.Println(addM(12.34, 345.34))

    var x myAlias = 42
    fmt.Println(addM(x, 23));	// WONT WORK, COMPILE ERROR !!
    // To fix this, we can redefine the myNumber interface :
    type myNumber interface {
        ~int | ~float64
    }
    // This tells go compiler to include all values of type (int/float) AND any values who has an underlying value of type (int/float).
}
```

__Package constraints :__

We need to import __"golang.org/x/exp/constraints"__ to use the constraints package.

```go
import "golang.org/x/exp/constraints"

type myNumber interface {
    constraints.Integer | constraints.Float
    // constraints.Integer is a type constraint that includes all integer types (int, int8, int16, int32, int64).
    // constraints.Float is a type constraint that includes all floating-point types (float32, float64).
    // This allows us to use any integer or floating-point type as a type parameter in our generic function.
}
```

__concrete types vs interface types.__

- __concrete type__ is a type that u can directly instantiate or create a value from.

- __Interface type__ is a type which defines contracts (set of methods or types) but does not represent specific data or instance.They represent behaviour or type but not specific set of values.

Comparable interface : [Check here](https://go.dev/blog/comparable)

The comparable constraint permits any type whose values may be compared using `==` and `!=`.

Meaning:
- __Allowed__: int, float64, string, bool, pointers, structs where all Fields are comparable, arrays of comparable elements, etc.
- __Not Allowed__: slices, maps, functions (because you can’t compare those with == and != in Go)

It’s a compiler-enforced constraint built into the type system.

__SUMMARY__ : It allows types that can be compared using == and !=.

__QUESTION : Why use 'K comparable' instead of 'K any' in generics?__
__ANSWER__ : Map key types in Go must be comparable by language rule.
Using 'K comparable' allows the compiler to enforce that the type

FYI :
```go
type person struct {
    name string
    age int
}
p1 := person{name: "Alice", age: 30}
p2 := person{name: "Alice", age: 30}
fmt.Println(p1 == p2) // This will print 'true' because both have the same values for name and age

// Official docs : https://go.dev/doc/tutorial/generics
```

__Array list implementation using generics :__
```go
// Define a generic type ArrayList with type parameter T
type ArrayList[T any] struct {
	items []T
}

// Add an element to the list
func (a *ArrayList[T]) Add(item T) {
	a.items = append(a.items, item)
}

// Get an element at a specific index
func (a *ArrayList[T]) Get(index int) T {
	return a.items[index]
}

// Get length of the list
func (a *ArrayList[T]) Len() int {
	return len(a.items)
}

func main() {
	// Create an ArrayList of int
	intList := ArrayList[int]{}
	intList.Add(10)
	intList.Add(20)

	fmt.Println("First int:", intList.Get(0))
	fmt.Println("Length:", intList.Len())

	// Create an ArrayList of string
	strList := ArrayList[string]{}
	strList.Add("Go")
	strList.Add("Lang")

	fmt.Println("First string:", strList.Get(0))
	fmt.Println("Length:", strList.Len())
}
```
***

#### JSON Operations : 

```go
import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId        int    `json:"userId"`
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Completed     bool   `json:"completed"`
	optionalField string `json:"-"` // This will ignore the Field completely

}
```

__Summary about JSON tags :__
1. Field appears in JSON as key `myName`.
```go
Field int `json:"myName"`
```

2. Field appears in JSON as key "myName" and
the Field is omitted from the object if its value is empty, as defined above.
```go
Field int `json:"myName,omitempty"`
```

3. Field appears in JSON as key "Field" (the default), but
the Field is skipped if empty. __Note the leading comma.__
```go
Field int `json:",omitempty"`
```

4. Field is ignored by this package.
```go
Field int `json:"-"`
```

5. Field appears in JSON as key "-".
```go
Field int `json:"-,"`
```

```go
func main() {

	jsonStr := `{
		"userId": 1,
		"id": 1,
		"title": "delectus aut autem",
		"completed": false
		"optionalField": "This is an optional field that is not in the struct"
	  }`

	var p User

	// Unmarshal JSON string into user struct
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the struct
	fmt.Printf("UserId: %d, Id: %d, Title: %s, Completed: %t\n", p.UserId, p.Id, p.Title, p.Completed)

	// Marshal struct back to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
    
        // Converting from strng to bytes :
	_byteData := []byte(jsonStr)

	// Converting from bytes to string :
	_stringData := string(_byteData)
	// uint8 is a alias for byte, so you can use it interchangeably.
}
```

__Difference between Marshal/Unmarshal vs Encode/Decode:__

__Marshal/Unmarshal:__ These functions are used to convert Go data structures to and from JSON format in memory. They work with byte slices and return the JSON representation as a byte slice or parse a byte slice into a Go data structure.

__Encode/Decode:__ These functions are used to read and write JSON data directly from/to an `io.Writer` or `io.Reader`. They are typically used for streaming JSON data
and can handle larger data sets more efficiently. Encode writes JSON to an `io.Writer`, and Decode reads JSON from an `io.Reader`. They do not return byte slices but instead
read from or write to the provided streams.

__In Summary__, 
- Encode and Decode is used to write Directly to the wire.
- Marshal & Unmarshal is used to convert data and store it in variable (memory).

The wire may refer to network connections, files, or any other.

***

#### Sorting slices
```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Less(i, j) says: “Should i come before j?”
// < → Ascending
// > → Descending

func main() {
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    fmt.Println(people)
    // There are two ways to sort a slice. First, one can define
    // a set of methods for the slice type, as with ByAge, and
    // call sort.Sort. In this first example we use that technique.
    sort.Sort(ByAge(people))
    fmt.Println(people)

    // The other way is to use sort.Slice with a custom Less
    // function, which can be provided as a closure. In this
    // case no methods are needed. (And if they exist, they
    // are ignored.) Here we re-sort in reverse order: compare
    // the closure with ByAge.Less.
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age > people[j].Age
    })
    fmt.Println(people)

}

```
__Multi field sorting :__
If you ever need multi-field sorting (like sort by Age, and if same Age,
sort by Name), you can chain conditions in Less:

```go
sort.Slice(people, func(i, j int) bool {
    if people[i].Age == people[j].Age {
        return people[i].Name < people[j].Name
    }
    return people[i].Age < people[j].Age
})
```
***
### Concurrency in go:
[Chatgpt Interaction](https://chatgpt.com/share/683c21db-db64-800f-b248-59e9254f739b)

```go
import "fmt"

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo-", i)
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar-", i)
	}
}
func main() {
	go foo()
	go bar()
        // main function will exit here.
	// main is also goroutine / main thread.
	// When main thread stops, every other goroutine stops as well.
}
```
__concept of wait groups :__

we declare a wait group at a package level and and use it in goroutine to notify main thread that this task is done.

```go
import "sync"

var wg sync.WaitGroup

func main(){
    wg.add(1) ; // what does this 1 represent? It represents the number of goroutines we are going to wait for.
    go foo();
    bar();
    wg.Wait(); // this will block the main thread until all goroutines are done.
}

func foo() {
    defer wg.Done() // this will notify the main thread that this goroutine is done.
    // do something
}
```

__Go concurrency tagline :__
> "Do not communicate by sharing memory; instead, share memory by communicating."

Golang is designed in such a way that only one goroutine will have access to a variable at a time. which is achieved by using channels.

__Simple illustration of channels :__

```go
func doSomething(x int) int {
    return x * 2
}

func main() {
    ch := make(chan int) // create a channel of type int
    go func() {
        ch <- doSomething(5) // send the result of doSomething to the channel
    }()
    fmt.Println(<-ch) // receive the result from the channel
}
```

__Simulating race condition__:
```go
import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())

    counter := 0

    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
            v := counter
            // time.Sleep(time.Second)
            runtime.Gosched() // this is used to yield the processor, allowing other goroutines to run
            v++
            counter = v
            wg.Done()
        }()
        fmt.Println("Goroutines:", runtime.NumGoroutine())
    }
    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println("count:", counter)
}
```

__NOTE__ : to check if race confition exists in a program, we can use the `go run -race` command.

__Use of Mutex :__
```go
func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())

    counter := 0

    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    var mu sync.Mutex

    for i := 0; i < gs; i++ {
        go func() {
            mu.Lock()		// this will lock the further code to be accessed by different goroutines
            // until the current goroutine is done with the code.
            v := counter
            runtime.Gosched()
            v++
            counter = v
            mu.Unlock()		// this will unlock the code, allowing other goroutines to access it.
            // NOTE : if we forget to unlock the code, it will cause a deadlock.
            // A deadlock is a situation where two or more goroutines are waiting for each other to release a lock, causing them to be stuck indefinitely.
            wg.Done()
        }()
        fmt.Println("Goroutines:", runtime.NumGoroutine())
    }
    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println("count:", counter)
}
```

__Concept of Automic :__
```go
func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())

    var counter int64

    const gs = 100
    var wg sync.WaitGroup
    wg.Add(gs)

    for i := 0; i < gs; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)	// this automic package is user to perform atomic operations on variables. It has methods like AddInt64, LoadInt64, StoreInt64, etc.
            // No need to use mutex when using atomic operations.
            // This means that the operation will be performed atomically, without any other goroutine being able to access the variable at the same time.

            runtime.Gosched()
            fmt.Println("Counter\t", atomic.LoadInt64(&counter)) // this will load the value of counter atomically. (its like a getter method)
            wg.Done()
        }()
        fmt.Println("Goroutines:", runtime.NumGoroutine())
    }
    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println("count:", counter)
}
```

__Concept of channels :__

this code wont work :
```go
func main(){
    ch := make(chan int) // create a channel of type int
    ch <- 5 // send the value 5 to the channel
    fmt.Println(<-ch) // receive the value from the channel
    // this will cause a deadlock because the channel is unbuffered and there is no goroutine to receive the value.
}
```
This will work :
```go
func main(){
    ch := make(chan int) // create a channel of type int
    go func(){
        ch <- 5 // send the value 5 to the channel
    }()
    fmt.Println(<-ch)
}
```

concept of __buffered channels__ :
```go
ch := make(chan int, 2) // create a buffered channel of type int with a capacity of 2
// This channel can hold up to 2 values before it blocks the sending goroutine.

// Below function will work fine because we are using a buffered channel.
func main(){
    ch := make(chan int, 1) // create a channel of type int, buffered with a capacity of 1 - this will allow 1 value to sit in there until its pulled out.
    ch <- 5
    fmt.Println(<-ch)
}

// Below function will NOT work fine because we are putting 2 values in a channel of capacity 1, which will cause a deadlock.
func main(){
    ch := make(chan int, 1) // create a channel of type int, buffered with a capacity of 1 - this will allow 1 value to sit in there until its pulled out.
    ch <- 5
    ch <- 10
    fmt.Println(<-ch)
}
```

__Directional channels :__
- A directional channel is a channel that can only be used to send or receive values, but not both.
- This is useful when you want to restrict the usage of a channel to a specific direction.

__Sending only channel__ :
```go
func main(){
    ch := make(chan <- int, 2) // create a channel of type int, buffered with a capacity of 2, but only for sending values.
    ch <- 5 // this will work fine
    ch <- 10 // this will also work fine
    // fmt.Println(<-ch) // this will cause a compile error because the channel is only for sending values, not receiving.
}
```

__Receiving channel__:
```go
func main(){
    rh := make(<-chan int, 2) // create a channel of type int, buffered with a capacity of 2, but only for receiving values.
    rh <- 5 // this will cause a compile error because the channel is only for receiving values, not sending.
    rh <- 10 // this will also cause a compile error because the channel is only for receiving values, not sending.
    fmt.Println(<-rh) // this will work fine because we are receiving values from the channel.
}
```

NOTE : __Assigning from specific channel type to general type wont work__

```go
func main(){
    c := make(chan int)
    cs := make(chan <- int) // sending only channel
    cr := make(<-chan int)  // receiving only channel

    c = cs // this will cause a compile error because cs is a sending only channel, not a general channel.
    c = cr // this will also cause a compile error because cr is a receiving only channel, not a general channel.

    // general to specific channel type will work fine
    cs = c // this will work fine because c is a general channel, and we can assign it to a sending only channel.
    cr = c // this will work fine because c is a general channel, and we can assign it to a receiving only channel.
}
```

__Sample concurrent program using channels :__
```go
func main(){
    ch  := make(chan int)
    go foo(ch)
    bar(ch)
    fmt.Println("about to exit")
}

func foo(ch chan <- int) { // general to specific type, works !
    ch <- 43
}

func bar(ch <- chan int) { // specific to general type, works !
    fmt.Println("value received from channel is : ", <-ch) // this will receive the value from the channel
}
```

Ranging over channels :
One who 'sends' should 'close' the channel.
```go
func main(){
    ch := make(chan int)

    go func(){
        for i := 0; i < 5; i++ {
            ch <- i * 2 // send values to the channel
        }
        close(ch) // IMPORTANT : close the channel after sending all values
        // without closing, deadlock will happen.
    }()

    for v := range ch { // Keep ranging over the channel until it is closed
        fmt.Println("Received from channel:", v) // this will print the values received from the channel
    }

    fmt.Println("Channel closed, exiting main function")
}
```
__Select statement :__

```go
func main(){
    eve, odd, quit := make(chan int), make(chan int), make(chan int)
    go send(eve, odd, quit)
    receive(eve, odd, quit)
    fmt.Println("Exiting main function")
}

func send(eve, odd, quit chan<- int) {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            eve <- i // send even numbers to the eve channel
        } else {
            odd <- i // send odd numbers to the odd channel
        }
    }
    quit <- 0 // send a signal to quit
}

func receive(eve, odd, quit <-chan int) {
    for {
        select {
        case v := <-eve:
            fmt.Println("Received from eve channel:", v) // this will print the values received from the eve channel
        case v := <-odd:
            fmt.Println("Received from odd channel:", v) // this will print the values received from the odd channel
        case <-quit:
            fmt.Println("Received quit signal, exiting receive function")
            return // exit the function when quit signal is received
        }
    }
    // NOTE : if we don't have a quit channel, the program will run indefinitely.
}
```

Implementation of same program __without using quit channel__ :
```go
func main() {
    eve, odd := make(chan int), make(chan int)
    go send(eve, odd)
    receive(eve, odd)
    fmt.Println("Exiting main function")
}

func send(eve, odd chan int) {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            eve <- i
        } else {
            odd <- i
        }
    }
    close(eve)
    close(odd)
}

func receive(eve, odd chan int) {
    for {
        select {
        case v, ok := <-eve:
            if ok {
                fmt.Println("Received from eve channel:", v)
            } else {
                eve = nil // disable this case in select
            }
        case v, ok := <-odd:
            if ok {
                fmt.Println("Received from odd channel:", v)
            } else {
                odd = nil // disable this case in select
            }
        }

        if eve == nil && odd == nil {
            fmt.Println("Both channels closed, exiting receive function")
            return
        }
    }
}
```
__IMPORTANT__ : close(eve) and close(odd) mark the channels as closed, __Receivers detect this via ok.__

Same code __using bool channel for quit__ :
```go
func main() {
    even, odd := make(chan int), make(chan int)
    quit := make(chan bool)

    go send(even, odd, quit)
    receive(even, odd, quit)

    fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int, quit chan<- bool) {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            even <- i
        } else {
            odd <- i
        }
    }
    close(quit) // this will send false which will be detected in comma-ok in receive.
    // NOTE : we can also send a value to quit channel, but here we are just closing it.
    // quit <- true // this will send true to the quit channel, which will be detected in receive.
}

// receive channel
func receive(even, odd <-chan int, quit <-chan bool) {
    for {
        select {
        case v := <-even:
            fmt.Println("the value received from the even channel:", v)
        case v := <-odd:
            fmt.Println("the value received from the odd channel:", v)
        case receivedVal, ok := <-quit:
            if !ok {
                fmt.Println("from comma ok", receivedVal, ok)
                return
            } else {
                fmt.Println("from comma ok", receivedVal)
            }
        }
    }
}
```

MISC code : 
```go
func main(){
    c := make(chan int)
    go func() {
        c <- 42
    }()
    val, ok := <-c // this will receive the value from the channel and also check if the channel is closed or not
    if ok {
        fmt.Println(val, ok) // output : 42 true
    }
}
```
***
#### FAN-IN & FAN-OUT DESIGN PATTERN
__Fan-in :__ multiple goroutines sending data to a single channel
__Fan-out :__ single goroutine receiving data from multiple channels

```go
func main() {
    even, odd, fanin := make(chan int), make(chan int), make(chan int)

    go send(even, odd)
    go receive(even, odd, fanin) // Note that even receive function is diff thread unlike above examples

    for v := range fanin {			// range until fanin is closed
        fmt.Println(v)
    }

    fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int) {
    for i := 0; i < 100; i++ {
        if i%2 == 0 {
            even <- i
        } else {
            odd <- i
        }
    }
    close(even)
    close(odd)
}

// receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for v := range even {
            fanin <- v
        }
        wg.Done()
    }()

    go func() {
        for v := range odd {
            fanin <- v
        }
        wg.Done()
    }()

    wg.Wait()
    close(fanin)
}
```

__ROB PIKES CODE FOR ILLUSTRATION OF FAN-IN AND FAN-OUT DESIGN PATTERN__
```go
func main() {
    c := fanIn(boring("Joe"), boring("Ann"))
    for i := 0; i < 10; i++ {
        fmt.Println(<-c)
    }
    fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
    c := make(chan string)
    go func() {
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
        }
    }()
    return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            c <- <-input1
        }
    }()
    go func() {
        for {
            c <- <-input2
        }
    }()
    return c
}
```

__MODIFICATION OF ROB PIKE'S CODE FOR GRACEFUL EXIT__

```go
func main() {
    done := make(chan struct{})
    c := fanIn(done, boring("Joe", done), boring("Ann", done))
    for i := 0; i < 10; i++ {
        fmt.Println(<-c)
    }
    fmt.Println("You're both boring; I'm leaving.")
    close(done)
    time.Sleep(time.Second) // give time for goroutines to exit gracefully
}

func boring(msg string, done <-chan struct{}) <-chan string {
    c := make(chan string)
    go func() {
        defer fmt.Println(msg, "cleanup done")
        for i := 0; ; i++ {
            select {
            case c <- fmt.Sprintf("%s %d", msg, i):
                time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
            case <-done:
                close(c)
                return
            }
        }
    }()
    return c
}

// FAN IN
func fanIn(done <-chan struct{}, input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case v, ok := <-input1:
                if !ok {
                    return
                }
                c <- v
            case <-done:
                return
            }
        }
    }()
    go func() {
        for {
            select {
            case v, ok := <-input2:
                if !ok {
                    return
                }
                c <- v
            case <-done:
                return
            }
        }
    }()
    return c
}
```

__FAN-OUT DESIGN PATTERN EXAMPLE :__

```go
func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    go populate(c1)

    go fanOutIn(c1, c2)

    for v := range c2 {
        fmt.Println(v)
    }

    fmt.Println("about to exit")
}

func populate(c chan int) {
    for i := 0; i < 100; i++ {
        c <- i
    }
    close(c)
}

func fanOutIn(c1, c2 chan int) {
    var wg sync.WaitGroup
    for v := range c1 {
        wg.Add(1)
        go func(v2 int) {
            c2 <- timeConsumingWork(v2)
            wg.Done()
        }(v)
    }
    wg.Wait()
    close(c2)
}

func timeConsumingWork(n int) int {
    time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
    return n + rand.Intn(1000)
}
```

__FAN OUT design pattern with Throttling  :__

```go
func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    go populate(c1)

    go fanOutIn(c1, c2)

    for v := range c2 {
        fmt.Println(v)
    }

    fmt.Println("about to exit")
}

func populate(c chan int) {
    for i := 0; i < 100; i++ {
        c <- i
    }
    close(c)
}

func fanOutIn(c1, c2 chan int) {
    var wg sync.WaitGroup
    const goroutines = 10
    wg.Add(goroutines)

    for i := 0; i < goroutines; i++ { 	// here we are using only 10 threads.
        go func() {
            for v := range c1 {
                func(v2 int) {			// Here we were using go routines in earlier example.
                    c2 <- timeConsumingWork(v2)
                }(v)
            }
            wg.Done()
        }()
    }
    wg.Wait()
    close(c2)
}

func timeConsumingWork(n int) int {
    time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
    return n + rand.Intn(1000)
}
```

__Context in go routines :__

Its used to manage the lifecycle of goroutines, especially for cancellation and timeouts.

In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.

__Example usage__ of how to gracefully terminate goroutines using context :

__Example-1:__
```go
// example-1 :
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    fmt.Println("error check 1:", ctx.Err())
    fmt.Println("num gortins 1:", runtime.NumGoroutine())

    go func() {
        n := 0
        for {
            select {
            case <-ctx.Done():
                return
            default:
                n++
                time.Sleep(time.Millisecond * 200)
                fmt.Println("working", n)
            }
        }
    }()

    time.Sleep(time.Second * 2)
    fmt.Println("error check 2:", ctx.Err())
    fmt.Println("num gortins 2:", runtime.NumGoroutine())

    fmt.Println("about to cancel context")
    cancel()
    fmt.Println("cancelled context")

    time.Sleep(time.Second * 2)
    fmt.Println("error check 3:", ctx.Err())
    fmt.Println("num gortins 3:", runtime.NumGoroutine())
}
```

__Example-2 :__ 
```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // cancel when we are finished

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}

func gen(ctx context.Context) <-chan int {
    dst := make(chan int)
    n := 1
    go func() {
        for {
            select {
            case <-ctx.Done():
                return // returning not to leak the goroutine
            case dst <- n:
                n++
            }
        }
    }()
    return dst
}
```

__Example-3:__

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go worker(ctx, "A")
    go worker(ctx, "B")

    time.Sleep(3 * time.Second)
    cancel() // cancel all child goroutines
    time.Sleep(1 * time.Second)
    fmt.Println("Main exiting")
}

func worker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("worker", name, "exiting")
            return
        default:
            fmt.Println("worker", name, "working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
```

__Question__: 
i read one concept - in web application, for every request, there will be diff goroutines, and these go routiunes may spin up further goroutines to get other jobs done like db accessing, authentication etc. so while doing sub authentication and authorisation, how can a sub go routine which was spun up by its parent goroutine, can access cookies data ? u can show me this using psuedo code also. basically innermost goroutine access data of its parent goroutine - is it via closure ? 

real world use case :

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Create a context derived from request context
    ctx := r.Context()

    // Extract cookie value
    sessionID, _ := r.Cookie("session_id")

    // Add sessionID to context so sub-goroutines can access it
    ctx = context.WithValue(ctx, "sessionID", sessionID.Value)

    // spawn auth worker goroutine
    go authWorker(ctx)

    // spawn db query worker goroutine
    go dbWorker(ctx)

    fmt.Fprintln(w, "Request handled")
}

func authWorker(ctx context.Context) {
    sessionID := ctx.Value("sessionID").(string)
    fmt.Println("Auth worker got sessionID:", sessionID)
}

func dbWorker(ctx context.Context) {
    sessionID := ctx.Value("sessionID").(string)
    fmt.Println("DB worker got sessionID:", sessionID)
}
```

__summary__ :
1. Each HTTP request gets a context.Context (r.Context())
2. You can add key-value pairs using context.WithValue
3. When spinning sub-goroutines, you pass the context along
4. Sub-goroutines can read values via ctx.Value("key")
5. No need for fragile closure captures, everything is safe via context

***

### Adhoc questions and resources  : 

why are we suppose to close the resp.body atfer we are done reading data ? 

__Answer__ :  Think of it like reading a book from the library:
1. You borrowed a book (resp.Body)
2. You’ve read all the pages (io.ReadAll)
3. But… you still have it in your bag! 😅

Should the library system automatically assume you're done and recall the book?
No — you must explicitly return it (resp.Body.Close()), so that:

1. Other people can borrow it
2. The library can clean up its records3
3. The shelf gets freed

#### Semaphore : 
A semaphore is a concurrency control mechanism that limits the number of goroutines (or threads) that can access a shared resource at the same time.
Think of it like a room with limited chairs:
* If 3 people can sit at once (capacity = 3),
* Others must wait until someone leaves.
If you're doing something heavy or external (e.g., API calls, DB hits, file writes), and want to limit concurrency, you use a semaphore.
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// This is our semaphore: a buffered channel of size 3
	semaphore := make(chan struct{}, 3)

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}

			fmt.Printf("Goroutine %d acquired semaphore\n", id)
			time.Sleep(1 * time.Second) // Simulate work
			fmt.Printf("Goroutine %d releasing semaphore\n", id)

			// Release semaphore
			<-semaphore
		}(i)
	}

	wg.Wait()
}
```
Here __When all 3 slots are full, other goroutines block until one is freed.__ 
Alternatively we have `import "golang.org/x/sync/semaphore"`
In summary : __Semaphore = a concurrency limiter__
