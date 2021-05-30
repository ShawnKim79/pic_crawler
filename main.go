package main

import (
	"fmt"
	"math/rand"
	"pic_crawler/modules/lotto"
	"time"
)

func main() {
	fmt.Println("Start Lotto!!")
	done := make(chan bool)

	go func() {
		for i := 0; i < 6; i++ {
			RandomSeed := rand.NewSource(time.Now().UnixNano())
			RandomObj := rand.New(RandomSeed)
			lotto.WriteLottoSet(i, RandomObj.Intn(45))
		}

		done <- true //채널에 true를 송신함
	}()

	<-done //수신함으로써 대기를 끝냄

	lotto.PrintLottoSet()

}
