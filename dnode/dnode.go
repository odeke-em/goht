// Author: Emmanuel Odeke <odeke@ualberta.ca>

package dnode

type DNode struct {
    size uint
    prev *DNode
    next *DNode
    Value interface{}
}

func New(v interface{}) *DNode {
    return new(DNode).Init(v)
}

func (d *DNode) Init(v interface{}) *DNode {
    d.size = 1
    d.prev = nil   
    d.next = nil
    d.Value = v

    return d
}

func (d *DNode) GetValue() interface{} {
    return d.Value
}

func (d *DNode) SetValue(v interface{}) (vacated interface{}) {
    vacated = d.Value
    d.Value = v
    return
}

func (d *DNode) GetPrev() *DNode {
    return d.prev
}

func (d *DNode) GetNext() *DNode {
    return d.next
}

func (d *DNode) SetPrev(p *DNode) {
    d.prev = p
}

func (d *DNode) SetNext(n *DNode) {
    d.next = n
}

func (d *DNode) Prepend(v interface{}) *DNode {
    fresh := New(v) 
    if d.prev != nil {
        fresh.prev = d.prev 
    }

    d.prev = fresh
    fresh.next = d

    d.size++

    return fresh
}

func (d *DNode) Append(v interface{}) *DNode {
    fresh := New(v) 
    if d.next != nil {
        fresh.next = d.next
    }

    d.next = fresh
    fresh.prev = d
    d.size++

    return fresh
}

func (d *DNode) Shift() *DNode {
    d = d.next
    return d
}

func (d *DNode) UnShift() *DNode {
    d = d.prev
    return d
}
