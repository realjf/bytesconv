# ##############################################################################
# # File: Makefile                                                             #
# # Project: bytesconv                                                         #
# # Created Date: 2023/08/15 09:38:00                                          #
# # Author: realjf                                                             #
# # -----                                                                      #
# # Last Modified: 2023/08/15 11:35:32                                         #
# # Modified By: realjf                                                        #
# # -----                                                                      #
# # Copyright (c) 2023                                                         #
# ##############################################################################


# test
.PHONY: test
test:
	@go test ./... -v


# lint
.PHONY: lint
lint:
	@golangci-lint run -v ./...

