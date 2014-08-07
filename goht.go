/*
 * Author: Emmanuel Odeke <odeke@ualberta.ca>
 *  Trie hashmap implementation
*/

package goht

import (
    "fmt"
    "math"
)

const (
    Base = 10
)

type Trie struct {
    children [] *Trie
    Value interface{}
    EOS bool
}

func New(v interface{}) *Trie {
   return new(Trie).Init(v)
}

func (t *Trie) Init(v interface{}) *Trie {
    t.EOS = false
    t.children = make([] *Trie, Base, Base)
    t.Value = v
    return t
}

func (t *Trie) SetTValue(v interface {}) (err error) {
    err = nil 
    if t == nil {
        err = fmt.Errorf("Cannot set value of a null node")
    } else {
        t.Value = v
        t.EOS = true
    }

    return
}

func (t *Trie) accessOp (h int, isPop bool) (retr interface{}) {
    // Handles 'Get' and 'Pop' operations. This is toggled by the param 'isPop'
    trav := t
    retr = nil
    
    powerCount := math.Ceil(math.Log10(float64(h)))
    for {
        if trav == nil || trav.children == nil {
            return
        } else if powerCount <= 0 {
            break
        } else {
            trav = trav.children[h % 10]
            h /= 10
            powerCount--
        }
    }

    if trav.EOS == true {
        retr = trav.Value
        if isPop == true {
            trav.EOS = false
            trav.Value = nil
        }
    }

    return
}

func (t *Trie) Pop(h int) (retr interface{}) {
    retr = nil
    if t == nil || h < 0 {
        return 
    } else {
        retr = t.accessOp(h, true)
    }

    return
}

func (t *Trie) Get(h int) (retr interface{}) {
    retr = nil
    if t == nil || h < 0 {
        return 
    } else {
        retr = t.accessOp(h, false)
    }

    return
}

func (t *Trie) Put(h int, v interface{}) (err error) {
    err = nil
    if t == nil || h < 0 {
        t = new(Trie).Init(v)
        return
    }

    trav := t
    radixCount := math.Ceil(math.Log10(float64(h)))

    for {
        if radixCount <= 0 {
            break
        }

        mod := h % Base
        if trav.children[mod] == nil {
            trav.children[mod] = New(nil)
        }

        trav = trav.children[mod]

        h /= Base
        radixCount--
    }

    if trav == nil {
        err = fmt.Errorf("Trav should not be nil")
    } else {
        err = trav.SetTValue(v)
    }

    return
}
