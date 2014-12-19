package main

import (
    "github.com/kr/beanstalk"
    "log"
    "time"
)

func StartBeanstalk() {
    c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
    if err != nil {
        panic(err)
    }

    for {
        id, body, err := c.Reserve(5 * time.Hour)
        if err != nil {
            log.Fatalf("Error: %s", err.Error())
        } else {
            c.Delete(id)
            log.Printf("%d %s", id, body)
        }    
    }    
}

