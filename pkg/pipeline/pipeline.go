package pipeline

import (
	"github.com/davecgh/go-spew/spew"
)

func Pipeline(fns ...func(...interface{}) []interface{}) func(...interface{}) []interface{} {
	return func(args ...interface{}) (results []interface{}) {
		results = args
		for _, fn := range fns {
			results = fn(results...)
		}
		return
	}
}

func A(args ...interface{}) []interface{} {
	main := args[0].(int)
	return []interface{}{main * 2}
}

func B(args ...interface{}) []interface{} {
	main := args[0].(int)
	return []interface{}{main - 2}
}

func Run(a int) {
	spew.Dump(Pipeline(A, B)(a))
}

func LiftString(orig func(string) string) func(...interface{}) []interface{} {
	return func(args ...interface{}) []interface{} {
		if v, ok := args[0].(string); ok {
			return []interface{}{orig(v)}
		}
		panic(nil)
	}
}
