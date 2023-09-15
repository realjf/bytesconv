// #############################################################################
// # File: bytesconv.go                                                        #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:16:46                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/15 18:14:20                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023                                                        #
// #############################################################################
package bytesconv

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Official Units
	iotaCount          = 0
	B                  = big.NewInt(1)                         // 字节 Byte 1
	KB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 千字节 KiloByte  2^10
	MB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 兆字节 MegaByte  2^20
	GB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 吉字节 GigaByte  2^30
	TB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 太字节 TeraByte  2^40
	PB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 拍字节 PetaByte  2^50
	EB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 艾字节 ExaByte   2^60
	ZB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 泽字节 ZettaByte 2^70
	YB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 尧字节 YottaByte 2^80
	RB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // RonnaByte       2^90
	QB        *big.Int = big.NewInt(1 << (10 * iotaCounter())) // QuettaByte      2^100

	// Unofficial Units
	// BB *big.Int = RB                                    // 珀字节 BrontoByte  2^90
	// NB *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 诺字节 NonaByte
	// DB *big.Int = big.NewInt(1 << (10 * iotaCounter())) // 刀字节 DoggaByte
)

func iotaCounter() int {
	iotaCount++
	return iotaCount
}

const (
	kilobyte   = "k"
	megabyte   = "M"
	gigabyte   = "G"
	terabyte   = "T"
	petabyte   = "P"
	exabyte    = "E"
	zettabyte  = "Z"
	yottabyte  = "Y"
	ronnabyte  = "R"
	quettabyte = "Q"

	kibibyte = "K"
	mebibyte = "M"
	gibibyte = "G"
	tebibyte = "T"
	pebibyte = "P"
	exbibyte = "E"
	zebibyte = "Z"
	yobibyte = "Y"
	robibyte = "R"
	qubibyte = "Q"
)

const (
	Units = "BKMGTPEZYRQ"
)

const (
	SIUnitValue  = 1000
	IECUnitValue = 1024
)

// ========================================================= common public method =================================================================

// size是数值，uint是单位, default is Byte
func ToBytes(size int64, unit string) *big.Int {
	switch unit {
	case "Bit", "b":
		return big.NewInt(0).Div(big.NewInt(size), big.NewInt(8))
	case "K", "KB":
		return big.NewInt(0).Mul(big.NewInt(size), KB)
	case "M", "MB":
		return big.NewInt(0).Mul(big.NewInt(size), MB)
	case "G", "GB":
		return big.NewInt(0).Mul(big.NewInt(size), GB)
	case "T", "TB":
		return big.NewInt(0).Mul(big.NewInt(size), TB)
	case "P", "PB":
		return big.NewInt(0).Mul(big.NewInt(size), PB)
	case "E", "EB":
		return big.NewInt(0).Mul(big.NewInt(size), EB)
	case "Z", "ZB":
		return big.NewInt(0).Mul(big.NewInt(size), ZB)
	case "Y", "YB":
		return big.NewInt(0).Mul(big.NewInt(size), YB)
	case "R", "RB":
		return big.NewInt(0).Mul(big.NewInt(size), RB)
	case "Q", "QB":
		return big.NewInt(0).Mul(big.NewInt(size), QB)
	case "B", "Byte":
		fallthrough
	default:
		return big.NewInt(size)
	}
}

// 以1000作为基数 Decimal Unit
func ByteCountSI(bytesSize *big.Int, precision int) string {
	div, exp := big.NewInt(1), 0
	for n := big.NewInt(0).Set(bytesSize); n.Cmp(big.NewInt(SIUnitValue)) != -1; n = big.NewInt(0).Div(n, big.NewInt(SIUnitValue)) {
		div = big.NewInt(0).Mul(div, big.NewInt(SIUnitValue))
		exp++
	}

	return printSI(bytesSize, div, exp, precision)
}

// 以1024作为基数 Binary Unit
func ByteCountIEC(bytesSize *big.Int, precision int) string {
	div, exp := big.NewInt(1), 0
	for n := big.NewInt(0).Set(bytesSize); n.Cmp(big.NewInt(IECUnitValue)) != -1; n = big.NewInt(0).Div(n, big.NewInt(IECUnitValue)) {
		div = big.NewInt(0).Mul(div, big.NewInt(IECUnitValue))
		exp++
	}
	return printIEC(bytesSize, div, exp, precision)
}

// ========================================================= private method =================================================================

func toPos(unit string) int {
	switch unit {
	case "B", "Byte":
		return strings.Index(Units, "B")
	case "k", "kB", "K", "KiB", "KB":
		return strings.Index(Units, "K")
	case "M", "MiB", "MB":
		return strings.Index(Units, "M")
	case "G", "GiB", "GB":
		return strings.Index(Units, "G")
	case "T", "TiB", "TB":
		return strings.Index(Units, "T")
	case "P", "PiB", "PB":
		return strings.Index(Units, "P")
	case "E", "EiB", "EB":
		return strings.Index(Units, "E")
	case "Z", "ZiB", "ZB":
		return strings.Index(Units, "Z")
	case "Y", "YiB", "YB":
		return strings.Index(Units, "Y")
	case "R", "RiB", "RB":
		return strings.Index(Units, "R")
	case "Q", "QiB", "QB":
		return strings.Index(Units, "Q")
	default:
		return -1
	}
}

func toSIUnit(pos int) (u string) {
	if pos > 0 && pos < len(Units) {
		return string(" kMGTPEZYRQ"[pos]) + "B"
	} else if pos == 0 {
		return "B"
	}
	return
}

func toIECUnit(pos int) (u string) {
	if pos > 0 && pos < len(Units) {
		return string(" KMGTPEZYRQ"[pos]) + "iB"
	} else if pos == 0 {
		return "B"
	}
	return
}

func printIEC(bytesSize, div *big.Int, exp, precision int) string {
	divddend, _ := big.NewFloat(0).SetString(bytesSize.String())
	divisor, _ := big.NewFloat(0).SetString(div.String())
	answer := big.NewFloat(0).Quo(divddend, divisor)
	return fmt.Sprintf("%s %s", answer.Text('f', precision), toIECUnit(exp))
}

func printSI(bytesSize, div *big.Int, exp, precision int) string {
	divddend, _ := big.NewFloat(0).SetString(bytesSize.String())
	divisor, _ := big.NewFloat(0).SetString(div.String())
	answer := big.NewFloat(0).Quo(divddend, divisor)
	return fmt.Sprintf("%s %s", answer.Text('f', precision), toSIUnit(exp))
}

func toIEC(sizeAndUnit string, precision int, outputUnit string) string {
	return convertUnit(sizeAndUnit, precision, IECUnitValue, outputUnit, printIEC)
}

func toSI(sizeAndUnit string, precision int, outputUnit string) string {
	return convertUnit(sizeAndUnit, precision, SIUnitValue, outputUnit, printSI)
}

func convertUnit(sizeAndUnit string, precision int, unitValue int64, outputUnit string, printFunc func(bytesSize, div *big.Int, exp, precision int) string) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	div, exp := big.NewInt(1), 0
	loop := 0
	for n := bytesSize; loop < toPos(outputUnit); n = big.NewInt(0).Div(n, big.NewInt(unitValue)) {
		div = big.NewInt(0).Mul(div, big.NewInt(unitValue))
		exp++
		loop++
	}
	return printFunc(bytesSize, div, exp, precision)
}

// ========================================================= IEC =================================================================

// 单位之间转换
func ToKiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, kibibyte)
}

func ToMiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, mebibyte)
}

func ToGiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, gibibyte)
}

func ToTiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, tebibyte)
}

func ToPiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, pebibyte)
}

func ToEiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, exbibyte)
}

func ToZiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, zebibyte)
}

func ToYiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, yobibyte)
}

func ToRiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, robibyte)
}

func ToQiB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, qubibyte)
}

// ========================================================= SI =================================================================

func ToKB(sizeAndUnit string, precision int) string {
	return toSI(sizeAndUnit, precision, kilobyte)
}

func ToMB(sizeAndUnit string, precision int) string {
	return toSI(sizeAndUnit, precision, megabyte)
}

func ToGB(sizeAndUnit string, precision int) string {
	return toSI(sizeAndUnit, precision, gigabyte)
}

func ToTB(sizeAndUnit string, precision int) string {
	return toSI(sizeAndUnit, precision, terabyte)
}

func ToPB(sizeAndUnit string, precision int) string {
	return toSI(sizeAndUnit, precision, petabyte)
}

func ToEB(sizeAndUnit string, precision int) string {
	return toSI(sizeAndUnit, precision, exabyte)
}

func ToZB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, zettabyte)
}

func ToYB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, yottabyte)
}

func ToRB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, ronnabyte)
}

func ToQB(sizeAndUnit string, precision int) string {
	return toIEC(sizeAndUnit, precision, quettabyte)
}
