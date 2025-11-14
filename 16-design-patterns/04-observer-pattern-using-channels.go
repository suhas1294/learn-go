package main

import (
	"fmt"
	"sync"
	"time"
)

// ============= TRADITIONAL OBSERVER (OOP Style) =============

type StockObserver interface {
	Update(symbol string, price float64)
}

type StockExchange struct {
	mu        sync.RWMutex
	observers map[string][]StockObserver
	prices    map[string]float64
}

func NewStockExchange() *StockExchange {
	return &StockExchange{
		observers: make(map[string][]StockObserver),
		prices:    make(map[string]float64),
	}
}

func (s *StockExchange) Subscribe(symbol string, observer StockObserver) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers[symbol] = append(s.observers[symbol], observer)
}

func (s *StockExchange) UpdatePrice(symbol string, price float64) {
	s.mu.Lock()
	s.prices[symbol] = price
	observers := s.observers[symbol]
	s.mu.Unlock()

	// Notify observers
	for _, observer := range observers {
		observer.Update(symbol, price)
	}
}

type TradingBot struct {
	name         string
	buyThreshold float64
}

func (t *TradingBot) Update(symbol string, price float64) {
	if price < t.buyThreshold {
		fmt.Printf("[%s] BUY signal for %s at $%.2f\n", t.name, symbol, price)
	}
}

type PriceAlertService struct{}

func (p *PriceAlertService) Update(symbol string, price float64) {
	fmt.Printf("[Alert] %s price updated to $%.2f\n", symbol, price)
}

// ============= CHANNEL-BASED OBSERVER (Idiomatic Go) =============
// This is more Go-like - uses channels and goroutines

type PriceUpdate struct {
	Symbol string
	Price  float64
	Time   time.Time
}

type ChannelStockExchange struct {
	mu          sync.RWMutex
	prices      map[string]float64
	subscribers map[string][]chan PriceUpdate
}

func NewChannelStockExchange() *ChannelStockExchange {
	return &ChannelStockExchange{
		prices:      make(map[string]float64),
		subscribers: make(map[string][]chan PriceUpdate),
	}
}

func (c *ChannelStockExchange) Subscribe(symbol string) <-chan PriceUpdate {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch := make(chan PriceUpdate, 10) // Buffered channel
	c.subscribers[symbol] = append(c.subscribers[symbol], ch)
	return ch
}

func (c *ChannelStockExchange) UpdatePrice(symbol string, price float64) {
	c.mu.Lock()
	c.prices[symbol] = price
	subscribers := c.subscribers[symbol]
	c.mu.Unlock()

	update := PriceUpdate{
		Symbol: symbol,
		Price:  price,
		Time:   time.Now(),
	}

	// Non-blocking send to all subscribers
	for _, ch := range subscribers {
		select {
		case ch <- update:
		default:
			fmt.Printf("Warning: Subscriber channel full for %s\n", symbol)
		}
	}
}

func (c *ChannelStockExchange) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channels := range c.subscribers {
		for _, ch := range channels {
			close(ch)
		}
	}
}

// ============= USING SELECT FOR MULTIPLE STREAMS =============

type MultiStockMonitor struct {
	stocks []string
}

func (m *MultiStockMonitor) Monitor(exchange *ChannelStockExchange) {
	channels := make(map[string]<-chan PriceUpdate)

	// Subscribe to multiple stocks
	for _, stock := range m.stocks {
		channels[stock] = exchange.Subscribe(stock)
	}

	// Monitor all stocks concurrently with select
	for {
		for stock, ch := range channels {
			select {
			case update, ok := <-ch:
				if !ok {
					fmt.Printf("Channel closed for %s\n", stock)
					delete(channels, stock)
					continue
				}
				fmt.Printf("[Monitor] %s: $%.2f at %s\n",
					update.Symbol, update.Price, update.Time.Format("15:04:05"))
			default:
				// Non-blocking
			}
		}

		if len(channels) == 0 {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Traditional Observer ===")
	exchange := NewStockExchange()

	bot1 := &TradingBot{name: "Bot1", buyThreshold: 100}
	bot2 := &TradingBot{name: "Bot2", buyThreshold: 95}
	alert := &PriceAlertService{}

	exchange.Subscribe("AAPL", bot1)
	exchange.Subscribe("AAPL", bot2)
	exchange.Subscribe("AAPL", alert)

	exchange.UpdatePrice("AAPL", 105)
	exchange.UpdatePrice("AAPL", 98)
	exchange.UpdatePrice("AAPL", 92)

	fmt.Println("\n=== Channel-based Observer ===")
	channelExchange := NewChannelStockExchange()

	// Subscriber 1: Trading bot
	go func() {
		aaplCh := channelExchange.Subscribe("AAPL")
		for update := range aaplCh {
			if update.Price < 100 {
				fmt.Printf("[TradingBot] BUY %s at $%.2f\n", update.Symbol, update.Price)
			}
		}
	}()

	// Subscriber 2: Logger
	go func() {
		aaplCh := channelExchange.Subscribe("AAPL")
		for update := range aaplCh {
			fmt.Printf("[Logger] Recorded %s: $%.2f\n", update.Symbol, update.Price)
		}
	}()

	// Publisher
	time.Sleep(100 * time.Millisecond)
	channelExchange.UpdatePrice("AAPL", 105)
	time.Sleep(100 * time.Millisecond)
	channelExchange.UpdatePrice("AAPL", 98)
	time.Sleep(100 * time.Millisecond)
	channelExchange.UpdatePrice("AAPL", 92)

	time.Sleep(500 * time.Millisecond)
	channelExchange.Close()
	time.Sleep(100 * time.Millisecond)
}
