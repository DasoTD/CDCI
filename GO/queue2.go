package main

import "fmt"

type Order2 struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

type Queue2 []*Order2

func (order *Order2) New(priority, quantity int, product, customerName string) {
	order.priority = priority
	order.quantity =quantity
	order.product = product
	order.customerName = customerName
}

func (queue * Queue2) Add(order *Order2){
	if len(*queue)==0{
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order2
		for i, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i],append(Queue2{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}

	}
}

func (queue *Queue2) DeQueue() Queue2{
	if len(*queue)==0 {
		return nil
	}
	removed := (*queue)[:1]

	*queue = (*queue)[1:]

	return removed

}

func mainy(){
	// var queue Queue
	queue := make(Queue2, 0)

	var order1 *Order2 = &Order2{}

	var priority1 int = 1
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"

	order1.New(priority1, quantity1, product1, customerName1)

	var order2 *Order2 = &Order2{}

	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "lomputer"
	var customerName2 string = "Lreg White"

	order2.New(priority2,quantity2,product2,customerName2)

	var order3 *Order2 = &Order2{}

	var priority3 int = 3
	var quantity3 int = 10
	var product3 string = "lomputer"
	var customerName3 string = "Lreg White"

	order3.New(priority3,quantity3,product3,customerName3)


	queue.Add(order1)
	queue.Add(order2)
	queue.Add(order3)

	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}

	queue.DeQueue()

	fmt.Println("dxhcgfct")

	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}

	fmt.Println(queue.DeQueue())

	fmt.Println("dddd")

	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}