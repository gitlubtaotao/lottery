package main

import (
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"sync"
	"testing"
)

func TestLotteryController_Get(t *testing.T) {
	e := httptest.New(t, newApp())
	var sync = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		sync.Add(1)
		go func(i int) {
			defer sync.Done()
			fmt.Println(e.GET("/").Expect().Status(httptest.StatusOK).Body(), i)
		}(i)
	}
	sync.Wait()
}
func TestLotteryController_GetPrize(t *testing.T) {
	e := httptest.New(t, newApp())
	var sync = sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		sync.Add(1)
		go func(i int) {
			defer sync.Done()
			fmt.Println(e.GET("/prize").Expect().Status(httptest.StatusOK).Body(), i)
		}(i)
	}
	sync.Wait()
}
