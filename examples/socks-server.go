package main

import (
	"time"

	"github.com/FlowerWrong/gosocks"
)

func main() {
	srv := gosocks.NewBasicServer(":10800", time.Minute)
	srv.ListenAndServe()
}
