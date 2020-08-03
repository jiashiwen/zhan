package commons

import (
	"fmt"
	"golang.org/x/text/width"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"unicode"
)

func TestLine2Column(t *testing.T) {

	content, err := os.Open("../resources/KunWeiDi.gua")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer content.Close()

	contentbytes, err := ioutil.ReadAll(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentstr := string(contentbytes)

	col := Line2Column(contentstr, 36)

	for _, v := range col {
		for _, word := range v {
			fmt.Print(string(word))
		}
		fmt.Println("")
	}

	fmt.Println("反")
	for _, v := range col {
		for i := len(v) - 1; i >= 0; i-- {
			fmt.Print(string(v[i]))
		}
		fmt.Println("")
	}

	s := `。，（）-1！@234567890abc１２３４５６７８９ａｂｃ`
	numConv := unicode.SpecialCase{
		unicode.CaseRange{
			Lo: 0x3002, // Lo 全角句号
			Hi: 0x3002, // Hi 全角句号
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x002e - 0x3002, // LowerCase 转成半角句号
				0,               // TitleCase
			},
		},
		//
		unicode.CaseRange{
			Lo: 0xFF01, // 从全角 1
			Hi: 0xFF19, // 到全角 9
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x0021 - 0xFF01, // LowerCase 转成半角
				0,               // TitleCase
			},
		},
	}

	fmt.Println(strings.ToLowerSpecial(numConv, s))

	fmt.Println(width.Widen.String(s))

}
