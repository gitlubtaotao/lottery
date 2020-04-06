package main

import (
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"sync"
	"testing"
)

func TestLotteryController_Get(t *testing.T) {
	e := httptest.New(t, newApp())
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 0\n")
}
func TestLotteryController_PostImport(t *testing.T) {
	e := httptest.New(t, newApp())
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e.POST("/import").WithFormField("users", fmt.Sprintf("test_u%d", i)).Expect().Status(httptest.StatusOK)
		}(i)
	}
	wg.Wait()
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 100\n")
	e.GET("/lucky").Expect().Status(httptest.StatusOK)
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 99\n")
}

func TestLotteryController_GetLucky(t *testing.T) {
	e := httptest.New(t, newApp())
	
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 99\n")
}
