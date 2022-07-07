package main

import (
	"github.com/bhmy-shm/orm-gdbs/examples"
	"sync"
)

var wg =sync.WaitGroup{}
//程序员在囧途(www.jtthink.com)出品 咨询群：98514334
func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		examples.TxStocktest(101,2,1)
	}()
	go func() {
		defer wg.Done()
		examples.TxStocktest(101,3,1)
	}()
	wg.Wait()
}
