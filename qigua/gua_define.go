package qigua

import (
	"errors"
	"strconv"
	"strings"
)

type YinYang string
type XianTianBaGua string
type HouTianBagua string
type YiGua string

const (
	YinYang_TaiJi    YinYang = "00000000"
	YinYang_Yang     YinYang = "1"
	YinYang_Yin      YinYang = "0"
	YinYang_TaiYang  YinYang = "11"
	YinYang_TaiYin   YinYang = "00"
	YinYang_ShaoYang YinYang = "10"
	YinYang_ShaoYin  YinYang = "01"
)

const (
	XianTianBaGua_Qian XianTianBaGua = "111"
	XianTianBaGua_Dui  XianTianBaGua = "011"
	XianTianBaGua_Li   XianTianBaGua = "101"
	XianTianBaGua_Zhen XianTianBaGua = "001"
	XianTianBaGua_Xun  XianTianBaGua = "110"
	XianTianBaGua_Kan  XianTianBaGua = "010"
	XianTianBaGua_Gen  XianTianBaGua = "100"
	XianTianBaGua_Kun  XianTianBaGua = "000"
)

const (
	HouTianBaGua_Kan  HouTianBagua = "010"
	HouTianBaGua_Kun  HouTianBagua = "000"
	HouTianBaGua_Zhen HouTianBagua = "001"
	HouTianBaGua_Xun  HouTianBagua = "110"
	HouTianBaGua_Qian HouTianBagua = "111"
	HouTianBaGua_Dui  HouTianBagua = "011"
	HouTianBaGua_Gen  HouTianBagua = "100"
	HouTianBaGua_Li   HouTianBagua = "101"
)

const (
	YiGua_KunWeiDi        = "000000"
	YiGua_ShanDiBo        = "000001"
	YiGua_ShuiDiBi        = "000010"
	YiGua_FengDiGuan      = "000011"
	YiGua_LeiDiYu         = "000100"
	YiGua_HuoDiJin        = "000101"
	YiGua_ZeDiCui         = "000110"
	YiGua_TianDiPi        = "000111"
	YiGua_DiShanQian      = "001000"
	YiGua_GenWeiShan      = "001001"
	YiGua_ShuiShanJian    = "001010"
	YiGua_FengShanJian    = "001011"
	YiGua_LeiShanXiaoGuo  = "001100"
	YiGua_HuoShanLv       = "001101"
	YiGua_ZeShanXian      = "001110"
	YiGua_TianDiDun       = "001111"
	YiGua_DiShuiShi       = "010000"
	YiGua_ShanShuiMeng    = "010001"
	YiGua_KanWeiShui      = "010010"
	YiGua_FengShuiHuan    = "010011"
	YiGua_LeiShuiXie      = "010100"
	YiGua_HuoShuiWeiJi    = "010101"
	YiGua_ZeShuiKun       = "010110"
	YiGua_TianShuiSong    = "010111"
	YiGua_DiFengSheng     = "011000"
	YiGua_ShanShuiGu      = "011001"
	YiGua_ShuiFengJing    = "011010"
	YiGua_XunWeiFeng      = "011011"
	YiGua_LeiFengHeng     = "011100"
	YiGua_HuoFengDing     = "011101"
	YiGua_ZeFengDaGuo     = "011110"
	YiGua_TianFengGou     = "011111"
	YiGua_DiLeiFu         = "100000"
	YiGua_ShanLeiYi       = "100001"
	YiGua_ShuiLeiZhun     = "100010"
	YiGua_FengLeiYi       = "100011"
	YiGua_ZhenWeiLei      = "100100"
	YiGua_HuoLeiShiHe     = "100101"
	YiGua_ZeLeiSui        = "100110"
	YiGua_TianLeiWuWang   = "100111"
	YiGua_DiHuoMingYi     = "101000"
	YiGua_ShanHuoBen      = "101001"
	YiGua_ShuiHuoJiJi     = "101010"
	YiGua_FengHuoJiaRen   = "101011"
	YiGua_LeiHuoFeng      = "101100"
	YiGua_LiWeiHuo        = "101101"
	YiGua_ZeHuoGe         = "101110"
	YiGua_TianHuoTongRen  = "101111"
	YiGua_DiZeLin         = "110000"
	YiGua_ShanZeSun       = "110001"
	YiGua_ShuiZeJie       = "110010"
	YiGua_FengZeZhongFu   = "110011"
	YiGua_LeiZeGuiMei     = "110100"
	YiGua_HuoZeKui        = "110101"
	YiGua_DuiWeiZe        = "110110"
	YiGua_TianZeLv        = "110111"
	YiGua_DiTianTai       = "111000"
	YiGua_ShanTianDaXu    = "111001"
	YiGua_ShuiTianXu      = "111010"
	YiGua_FengTianXiaoXu  = "111011"
	YiGua_LeiTianDaZhuang = "111100"
	YiGua_HuoTianDaYou    = "111101"
	YiGua_ZeTianGuai      = "111110"
	YiGua_QianWeiTian     = "111111"
)

func (yinyang YinYang) String() string {
	switch yinyang {
	case YinYang_TaiJi:
		return "☯ 太极"
	case YinYang_Yang:
		return "⚊ 阳"
	case YinYang_Yin:
		return "⚋ 阴"
	case YinYang_TaiYang:
		return "⚌ 太阳"
	case YinYang_TaiYin:
		return "⚏ 太阴"
	case YinYang_ShaoYang:
		return "⚎ 少阳"
	case YinYang_ShaoYin:
		return "⚍ 少阴"
	default:
		return ""
	}
}

func (xiantianbagua XianTianBaGua) String() string {
	switch xiantianbagua {
	case XianTianBaGua_Qian:
		return "☰ 乾"
	case XianTianBaGua_Dui:
		return "☱ 兑"
	case XianTianBaGua_Li:
		return "☲ 离"
	case XianTianBaGua_Zhen:
		return "☳ 震"
	case XianTianBaGua_Xun:
		return "☴ 巽"
	case XianTianBaGua_Kan:
		return "☵ 坎"
	case XianTianBaGua_Gen:
		return "☶ 艮"
	case XianTianBaGua_Kun:
		return "☷ 坤"
	default:
		return ""
	}
}

func (xiantianbagua XianTianBaGua) Number() int {
	switch xiantianbagua {
	case XianTianBaGua_Qian:
		return 1
	case XianTianBaGua_Dui:
		return 2
	case XianTianBaGua_Li:
		return 3
	case XianTianBaGua_Zhen:
		return 4
	case XianTianBaGua_Xun:
		return 5
	case XianTianBaGua_Kan:
		return 6
	case XianTianBaGua_Gen:
		return 7
	case XianTianBaGua_Kun:
		return 8
	default:
		return 0
	}
}

func (houtianbagua HouTianBagua) String() string {
	switch houtianbagua {
	case HouTianBaGua_Qian:
		return "☰ 乾"
	case HouTianBaGua_Dui:
		return "☱ 兑"
	case HouTianBaGua_Li:
		return "☲ 离"
	case HouTianBaGua_Zhen:
		return "☳ 震"
	case HouTianBaGua_Xun:
		return "☴ 巽"
	case HouTianBaGua_Kan:
		return "☵ 坎"
	case HouTianBaGua_Gen:
		return "☶ 艮"
	case HouTianBaGua_Kun:
		return "☷ 坤"
	default:
		return ""
	}
}

func (houtianbagua HouTianBagua) Number() int {
	switch houtianbagua {
	case HouTianBaGua_Kan:
		return 1
	case HouTianBaGua_Kun:
		return 2
	case HouTianBaGua_Zhen:
		return 3
	case HouTianBaGua_Xun:
		return 4
	case HouTianBaGua_Qian:
		return 6
	case HouTianBaGua_Dui:
		return 7
	case HouTianBaGua_Gen:
		return 8
	case HouTianBaGua_Li:
		return 9
	default:
		return 0
	}
}

func (yigua YiGua) String() string {
	switch yigua {
	case YiGua_KunWeiDi:
		return "䷁ 坤 0 坤為地 KunWeiDi"
	case YiGua_ShanDiBo:
		return "䷖ 剥 1 山地剥 ShanDiBo"
	case YiGua_ShuiDiBi:
		return "䷇ 比 2 水地比 ShuiDiBi"
	case YiGua_FengDiGuan:
		return "䷓ 观 3 風地观 FengDiGuan"
	case YiGua_LeiDiYu:
		return "䷏ 豫 4 雷地豫 LeiDiYu"
	case YiGua_HuoDiJin:
		return "䷢ 晋 5 火地晋 HuoDiJin"
	case YiGua_ZeDiCui:
		return "䷬ 萃 6 泽地萃 ZeDiCui"
	case YiGua_TianDiPi:
		return "䷋ 否 7 天地否 TianDiPi"
	case YiGua_DiShanQian:
		return "䷎ 謙 8 地山謙 DiShanQian"
	case YiGua_GenWeiShan:
		return "䷳ 艮 9 艮為山 GenWeiShan"
	case YiGua_ShuiShanJian:
		return "䷦ 蹇 10 水山蹇 ShuiShanJian"
	case YiGua_FengShanJian:
		return "䷴ 漸 11 風山漸 FengShanJian"
	case YiGua_LeiShanXiaoGuo:
		return "䷽ 小過 12 雷山小過 LeiShanXiaoGuo"
	case YiGua_HuoShanLv:
		return "䷷ 旅 13 火山旅 HuoShanLv"
	case YiGua_ZeShanXian:
		return "䷞ 咸 14 泽山咸 ZeShanXian"
	case YiGua_TianDiDun:
		return "䷠ 遯 15 天山遯 TianDiDun"
	case YiGua_DiShuiShi:
		return "䷆ 師 16 地水師 DiShuiShi"
	case YiGua_ShanShuiMeng:
		return "䷃ 蒙 17 山水蒙 ShanShuiMeng"
	case YiGua_KanWeiShui:
		return "䷜ 坎 18 坎為水 KanWeiShui"
	case YiGua_FengShuiHuan:
		return "䷺ 渙 19 風水渙 FengShuiHuan"
	case YiGua_LeiShuiXie:
		return "䷧ 解 20 雷水解 LeiShuiXie"
	case YiGua_HuoShuiWeiJi:
		return "䷿ 未济 21 火水未济 HuoShuiWeiJi"
	case YiGua_ZeShuiKun:
		return "䷮ 困 22 泽水困 ZeShuiKun"
	case YiGua_TianShuiSong:
		return "䷅ 訟 23 天水訟 TianShuiSong"
	case YiGua_DiFengSheng:
		return "䷭ 升 24 地風升 DiFengSheng"
	case YiGua_ShanShuiGu:
		return "䷑ 蠱 25 山風蠱 ShanShuiGu"
	case YiGua_ShuiFengJing:
		return "䷯ 井 26 水風井 ShuiFengJing"
	case YiGua_XunWeiFeng:
		return "䷸ 巽 27 巽為風 XunWeiFeng"
	case YiGua_LeiFengHeng:
		return "䷟ 恒 28 雷風恒 LeiFengHeng"
	case YiGua_HuoFengDing:
		return "䷱ 鼎 29 火風鼎 HuoFengDing"
	case YiGua_ZeFengDaGuo:
		return "䷛ 大過 30 泽風大過 ZeFengDaGuo"
	case YiGua_TianFengGou:
		return "䷫ 姤 31 天風姤 TianFengGou"
	case YiGua_DiLeiFu:
		return "䷗ 復 32 地雷復 DiLeiFu"
	case YiGua_ShanLeiYi:
		return "䷚ 頤 33 山雷頤 ShanLeiYi"
	case YiGua_ShuiLeiZhun:
		return "䷂ 屯 34 水雷屯 ShuiLeiZhun"
	case YiGua_FengLeiYi:
		return "䷩ 益 35 風雷益 FengLeiYi"
	case YiGua_ZhenWeiLei:
		return "䷲ 震 36 震為雷 ZhenWeiLei"
	case YiGua_HuoLeiShiHe:
		return "䷔ 噬嗑 37 火雷噬嗑 HuoLeiShiHe"
	case YiGua_ZeLeiSui:
		return "䷐ 随 38 泽雷随 ZeLeiSui"
	case YiGua_TianLeiWuWang:
		return "䷘ 无妄 39 天雷无妄 TianLeiWuWang"
	case YiGua_DiHuoMingYi:
		return "䷣ 明夷 40 地火明夷 DiHuoMingYi"
	case YiGua_ShanHuoBen:
		return "䷕ 賁 41 山火賁 ShanHuoBen"
	case YiGua_ShuiHuoJiJi:
		return "䷾ 既济 42 水火既济 ShuiHuoJiJi"
	case YiGua_FengHuoJiaRen:
		return "䷤ 家人 43 風火家人 FengHuoJiaRen"
	case YiGua_LeiHuoFeng:
		return "䷶ 豊 44 雷火豊 LeiHuoFeng"
	case YiGua_LiWeiHuo:
		return "䷝ 火 45 離為火 LiWeiHuo"
	case YiGua_ZeHuoGe:
		return "䷰ 革 46 泽火革 ZeHuoGe"
	case YiGua_TianHuoTongRen:
		return "䷌ 同人 47 天火同人 TianHuoTongRen"
	case YiGua_DiZeLin:
		return "䷒ 臨 48 地泽臨 DiZeLin"
	case YiGua_ShanZeSun:
		return "䷨ 損 49 山泽損 ShanZeSun"
	case YiGua_ShuiZeJie:
		return "䷻ 節 50 水泽節 ShuiZeJie"
	case YiGua_FengZeZhongFu:
		return "䷼ 中孚 51 風泽中孚 FengZeZhongFu"
	case YiGua_LeiZeGuiMei:
		return "䷵ 归妹 52 雷泽归妹 LeiZeGuiMei"
	case YiGua_HuoZeKui:
		return "䷥ 睽 53 火泽睽 HuoZeKui"
	case YiGua_DuiWeiZe:
		return "䷹ 泽 54 兌為泽 DuiWeiZe"
	case YiGua_TianZeLv:
		return "䷉ 履 55 天泽履 TianZeLv"
	case YiGua_DiTianTai:
		return "䷊ 泰 56 地天泰 DiTianTai"
	case YiGua_ShanTianDaXu:
		return "䷙ 大畜 57 山天大畜 ShanTianDaXu"
	case YiGua_ShuiTianXu:
		return "䷄ 需 58 水天需 ShuiTianXu"
	case YiGua_FengTianXiaoXu:
		return "䷈ 小畜 59 風天小畜 FengTianXiaoXu"
	case YiGua_LeiTianDaZhuang:
		return "䷡ 大壮 60 雷天大壮 LeiTianDaZhuang"
	case YiGua_HuoTianDaYou:
		return "䷍ 大有 61 火天大有 HuoTianDaYo"
	case YiGua_ZeTianGuai:
		return "䷪ 夬 62 泽天夬 ZeTianGuai"
	case YiGua_QianWeiTian:
		return "䷀ 天 63 乾為天 QianWeiTian"
	default:
		return ""
	}
}

func (yigua YiGua) Number() (int, error) {
	numstr := strings.Split(yigua.String(), " ")[2]
	if numstr == "" {
		return 0, errors.New("Gua not exists")
	}

	return strconv.Atoi(numstr)
}

func (yigua YiGua) Rune() string {
	return strings.Split(yigua.String(), " ")[0]
}

func (yigua YiGua) Name() string {
	return strings.Split(yigua.String(), " ")[1]
}

func (yigua YiGua) FullName() string {
	return strings.Split(yigua.String(), " ")[3]
}

func (yigua YiGua) FileName() string {
	return strings.Split(yigua.String(), " ")[4] + ".gua"
}
