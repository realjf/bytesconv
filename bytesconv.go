// #############################################################################
// # File: bytesconv.go                                                        #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:16:46                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/15 18:00:04                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023                                                        #
// #############################################################################
package bytesconv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

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
	Units = "BKMGTPE"
)

const (
	SIUnitValue  = 1000
	IECUnitValue = 1024
)

// ========================================================= common public method =================================================================

// size是数值，uint是单位, default is Byte
func ToBytes(size int64, unit string) int64 {
	switch unit {
	case "KB", "K":
		return size * KB
	case "MB", "M":
		return size * MB
	case "GB", "G":
		return size * GB
	case "TB", "T":
		return size * TB
	case "PB", "P":
		return size * PB
	case "EB", "E":
		return size * EB
	case "Bit":
		fallthrough
	case "b":
		return size / 8
	case "B", "Byte":
		fallthrough
	default:
		return size
	}
}

// 以1000作为基数 Decimal Unit
func ByteCountSI(bytesSize int64, precision int) string {
	div, exp := int64(1), 0
	for n := bytesSize; n >= SIUnitValue; n /= SIUnitValue {
		div *= SIUnitValue
		exp++
	}

	return printSI(bytesSize, div, exp, precision)
}

// 以1024作为基数 Binary Unit
func ByteCountIEC(bytesSize int64, precision int) string {
	div, exp := int64(1), 0
	for n := bytesSize; n >= IECUnitValue; n /= IECUnitValue {
		div *= IECUnitValue
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
	// fmt.Printf("pos: %d %s\n", pos, string(" kMGTPEZYRQ"[pos]))
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

func printIEC(bytesSize, div int64, exp, precision int) string {
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %s", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), toIECUnit(exp))
}

func printSI(bytesSize, div int64, exp, precision int) string {
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %s", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), toSIUnit(exp))
}

func toIEC(sizeAndUnit string, precision int, outputUnit string) string {
	return convertUnit(sizeAndUnit, precision, IECUnitValue, outputUnit, printIEC)
}

func toSI(sizeAndUnit string, precision int, outputUnit string) string {
	return convertUnit(sizeAndUnit, precision, SIUnitValue, outputUnit, printSI)
}

func convertUnit(sizeAndUnit string, precision int, unitValue int64, outputUnit string, printFunc func(bytesSize, div int64, exp, precision int) string) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	unit := unitValue
	div, exp := int64(1), 0
	loop := 0
	for n := bytesSize; loop < toPos(outputUnit); n /= unit {
		div *= unit
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
