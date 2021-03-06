package ordset

import (
	"testing"
)

func TestInitialization(t *testing.T) {
	od := New()
	if od == nil {
	}

	if od.dl == nil {
		t.Errorf("Head node cannot be nil")
	}

	if od.jar == nil {
		t.Errorf("Jar cannot be nil")
	}
}

func TestInsertion(t *testing.T) {
	attrMap := map[uint]interface{}{
		100: "Amber", 2938: "tribe",
		94585: "iron", 4586: nil, 193851: t,
	}

	od := New()
	for h, v := range attrMap {
		vac, isFirstInsert := od.Add(h, v)
		if isFirstInsert != true {
			t.Errorf("Expected a unique insert, vacated: %v", vac)
		}
	}

	// Now doing second phase of insertions
	for h, v := range attrMap {
		vac, isFirstInsert := od.Add(h, v)
		if isFirstInsert != false {
			t.Errorf("Expected a unique insert, vacated: %v", vac)
		}
	}

}

func TestRemove(t *testing.T) {
	elemMap := map[uint]interface{}{
		999: "ignot", 1820: "feng",
		1924: "Ambivalent", 28457: 2948.18,
	}

	od := New()
	for h, v := range elemMap {
		od.Add(h, v)
	}

	for h, _ := range elemMap {
		_, existed := od.Remove(h)
		if existed != true {
			t.Errorf("Should have been vacated")
		}
	}
}

func TestValuesOrderedByKeys(t *testing.T) {
	max := uint(100)
	implicitlySortedValues := make([]uint, 0, max)
	od := New()

	for i := uint(0); i < max; i++ {
		od.Add(i, i)
		implicitlySortedValues = append(implicitlySortedValues, i)
	}

	if expected, got := len(implicitlySortedValues), od.Length(); expected != got {
		t.Errorf("expected length %v got %v", expected, got)
	}

	valuesSortedByKeys := od.ValuesOrderedByKeys()

	if expected, got := len(implicitlySortedValues), len(valuesSortedByKeys); expected != got {
		t.Errorf("expected length %v got %v", expected, got)
	}

	// Element by element comparison since []uint != []inteface{}
	for i, impValue := range implicitlySortedValues {
		retr := valuesSortedByKeys[i]
		if retr != impValue {
			t.Errorf("element %v:: %v != %v", i, retr, impValue)
		}
	}
}
