package main

import (
	"fmt" //to print messages to stdout
	"log" //logging :)
	//our web server that will host the mock
	"github.com/buaazp/fasthttprouter" 
	"github.com/valyala/fasthttp"
	"os"
	"io/ioutil"
)

var configuration []byte

func ReadConfig(){
	fmt.Println("Reading config...")
	config, e := ioutil.ReadFile("/configs/config.json")
	if e != nil {
		fmt.Printf("Error reading cofig file %v\n", e)
		os.Exit(1)
	}
	configuration = config
	fmt.Println("Config loaded!")

}

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello") 
}
func main() {
    
	fmt.Println("Starting.")
	//ReadConfig()
	router := fasthttprouter.New()
	router.GET("/", Response)
	
	log.Fatal(fasthttp.ListenAndServe(":5000", router.Handler))
}