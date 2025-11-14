package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ============= DOMAIN MODEL =============

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// ============= REPOSITORY INTERFACE =============

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}

// ============= IN-MEMORY IMPLEMENTATION =============

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) FindAll(ctx context.Context) ([]*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.ID == "" {
		return errors.New("user ID cannot be empty")
	}

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

// ============= POSTGRES IMPLEMENTATION =============

type PostgresUserRepository struct {
	connString string
	// db *sql.DB (in real implementation)
}

func NewPostgresUserRepository(connString string) *PostgresUserRepository {
	return &PostgresUserRepository{
		connString: connString,
	}
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	fmt.Printf("SELECT * FROM users WHERE id = '%s'\n", id)
	// db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id)
	return &User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	fmt.Printf("SELECT * FROM users WHERE email = '%s'\n", email)
	return &User{ID: "1", Name: "John", Email: email}, nil
}

func (r *PostgresUserRepository) FindAll(ctx context.Context) ([]*User, error) {
	fmt.Println("SELECT * FROM users")
	return []*User{}, nil
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *User) error {
	fmt.Printf("INSERT INTO users VALUES ('%s', '%s', '%s')\n", user.ID, user.Name, user.Email)
	return nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id string) error {
	fmt.Printf("DELETE FROM users WHERE id = '%s'\n", id)
	return nil
}

// ============= CACHING DECORATOR (Composition) =============

type CachedUserRepository struct {
	repo  UserRepository
	cache map[string]*User
	mu    sync.RWMutex
	ttl   time.Duration
}

func NewCachedUserRepository(repo UserRepository, ttl time.Duration) *CachedUserRepository {
	return &CachedUserRepository{
		repo:  repo,
		cache: make(map[string]*User),
		ttl:   ttl,
	}
}

func (c *CachedUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
	// Check cache first
	c.mu.RLock()
	if user, exists := c.cache[id]; exists {
		c.mu.RUnlock()
		fmt.Println("Cache hit!")
		return user, nil
	}
	c.mu.RUnlock()

	// Cache miss - fetch from repo
	fmt.Println("Cache miss - fetching from database")
	user, err := c.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update cache
	c.mu.Lock()
	c.cache[id] = user
	c.mu.Unlock()

	// Schedule cache invalidation
	go func() {
		time.Sleep(c.ttl)
		c.mu.Lock()
		delete(c.cache, id)
		c.mu.Unlock()
	}()

	return user, nil
}

func (c *CachedUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	return c.repo.FindByEmail(ctx, email)
}

func (c *CachedUserRepository) FindAll(ctx context.Context) ([]*User, error) {
	return c.repo.FindAll(ctx)
}

func (c *CachedUserRepository) Save(ctx context.Context, user *User) error {
	// Invalidate cache on save
	c.mu.Lock()
	delete(c.cache, user.ID)
	c.mu.Unlock()

	return c.repo.Save(ctx, user)
}

func (c *CachedUserRepository) Delete(ctx context.Context, id string) error {
	// Invalidate cache on delete
	c.mu.Lock()
	delete(c.cache, id)
	c.mu.Unlock()

	return c.repo.Delete(ctx, id)
}

// ============= SERVICE LAYER =============

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email string) (*User, error) {
	// Check if email already exists
	existing, _ := s.repo.FindByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	// Create new user
	user := &User{
		ID:        generateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func generateID() string {
	return fmt.Sprintf("user_%d", time.Now().UnixNano())
}

func main() {
	ctx := context.Background()

	fmt.Println("=== In-Memory Repository ===")
	memRepo := NewInMemoryUserRepository()
	service := NewUserService(memRepo)

	user, _ := service.RegisterUser(ctx, "John Doe", "john@example.com")
	fmt.Printf("Registered: %+v\n", user)

	retrieved, _ := service.GetUser(ctx, user.ID)
	fmt.Printf("Retrieved: %+v\n", retrieved)

	fmt.Println("\n=== Postgres Repository ===")
	pgRepo := NewPostgresUserRepository("postgres://localhost/mydb")
	pgService := NewUserService(pgRepo)

	pgService.RegisterUser(ctx, "Jane Doe", "jane@example.com")
	pgService.GetUser(ctx, "1")

	fmt.Println("\n=== Cached Repository ===")
	cachedRepo := NewCachedUserRepository(memRepo, 5*time.Second)
	cachedService := NewUserService(cachedRepo)

	cachedService.GetUser(ctx, user.ID) // Cache miss
	cachedService.GetUser(ctx, user.ID) // Cache hit
	cachedService.GetUser(ctx, user.ID) // Cache hit
}
