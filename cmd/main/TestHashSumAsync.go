package main

import (
	"fmt"
	"time"

	"gitlab.com/auk-go/pathhelper/checksummer"
	"gitlab.com/auk-go/pathhelper/hashas"
)

func TestHashSumAsync() {
	start := time.Now()
	c := checksummer.NewAsync(true, "D:\\vm", hashas.Md5)
	elapsed := time.Since(start)
	prettyPrint(c.GetMap())
	fmt.Printf("Elapsed Async: %s\n", elapsed)
}
