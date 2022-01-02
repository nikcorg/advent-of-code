package main

import "strings"

func concat[T any](as []T, bs []T) []T {
	return append(as, bs...)
}

func split(sep string) func(string) []string {
	return func(x string) []string {
		return strings.Split(x, sep)
	}
}

func fmap[T any, U any](f func(T) U, xs []T) []U {
	ys := make([]U, len(xs))
	for i := range xs {
		ys[i] = f(xs[i])
	}
	return ys
}

func foldl[T any, Tacc any](f func(Tacc, T) Tacc, acc0 Tacc, xs []T) Tacc {
	acc := acc0
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}
