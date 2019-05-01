package sleep

import(
"testing"
)

func TestSafetyInPositiveInput(t *testing.T) {
	testData := []uint64{5,4,2,1,4}
	t.Log("Input Data:", testData)
	r := Sort(testData)

	for i:=0 ;i<len(r); i++{
		next := i + 1
		if next >= len(r){
			break
		}
		if r[i] > r[next] {
			t.Fatal("Failed!")
		}
	}
	t.Log("Result,",r, "Pass!")
}