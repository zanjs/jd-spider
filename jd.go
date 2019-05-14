package main

import (
	"fmt"
)

// BaseCraw is
type BaseCraw struct {
	domain string
}

// Get is
func (b *BaseCraw) Get(s string) {
	fmt.Println(s)
	fmt.Println(b.domain)
}

// GetJD is
func (b *BaseCraw) GetJD() string {
	fmt.Println(b)
	return b.domain
}

// MethodMapFn is
func MethodMapFn(b BaseCraw) map[string]interface{} {
	return map[string]interface{}{
		"Get": b.Get,
	}
}
