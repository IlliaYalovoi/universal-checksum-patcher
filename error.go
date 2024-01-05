package main

import "errors"

var (
	ErrNoMatch    = errors.New("cannot detect bytes pattern to patch. Most likely patcher are outdated due to game updates")
	ErrCantLocate = errors.New("cannot locate file in current directory")
)
