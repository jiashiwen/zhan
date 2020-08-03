package qigua

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetYaoNumber(t *testing.T) {
	GetYaoNumber()
}

func TestGetGuaNumbers(t *testing.T) {
	gua := GetGuaNumbers()
	fmt.Println(gua)
}

func TestGetGuaAndZhigua(t *testing.T) {
	gua, zhigua, guanum := GetGuaAndZhigua()
	unicodkun := "EBD0"
	fmt.Println(YiGua(gua).String(), YiGua(zhigua).String(), guanum)

	temp, _ := strconv.ParseInt(unicodkun, 16, 32)

	fmt.Print(fmt.Sprintf("%c", temp))
}
