package main

import (
	"log"
	"testing"
)

func Test_starting(t *testing.T) {
	log.SetFlags(0)
	for _, p := range starting(7, []int{2, 1}) {
		log.Println(string(p))
	}
}
