package all

func getFixedStr(num int) string {
	RomanMap := map[int]string{
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}
	if str, ok := RomanMap[num]; ok {
		return str
	}
	return ""
}

func intToRoman(num int) string {
	str := getFixedStr(num)
	if str != "" {
		return str
	}

	res := num
	for res != 0 {
		if res >= 1000 {
			times := res / 1000
			res = res % 1000
			for i := 0 ; i < times; i++ {
				str += "M"
			}
		} else if res >= 900 {
			res = res - 900
			str += getFixedStr(900)

		} else if res >= 500 {
			times := res / 500
			res = res % 500
			for i := 0 ; i < times; i++ {
				str += "D"
			}
		} else if res >= 400 {
			res = res - 400
			str += getFixedStr(400)
		} else if res >= 100 {
			times := res / 100
			res = res % 100
			for i := 0 ; i < times; i++ {
				str += "C"
			}
		} else if res >= 90 {
			res = res - 90
			str += getFixedStr(90)
		} else if res >= 50 {
			times := res / 50
			res = res % 50
			for i := 0 ; i < times; i++ {
				str += "L"
			}
		} else if res >= 40 {
			res = res - 40
			str += getFixedStr(40)
		} else if res >= 10 {
			times := res / 10
			res = res % 10
			for i := 0 ; i < times; i++ {
				str += "X"
			}
		} else if res >= 9 {
			res = res - 9
			str += getFixedStr(9)
		} else if res >= 5 {
			times := res / 5
			res = res % 5
			for i := 0 ; i < times; i++ {
				str += "V"
			}
		} else if res >= 4 {
			res = res - 4
			str += getFixedStr(4)
		}else if res >= 1 {
			times := res / 1
			res = res % 1
			for i := 0 ; i < times; i++ {
				str += "I"
			}
		}

	}

	return str
}
