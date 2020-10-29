package pascaltriangle

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	var trianglehigh = 30
	for i := 0; i < trianglehigh; i++ {
		// Just add white space for view
		for k := trianglehigh - i; k > 0; k-- {
			fmt.Printf(" ")
		} // for
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", combination(i, j))
		}
		fmt.Printf("\n")
	}
}

func TestDP(t *testing.T) {
	var trianglehigh = 15
	for i := 0; i < trianglehigh; i++ {
		// Just add white space for view
		for k := trianglehigh - i; k > 0; k-- {
			fmt.Printf(" ")
		} // for
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", combinationDP(i, j))
		}
		fmt.Printf("\n")
	}

}

func BenchmarkBasic(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var trianglehigh = 15
		for i := 0; i < trianglehigh; i++ {
			// Just add white space for view
			for j := 0; j < i+1; j++ {
				combination(i, j)
			}
		}
	}
}

func BenchmarkDP(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var trianglehigh = 15
		for i := 0; i < trianglehigh; i++ {
			// Just add white space for view
			for j := 0; j < i+1; j++ {
				combinationDP(i, j)
			}
		}
	}
	// b.Log(mpt)
}

func TestGenerate(t *testing.T) {

	t.Log(generate2(100))
}
