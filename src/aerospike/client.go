// Copyright FreeWheel Inc. All Rights Reserved.
// Author: jjli@freewheel.tv(Li JingJie)
// Created Time: Thu Jan  5 10:11:41 2017

package aerospike

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
)

func Client() {
	client, err := as.NewClient("172.24.10.101", 3000)
	if err != nil {
		fmt.Println("Open aerospike failed!")
		return
	} else {
		fmt.Println("Open aerospike success!")
	}

	// key, err := as.NewKey("namespace", "set",
	// 	"key value goes here and can be any supported primitive")

	// bin1 := as.NewBin("bin1", "value1")
	// bin2 := as.NewBin("bin2", "value2")

	// // Write a record
	// err = client.PutBins(nil, key, bin1, bin2)

	// // Read a record
	// record, err := client.Get(nil, key)

	client.Close()
}
