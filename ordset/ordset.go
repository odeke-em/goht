/*
    Author: Emmanuel Odeke <odeke@ualberta.ca>
    Description: Data structure that allows for set operations:
        put, pop, get
    however in O(1) time.
    The caveat however is that extra memory is used since under the
    hood, it consists of a doubly linked list as well as a hashmap.
    The hashmap is a item-hash to doubly linked list node map.
*/

package ordset

import (
    "github.com/odeke-em/goht/dnode"
)

type OrdSet struct {
    dl *dnode.DNode
    jar map[uint] *dnode.DNode
}

func New() *OrdSet {
    od := new(OrdSet)
    od.dl = dnode.New(nil)
    od.jar = map[uint] *dnode.DNode {}

    return od
}

func (o *OrdSet) getDNode(h uint) *dnode.DNode {
    if o.jar == nil {
        return nil
    } else {
        retr, _ := o.jar[h]
        return retr
    }
}

func (o *OrdSet) Get(h uint) (interface{}, bool) {
    if o.jar == nil {
        return nil, false
    } else {
        retr, check := o.jar[h]
        if check == false || retr == nil{
            return nil, false
        } else {
            return retr.Value, true
        }
    }
}

func (o *OrdSet) putIntoJar(k uint, v *dnode.DNode) {
    o.jar[k] = v
}

func (o *OrdSet) Add(h uint, item interface{}) (vacated interface{}, fIns bool) {
    vacated = nil
    fIns = true // First time insert

    retr := o.getDNode(h)
    if retr == nil { // Time to add it
        nd := o.dl.Prepend(item)
        o.putIntoJar(h, nd)
    } else { // Update that value
        vacated = retr.SetValue(item)
        fIns = false
    }

    return
}

func (o *OrdSet) Remove(h uint) (vacated interface{}, exists bool) {
    retr := o.getDNode(h)

    if retr == nil {
        return nil, false 
    } else {
        prev := retr.GetPrev()
        next := retr.GetNext()

        if prev != nil {
            prev.SetNext(next)
        }

        if next != nil {
            next.SetPrev(prev)
        }

        delete(o.jar, h)
        return retr.Value, true
    }
}

func (o *OrdSet) Iterator() *dnode.DNode {
    return o.dl
}
