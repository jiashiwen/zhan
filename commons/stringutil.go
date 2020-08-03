package commons

import (
	"golang.org/x/text/width"
	"strings"
)

const replacement = ""

var replacer = strings.NewReplacer(
	"\r\n", replacement,
	"\r", replacement,
	"\n", replacement,
	"\v", replacement,
	"\f", replacement,
	"\u0085", replacement,
	"\u2028", replacement,
	"\u2029", replacement,
)

func Replacer(s string) string {
	return replacer.Replace(s)
}

//行转列，竖版输出中文
//colsize 列的容量
func Line2Column(contentstr string, colsize int) [][]rune {

	//content, err := os.Open(file)
	//if err != nil {
	//	return nil, err
	//}
	//defer content.Close()
	//
	//contentbytes, err := ioutil.ReadAll(content)
	//if err != nil {
	//	return nil, err
	//}
	//contentstr := string(contentbytes)

	//分段
	paragraph := strings.Split(contentstr, "\n")
	document := string("")

	//段落处理，包括空格补齐，去掉空段落，段落合并
	for _, v := range paragraph {
		if len(v) == 0 {
			continue
		}

		//去掉换行符
		v = Replacer(v)

		//转全角
		v = width.Widen.String(v)

		//计算并补全，使用全角空格填充不满的列
		rs := []rune(v)
		quotient := len(rs) / colsize
		remainder := len(rs) % colsize

		if remainder != 0 {
			quotient = quotient + 1
			for i := 0; i < (colsize - remainder); i++ {
				//全角空格补位
				v = v + "　"
			}
		}

		//生成加工后文档
		document = document + v
	}

	docrune := []rune(document)
	quotient := len(docrune) / colsize

	article := make([][]rune, colsize)
	for i := 0; i < colsize; i++ {
		article[i] = make([]rune, quotient)
	}

	i := 0
	j := 0
	for _, v := range docrune {
		article[i][j] = v

		if i < (colsize - 1) {
			i++
		} else {
			i = 0
			j++

		}
	}
	return article

}
