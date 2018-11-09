package main

import "go/types"

type objType int

const (
	booleanType objType = iota + 1
	integerType
	realType
	stringType
	nameType
	arrayType
	dictionaryType
	streamType
	nullType
)

type Object interface {
	getNumber() int
	getGeneration() int
	getType() objType
}

type BaseObj struct {
	Number     int
	Generation int
	Type       objType
}

type BooleanObj struct {
	BaseObj
	Content bool
}

func (o *BooleanObj) getNumber() int {
	return o.Number
}

func (o *BooleanObj) getGeneration() int {
	return o.Generation
}

func (o *BooleanObj) getType() objType {
	return o.Type
}

func NewBooleanObj(Number int, Gen int, Content bool) BooleanObj {
	return BooleanObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       booleanType,
		},
		Content: Content,
	}
}

type IntegerObj struct {
	BaseObj
	Content int
}

func (o *IntegerObj) getNumber() int {
	return o.Number
}

func (o *IntegerObj) getGeneration() int {
	return o.Generation
}

func (o *IntegerObj) getType() objType {
	return o.Type
}

func NewIntegerObj(Number int, Gen int, Content int) IntegerObj {
	return IntegerObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       integerType,
		},
		Content: Content,
	}
}

type RealObj struct {
	BaseObj
	Content float64
}

func (o *RealObj) getNumber() int {
	return o.Number
}

func (o *RealObj) getGeneration() int {
	return o.Generation
}

func (o *RealObj) getType() objType {
	return o.Type
}

func NewRealObj(Number int, Gen int, Content float64) RealObj {
	return RealObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       realType,
		},
		Content: Content,
	}
}

type StringObj struct {
	BaseObj
	Content string
}

func (o *StringObj) getNumber() int {
	return o.Number
}

func (o *StringObj) getGeneration() int {
	return o.Generation
}

func (o *StringObj) getType() objType {
	return o.Type
}

func NewStringObj(Number int, Gen int, Content string) StringObj {
	return StringObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       stringType,
		},
		Content: Content,
	}
}

type NameObj struct {
	BaseObj
	Content string
}

func (o *NameObj) getNumber() int {
	return o.Number
}

func (o *NameObj) getGeneration() int {
	return o.Generation
}

func (o *NameObj) getType() objType {
	return o.Type
}

func NewNameObj(Number int, Gen int, Content string) NameObj {
	return NameObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       nameType,
		},
		Content: Content,
	}
}

type ArrayObj struct {
	BaseObj
	Content []Object
}

func (o *ArrayObj) getNumber() int {
	return o.Number
}

func (o *ArrayObj) getGeneration() int {
	return o.Generation
}

func (o *ArrayObj) getType() objType {
	return o.Type
}

func NewArrayObj(Number int, Gen int, Content []Object) ArrayObj {
	return ArrayObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       arrayType,
		},
		Content: Content,
	}
}

type DictionaryObj struct {
	BaseObj
	Content map[string]Object
}

func (o *DictionaryObj) getNumber() int {
	return o.Number
}

func (o *DictionaryObj) getGeneration() int {
	return o.Generation
}

func (o *DictionaryObj) getType() objType {
	return o.Type
}

func NewDictionaryObj(Number int, Gen int, Content map[string]Object) DictionaryObj {
	return DictionaryObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       dictionaryType,
		},
		Content: Content,
	}
}

type StremObj struct {
	BaseObj
	Content string
}

func (o *StremObj) getNumber() int {
	return o.Number
}

func (o *StremObj) getGeneration() int {
	return o.Generation
}

func (o *StremObj) getType() objType {
	return o.Type
}

func NewStreamObj(Number int, Gen int, Content string) StremObj {
	return StremObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       streamType,
		},
		Content: Content,
	}
}

type NullObj struct {
	BaseObj
	Content types.Nil
}

func (o *NullObj) getNumber() int {
	return o.Number
}

func (o *NullObj) getGeneration() int {
	return o.Generation
}

func (o *NullObj) getType() objType {
	return o.Type
}

func NewNullObj(Number int, Gen int, Content types.Nil) NullObj {
	return NullObj{
		BaseObj: BaseObj{
			Number:     Number,
			Generation: Gen,
			Type:       nullType,
		},
		Content: Content,
	}
}
