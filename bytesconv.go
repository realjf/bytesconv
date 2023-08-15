// #############################################################################
// # File: bytesconv.go                                                        #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:16:46                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/08/15 09:32:17                                        #
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

// size是数值，uint是单位
func ToBytes(size int64, unit string) int64 {
	switch strings.ToLower(unit) {
	case "kb":
		fallthrough
	case "k":
		return size * KB
	case "mb":
		fallthrough
	case "m":
		return size * MB
	case "gb":
		fallthrough
	case "g":
		return size * GB
	case "tb":
		fallthrough
	case "t":
		return size * TB
	case "pb":
		fallthrough
	case "p":
		return size * PB
	case "eb":
		fallthrough
	case "e":
		return size * EB
	default:
		return size
	}
}

// 以1000作为基数
func ByteCountSI(bytesSize int64, precision int) string {
	const unit = 1000
	if bytesSize < unit {
		return fmt.Sprintf("%d B", bytesSize)
	}
	div, exp := int64(unit), 0
	for n := bytesSize / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "kMGTPE"[exp])
}

// 以1024作为基数
func ByteCountIEC(bytesSize int64, precision int) string {
	const unit = 1024
	if bytesSize < unit {
		return fmt.Sprintf("%d B", bytesSize)
	}
	div, exp := int64(unit), 0
	for n := bytesSize / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	if unit_str := "KMGTPE"[exp]; string(unit_str) != "" {
		return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), unit_str)
	} else {
		return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), unit_str)
	}
}

// 单位之间转换
func ToKiB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 1; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToMiB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 2; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToGiB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 3; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToTiB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 4; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToPiB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 5; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToEiB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 6; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToKB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 1; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToMB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 2; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToGB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 3; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToTB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 4; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToPB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 5; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}

func ToEB(sizeAndUnit string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(sizeAndUnit)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	bytesSize := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := bytesSize / unit; loop < 6; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(bytesSize)/float64(div)), "KMGTPE"[exp])
}
