// #############################################################################
// # File: bytesconv_test.go                                                   #
// # Project: bytesconv                                                        #
// # Created Date: 2023/08/11 09:22:16                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/08/11 10:14:54                                        #
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
	fmt.Println(bytesconv.ByteCountIEC(1000, 2))
	fmt.Println(bytesconv.ByteCountIEC(987654321, 3))
	fmt.Println(bytesconv.ToKiB("10M", 1))
	fmt.Println(bytesconv.ToKiB("1b", 6))
}
