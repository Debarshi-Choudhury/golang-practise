//go:build pro || enterprise

package main

func init() {
	features = append(features,
		"Pro feature #1",
		"Pro feature #2",
	)
}

//pro features will be available with both pro and enterprise build tags
