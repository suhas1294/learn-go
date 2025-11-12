# HTTP Request Handling in Go

Complete guide to handling HTTP requests with and without frameworks, implementing middleware, and understanding execution flow.

---

## Table of Contents
1. [Handling Requests Without Libraries](#handling-requests-without-libraries)
2. [Handling Requests With Gin Framework](#handling-requests-with-gin-framework)
3. [Middleware Without Libraries](#middleware-without-libraries)
4. [Middleware With Gin](#middleware-with-gin)
5. [Middleware Execution Workflow (Without Libraries)](#middleware-execution-workflow-without-libraries)
6. [Middleware Execution Workflow (With Gin)](#middleware-execution-workflow-with-gin)
7. [Pattern Matching in Go](#pattern-matching-in-go)

---

## Handling Requests Without Libraries

### Basic HTTP Server
```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Register handler for specific route
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/users", usersHandler)

    // Start server on port 8080
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        w.Write([]byte("Get all users"))
    } else if r.Method == "POST" {
        w.Write([]byte("Create user"))
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Method not allowed"))
    }
}
```

### Handling Different HTTP Methods
```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        w.Write([]byte("GET: Fetch user"))
    case http.MethodPost:
        w.Write([]byte("POST: Create user"))
    case http.MethodPut:
        w.Write([]byte("PUT: Update user"))
    case http.MethodDelete:
        w.Write([]byte("DELETE: Remove user"))
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

### Reading Request Body and Query Parameters
```go
import (
    "encoding/json"
    "io"
    "net/http"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    // Read query parameters
    id := r.URL.Query().Get("id")
    active := r.URL.Query().Get("active")

    // Read request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Cannot read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Parse JSON
    var user User
    err = json.Unmarshal(body, &user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Send JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User created",
        "user":    user,
        "id":      id,
        "active":  active,
    })
}
```

### Using Custom ServeMux (Router)
```go
func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/users", usersHandler)
    mux.HandleFunc("/api/products", productsHandler)
    mux.HandleFunc("/health", healthHandler)

    http.ListenAndServe(":8080", mux)
}
```

---

## Handling Requests With Gin Framework

### Installation
```bash
go get -u github.com/gin-gonic/gin
```

### Basic Gin Server
```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create Gin router
    r := gin.Default()

    // Simple route
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Start server
    r.Run(":8080")
}
```

### Handling Different HTTP Methods
```go
func main() {
    r := gin.Default()

    r.GET("/users", getUsers)
    r.POST("/users", createUser)
    r.PUT("/users/:id", updateUser)
    r.DELETE("/users/:id", deleteUser)

    r.Run(":8080")
}

func getUsers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Get all users"})
}

func createUser(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func updateUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "User updated", "id": id})
}

func deleteUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
}
```

### Reading Request Data with Gin
```go
type User struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=130"`
}

func createUserHandler(c *gin.Context) {
    // Read query parameters
    active := c.Query("active")
    page := c.DefaultQuery("page", "1")

    // Read path parameters
    id := c.Param("id")

    // Bind JSON to struct with validation
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Send response
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created",
        "user":    user,
        "active":  active,
        "page":    page,
    })
}
```

### Route Grouping
```go
func main() {
    r := gin.Default()

    // API v1 group
    v1 := r.Group("/api/v1")
    {
        v1.GET("/users", getUsers)
        v1.POST("/users", createUser)

        // Nested group
        admin := v1.Group("/admin")
        {
            admin.GET("/dashboard", adminDashboard)
            admin.POST("/settings", updateSettings)
        }
    }

    r.Run(":8080")
}
```

---

## Middleware Without Libraries

### Basic Middleware Pattern
```go
// Middleware is a function that wraps an http.Handler
type Middleware func(http.Handler) http.Handler

// Logger middleware
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Authentication middleware
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        fmt.Println("Auth token:", token)
        next.ServeHTTP(w, r)
    })
}

// CORS middleware
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

### Applying Single Middleware
```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", usersHandler)

    // Wrap handler with middleware
    http.ListenAndServe(":8080", LoggerMiddleware(mux))
}
```

### Chaining Multiple Middleware (Method 1: Manual)
```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", usersHandler)

    // Chain: CORS -> Auth -> Logger -> Handler
    wrapped := CORSMiddleware(AuthMiddleware(LoggerMiddleware(mux)))

    http.ListenAndServe(":8080", wrapped)
}
```

### Chaining Multiple Middleware (Method 2: Chain Helper)
```go
// Chain helper function
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
    // Apply middlewares in reverse order
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler)
    }
    return handler
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", usersHandler)

    // Chain middlewares
    wrapped := Chain(mux, LoggerMiddleware, AuthMiddleware, CORSMiddleware)

    http.ListenAndServe(":8080", wrapped)
}
```

### Per-Route Middleware
```go
func main() {
    mux := http.NewServeMux()

    // Public route (no auth)
    mux.Handle("/public", LoggerMiddleware(http.HandlerFunc(publicHandler)))

    // Protected route (with auth)
    mux.Handle("/protected", Chain(
        http.HandlerFunc(protectedHandler),
        LoggerMiddleware,
        AuthMiddleware,
    ))

    http.ListenAndServe(":8080", mux)
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Public content"))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Protected content"))
}
```

### Advanced Middleware with Context
```go
import (
    "context"
    "net/http"
)

type contextKey string

const UserContextKey contextKey = "user"

// Middleware that adds user to context
func UserMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract user from token/session
        user := "john_doe"

        // Add to context
        ctx := context.WithValue(r.Context(), UserContextKey, user)

        // Pass request with new context
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Handler that reads from context
func profileHandler(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value(UserContextKey).(string)
    w.Write([]byte("User: " + user))
}
```

---

## Middleware With Gin

### Built-in Middleware
```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.New()

    // Use built-in middleware
    r.Use(gin.Logger())   // Logger middleware
    r.Use(gin.Recovery()) // Recovery middleware (handles panics)

    r.GET("/users", getUsers)

    r.Run(":8080")
}
```

### Custom Middleware in Gin
```go
// Logger middleware
func CustomLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path

        // Process request
        c.Next()

        // After request
        latency := time.Since(start)
        statusCode := c.Writer.Status()
        fmt.Printf("[%s] %s %s %d %v\n",
            time.Now().Format("2006-01-02 15:04:05"),
            c.Request.Method,
            path,
            statusCode,
            latency,
        )
    }
}

// Auth middleware
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")

        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort() // Stop execution
            return
        }

        // Validate token
        if token != "valid-token" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Add user to context
        c.Set("userID", "12345")

        c.Next() // Continue to next middleware/handler
    }
}

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }

        c.Next()
    }
}
```

### Applying Global Middleware
```go
func main() {
    r := gin.New()

    // Apply to all routes
    r.Use(CustomLogger())
    r.Use(CORSMiddleware())
    r.Use(AuthMiddleware())

    r.GET("/users", getUsers)
    r.POST("/users", createUser)

    r.Run(":8080")
}
```

### Applying Middleware to Route Groups
```go
func main() {
    r := gin.Default()

    // Public routes (no auth)
    r.GET("/health", healthCheck)

    // Protected routes (with auth)
    protected := r.Group("/api")
    protected.Use(AuthMiddleware())
    {
        protected.GET("/users", getUsers)
        protected.POST("/users", createUser)
    }

    // Admin routes (with auth + admin check)
    admin := r.Group("/admin")
    admin.Use(AuthMiddleware(), AdminMiddleware())
    {
        admin.GET("/dashboard", adminDashboard)
        admin.DELETE("/users/:id", deleteUser)
    }

    r.Run(":8080")
}

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("userID")

        // Check if user is admin
        if userID != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            c.Abort()
            return
        }

        c.Next()
    }
}
```

### Per-Route Middleware
```go
func main() {
    r := gin.Default()

    // Apply middleware to specific route only
    r.GET("/protected", AuthMiddleware(), RateLimitMiddleware(), protectedHandler)

    // Multiple routes with same middleware
    r.POST("/users", AuthMiddleware(), createUser)
    r.PUT("/users/:id", AuthMiddleware(), updateUser)

    r.Run(":8080")
}

func RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Rate limiting logic
        c.Next()
    }
}
```

---

## Middleware Execution Workflow (Without Libraries)

### Complete Example
```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

type Middleware func(http.Handler) http.Handler

// Middleware 1: Logger
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("→ [Logger] Before handler - Request received")
        start := time.Now()

        next.ServeHTTP(w, r) // Call next middleware/handler

        fmt.Printf("← [Logger] After handler - Took %v\n", time.Since(start))
    })
}

// Middleware 2: Auth
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("  → [Auth] Before handler - Checking token")

        token := r.Header.Get("Authorization")
        if token == "" {
            fmt.Println("  ✗ [Auth] No token - Stopping execution")
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return // Stop here, don't call next
        }

        fmt.Println("  ✓ [Auth] Token valid - Continuing")
        next.ServeHTTP(w, r)

        fmt.Println("  ← [Auth] After handler")
    })
}

// Middleware 3: CORS
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("    → [CORS] Before handler - Setting headers")

        w.Header().Set("Access-Control-Allow-Origin", "*")

        next.ServeHTTP(w, r)

        fmt.Println("    ← [CORS] After handler")
    })
}

// Chain helper
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler)
    }
    return handler
}

// Final handler
func usersHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("      ⚙ [Handler] Processing request")
    w.Write([]byte("Users data"))
    fmt.Println("      ⚙ [Handler] Response sent")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", usersHandler)

    // Chain: Logger -> Auth -> CORS -> Handler
    wrapped := Chain(mux, LoggerMiddleware, AuthMiddleware, CORSMiddleware)

    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", wrapped)
}
```

### Execution Flow (Step-by-Step)

#### Request: `GET /users` with Authorization header

```
Step 1: Request arrives
        ↓
Step 2: LoggerMiddleware receives request
        → Prints: "→ [Logger] Before handler"
        → Calls next.ServeHTTP() → Goes to AuthMiddleware
        ↓
Step 3: AuthMiddleware receives request
        → Prints: "  → [Auth] Before handler - Checking token"
        → Token found: "✓ [Auth] Token valid"
        → Calls next.ServeHTTP() → Goes to CORSMiddleware
        ↓
Step 4: CORSMiddleware receives request
        → Prints: "    → [CORS] Before handler"
        → Sets CORS headers
        → Calls next.ServeHTTP() → Goes to Handler
        ↓
Step 5: usersHandler receives request
        → Prints: "      ⚙ [Handler] Processing request"
        → Processes business logic
        → Sends response: "Users data"
        → Prints: "      ⚙ [Handler] Response sent"
        → Returns to CORSMiddleware
        ↓
Step 6: Back in CORSMiddleware (after next.ServeHTTP())
        → Prints: "    ← [CORS] After handler"
        → Returns to AuthMiddleware
        ↓
Step 7: Back in AuthMiddleware (after next.ServeHTTP())
        → Prints: "  ← [Auth] After handler"
        → Returns to LoggerMiddleware
        ↓
Step 8: Back in LoggerMiddleware (after next.ServeHTTP())
        → Calculates duration
        → Prints: "← [Logger] After handler - Took 245µs"
        → Returns to client
        ↓
Step 9: Response sent to client
```

**Console Output:**
```
→ [Logger] Before handler - Request received
  → [Auth] Before handler - Checking token
  ✓ [Auth] Token valid - Continuing
    → [CORS] Before handler - Setting headers
      ⚙ [Handler] Processing request
      ⚙ [Handler] Response sent
    ← [CORS] After handler
  ← [Auth] After handler
← [Logger] After handler - Took 245µs
```

#### Request: `GET /users` WITHOUT Authorization header

```
Step 1: Request arrives
        ↓
Step 2: LoggerMiddleware receives request
        → Prints: "→ [Logger] Before handler"
        → Calls next.ServeHTTP() → Goes to AuthMiddleware
        ↓
Step 3: AuthMiddleware receives request
        → Prints: "  → [Auth] Before handler - Checking token"
        → No token found: "✗ [Auth] No token"
        → Returns error (401 Unauthorized)
        → DOES NOT call next.ServeHTTP()
        → Returns to LoggerMiddleware
        ↓
Step 4: Back in LoggerMiddleware
        → Prints: "← [Logger] After handler - Took 89µs"
        → Returns to client
        ↓
Step 5: Response sent to client (401 Unauthorized)
```

**Console Output:**
```
→ [Logger] Before handler - Request received
  → [Auth] Before handler - Checking token
  ✗ [Auth] No token - Stopping execution
← [Logger] After handler - Took 89µs
```

**Key Points:**
- Middleware executes in order: Logger → Auth → CORS → Handler
- After handler completes, execution returns back through each middleware
- If middleware doesn't call `next.ServeHTTP()`, execution stops
- Each middleware can run code before AND after the next handler

---

## Middleware Execution Workflow (With Gin)

### Complete Example
```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// Middleware 1: Logger
func CustomLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("→ [Logger] Before handler")
        start := time.Now()

        c.Next() // Execute remaining handlers

        latency := time.Since(start)
        statusCode := c.Writer.Status()
        fmt.Printf("← [Logger] After handler - Status: %d, Latency: %v\n", statusCode, latency)
    }
}

// Middleware 2: Auth
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("  → [Auth] Before handler - Checking token")

        token := c.GetHeader("Authorization")
        if token == "" {
            fmt.Println("  ✗ [Auth] No token - Aborting")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort() // Stop execution, skip remaining handlers
            return
        }

        if token != "valid-token" {
            fmt.Println("  ✗ [Auth] Invalid token - Aborting")
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        fmt.Println("  ✓ [Auth] Token valid - Setting userID in context")
        c.Set("userID", "12345")

        c.Next() // Continue to next middleware/handler

        fmt.Println("  ← [Auth] After handler")
    }
}

// Middleware 3: CORS
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("    → [CORS] Before handler - Setting headers")

        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

        c.Next()

        fmt.Println("    ← [CORS] After handler")
    }
}

// Handler
func getUsers(c *gin.Context) {
    fmt.Println("      ⚙ [Handler] Processing request")

    userID := c.GetString("userID")
    fmt.Printf("      ⚙ [Handler] User ID from context: %s\n", userID)

    c.JSON(http.StatusOK, gin.H{
        "message": "Users data",
        "userID":  userID,
    })

    fmt.Println("      ⚙ [Handler] Response sent")
}

func main() {
    r := gin.New()

    // Apply middleware globally
    r.Use(CustomLogger())
    r.Use(AuthMiddleware())
    r.Use(CORSMiddleware())

    r.GET("/users", getUsers)

    fmt.Println("Server starting on :8080")
    r.Run(":8080")
}
```

### Execution Flow (Step-by-Step)

#### Request: `GET /users` with valid Authorization header

```
Step 1: Request arrives at Gin router
        ↓
Step 2: CustomLogger middleware
        → Prints: "→ [Logger] Before handler"
        → Records start time
        → Calls c.Next() → Executes next middleware
        ↓
Step 3: AuthMiddleware
        → Prints: "  → [Auth] Before handler - Checking token"
        → Gets Authorization header: "valid-token"
        → Token is valid
        → Prints: "  ✓ [Auth] Token valid - Setting userID"
        → Sets userID in context: c.Set("userID", "12345")
        → Calls c.Next() → Executes next middleware
        ↓
Step 4: CORSMiddleware
        → Prints: "    → [CORS] Before handler - Setting headers"
        → Sets CORS headers
        → Calls c.Next() → Executes handler
        ↓
Step 5: getUsers handler
        → Prints: "      ⚙ [Handler] Processing request"
        → Gets userID from context: "12345"
        → Prints: "      ⚙ [Handler] User ID from context: 12345"
        → Sends JSON response
        → Prints: "      ⚙ [Handler] Response sent"
        → Returns (c.Next() in CORS completes)
        ↓
Step 6: Back in CORSMiddleware (after c.Next())
        → Prints: "    ← [CORS] After handler"
        → Returns (c.Next() in Auth completes)
        ↓
Step 7: Back in AuthMiddleware (after c.Next())
        → Prints: "  ← [Auth] After handler"
        → Returns (c.Next() in Logger completes)
        ↓
Step 8: Back in CustomLogger (after c.Next())
        → Calculates latency
        → Gets status code: 200
        → Prints: "← [Logger] After handler - Status: 200, Latency: 312µs"
        → Returns to Gin framework
        ↓
Step 9: Gin sends response to client
```

**Console Output:**
```
→ [Logger] Before handler
  → [Auth] Before handler - Checking token
  ✓ [Auth] Token valid - Setting userID in context
    → [CORS] Before handler - Setting headers
      ⚙ [Handler] Processing request
      ⚙ [Handler] User ID from context: 12345
      ⚙ [Handler] Response sent
    ← [CORS] After handler
  ← [Auth] After handler
← [Logger] After handler - Status: 200, Latency: 312µs
```

#### Request: `GET /users` WITHOUT Authorization header

```
Step 1: Request arrives at Gin router
        ↓
Step 2: CustomLogger middleware
        → Prints: "→ [Logger] Before handler"
        → Records start time
        → Calls c.Next() → Executes next middleware
        ↓
Step 3: AuthMiddleware
        → Prints: "  → [Auth] Before handler - Checking token"
        → No Authorization header found
        → Prints: "  ✗ [Auth] No token - Aborting"
        → Sends JSON error: {"error": "Unauthorized"}
        → Calls c.Abort() → Stops middleware chain
        → Returns (does NOT call c.Next())
        → Returns (c.Next() in Logger completes)
        ↓
Step 4: Back in CustomLogger (after c.Next())
        → Calculates latency
        → Gets status code: 401
        → Prints: "← [Logger] After handler - Status: 401, Latency: 156µs"
        → Returns to Gin framework
        ↓
Step 5: Gin sends error response to client (401)
```

**Console Output:**
```
→ [Logger] Before handler
  → [Auth] Before handler - Checking token
  ✗ [Auth] No token - Aborting
← [Logger] After handler - Status: 401, Latency: 156µs
```

**Key Points:**
- `c.Next()` executes all remaining handlers in the chain
- After all handlers complete, execution returns back through each middleware
- `c.Abort()` stops the chain (remaining handlers are skipped)
- Even after `c.Abort()`, code after `c.Next()` still executes
- Context can store and retrieve values across middleware/handlers

---

## Key Differences: Native vs Gin

### Without Libraries (Native Go)
- **Wrapping Pattern**: Each middleware wraps the next handler
- **Type**: `func(http.Handler) http.Handler`
- **Flow Control**: Return early to stop execution
- **Context**: Use `r.Context()` to pass data
- **More Control**: Direct access to ResponseWriter and Request
- **More Verbose**: More boilerplate code

### With Gin
- **Chain Pattern**: Middleware forms a chain with explicit `c.Next()`
- **Type**: `func(*gin.Context)`
- **Flow Control**: Use `c.Abort()` to stop execution
- **Context**: Use `c.Set()` and `c.Get()` for data
- **Convenience**: Helper methods (JSON, Param, Query, etc.)
- **Less Verbose**: Cleaner, more intuitive API

---

## Best Practices

1. **Order matters**: Place authentication before authorization, logging first
2. **Error handling**: Always handle errors in middleware
3. **Performance**: Keep middleware lightweight
4. **Context usage**: Use context to pass request-scoped data
5. **Early returns**: Stop execution early for auth failures
6. **Recovery**: Always use recovery middleware to handle panics
7. **CORS**: Place CORS middleware early in the chain
8. **Testing**: Test middleware independently

---

## Pattern Matching in Go

Go doesn't have built-in pattern matching like Rust or Scala, but provides regex support and alternative approaches.

### Basic Regex Matching

```go
import (
    "fmt"
    "regexp"
)

func main() {
    // Compile regex pattern
    pattern := regexp.MustCompile(`^[a-z]+$`)

    // Check if string matches
    matched := pattern.MatchString("hello")
    fmt.Println(matched) // true

    matched = pattern.MatchString("Hello123")
    fmt.Println(matched) // false
}
```

### Find Matches in String

```go
import (
    "fmt"
    "regexp"
)

func main() {
    text := "My email is john@example.com and jane@test.org"

    // Email pattern
    pattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

    // Find first match
    firstMatch := pattern.FindString(text)
    fmt.Println(firstMatch) // john@example.com

    // Find all matches
    allMatches := pattern.FindAllString(text, -1)
    fmt.Println(allMatches) // [john@example.com jane@test.org]
}
```

### Extract Captured Groups

```go
import (
    "fmt"
    "regexp"
)

func main() {
    text := "Date: 2024-03-15"

    // Pattern with capture groups
    pattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

    // Find submatch
    matches := pattern.FindStringSubmatch(text)
    if len(matches) > 0 {
        fmt.Println("Full match:", matches[0]) // 2024-03-15
        fmt.Println("Year:", matches[1])       // 2024
        fmt.Println("Month:", matches[2])      // 03
        fmt.Println("Day:", matches[3])        // 15
    }
}
```

### Validate Input in HTTP Handler

```go
import (
    "net/http"
    "regexp"
)

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var phonePattern = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

func validateUserHandler(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    phone := r.URL.Query().Get("phone")

    // Validate email
    if !emailPattern.MatchString(email) {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        return
    }

    // Validate phone
    if !phonePattern.MatchString(phone) {
        http.Error(w, "Invalid phone format", http.StatusBadRequest)
        return
    }

    w.Write([]byte("Validation passed"))
}
```

### Route Pattern Matching (Without Libraries)

```go
import (
    "fmt"
    "net/http"
    "regexp"
)

type Route struct {
    Pattern *regexp.Regexp
    Handler http.HandlerFunc
}

type Router struct {
    routes []Route
}

func (rt *Router) AddRoute(pattern string, handler http.HandlerFunc) {
    rt.routes = append(rt.routes, Route{
        Pattern: regexp.MustCompile(pattern),
        Handler: handler,
    })
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    for _, route := range rt.routes {
        if route.Pattern.MatchString(r.URL.Path) {
            route.Handler(w, r)
            return
        }
    }
    http.NotFound(w, r)
}

func main() {
    router := &Router{}

    // Match exact path
    router.AddRoute("^/users$", usersHandler)

    // Match path with ID (numeric)
    router.AddRoute("^/users/\\d+$", userByIDHandler)

    // Match path with slug (alphanumeric)
    router.AddRoute("^/posts/[a-z0-9-]+$", postBySlugHandler)

    http.ListenAndServe(":8080", router)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("All users"))
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("User by ID"))
}

func postBySlugHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Post by slug"))
}
```

### Extract Path Parameters with Regex

```go
import (
    "fmt"
    "net/http"
    "regexp"
)

var userPattern = regexp.MustCompile(`^/users/(\d+)$`)
var postPattern = regexp.MustCompile(`^/posts/([a-z0-9-]+)$`)

func routerHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path

    // Match user by ID
    if matches := userPattern.FindStringSubmatch(path); matches != nil {
        userID := matches[1]
        fmt.Fprintf(w, "User ID: %s", userID)
        return
    }

    // Match post by slug
    if matches := postPattern.FindStringSubmatch(path); matches != nil {
        slug := matches[1]
        fmt.Fprintf(w, "Post slug: %s", slug)
        return
    }

    http.NotFound(w, r)
}

func main() {
    http.HandleFunc("/", routerHandler)
    http.ListenAndServe(":8080", nil)
}

// GET /users/123 → "User ID: 123"
// GET /posts/hello-world → "Post slug: hello-world"
```

### Replace Using Regex

```go
import (
    "fmt"
    "regexp"
)

func main() {
    text := "Contact: john@example.com or call 123-456-7890"

    // Replace emails with [EMAIL]
    emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
    result := emailPattern.ReplaceAllString(text, "[EMAIL]")
    fmt.Println(result)
    // Output: Contact: [EMAIL] or call 123-456-7890

    // Replace phone numbers with [PHONE]
    phonePattern := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
    result = phonePattern.ReplaceAllString(result, "[PHONE]")
    fmt.Println(result)
    // Output: Contact: [EMAIL] or call [PHONE]
}
```

### Split String Using Regex

```go
import (
    "fmt"
    "regexp"
)

func main() {
    text := "apple,banana;orange:grape|mango"

    // Split by multiple delimiters
    pattern := regexp.MustCompile(`[,;:|]`)
    parts := pattern.Split(text, -1)
    fmt.Println(parts)
    // Output: [apple banana orange grape mango]
}
```

### Validate Request Body Fields (Gin)

```go
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "regexp"
)

type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

var (
    usernamePattern = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
    emailPattern    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    passwordPattern = regexp.MustCompile(`^.{8,}$`) // At least 8 chars
)

func createUserHandler(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate username
    if !usernamePattern.MatchString(user.Username) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Username must be 3-20 alphanumeric characters",
        })
        return
    }

    // Validate email
    if !emailPattern.MatchString(user.Email) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid email format",
        })
        return
    }

    // Validate password
    if !passwordPattern.MatchString(user.Password) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Password must be at least 8 characters",
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
```

### Common Regex Patterns

```go
// Email validation
emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Phone number (US format)
phoneRegex := regexp.MustCompile(`^\+?1?\d{10}$`)

// URL validation
urlRegex := regexp.MustCompile(`^https?://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

// IPv4 address
ipv4Regex := regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)

// Date (YYYY-MM-DD)
dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

// Alphanumeric only
alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Password (min 8 chars, 1 uppercase, 1 lowercase, 1 digit)
strongPasswordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`)

// Username (3-20 chars, alphanumeric and underscore)
usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

// Credit card number
creditCardRegex := regexp.MustCompile(`^\d{4}-?\d{4}-?\d{4}-?\d{4}$`)

// UUID
uuidRegex := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)
```

### Performance Tips

```go
// BAD: Compile regex in hot path (every request)
func handler(w http.ResponseWriter, r *http.Request) {
    pattern := regexp.MustCompile(`pattern`) // Compiles every time
    pattern.MatchString("text")
}

// GOOD: Compile once at package level
var pattern = regexp.MustCompile(`pattern`)

func handler(w http.ResponseWriter, r *http.Request) {
    pattern.MatchString("text") // Reuses compiled pattern
}
```

### Alternative: Type Switch (Pattern Matching on Types)

```go
func processValue(val interface{}) {
    switch v := val.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case []int:
        fmt.Printf("Int slice with %d elements\n", len(v))
    case map[string]interface{}:
        fmt.Printf("Map with %d keys\n", len(v))
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Pattern Matching Alternative: Switch on Conditions

```go
func categorizeAge(age int) string {
    switch {
    case age < 0:
        return "Invalid"
    case age < 13:
        return "Child"
    case age < 20:
        return "Teenager"
    case age < 60:
        return "Adult"
    default:
        return "Senior"
    }
}
```

---

**Note**: While Go doesn't have algebraic pattern matching like functional languages, regex provides powerful text pattern matching, and type switches/conditional switches offer structural pattern matching alternatives.

---

**Complete examples demonstrate real execution flow with console output for understanding.**
