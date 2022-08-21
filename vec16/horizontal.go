package vec16

import "github.com/duncanpierce/vector/constraintsExt"

/*
IsDuplicate checks elements of vector a for values which are equal to one another. If several elements have the same value,
exactly one is chosen as the original. Elements other than the original are deemed to be duplicates and have the corresponding bit set
in the boolean vector result.
It is not defined which value is chosen as the original. Your code should assume it is chosen at random.
*/
func IsDuplicate[Element constraintsExt.Number](m Bool, a [Length]Element) (r Bool) {
	seen := map[Element]bool{}
	m.ForTrue(func(i int) {
		r.Set(i, seen[a[i]])
		seen[a[i]] = true
	})
	return
}

/*
ElementSum adds each element selected by the Bool mask and returns their sum. If no elements are selected by the Bool mask, the result is 0.
*/
func ElementSum[Element constraintsExt.Number](m Bool, a [Length]Element) (r Element) {
	m.ForTrue(func(i int) {
		r += a[i]
	})
	return
}

/*
ElementProduct multiplies each element selected by the Bool mask and returns their product. If no elements are selected by the Bool mask, the result is 1.
*/
func ElementProduct[Element constraintsExt.Number](m Bool, a [Length]Element) (r Element) {
	r = 1
	m.ForTrue(func(i int) {
		r *= a[i]
	})
	return
}

/*
ElementMin returns the minimum of the elements selected by the Bool mask. If no elements are selected by the Bool mask, the result ok is false.
*/
func ElementMin[Element constraintsExt.Number](m Bool, a [Length]Element) (min Element, ok bool) {
	return ReduceElements(m, a, func(x, y Element) Element {
		if x < y {
			return x
		}
		return y
	})
}

/*
ElementMax returns the maximum of the elements selected by the Bool mask. If no elements are selected by the Bool mask, the result ok is false.
*/
func ElementMax[Element constraintsExt.Number](m Bool, a [Length]Element) (min Element, ok bool) {
	return ReduceElements(m, a, func(x, y Element) Element {
		if x > y {
			return x
		}
		return y
	})
}
