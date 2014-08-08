package goht

import (
    "testing"
)

func TestInitialization(t *testing.T) {
    hv := New("Emmanuel")
    if hv == nil {
        t.Error("Expecting initialized element")
    }

    if hv.Value != "Emmanuel" {
        t.Error("Expecting Emmanuel back")
    }

    if hv.Pop(10) != nil {
        t.Error("Not expecting any content in the map")
    }
}

func TestRetrieval(t *testing.T) {
    hp := New(nil)
    strMap := map[interface{}]int {
        "combo": 23, "anglo": 25, "sanglo": 29,
        "creme": 123, "pene": 678, "ignot": 97,
         "frugal": 99, 1238: 1238,
    }

    for obj, h := range strMap {
        err := hp.Put(h, obj)
        if err != nil {
            t.Error("Error encountered while putting: %v %v", h, obj)
        }
    }

    var retr interface{}
    for rObj, rH := range strMap {
        retr = hp.Get(rH)
        if retr != rObj {
            t.Error("Retrieval failed for", rObj)
        }
    }
}
