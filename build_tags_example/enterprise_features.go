//go:build enterprise

package main

func init() {
	features = append(features,
		"Enterprise feature #1",
		"Enterprise feature #2",
	)
}

//pro features will be available with only enterprise build tag
