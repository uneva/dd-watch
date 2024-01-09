//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func wireApp() {
	panic(wire.Build())
}
