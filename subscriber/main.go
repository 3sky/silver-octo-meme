package main

import (
    "encoding/json"
    "fmt"
    "os"
    "log"

    zmq "github.com/pebbe/zmq4"
)

type Message struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func main() {

    zctx, _ := zmq.NewContext()
    s, _ := zctx.NewSocket(zmq.SUB)
    
    port :=  os.Getenv("PUB_PORT")
    host :=  os.Getenv("PUB_HOST")

    if port == "" { port = "5553"} 
    if host == "" { port = "localhost"} 

    s.Connect("tcp://" + host + ":" + port)
    s.SetSubscribe("example")

    log.Println("Ready and tuned!")
    for {
        address, err := s.Recv(0)
        if err != nil {
            panic(err)
        }

        if msg, err := s.Recv(0); err != nil {
            panic(err)
        } else {
            contents := Message{}
            if err := json.Unmarshal([]byte(msg), &contents); err != nil {
                panic(err)
            }
            fmt.Println("Received message from " + address + " channel.")
            fmt.Printf("%+v\n", contents)
        }
    }
}
