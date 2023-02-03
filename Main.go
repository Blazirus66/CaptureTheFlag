package main

import (
// //"io/ioutil"
// "net/http"
// "strings"
"time"
)

func main() {
	ch := make(chan int)

	for i := 3000; i <= 4000; i++ {
		go findSecretKey(i, ch)
	}

	time.Sleep(time.Second)
	DownloadFile("File1", "ca32652906af8dd747e741cd3e960338138099b0615e62b4f23366cf65f52646", 3941, "secretKey")
	DownloadFile("File2", "8116fdd3f12b6d7c4b136cbdaa3360a57eb4eb676ae63294450ee1f4f34b36f3", 3610, "finalKey")
}