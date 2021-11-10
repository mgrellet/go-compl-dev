package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make (chan string)

	for _,link := range links {
		go checkLink(link, c)
	}

	for l := range c { //wait for the channel		
		//go checkLink(l, c)
		go func (lnk string){
			time.Sleep(5 * time.Second)
			checkLink(lnk, c)
		}(l)
	}
	
	//fmt.Println(<-c)

}

func checkLink(link string, c chan string){

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}