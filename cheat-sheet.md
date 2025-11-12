# Go Cheatsheet

Quick reference for common Go patterns and operations.

---

## Table of Contents
1. [Arrays & Slices](#arrays--slices)
2. [Sorting](#sorting)
3. [Grouping & Filtering](#grouping--filtering)
4. [Loops & Iteration](#loops--iteration)
5. [Structs](#structs)
6. [Maps](#maps)
7. [Strings](#strings)
8. [Type Conversions](#type-conversions)
9. [Concurrency](#concurrency)
10. [File Operations](#file-operations)
11. [HTTP Requests](#http-requests)
12. [Database Operations](#database-operations)
13. [JSON Operations](#json-operations)
14. [Error Handling](#error-handling)
15. [Functions](#functions)
16. [Generics](#generics)
17. [Design Patterns](#design-patterns)
18. [Logging](#logging)
19. [Time & Date](#time--date)
20. [Context](#context)
21. [Testing](#testing)
22. [Environment Variables](#environment-variables)
23. [Input/Output](#inputoutput)
24. [Control Flow](#control-flow)
25. [Make Command](#make-command)

---

## Arrays & Slices

### Add element to last position
```go
arr = append(arr, newElement)
```

### Add element to first position
```go
arr = append([]Person{newElement}, arr...)
```

### Replace n'th element
```go
arr[n] = newElement
```

### Remove n'th element
```go
arr = append(arr[:n], arr[n+1:]...)
```

### Remove first 3 elements
```go
arr = arr[3:]
```

### Remove last 3 elements
```go
arr = arr[:len(arr)-3]
```

### Extract from n'th to m'th element
```go
subArr := arr[n:m]
```

### Merge two slices
```go
merged := append(slice1, slice2...)
```

### Check if slice includes element
```go
func contains(arr []Person, target Person) bool {
    for _, item := range arr {
        if item == target {
            return true
        }
    }
    return false
}
```

---

## Sorting

### Sort number slice
```go
import "sort"

numbers := []int{3, 1, 4, 1, 5, 9}
sort.Ints(numbers)
```

### Sort custom type slice
```go
import "sort"

type Person struct {
    Name  string
    Age   int
    Email string
}

people := []Person{
    {"Alice", 30, "alice@example.com"},
    {"Bob", 25, "bob@example.com"},
}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

---

## Grouping & Filtering

### Group by field (e.g., email domain)
```go
import "strings"

func groupByDomain(people []Person) map[string][]Person {
    groups := make(map[string][]Person)
    for _, p := range people {
        domain := strings.Split(p.Email, "@")[1]
        groups[domain] = append(groups[domain], p)
    }
    return groups
}
```

---

## Loops & Iteration

### Loop by index (can mutate)
```go
for i := 0; i < len(people); i++ {
    people[i].Age++  // Mutates actual object
}
```

### Loop with range (gets copy)
```go
for i, person := range people {
    person.Age++  // Only modifies copy
    people[i].Age++  // Use index to mutate actual
}
```

### Loop with range (ignore index)
```go
for _, person := range people {
    fmt.Println(person.Name)
}
```

---

## Structs

### Create dynamic struct with values
```go
type Config struct {
    Host string
    Port int
}

config := Config{Host: "localhost", Port: 8080}
```

### Anonymous struct with values
```go
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  30,
}
```

### Nested struct with values
```go
type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Address Address
}

p := Person{
    Name: "Alice",
    Address: Address{
        City:  "NYC",
        State: "NY",
    },
}
```

---

## Maps

### Initialize and use map
```go
m := make(map[string]int)
m["key"] = 42

// Or with literal
m := map[string]int{"key": 42}
```

### Check if key exists (comma ok idiom)
```go
value, ok := m["key"]
if ok {
    fmt.Println("Found:", value)
}
```

### Iterate over map
```go
for key, value := range m {
    fmt.Println(key, value)
}
```

### Delete key from map
```go
delete(m, "key")
```

---

## Strings

### Split string into characters
```go
s := "hello"
chars := []rune(s)  // For Unicode support
```

### Split by delimiter
```go
import "strings"

sentence := "go is awesome"
words := strings.Split(sentence, " ")
```

### String manipulation
```go
import "strings"

// Trim whitespace
trimmed := strings.TrimSpace("  hello  ")

// Case conversion
upper := strings.ToUpper("hello")
lower := strings.ToLower("HELLO")

// Check contains
if strings.Contains("hello world", "world") { }

// Check prefix/suffix
if strings.HasPrefix("hello", "he") { }
if strings.HasSuffix("hello", "lo") { }
```

---

## Type Conversions

### String to byte slice
```go
str := "hello"
bytes := []byte(str)
```

### Byte slice to string
```go
bytes := []byte{104, 101, 108, 108, 111}
str := string(bytes)
```

### String to int
```go
import "strconv"

num, err := strconv.Atoi("123")
```

### Int to string
```go
import "strconv"

str := strconv.Itoa(123)
```

### Parse float
```go
import "strconv"

f, err := strconv.ParseFloat("3.14", 64)
```

### Type assertion
```go
var i interface{} = "hello"
s, ok := i.(string)
```

### Type switch
```go
switch v := i.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
default:
    fmt.Println("Unknown type")
}
```

---

## Concurrency

### Worker pool with channels (no WaitGroup)
```go
func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}
```

### Concurrency with WaitGroup only
```go
import "sync"

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Println("Worker", id)
        }(i)
    }

    wg.Wait()
}
```

### Mutex usage
```go
import "sync"

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}
```

### Atomic operations
```go
import "sync/atomic"

var counter int64

atomic.AddInt64(&counter, 1)
value := atomic.LoadInt64(&counter)
```

### Select with channels
```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout")
}
```

### Select with range
```go
done := make(chan bool)
messages := make(chan string)

go func() {
    for {
        select {
        case msg := <-messages:
            fmt.Println(msg)
        case <-done:
            return
        }
    }
}()
```

---

## File Operations

### Read file as string
```go
import "os"

data, err := os.ReadFile("file.txt")
if err != nil {
    panic(err)
}
content := string(data)
```

### Read file as JSON
```go
import (
    "encoding/json"
    "os"
)

type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

data, err := os.ReadFile("config.json")
if err != nil {
    panic(err)
}

var config Config
json.Unmarshal(data, &config)
```

### Read large file with buffer
```go
import (
    "bufio"
    "os"
)

file, err := os.Open("largefile.txt")
if err != nil {
    panic(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    // Process line
}
```

---

## HTTP Requests

### Simple GET request (print raw JSON)
```go
import (
    "io"
    "net/http"
)

resp, err := http.Get("https://api.example.com/data")
if err != nil {
    panic(err)
}
defer resp.Body.Close()

body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

### GET request with JSON parsing
```go
import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

resp, err := http.Get("https://api.example.com/user")
if err != nil {
    panic(err)
}
defer resp.Body.Close()

var user User
json.NewDecoder(resp.Body).Decode(&user)
```

### POST request with headers, query params, and body
```go
import (
    "bytes"
    "encoding/json"
    "net/http"
)

type Person struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

person := Person{Name: "Alice", Email: "alice@example.com", Age: 30}
jsonData, _ := json.Marshal(person)

req, _ := http.NewRequest("POST", "https://api.example.com/users?active=true", bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer token123")

client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    panic(err)
}
defer resp.Body.Close()
```

### PUT request with payload
```go
import (
    "bytes"
    "encoding/json"
    "net/http"
)

person := Person{Name: "Bob", Email: "bob@example.com", Age: 25}
jsonData, _ := json.Marshal(person)

req, _ := http.NewRequest("PUT", "https://api.example.com/users/1", bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")

client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

### PATCH request with payload
```go
import (
    "bytes"
    "encoding/json"
    "net/http"
)

update := map[string]interface{}{"age": 31}
jsonData, _ := json.Marshal(update)

req, _ := http.NewRequest("PATCH", "https://api.example.com/users/1", bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")

client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

### DELETE request with payload
```go
import (
    "bytes"
    "encoding/json"
    "net/http"
)

data := map[string]string{"reason": "duplicate"}
jsonData, _ := json.Marshal(data)

req, _ := http.NewRequest("DELETE", "https://api.example.com/users/1", bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")

client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

### POST request with single file
```go
import (
    "bytes"
    "io"
    "mime/multipart"
    "net/http"
    "os"
)

file, _ := os.Open("document.pdf")
defer file.Close()

body := &bytes.Buffer{}
writer := multipart.NewWriter(body)
part, _ := writer.CreateFormFile("file", "document.pdf")
io.Copy(part, file)
writer.Close()

req, _ := http.NewRequest("POST", "https://api.example.com/upload", body)
req.Header.Set("Content-Type", writer.FormDataContentType())

client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

### POST request with multiple files
```go
import (
    "bytes"
    "io"
    "mime/multipart"
    "net/http"
    "os"
)

body := &bytes.Buffer{}
writer := multipart.NewWriter(body)

files := []string{"file1.pdf", "file2.pdf"}
for _, filename := range files {
    file, _ := os.Open(filename)
    part, _ := writer.CreateFormFile("files", filename)
    io.Copy(part, file)
    file.Close()
}
writer.Close()

req, _ := http.NewRequest("POST", "https://api.example.com/upload", body)
req.Header.Set("Content-Type", writer.FormDataContentType())

client := &http.Client{}
resp, _ := client.Do(req)
defer resp.Body.Close()
```

---

## Database Operations

### Connect to MySQL (no ORM)
```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
if err != nil {
    panic(err)
}
defer db.Close()

err = db.Ping()
if err != nil {
    panic(err)
}
```

### Connect to SQLite (no ORM)
```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

db, err := sql.Open("sqlite3", "./database.db")
if err != nil {
    panic(err)
}
defer db.Close()
```

### SQL CRUD operations

#### Create
```go
result, err := db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", "Alice", "alice@example.com", 30)
id, _ := result.LastInsertId()
```

#### Read (single row)
```go
var user User
err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", 1).Scan(&user.ID, &user.Name, &user.Email)
```

#### Read (multiple rows)
```go
rows, err := db.Query("SELECT id, name, email FROM users")
if err != nil {
    panic(err)
}
defer rows.Close()

var users []User
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Name, &u.Email)
    users = append(users, u)
}
```

#### Update
```go
result, err := db.Exec("UPDATE users SET age = ? WHERE id = ?", 31, 1)
rowsAffected, _ := result.RowsAffected()
```

#### Delete
```go
result, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
rowsAffected, _ := result.RowsAffected()
```

---

## JSON Operations

### Marshal (struct to JSON)
```go
import "encoding/json"

person := Person{Name: "Alice", Age: 30}
jsonData, err := json.Marshal(person)
jsonString := string(jsonData)
```

### Marshal with indentation
```go
jsonData, err := json.MarshalIndent(person, "", "  ")
```

### Unmarshal (JSON to struct)
```go
import "encoding/json"

jsonStr := `{"name":"Alice","age":30}`
var person Person
err := json.Unmarshal([]byte(jsonStr), &person)
```

### Encode (write JSON to writer)
```go
import (
    "encoding/json"
    "os"
)

person := Person{Name: "Alice", Age: 30}
encoder := json.NewEncoder(os.Stdout)
encoder.Encode(person)
```

### Decode (read JSON from reader)
```go
import (
    "encoding/json"
    "strings"
)

jsonStr := `{"name":"Alice","age":30}`
decoder := json.NewDecoder(strings.NewReader(jsonStr))
var person Person
decoder.Decode(&person)
```

### JSON tags
```go
type Person struct {
    Name      string `json:"name"`
    Email     string `json:"email,omitempty"`
    Age       int    `json:"age"`
    SecretKey string `json:"-"`  // Never serialized
}
```

### Custom marshal/unmarshal
```go
import (
    "encoding/json"
    "time"
)

type CustomDate time.Time

func (c CustomDate) MarshalJSON() ([]byte, error) {
    return json.Marshal(time.Time(c).Format("2006-01-02"))
}

func (c *CustomDate) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }
    t, err := time.Parse("2006-01-02", s)
    *c = CustomDate(t)
    return err
}
```

---

## Error Handling

### Basic error handling
```go
if err != nil {
    return err
}
```

### Custom error
```go
import "errors"

var ErrNotFound = errors.New("item not found")
```

### Custom error with formatting
```go
import "fmt"

err := fmt.Errorf("user %s not found", username)
```

### Error wrapping
```go
import "fmt"

if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

### Error unwrapping
```go
import "errors"

if errors.Is(err, ErrNotFound) {
    // Handle specific error
}

var customErr *CustomError
if errors.As(err, &customErr) {
    // Handle custom error type
}
```

### Defer, panic, and recover
```go
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
            result = 0
        }
    }()
    return a / b  // Will panic if b is 0
}
```

---

## Functions

### Function as parameter
```go
func process(data []int, fn func(int) int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        result[i] = fn(v)
    }
    return result
}

// Usage
doubled := process([]int{1, 2, 3}, func(x int) int { return x * 2 })
```

### Variadic function
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Usage
result := sum(1, 2, 3, 4, 5)
```

### Multiple return values
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

---

## Generics

### Simple generic function
```go
func Max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Usage
maxInt := Max(10, 20)
maxFloat := Max(3.14, 2.71)
```

### Complex generic example (generic data structure)
```go
import "fmt"

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
value, ok := intStack.Pop()
```

---

## Design Patterns

### Singleton pattern
```go
import "sync"

type Singleton struct {
    data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{data: "singleton"}
    })
    return instance
}
```

### Factory pattern
```go
type Animal interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow" }

func AnimalFactory(animalType string) Animal {
    switch animalType {
    case "dog":
        return Dog{}
    case "cat":
        return Cat{}
    default:
        return nil
    }
}

// Usage
animal := AnimalFactory("dog")
fmt.Println(animal.Speak())
```

---

## Logging

### Different log levels
```go
import "log"

// Standard logging
log.Println("Info: Application started")

// With prefix
logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
logger.Println("This is an info message")

// Fatal (logs and exits)
log.Fatal("Fatal error occurred")

// Panic (logs and panics)
log.Panic("Panic error occurred")
```

### Using slog (structured logging - Go 1.21+)
```go
import "log/slog"

// Info level
slog.Info("User logged in", "user", "alice", "ip", "192.168.1.1")

// Debug level
slog.Debug("Debug message")

// Warning level
slog.Warn("Warning message", "reason", "high memory usage")

// Error level
slog.Error("Error occurred", "error", err)
```

---

## Time & Date

### Current time
```go
import "time"

now := time.Now()
```

### Format time
```go
// Format: "2006-01-02 15:04:05" is the reference time
formatted := now.Format("2006-01-02 15:04:05")
formatted = now.Format("Jan 02, 2006")
formatted = now.Format(time.RFC3339)
```

### Parse time
```go
t, err := time.Parse("2006-01-02", "2024-03-15")
```

### Duration
```go
duration := 5 * time.Second
time.Sleep(duration)

// Calculate difference
diff := time.Since(startTime)
elapsed := time.Until(futureTime)
```

### Add/Subtract time
```go
tomorrow := now.Add(24 * time.Hour)
lastWeek := now.AddDate(0, 0, -7)
```

---

## Context

### Context with timeout
```go
import (
    "context"
    "time"
)

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use in HTTP request
req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
```

### Context with cancellation
```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    // Some work
    cancel()  // Cancel when done
}()

select {
case <-ctx.Done():
    fmt.Println("Context cancelled:", ctx.Err())
}
```

### Context with value
```go
ctx := context.WithValue(context.Background(), "userID", 12345)

// Retrieve value
if userID, ok := ctx.Value("userID").(int); ok {
    fmt.Println("User ID:", userID)
}
```

---

## Testing

### Basic test
```go
import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
```

### Table-driven test
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, result)
            }
        })
    }
}
```

### Benchmark
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

---

## Environment Variables

### Read environment variable
```go
import "os"

value := os.Getenv("DATABASE_URL")
if value == "" {
    value = "default_value"
}
```

### Set environment variable
```go
os.Setenv("API_KEY", "secret123")
```

### Using godotenv
```go
import "github.com/joho/godotenv"

err := godotenv.Load()  // Loads .env file
if err != nil {
    log.Fatal("Error loading .env file")
}

dbURL := os.Getenv("DATABASE_URL")
```

---

## Input/Output

### Accept user input (scanf equivalent)
```go
import "fmt"

var name string
var age int

fmt.Print("Enter name: ")
fmt.Scan(&name)

fmt.Print("Enter age: ")
fmt.Scan(&age)
```

### Scan full line with spaces
```go
import (
    "bufio"
    "os"
)

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
```

### Print formatting
```go
name := "Alice"
age := 30
price := 19.99

fmt.Printf("Name: %s\n", name)           // String
fmt.Printf("Age: %d\n", age)             // Integer
fmt.Printf("Price: %.2f\n", price)       // Float with 2 decimals
fmt.Printf("Value: %v\n", person)        // Default format
fmt.Printf("Type: %T\n", person)         // Type
fmt.Printf("Pointer: %p\n", &person)     // Pointer
fmt.Printf("Boolean: %t\n", true)        // Boolean
fmt.Printf("Binary: %b\n", 10)           // Binary
fmt.Printf("Hex: %x\n", 255)             // Hexadecimal
```

---

## Control Flow

### Comma ok idiom
```go
// Map lookup
if value, ok := myMap["key"]; ok {
    fmt.Println("Found:", value)
}

// Type assertion
if str, ok := myInterface.(string); ok {
    fmt.Println("Is string:", str)
}

// Channel receive
if value, ok := <-myChan; ok {
    fmt.Println("Received:", value)
}
```

### Defer execution order
```go
func example() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    // Prints: Third, Second, First
}
```

### Interfaces
```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}

// Empty interface
var anything interface{} = "can hold any type"
```

---

## Make Command

### Common Makefile targets
```makefile
# Build the project
.PHONY: build
build:
	go build -o bin/app ./cmd/app

# Run the application
.PHONY: run
run:
	go run ./cmd/app

# Run tests
.PHONY: test
test:
	go test ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -cover ./...

# Run tests with verbose output
.PHONY: test-verbose
test-verbose:
	go test -v ./...

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy

# Download dependencies
.PHONY: deps
deps:
	go mod download

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf bin/

# Install tools
.PHONY: tools
tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Build for multiple platforms
.PHONY: build-all
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/app-linux ./cmd/app
	GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o bin/app.exe ./cmd/app

# Run with hot reload (using air)
.PHONY: dev
dev:
	air

# Generate code
.PHONY: generate
generate:
	go generate ./...

# Docker build
.PHONY: docker-build
docker-build:
	docker build -t myapp:latest .

# Docker run
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 myapp:latest
```

---

## Additional Notes

- Always handle errors appropriately in production code
- Use `defer` for cleanup operations (closing files, connections, etc.)
- Prefer slices over arrays for flexibility
- Use `context.Context` for cancellation and timeouts in long-running operations
- Follow Go naming conventions (exported names start with capital letter)
- Use `go fmt` to format code consistently
- Run `go vet` to catch common mistakes
- Consider using `golangci-lint` for comprehensive linting

---

**Generated for quick reference. Test all snippets in your environment.**
