package qigua

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

func GetYaoNumber() int64 {
	//rand.Seed(time.Now().UnixNano())
	//大衍之数五十，其用四十有九
	var dayan = int64(50)
	dayan = dayan - 1

	//分而为二以像两
	//right := rand.Intn(dayan-6) + 6
	n, _ := crand.Int(crand.Reader, big.NewInt(dayan-12))
	right := n.Int64() + int64(6)
	left := dayan - right

	//挂一以像三
	//sancai := 1
	right = right - int64(1)

	//揲之以四以象四时,归奇于扐以象闰
	run := left % 4
	if run == 0 {
		run = 4
	}

	left = left - run

	//五岁再闰，故再扐而后挂
	zairun := right % 4
	if zairun == 0 {
		zairun = 4
	}

	right = right - zairun

	//二变

	erbian := right + left
	erright, _ := crand.Int(crand.Reader, big.NewInt(erbian-12))
	right = erright.Int64() + int64(6)
	left = erbian - right
	right = right - 1
	run = left % 4
	if run == 0 {
		run = 4
	}

	left = left - run

	zairun = right % 4
	if zairun == 0 {
		zairun = 4
	}

	right = right - zairun

	//三变
	sanbian := left + right
	//right = rand.Intn(sanbian-6) + 6
	sanright, _ := crand.Int(crand.Reader, big.NewInt(erbian-12))
	right = sanright.Int64() + int64(6)
	left = sanbian - right
	right = right - 1
	run = left % 4
	if run == 0 {
		run = 4
	}

	left = left - run

	zairun = right % 4
	if zairun == 0 {
		zairun = 4
	}

	right = right - zairun

	heshu := right + left
	yaonum := heshu / 4

	return yaonum

}

func GetYaoNumberWithDataNow() int {
	rand.Seed(time.Now().UnixNano())
	//大衍之数五十，其用四十有九
	var dayan = 50
	dayan = dayan - 1

	//分而为二以像两
	right := rand.Intn(dayan-12) + 6
	//n, _ := rand.Int(rand.Reader, big.NewInt(dayan-12))
	//right := n.Int64() + int64(6)
	left := dayan - right

	//挂一以像三
	//sancai := 1
	right = right - 1

	//揲之以四以象四时,归奇于扐以象闰
	run := left % 4
	if run == 0 {
		run = 4
	}

	left = left - run

	//五岁再闰，故再扐而后挂
	zairun := right % 4
	if zairun == 0 {
		zairun = 4
	}

	right = right - zairun

	//二变

	erbian := right + left

	//erright, _ := rand.Int(rand.Reader, big.NewInt(erbian-12))
	//right = erright.Int64() + int64(6)
	left = erbian - right

	run = left % 4
	if run == 0 {
		run = 4
	}

	left = left - run

	zairun = right % 4
	if zairun == 0 {
		zairun = 4
	}

	right = right - zairun

	//三变
	sanbian := left + right
	right = rand.Intn(sanbian-12) + 6
	//sanright, _ := rand.Int(rand.Reader, big.NewInt(erbian-12))
	//right = sanright.Int64() + int64(6)
	left = sanbian - right

	run = left % 4
	if run == 0 {
		run = 4
	}

	left = left - run

	zairun = right % 4
	if zairun == 0 {
		zairun = 4
	}

	right = right - zairun

	heshu := right + left
	yaonum := heshu / 4

	return yaonum

}

func GetGuaNumbers() []int64 {
	var guaNumbers = make([]int64, 6)

	for i := 0; i < 6; i++ {
		guaNumbers[5-i] = GetYaoNumber()
		//guaNumbers[5-i] = int64(GetYaoNumberWithDataNow())
	}
	return guaNumbers
}

func GetGuaAndZhigua() (string, string, string) {
	gua := make([]int, 6)
	zhigua := make([]int, 6)
	guanum := GetGuaNumbers()

	for k, v := range guanum {
		if v == 6 {
			gua[k] = 0
			zhigua[k] = 1
		}

		if v == 9 {
			gua[k] = 1
			zhigua[k] = 0
		}

		if v == 7 {
			gua[k] = 1
			zhigua[k] = 1
		}

		if v == 8 {
			gua[k] = 0
			zhigua[k] = 0
		}
	}

	yigua := ""
	yizhigua := ""
	guanumstr := ""

	for _, v := range gua {
		yigua = yigua + strconv.Itoa(v)
	}

	for _, v := range zhigua {
		yizhigua = yizhigua + strconv.Itoa(v)
	}

	for _, v := range guanum {
		if v == 6 {
			guanumstr = guanumstr + "六"
		}

		if v == 9 {
			guanumstr = guanumstr + "九"
		}

		if v == 7 {
			guanumstr = guanumstr + "七"
		}

		if v == 8 {
			guanumstr = guanumstr + "八"
		}
	}
	return yigua, yizhigua, guanumstr
}
