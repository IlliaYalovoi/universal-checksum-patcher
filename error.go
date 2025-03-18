package main

import "errors"

var (
	errNoMatch    = errors.New("cannot detect bytes pattern to patch. Most likely patcher are outdated due to game updates or pather already was applied")
	errCantLocate = errors.New("cannot locate file in current directory")
)
