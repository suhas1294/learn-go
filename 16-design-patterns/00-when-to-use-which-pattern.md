## When to Use Which Paradigm in Go?

### **OOP-Style (Interfaces + Structs)**
```go
type PaymentProcessor interface {
    Process(amount float64) error
}

type StripeProcessor struct {
    apiKey string
}

func (s *StripeProcessor) Process(amount float64) error {
    // Implementation
}
```

**Use when:**
- ✅ Need polymorphism (multiple implementations of same interface)
- ✅ State management is important (struct fields)
- ✅ Clear domain modeling (User, Order, Payment entities)
- ✅ Building libraries/frameworks that others will extend
- ✅ Testing with mocks/stubs (easy to mock interfaces)

**Examples:** Repositories, Services, Domain models

---

### **Functional Style (Functions + Closures)**
```go
type Middleware func(HandlerFunc) HandlerFunc

func AuthMiddleware(next HandlerFunc) HandlerFunc {
    return func(req *Request) *Response {
        // Implementation
    }
}
```

**Use when:**
- ✅ Composing behavior dynamically
- ✅ Middleware/decorator patterns
- ✅ Minimal state needed
- ✅ Higher-order functions make sense
- ✅ Pipeline transformations

**Examples:** HTTP middleware, data pipelines, event handlers

---

### **Hybrid Approach (Best of Both)**
```go
// Interface for polymorphism
type Storage interface {
    Save(data string) error
}

// Struct for state
type FileStorage struct {
    path string
}

// Method for behavior
func (f *FileStorage) Save(data string) error {
    return nil
}

// Functional options for configuration
type StorageOption func(*FileStorage)

func WithCompression() StorageOption {
    return func(s *FileStorage) {
        // Enable compression
    }
}

func NewFileStorage(path string, opts ...StorageOption) *FileStorage {
    s := &FileStorage{path: path}
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

**Use when:**
- ✅ Need both state AND composability
- ✅ Complex configuration (functional options pattern)
- ✅ Real-world production code

**This is the most common in Go standard library and popular frameworks**

---

## **Decision Matrix**

| **Pattern** | **Go Idiom** | **Why?** |
|-------------|--------------|----------|
| **Strategy** | Interface-based | Polymorphism, testable |
| **Factory** | Function returning interface | Simple, no classes needed |
| **Builder** | Functional options | More flexible than traditional builder |
| **Observer** | Channels + goroutines | Leverage Go's concurrency primitives |
| **Decorator** | Middleware functions | Composable, used everywhere in Go |
| **Chain of Responsibility** | Middleware chain | HTTP handlers, processing pipelines |
| **Repository** | Interface + multiple implementations | Testability, swap implementations |

---

## **Go Philosophy**

1. **Composition over inheritance** - Go has no inheritance
2. **Interfaces are implicit** - No "implements" keyword
3. **Small interfaces** - Prefer `interface { Method() }` over large interfaces
4. **Channels for communication** - "Share memory by communicating"
5. **Simplicity** - Don't overengineer with patterns

**Rob Pike's advice:** "The bigger the interface, the weaker the abstraction"

Use patterns when they solve real problems, not for the sake of patterns!
