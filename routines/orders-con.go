package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Amount int
	Status string
}

func main() {
	start := time.Now()

	orders := generateOrders(15)

	go processOrders(orders)
	go updateOrderStatus(orders)
	reportOrders(orders)
	fmt.Printf("Concurrent Execution: processOrders & updateOrderStatus: %v\n", time.Since(start))

}

// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
func generateOrders(n int) []*Order {
	orders := make([]*Order, n)
	for i := 0; i < n; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Amount: rand.Intn(1000),
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

func updateOrderStatus(order []*Order) {
	for _, order := range order {
		time.Sleep(time.Duration(order.Amount) * time.Millisecond)
		order.Status = []string{"DELIVERED", "TRANSIT"}[rand.Intn(2)]
		fmt.Printf("Order %d status updated: %s\n", order.ID, order.Status)
	}
}

func processOrder(order *Order) {
	time.Sleep(time.Duration(order.Amount) * time.Millisecond)
	order.Status = "PROCESSED"
	fmt.Printf("Order %d processed: %s\n", order.ID, order.Status)
}

func reportOrders(orders []*Order) {
	fmt.Println("Reporting orders")
	fmt.Println("=================")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}
}
