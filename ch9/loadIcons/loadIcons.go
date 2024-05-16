package main

import (
	"sync"
	"time"
)

var loadOnce sync.Once
var icons map[string]string

func loadIcon(name string) string {
	// long time operation
	// block 1 second
	time.Sleep(1 * time.Second)
	return name
}

func loadIcons() {
	icons = make(map[string]string)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
}

func Icon(name string) string {
	loadOnce.Do(loadIcons) // Concurrency-safe
	return icons[name]
}
