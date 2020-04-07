package main

import (
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"sync"
	"testing"
)

func TestLotteryController_Get(t *testing.T) {
	e := httptest.New(t, newApp())
	InitLog()
	InitGift()
	var sy = sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		sy.Add(1)
		go func(i int) {
			defer sy.Done()
			fmt.Println(e.GET("/").Expect().Status(httptest.StatusOK).Body(), i)
		}(i)
	}
	sy.Wait()
}

/*
 *
 */
func TestLotteryController_GetLucky(t *testing.T) {
	e := httptest.New(t, newApp())
	InitLog()
	InitGift()
	var sy = sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		sy.Add(1)
		go func(i int) {
			defer sy.Done()
			fmt.Println(e.GET("/lucky").Expect().Status(httptest.StatusOK).Body(), i)
		}(i)
	}
	sy.Wait()
}
