package list

import (
	"math/rand"
	"reflect"
	"sort"
	"time"
)

type List struct {
	Items []interface{}
	Index int
}

func NewEmptyList() *List {
	return &List{}
}

func NewList(items []interface{}) *List {
	return &List{Items: items, Index: 0}
}

func (p *List) Remove(item interface{}) {
	for i, v := range p.Items {
		if item == v {
			p.Items = append(p.Items[:i], p.Items[i+1:]...)
		}
	}
}

func (p *List) HasNext() bool {
	return len(p.Items) > p.Index
}

func (p *List) Next() interface{} {
	if p.HasNext() {
		item := p.Items[p.Index]
		p.Index += 1
		return item
	}
	return nil
}

func (p *List) Get(i int) interface{} {
	if i > len(p.Items) {
		return nil
	}
	return p.Items[i]
}

func (p *List) Len() int {
	return len(p.Items)
}

func (p *List) Add(item interface{}) {
	p.Items = append(p.Items, item)
	return
}

func (p *List) Random() interface{} {
	i := rand.New(rand.NewSource(time.Now().Unix())).Intn(p.Len())
	return p.Get(i)
}

func (p *List) Cols(col string) []interface{} {
	cols := make([]interface{}, 0)
	for _, item := range p.Items {
		cols = append(cols, reflect.ValueOf(item).FieldByName(col).Interface())
	}
	return cols
}

func (p *List) ColChains(col string) *List {
	cols := make([]interface{}, 0)
	for _, item := range p.Items {
		cols = append(cols, reflect.ValueOf(item).FieldByName(col).Interface())
	}
	return NewList(cols)
}

func (p *List) Filter(exp func(item interface{}) bool) []interface{} {
	filters := make([]interface{}, 0)
	for _, item := range p.Items {
		if exp(item) {
			filters = append(filters, item)
		}
	}
	return filters
}

func (p *List) FilterChains(exp func(item interface{}) bool) *List {
	filters := make([]interface{}, 0)
	for _, item := range p.Items {
		if exp(item) {
			filters = append(filters, item)
		}
	}
	return NewList(filters)
}

func (p *List) Sort(exp func(i, j int) bool) *List {
	sort.Slice(p.Items, exp)
	return p
}

func (p *List) Foreach(exp func(item interface{})) {
	for _, item := range p.Items {
		exp(item)
	}
}
