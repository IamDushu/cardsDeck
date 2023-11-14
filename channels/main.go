package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//After creating go routines for all above links;
	//The Main Go Routine waits here for a message from channel
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) //Blocks the for loop from continuing on until receives a messsage from channel
	// }

	//Repeating Routines
	// for {
	// 	go checkLink(<-c, c) //Though it's infinite for loop; the loop gets blocked here as we're waiting for <-c
	// }

	//For other devs, the above code syntax might be confusing, so alt syntax
	//Here we are using range with channel -> wait for channel to return some value
	//After it has returned assign to value l; then run the body of for loop
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	//Sleeping Go Routines
	// for l := range c {
	// 	time.Sleep(5 * time.Second) //:( This will pause the Main Routine for 5 secs and the messages in channel get throttled from other go routines
	// 	go checkLink(l, c)
	// }

	// for l := range c {
	// 	go func() { //Function Literal or Anonymous Function
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(l, c) //Warning: loop variable l captured by func literal
	// 		//OUTPUT
	// 		// http://google.com is live!
	// 		// http://stackoverflow.com is live!
	// 		// http://facebook.com is live!
	// 		// http://golang.org is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		// http://amazon.com is live!
	// 		//Looks like l variable is receiving only amazon.com after 1st five requests
	//    // Concept of Closures
	// 	}()
	// }

	for l := range c {
		go func(link string) { //To avoid the effects we saw above
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is live!")
	c <- link //here link is also a string "http://"
}
