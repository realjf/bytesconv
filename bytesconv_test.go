// #############################################################################
// # File: bytesconv_test.go                                                   #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:22:16                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/15 17:42:42                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023                                                        #
// #############################################################################
package bytesconv_test

import (
	"fmt"
	"testing"

	"github.com/realjf/bytesconv"
)

func TestByteCountIEC(t *testing.T) {
	fmt.Println(bytesconv.ByteCountSI(1, 2))
	fmt.Println(bytesconv.ByteCountSI(10, 2))
	fmt.Println(bytesconv.ByteCountSI(1000, 2))
	fmt.Println(bytesconv.ByteCountSI(10000, 2))
	fmt.Println(bytesconv.ByteCountSI(100000, 2))
	fmt.Println(bytesconv.ByteCountSI(1000000, 2))
	fmt.Println(bytesconv.ByteCountIEC(98, 2))
	fmt.Println(bytesconv.ByteCountIEC(1024, 2))
	fmt.Println(bytesconv.ByteCountIEC(98765, 3))
	fmt.Println(bytesconv.ByteCountIEC(987654321, 3))
	fmt.Println(bytesconv.ToKiB("1B", 6))
	fmt.Println(bytesconv.ToKiB("10B", 4))
	fmt.Println(bytesconv.ToKiB("10K", 1))
	fmt.Println(bytesconv.ToKiB("10M", 1))
	fmt.Println(bytesconv.ToKiB("10G", 1))
	fmt.Println(bytesconv.ToKiB("10T", 1))
	fmt.Println(bytesconv.ToKiB("10P", 1))
	fmt.Println(bytesconv.ToKB("1B", 6))
	fmt.Println(bytesconv.ToKB("10B", 4))
	fmt.Println(bytesconv.ToKB("10K", 1))
	fmt.Println(bytesconv.ToKB("10M", 1))
	fmt.Println(bytesconv.ToKB("10G", 1))
	fmt.Println(bytesconv.ToKB("10T", 1))
	fmt.Println(bytesconv.ToKB("10P", 1))
}
