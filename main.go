package main

import (
	"context"
	//"encoding/json"
	"io/ioutil"

	//"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	//flag.DurationVar(&wait, "graceful-timeout", 60 * time.Second, "the duration for which the server gracefully waits")
	//flag.Parse()

	fmt.Println("Starting the web endpoint...")

	router := mux.NewRouter()
	router.HandleFunc("/",homePage)

	server:= &http.Server{
		Handler: router,
		Addr:"0.0.0.0:7779",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)
	fmt.Println("Shutting down the web endpoint...")
	os.Exit(0)
}

type message struct{
	Message string
}

func callOneAndTwo() string{

	url := "http://one-myproject.192.168.64.2.nip.io"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)



	url2 := "http://two-myproject.192.168.64.2.nip.io"
	req2, _ := http.NewRequest("GET", url2, nil)
	res2, _ := http.DefaultClient.Do(req2)
	defer res.Body.Close()
	body2, _ := ioutil.ReadAll(res2.Body)

	return string(body)+string(body2)

}

func createFile() {
    // check if file exists
    var _, err = os.Stat(path)

    // create file if not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) {
            return
        }
        defer file.Close()
    }

    fmt.Println("File Created Successfully", path)
}
func homePage(writer http.ResponseWriter, request *http.Request) {
	//response := message{Message:"From Three"}
	//data,err := json.Marshal(response)
	//
	//if err !=nil{
	//	panic("Erorr at One JSON MARSHALL")
	//}

	//fmt.Fprint(writer,string(data))
	createFile()
	fmt.Fprint(writer, callOneAndTwo())
}
