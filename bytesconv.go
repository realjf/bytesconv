// #############################################################################
// # File: bytesconv.go                                                        #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:16:46                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/08/11 10:14:00                                        #
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
func ByteCountSI(b int64, precision int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "kMGTPE"[exp])
}

// 以1024作为基数
func ByteCountIEC(b int64, precision int) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	if unit_str := "KMGTPE"[exp]; string(unit_str) != "" {
		return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), unit_str)
	} else {
		return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), unit_str)
	}
}

// 单位之间转换
func ToKiB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 1; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToMiB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 2; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToGiB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 3; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToTiB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 4; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToPiB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 5; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToEiB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1024
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 6; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %ciB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToKB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 1; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToMB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 2; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToGB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 3; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToTB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 4; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToPB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 5; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}

func ToEB(in string, precision int) string {
	re := regexp.MustCompile(`^([0-9]+)\s*([a-zA-Z]+)$`)
	substring := re.FindStringSubmatch(in)
	if len(substring) == 0 {
		return ""
	}
	size, _ := strconv.ParseInt(substring[1], 10, 64)
	b := ToBytes(size, substring[2])

	const unit = 1000
	div, exp := int64(unit), -1
	loop := 0
	for n := b / unit; loop < 6; n /= unit {
		div *= unit
		exp++
		loop++
	}
	prec := "%." + strconv.Itoa(precision) + "f"
	return fmt.Sprintf("%s %cB", fmt.Sprintf(prec, float64(b)/float64(div)), "KMGTPE"[exp])
}
