package lotto

import (
	"fmt"
	"math/rand"
	"time"
)

// 6/45 기반 로또 번호 생성기.

var lotto_set [6]int = [6]int{0, 0, 0, 0, 0, 0}

func PrintLottoSet() bool {

	fmt.Println(lotto_set)
	return true
}

func LotsNumber(LimitValue int) (r int) {
	RandomSeed := rand.NewSource(time.Now().UnixNano())
	RandomObj := rand.New(RandomSeed)

	if LimitValue < 0 {
		LimitValue = LimitValue * -1
	}

	r = RandomObj.Intn(LimitValue + 100000) // 랜덤한 숫자 생성
	r = (r % 45) + 1

	return
}

func WriteLottoSet(index int, value int) {
	RandomSeed := rand.NewSource(time.Now().UnixNano())
	RandomObj := rand.New(RandomSeed)
	SleepCount := RandomObj.Intn(2) + 1
	fmt.Printf("number generation start. Wait %d second\n", SleepCount)
	time.Sleep(time.Duration(SleepCount) * time.Second)

	lotto_set[index] = LotsNumber(value)

	PrintLottoSet()
}
