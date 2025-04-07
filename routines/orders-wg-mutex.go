package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount int
	Status string
	mu     sync.Mutex
}

var totalUpdates int
var updateMutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	orders := generateOrders(20)

	wg.Add(1)
	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	// this is updating status 3 times = 3*20 = 60
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateOrdersStatus(orders)
		}()
	}

	wg.Wait()
	reportOrders(orders)

	fmt.Printf("Concurrent Execution: processOrders & updateOrderStatus: %v\n", time.Since(start))

}

func generateOrders(n int) []*Order {
	orders := make([]*Order, n)
	for i := 0; i < n; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Amount: rand.Intn(300),
			Status: "NEW",
		}
	}
	return orders
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		processOrder(order)
	}
}

func processOrder(order *Order) {
	time.Sleep(time.Duration(order.Amount) * time.Millisecond)
	order.Status = "Received"
	fmt.Printf("Order %d processed: %s\n", order.ID, order.Status)
}

func updateOrdersStatus(orders []*Order) {
	for _, order := range orders {
		updateOrderStatus(order)
	}
}

func updateOrderStatus(order *Order) {
	order.mu.Lock()
	time.Sleep(time.Duration(order.Amount) * time.Millisecond)
	order.Status = []string{"Delivered", "InTransit", "Processed", "AtHub"}[rand.Intn(4)]
	fmt.Printf("Order %d status updated: %s\n", order.ID, order.Status)
	defer order.mu.Unlock()

	//updateMutex.Lock()
	//defer updateMutex.Unlock()
	curentUpdate := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = curentUpdate + 1
}

func reportOrders(orders []*Order) {
	fmt.Println("Reporting orders")
	fmt.Println("=================")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}

	fmt.Println("Total Updates: ", totalUpdates)
}
