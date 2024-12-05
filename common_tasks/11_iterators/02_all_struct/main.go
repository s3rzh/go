package main

import (
	"fmt"
	"github.com/s3rzh/cloud_native_go/iter/iter1/secret"
)

func main() {
	ints := secret.Ints{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, false},
		{6, false},
		{7, true},
		{8, false},
	}

	for n := range ints.All {
		if !secret.Int(n).Visible {
			continue
		}

		fmt.Println(n)
	}
}
