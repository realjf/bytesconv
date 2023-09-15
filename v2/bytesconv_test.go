// #############################################################################
// # File: bytesconv_test.go                                                   #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:22:16                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/15 18:15:54                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023                                                        #
// #############################################################################
package bytesconv_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/realjf/bytesconv/v2"
)

func TestByteCountIEC(t *testing.T) {
	fmt.Println(bytesconv.ByteCountSI(big.NewInt(1), 2))
	fmt.Println(bytesconv.ByteCountSI(big.NewInt(10), 2))
	fmt.Println(bytesconv.ByteCountSI(big.NewInt(1000), 2))
	fmt.Println(bytesconv.ByteCountSI(big.NewInt(10000), 2))
	fmt.Println(bytesconv.ByteCountSI(big.NewInt(1000000), 2))
	fmt.Println(bytesconv.ByteCountSI(big.NewInt(1000000000), 2))
	fmt.Println(bytesconv.ByteCountIEC(big.NewInt(10), 2))
	fmt.Println(bytesconv.ByteCountIEC(big.NewInt(1024), 2))
	fmt.Println(bytesconv.ByteCountIEC(big.NewInt(9104342), 3))
	fmt.Println(bytesconv.ByteCountIEC(big.NewInt(143421343), 3))
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
