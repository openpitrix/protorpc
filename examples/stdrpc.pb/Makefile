# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

PROTO_FILES=$(sort $(wildcard ./*.proto))

default: $(PROTO_FILES) Makefile
	go install github.com/chai2010/protorpc/protoc-gen-stdrpc
	protoc --go_out=. ${PROTO_FILES}
	protoc --stdrpc_out=. ${PROTO_FILES}
	go test

clean:
