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
	//mu     sync.Mutex
}

var orderChan = make(chan *Order)

func main() {
	var wg sync.WaitGroup

	// running 2 routines
	wg.Add(2)

	start := time.Now()

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}
		close(orderChan)
		fmt.Println("All orders generated")
	}()

	go updateOrdersStatus(orderChan, &wg)

	wg.Wait()
	fmt.Printf("Concurrent Execution: processOrders & updateOrderStatus: %v\n", time.Since(start))

}

func generateOrders(n int) []*Order {
	orders := make([]*Order, n)
	for i := 0; i < n; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		orders[i] = &Order{
			ID:     i + 1,
			Amount: rand.Intn(500),
			Status: "NEW",
		}
		fmt.Println("Order", orders[i].ID, "generated")
	}
	return orders
}

func updateOrdersStatus(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		time.Sleep(time.Duration(order.Amount) * time.Millisecond)
		order.Status = []string{"Delivered", "InTransit", "Processed", "AtHub"}[rand.Intn(4)]
		fmt.Printf("Order %d status updated: %s\n", order.ID, order.Status)
	}
}

func reportOrders(orders []*Order) {
	fmt.Println("Reporting orders")
	fmt.Println("=================")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}

}
