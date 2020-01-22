package models

import (
	"reflect"
)

// Reflected stores some critical information during reflection loops
type Reflected struct {
	Reflection reflect.Value
	Field      reflect.Value
	Type       reflect.StructField
	Tag        reflect.StructTag
}

// Configure is called inside reflection NumField loops to instantiate that field
func (r *Reflected) Configure(i int) {
	r.Field = r.Reflection.Field(i)
	r.Type = r.Reflection.Type().Field(i)
	r.Tag = r.Type.Tag
}

// Valid returns true if the field is fully valid
func (r *Reflected) Valid() bool {
	if !r.Field.IsValid() || r.Field.IsNil() || r.Field.IsZero() {
		return false
	}
	return true
}

// Field encapsulates the underlying interface of a sub model
// ex. Compression Enabled / GZIP / Mime
type Field struct {
	Name    string
	Content interface{}
}

// Parent is an encapsulation of several sub models (*Compression)
type Parent struct {
	Name   string
	Fields []*Field
}

// Exists ...
func (p *Parent) Exists(name string) bool {
	for _, field := range p.Fields {
		if field.Name == name {
			return true
		}
	}
	return false
}

// GetIndex ...
func (p *Parent) GetIndex(name string) (*int, bool) {
	for i, field := range p.Fields {
		if field.Name == name {
			return &i, true
		}
	}
	return nil, false
}

// Get ...
func (p *Parent) Get(name string) *Field {
	for _, field := range p.Fields {
		if field.Name == name {
			return field
		}
	}
	return nil
}

// Update ...
func (p *Parent) Update(name string, f *Field) {
	if i, ok := p.GetIndex(name); ok {
		p.Fields[*i] = f
	}
}

// Add ...
func (p *Parent) Add(name string) *Field {
	if !(p.Exists(name)) {
		f := &Field{
			Name: name,
		}
		p.Fields = append(p.Fields, f)
		return f
	}
	return nil
}

// Dump ...
func (p *Parent) Dump() map[string]interface{} {
	ret := make(map[string]interface{})

	for _, field := range p.Fields {
		ret[field.Name] = []interface{}{field.Content}
	}

	return ret
}

// Parents is an encapsulation of several parents
// ex. Delivery
type Parents struct {
	Queue []*Parent
}

// Dump ...
func (p *Parents) Dump() map[string]interface{} {
	tmp := make(map[string]interface{})

	for _, parent := range p.Queue {
		tmp[parent.Name] = parent.Dump()
	}

	return tmp

}

// Exists ...
func (p *Parents) Exists(name string) bool {
	for _, parent := range p.Queue {
		if parent.Name == name {
			return true
		}
	}
	return false
}

// GetIndex ...
func (p *Parents) GetIndex(name string) (*int, bool) {
	for i, parent := range p.Queue {
		if parent.Name == name {
			return &i, true
		}
	}
	return nil, false
}

// Get ...
func (p *Parents) Get(name string) *Parent {
	for _, parent := range p.Queue {
		if parent.Name == name {
			return parent
		}
	}
	return nil
}

// Update ...
func (p *Parents) Update(name string, parent *Parent) {
	if i, ok := p.GetIndex(name); ok {
		p.Queue[*i] = parent
	}
}

// AddParent ...
func (p *Parents) AddParent(name string) *Parent {
	if !(p.Exists(name)) {
		parent := &Parent{
			Name: name,
		}
		p.Queue = append(p.Queue, parent)
		return parent
	}
	return nil
}
