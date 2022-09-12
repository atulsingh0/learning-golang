package simplemath

// function mathExp return another function
func MathExp() func(float64, float64) float64 {
	return func(f1 float64, f2 float64) float64 {
		return f1 + f2
	}
}

// func MathExpression(expr string) func(float64, float64) float64 {
// 	switch expr {
// 	case "add":
// 		return Add
// 	case "divide":
// 	}
// }
