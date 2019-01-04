package main 
import "fmt"
import "net/http"
import "time"

func main() {
	
    links := []string{
    	"http://google.com",
    	"http://stackoverflow.com",
    	"http://facebook.com",
    	"http://golang.org",
    	"http://amazon.com",
    }

    c := make(chan string)

    for _, link := range links{
    	go getStatus(link, c)
    }

    for l := range c{
    	go func (link string) {
    	    time.Sleep(5*time.Second)	
    	    getStatus(link,c)

    	}(l)
       
    }


}

func getStatus(link string, c chan string) {
	
	 _, err := http.Get(link)
	 if err != nil{
	 	fmt.Println(link, " is down")
	 	c <- link
	 	return
	 }

	 fmt.Println(link, " is working properly")
	 c <- link
	 fmt.Println()
} 