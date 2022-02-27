package main

import (
	"cli_converter/functions"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	getCmd := flag.NewFlagSet("convert", flag.ExitOnError)

	decimalToBinary := getCmd.String("decimalToBinary", "", "converts decimal number to binary number")
	binaryToDecimal := getCmd.String("binaryToDecimal", "", "converts binary number to decimal number")
	decimalToOctal := getCmd.String("decimalToOctal", "", "converts decimal number to octal number")
	octalToDecimal := getCmd.String("octalToDecimal", "", "converts octal number to decimal number")
	decimalToHex := getCmd.String("decimalToHex", "", "converts decimal number to hex number")
	hexToDecimal := getCmd.String("hexToDecimal", "", "converts hex number to decimal number")

	getOps := flag.NewFlagSet("addition", flag.ExitOnError)

	binaryAdd := getOps.String("binaryAdd", "", "adds binaries")
	octalAdd := getOps.String("octalAdd", "", "adds octals")
	hexAdd := getOps.String("hexAdd", "", "adds hexs")

	if len(os.Args) < 2 {
		fmt.Println("Expected the name of the function and it's value")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "convert":
		convert(getCmd, decimalToBinary, binaryToDecimal,
			decimalToOctal, octalToDecimal, decimalToHex, hexToDecimal)
	case "addition":
		addition(getOps, binaryAdd, octalAdd, hexAdd)
	default:
		fmt.Println("Smth went wrong")
	}
}

func convert(getCmd *flag.FlagSet, decimalToBinary *string, binaryToDecimal *string,
	decimalToOctal *string, octalToDecimal *string, decimalToHex *string, hexToDecimal *string) {

	getCmd.Parse(os.Args[2:])

	if *decimalToBinary == "" && *binaryToDecimal == "" && *decimalToOctal == "" &&
		*octalToDecimal == "" && *decimalToHex == "" && *hexToDecimal == "" {
		fmt.Println("provide more arguments")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *decimalToBinary != "" {
		res, _ := functions.DecimalToBinary(*decimalToBinary)
		fmt.Println(res)
		return
	} else if *binaryToDecimal != "" {
		res, _ := functions.BinaryToDecimal(*binaryToDecimal)
		fmt.Println(res)
		return
	} else if *decimalToOctal != "" {
		res, _ := functions.DecimalToOctal(*decimalToOctal)
		fmt.Println(res)
		return
	} else if *octalToDecimal != "" {
		res, _ := functions.OctalToDecimal(*octalToDecimal)
		fmt.Println(res)
		return
	} else if *decimalToHex != "" {
		res, _ := functions.DecimalToHex(*decimalToHex)
		fmt.Println(res)
		return
	} else if *hexToDecimal != "" {
		res, _ := functions.HexToDecimal(*hexToDecimal)
		fmt.Println(res)
		return
	}

}

func addition(getOps *flag.FlagSet, binaryAdd *string, octalAdd *string, hexAdd *string) {
	getOps.Parse(os.Args[2:])

	if *binaryAdd == "" && *octalAdd == "" && *hexAdd == "" {
		fmt.Println("Provide more arguments")
		getOps.PrintDefaults()
		os.Exit(1)
	}

	if *binaryAdd != "" {
		slice := strings.Split(strings.TrimSpace(*binaryAdd), ",")
		res, _ := functions.AdditionOverBinary(slice[0], slice[1])
		fmt.Println(res)
		return
	} else if *octalAdd != "" {
		slice := strings.Split(strings.TrimSpace(*octalAdd), ",")
		res, _ := functions.AddtionOverOctal(slice[0], slice[1])
		fmt.Println(res)
		return
	} else if *hexAdd != "" {
		slice := strings.Split(strings.TrimSpace(*hexAdd), ",")
		res, _ := functions.AdditionOverHex(slice[0], slice[1])
		fmt.Println(res)
		return
	}
}
