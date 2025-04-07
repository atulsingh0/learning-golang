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
	orders1 := generateOrders(15)

	seqRun(orders1)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("")

	// create copy of orders
	orders2 := make([]*Order, len(orders1))
	copy(orders2, orders1)
	processOrdersRun(orders2)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("")

	// create copy of orders
	orders3 := make([]*Order, len(orders1))
	copy(orders3, orders1)
	processNupdateOrdersRun(orders3)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("")
}

func seqRun(orders []*Order) {
	start := time.Now()
	processOrders(orders)
	updateOrderStatus(orders)
	reportOrders(orders)
	fmt.Printf("Sequential Execution: %v\n", time.Since(start))
}

func processOrdersRun(orders []*Order) {
	start := time.Now()
	go processOrders(orders)
	updateOrderStatus(orders)
	reportOrders(orders)
	fmt.Printf("Concurrent Execution: processOrders: %v\n", time.Since(start))
}

func processNupdateOrdersRun(orders []*Order) {
	start := time.Now()
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
		order.Status = []string{"PROCESSING", "DELIVERED", "TRANSIT"}[rand.Intn(3)]
		//fmt.Printf("Order %d status updated: %s\n", order.ID, order.Status)
	}
}

func processOrder(order *Order) {
	time.Sleep(time.Duration(order.Amount) * time.Millisecond)
	//fmt.Printf("Order %d processed\n", order.ID)
}

func reportOrders(orders []*Order) {
	fmt.Println("Reporting orders")
	fmt.Println("=================")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}
}
