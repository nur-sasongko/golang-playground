package main

import "sync"

var urlStore = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}
