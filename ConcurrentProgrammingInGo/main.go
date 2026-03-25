package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	receivedOrderCh := receiveOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)
	reservedInventoryCh := reserveInventory(validOrderCh)
	fillOrders(reservedInventoryCh, &wg)

	wg.Add(1)

	// go func(validOrderCh <-chan order) {
	// 	order := <-validOrderCh
	// 	fmt.Printf("Valid order received: %v\n", order)
	// 	wg.Done()
	// }(validOrderCh)
	// go func(invalidOrderCh <-chan invalidOrder) {
	// 	order := <-invalidOrderCh
	// 	fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
	// 	wg.Done()
	// }(invalidOrderCh)

	// go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	// loop:
	// 	for {
	// 		select {
	// 		case order, ok := <-validOrderCh:
	// 			if ok {
	// 				fmt.Printf("Valid order received: %v\n", order)
	// 			} else {
	// 				break loop
	// 			}
	// 		case order, ok := <-invalidOrderCh:
	// 			if ok {
	// 				fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
	// 			} else {
	// 				break loop
	// 			}
	// 		}
	// 	}
	// 	wg.Done()
	// }(validOrderCh, invalidOrderCh)

	go func(invalidOrderCh <-chan invalidOrder) {
		for order := range invalidOrderCh {
			fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrderCh)

	// const workers = 3
	// wg.Add(workers)

	// for i := 0; i < workers; i++ {
	// 	go func(reservedInventoryCh <-chan order) {
	// 		for order := range reservedInventoryCh {
	// 			fmt.Printf("Inventory reserved for: %v\n", order)
	// 		}
	// 		wg.Done()
	// 	}(reservedInventoryCh)
	// }

	// go func(fillOrdersCh <-chan order) {
	// 	for order := range fillOrdersCh {
	// 		fmt.Printf("Order has been completed: %v\n", order)
	// 	}
	// 	wg.Done()
	// }(fillOrdersCh)

	wg.Wait()
}

func fillOrders(in <-chan order, wg *sync.WaitGroup) {
	const workers = 3
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for o := range in {
				o.Status = filled
				fmt.Printf("Order has been completed: %v\n", o)
			}
			wg.Done()
		}()
	}
}

func reserveInventory(in <-chan order) <-chan order {
	out := make(chan order)
	var wg sync.WaitGroup

	const workers = 3
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for o := range in {
				o.Status = reserved
				out <- o
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder) {
	out := make(chan order)
	errCh := make(chan invalidOrder, 1)

	go func() {
		// order := <-in
		for order := range in {
			if order.Quantity <= 0 {
				errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
			} else {
				out <- order
			}
		}
		close(out)
		close(errCh)
	}()
	return out, errCh
}

func receiveOrders() <-chan order {
	out := make(chan order)
	go func() {
		for _, rawOrder := range rawOrders {
			var newOrder order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Print(err)
				continue
			}
			out <- newOrder
		}
		close(out)
	}()
	return out
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status":1}`,
	`{"productCode": 2222, "quantity": 42.3, "status":1}`,
	`{"productCode": 3333, "quantity": 19, "status":1}`,
	`{"productCode": 4444, "quantity": 8, "status":1}`,
}
