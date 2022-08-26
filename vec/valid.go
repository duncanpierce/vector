package vec

import "fmt"

func assignable[E any](zSlice, xSlice []E, xBroad bool, xIdent string) {
	if xBroad {
		if len(xSlice) > len(zSlice) {
			panic(fmt.Sprintf("broadcastable vector %s cannot be longer than vector z", xIdent))
		}
	} else {
		if len(xSlice) != len(zSlice) {
			panic(fmt.Sprintf("vector %s must be the same length as vector z", xIdent))
		}
	}
}
