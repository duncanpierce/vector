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
