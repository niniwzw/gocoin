package main

/*
  If you prefer to use OpenSSL implementation for verifying
  transaction signatures:
   1) Copy this file one level up (to the "./client" folder)
   2) Remove "speedup.go" from the client folder
   3) Redo "go build"
*/

import (
	"fmt"
	"github.com/niniwzw/gocoin/lib/others/cgo/openssl"
	"github.com/niniwzw/gocoin/lib/btc"
)

func EC_Verify(k, s, h []byte) bool {
	return openssl.EC_Verify(k, s, h)==1
}

func init() {
	fmt.Println("Using OpenSSL wrapper for EC_Verify")
	btc.EC_Verify = EC_Verify
}
