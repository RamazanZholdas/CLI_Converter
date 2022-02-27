package functions

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	dictionary = []string{"A", "B", "C", "D", "E", "F"}
)

// 17  = 17 / 2; 8 / 2; 4 / 2; 2 / 2, remainders together = 10001
func DecimalToBinary(str string) (string, error) {
	digit, err := strconv.Atoi(str)
	if err != nil {
		return "", err
	}

	var temp, res string

	for digit != 0 {
		temp += fmt.Sprint(digit % 2)
		digit = digit / 2
	}

	if len(temp) == 0 {
		res += "0"
	} else {
		leng := len(temp) - 1
		for i := leng; i >= 0; i-- {
			res += string(temp[i])
		}
	}

	return res, nil
}

//(d0 × 2^0 )+ (d1× 2^1 )+ (d2 × 2^2 )+ ..... + (dn−1 × 2^n-1)
func BinaryToDecimal(str string) (string, error) {
	for i := range str {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return "", err
		}
		if num > 1 && num < 0 {
			return "", errors.New("this is not a valid binary")
		}
	}
	res, i, j := 0, 0, len(str)-1
	for i <= len(str)-1 {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return "", err
		}
		res += num * int(math.Pow(2, float64(j)))
		i += 1
		j -= 1
	}
	return fmt.Sprint(res), nil
}

// 127  = 127 / 8; 15 / 8; 1 / 8, remainders together = 177
func DecimalToOctal(str string) (string, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return "", err
	}
	res, rem, quot := "", 0, 0
	for {
		rem = num % 8
		quot = num / 8
		if quot < 8 {
			res = fmt.Sprint(quot) + fmt.Sprint(rem) + res
			break
		}
		num = quot
		res = fmt.Sprint(rem) + res
	}
	return res, nil
}

//(d0 × 8^0 )+ (d1× 8^1 )+ (d2 × 8^2 )+ ..... + (dn−1 × 8^n-1)
func OctalToDecimal(str string) (string, error) {
	for i := range str {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return "", err
		}
		if num > 8 && num < 0 {
			return "", errors.New("this is not a valid octal")
		}
	}
	res, i, j := 0, 0, len(str)-1
	for i <= len(str)-1 {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return "", err
		}
		res += num * int(math.Pow(8, float64(j)))
		i += 1
		j -= 1
	}
	return fmt.Sprint(res), nil
}

//convert dec to bin, and then group it of 4 bins and convert it to hex using
//decimalToHexHelper(string) function
func DecimalToHex(str string) (string, error) {
	slice := []string{}
	bin, err := DecimalToBinary(str)
	if err != nil {
		return "", err
	}
	if len(bin) < 4 {
		newStr := strings.Repeat("0", 4-len(bin)) + bin
		dec, _ := BinaryToDecimal(newStr)
		slice = append(slice, dec)
		return DecimalToHexHelper(slice), nil
	}
	for len(bin)%4 != 0 {
		bin = "0" + bin
	}
	temp := 0
	for i := 0; i < len(bin); i++ {
		if (i+1)%4 == 0 {
			slice = append(slice, bin[temp:i+1])
			temp = i + 1
		}
	}
	return DecimalToHexHelper(slice), nil
}

func DecimalToHexHelper(slice []string) string {
	res := ""
	for i := range slice {
		str, _ := BinaryToDecimal(slice[i])
		num, _ := strconv.Atoi(str)
		if num >= 10 {
			res += dictionary[num-10]
			continue
		}
		res += str
	}
	return res
}

//(d0 × 16^0 )+ (d1× 16^1 )+ (d2 × 16^2 )+ ..... + (dn−1 × 16^n-1)
//converting letters to dec numbers using dictionary slice
func HexToDecimal(str string) (string, error) {
	pass, num, resNum, res, i, j := false, 0, 0, 0, 0, len(str)-1
	for i <= len(str)-1 {
		for k := range dictionary {
			if dictionary[k] == string(str[i]) {
				resNum = k + 10
				pass = true
				break
			}
		}
		if !pass {
			num, _ = strconv.Atoi(string(str[i]))
		} else {
			num = resNum
			pass = false
		}
		res += num * int(math.Pow(16, float64(j)))
		i += 1
		j -= 1
	}
	return fmt.Sprint(res), nil
}

func AdditionOverBinary(bin1, bin2 string) (string, error) {
	str := fmt.Sprint(bin1 + bin2)
	for i := range str {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return "", err
		}
		if num > 1 && num < 0 {
			return "", fmt.Errorf("this is not a valid binary")
		}
	}

	i, j, carry, res := len(bin1)-1, len(bin2)-1, 0, ""

	for i >= 0 || j >= 0 {
		num1, num2 := 0, 0
		var err error
		ans := carry
		if i >= 0 {
			num1, err = strconv.Atoi(string(bin1[i]))
			if err != nil {
				return "", err
			}
			ans += num1
			i--
		}

		if j >= 0 {
			num2, err = strconv.Atoi(string(bin1[i]))
			if err != nil {
				return "", err
			}
			ans += num2
			j--
		}

		if ans > 2 {
			res = "1" + res
			carry = 1
		} else if ans == 2 {
			res = "0" + res
			carry = 1
		} else {
			res = fmt.Sprint(ans) + res
			carry = 0
		}
	}

	if carry > 0 {
		res = fmt.Sprint(carry) + res
	}

	return res, nil
}

func AddtionOverOctal(octal1, octal2 string) (string, error) {
	str := fmt.Sprint(octal1 + octal2)
	for i := range str {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return "", err
		}
		if num > 8 && num < 0 {
			return "", errors.New("this is not a valid octal")
		}
	}
	i, j, carry, res := len(octal1)-1, len(octal2)-1, 0, ""
	for i >= 0 || j >= 0 {
		num1, num2 := 0, 0
		var err error
		ans := carry
		if i >= 0 {
			num1, err = strconv.Atoi(string(octal1[i]))
			if err != nil {
				return "", err
			}
			ans += num1
			i--
		}

		if j >= 0 {
			num2, err = strconv.Atoi(string(octal2[j]))
			if err != nil {
				return "", err
			}
			ans += num2
			j--
		}

		if ans > 7 {
			tempRes := strconv.Itoa(ans + 2)
			res = string(tempRes[1]) + res
			carry = 1
		} else {
			res = fmt.Sprint(ans) + res
			carry = 0
		}
	}

	if carry > 0 {
		res = fmt.Sprint(carry) + res
	}

	return res, nil
}

func AdditionOverHex(hex1, hex2 string) (string, error) {
	i, j, carry, res := len(hex1)-1, len(hex2)-1, 0, ""
	for i >= 0 || j >= 0 {
		num1, num2 := 0, 0
		var err error
		ans := carry
		if i >= 0 {
			if index, pass := contain(string(hex1[i])); pass {
				num1 = index + 10
			} else {
				num1, err = strconv.Atoi(string(hex1[i]))
				if err != nil {
					return "", err
				}
			}
			ans += num1
			i--
		}

		if j >= 0 {
			if index, pass := contain(string(hex2[j])); pass {
				num2 = index + 10
			} else {
				num2, err = strconv.Atoi(string(hex2[j]))
				if err != nil {
					return "", err
				}
			}
			ans += num2
			j--
		}

		if ans >= 10 && ans < 16 {
			res = dictionary[ans-10] + res
			carry = 0
		} else if ans > 15 {
			res = fmt.Sprint(ans-16) + res
			carry = 1
		} else {
			res = fmt.Sprint(ans) + res
			carry = 0
		}

	}

	if carry > 0 {
		res = fmt.Sprint(carry) + res
	}

	return res, nil
}

func SubstractionOverHex(hex1, hex2 string) string {
	i, j, res, borrow := len(hex1)-1, len(hex2)-1, "", 0
	for i >= 0 || j >= 0 {
		num1, num2 := 0, 0
		if i >= 0 {
			if index, pass := contain(string(hex1[i])); pass {
				num1 = index + 10
			} else {
				num1, _ = strconv.Atoi(string(hex1[i]))
			}
			num1 -= borrow
			i--
		}

		if j >= 0 {
			if index, pass := contain(string(hex2[j])); pass {
				num2 = index + 10
			} else {
				num2, _ = strconv.Atoi(string(hex2[j]))
			}
			j--
		}

		if ans := num1 - num2; ans >= 0 {
			if ans >= 10 {
				res = fmt.Sprint(dictionary[ans-10]) + res
			} else {
				res = fmt.Sprint(ans) + res
			}
			borrow = 0
		} else {
			ans := 16 + num1 - num2
			if ans >= 10 {
				res = fmt.Sprint(dictionary[ans-10]) + res
			} else {
				res = fmt.Sprint(ans) + res
			}
			borrow = 1
		}
	}
	return res
}

func SubstractionOverOctal(octal1, octal2 string) string {
	i, j, res, borrow := len(octal1)-1, len(octal2)-1, "", 0
	for i >= 0 || j >= 0 {
		num1, num2 := 0, 0
		if i >= 0 {
			num1, _ = strconv.Atoi(string(octal1[i]))
			num1 -= borrow
			i--
		}

		if j >= 0 {
			num2, _ = strconv.Atoi(string(octal2[j]))
			j--
		}

		if num1-num2 >= 0 {
			res = fmt.Sprint(num1-num2) + res
			borrow = 0
		} else {
			res = fmt.Sprint(8+num1-num2) + res
			borrow = 1
		}
	}
	return res
}

func SubstractionOverBinary(bin1, bin2 string) string {
	i, j, res, borrow := len(bin1)-1, len(bin2)-1, "", 0
	for i >= 0 || j >= 0 {
		num1, num2 := 0, 0
		if i >= 0 {
			num1, _ = strconv.Atoi(string(bin1[i]))
			num1 -= borrow
			i--
		}

		if j >= 0 {
			num2, _ = strconv.Atoi(string(bin2[j]))
			j--
		}

		if num1-num2 >= 0 {
			res = fmt.Sprint(num1-num2) + res
			borrow = 0
		} else {
			res = "1" + res
			borrow = 1
		}
	}
	return res
}

func contain(str string) (int, bool) {
	for i := range dictionary {
		if str == dictionary[i] {
			return i, true
		}
	}
	return 0, false
}
