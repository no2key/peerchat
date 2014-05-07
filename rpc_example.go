package main

import "example"
import "log"
import "strconv"
import "lang"

func main() {	
	
	//sync()
	//async()
	stack_ex()
}

func sync() {
	one := example.MakeNode("127.0.0.1", "55555")
	log.Printf("making one...")
	two := example.MakeNode("127.0.0.1", "55554")
	log.Printf("making two...")
	
	one.Ping(two.Address, "Hello, this is one!")
	two.Ping(one.Address, "This is two...")
	one.Ping(two.Address, "Cool")
	two.Ping(one.Address, "Last message.")
}

func async() {
	
	one := example.MakeNode("127.0.0.1", "55555")
	log.Printf("making one...")
	two := example.MakeNode("127.0.0.1", "55554")
	log.Printf("making two...")
	
	alpha := 25
	
	doneChannel := make(chan *example.PingReply, alpha)
	for i := 0; i < alpha; i++ {
		go one.AsyncPing(two.Address, "Async message " + strconv.Itoa(i), doneChannel)
	}
	
	for i := 0; i < alpha; i++ {
		reply := <-doneChannel
		log.Printf("Message recieved! %s", reply.OK)
	} 
}

func stack_ex() {
	
	// create a stack
	stack := lang.NewStack()
	
	// add elements
	stack.Push("the first")
	stack.Push("the second")
	stack.Push("and the last")
	
	// pop them off
	log.Printf(stack.Pop().(string))
	log.Printf(stack.Pop().(string))
	log.Printf(stack.Pop().(string))
}
