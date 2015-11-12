// Author: Emmanuel Odeke <odeke@ualberta.ca>

package dnode

import "testing"

func TestInit(t *testing.T) {
	v := int(10)
	d := New(v)
	if d == nil {
		t.Errorf("Expecting non nil dnode")
	}

	if d.Value != v {
		t.Errorf("Expecting %v got: %v", v, d.Value)
	}
}

func TestPrepend(t *testing.T) {
	keys := []interface{}{
		"TotalRecall", "IntoTheChopper",
		"axiom", "Aloha", nil, 2395.8, "Bonjourne",
	}

	d := New(nil)

	for _, v := range keys {
		d = d.Prepend(v)
		if d == nil {
			t.Errorf("Expected non nill node")
		}
	}

	trav := d
	// Remember first element is at nil
	for i := len(keys); i >= 1 && trav != nil; i-- {
		key := keys[i-1]
		if key != trav.Value {
			t.Errorf("Expected %v got %v", key, trav.Value)
		}

		trav = trav.Shift()
	}
}

func TestAppend(t *testing.T) {
	keys := []interface{}{
		"Anglo", t, "Key & Peele", 0xff,
		100.0, "Unity", nil, "Integral",
	}

	d := New(nil)

	for _, v := range keys {
		d = d.Append(v)
		if d == nil {
			t.Errorf("Expected non nill node")
		}
	}

	trav := d.next
	// Remember first element is at nil
	for i := 0; i < len(keys) && trav != nil; i++ {
		key := keys[i]
		if key != trav.Value {
			t.Errorf("Expected %v got %v", key, trav.Value)
		}

		trav = trav.Shift()
	}
}
