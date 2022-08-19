package vec2

/*
These are trickier to write within Go's generics:

FloatBits, FloatFromBits - need to specify float32 or float64 because there is no way to deduce uint64 <-> float64, uint32 <-> float32
Mod/Rem - can be done with 2 different operation names but ints need x%y while floats need math.Mod(x,y)

Horizontal reductions
vec2.Split should return T not [1]T
math.Round
math.RoundToEven
math.Signbit
math.Copysign
*/
