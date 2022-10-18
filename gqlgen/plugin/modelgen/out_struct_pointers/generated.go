// Code generated by github.com/vmware-tanzu/graph-framework-for-microservices/gqlgen, DO NOT EDIT.

package out_struct_pointers

import (
	"fmt"
	"io"
	"strconv"
)

type A interface {
	IsA()
}

type ArrayOfA interface {
	IsArrayOfA()
}

type B interface {
	IsB()
}

type C interface {
	IsA()
	IsC()
}

type D interface {
	IsA()
	IsB()
	IsD()
}

type FooBarer interface {
	IsFooBarer()
}

// InterfaceWithDescription is an interface with a description
type InterfaceWithDescription interface {
	IsInterfaceWithDescription()
}

type MissingInterface interface {
	IsMissingInterface()
}

type MissingUnion interface {
	IsMissingUnion()
}

// UnionWithDescription is an union with a description
type UnionWithDescription interface {
	IsUnionWithDescription()
}

type CDImplemented struct {
	A string  `json:"a" database:"CDImplementeda"`
	B int     `json:"b" database:"CDImplementedb"`
	C bool    `json:"c" database:"CDImplementedc"`
	D *string `json:"d" database:"CDImplementedd"`
}

func (CDImplemented) IsC() {}

func (CDImplemented) IsA() {}

func (CDImplemented) IsD() {}

func (CDImplemented) IsB() {}

type CyclicalA struct {
	FieldOne   *CyclicalB `json:"field_one" database:"CyclicalAfield_one"`
	FieldTwo   *CyclicalB `json:"field_two" database:"CyclicalAfield_two"`
	FieldThree *CyclicalB `json:"field_three" database:"CyclicalAfield_three"`
	FieldFour  string     `json:"field_four" database:"CyclicalAfield_four"`
}

type CyclicalB struct {
	FieldOne   *CyclicalA `json:"field_one" database:"CyclicalBfield_one"`
	FieldTwo   *CyclicalA `json:"field_two" database:"CyclicalBfield_two"`
	FieldThree *CyclicalA `json:"field_three" database:"CyclicalBfield_three"`
	FieldFour  *CyclicalA `json:"field_four" database:"CyclicalBfield_four"`
	FieldFive  string     `json:"field_five" database:"CyclicalBfield_five"`
}

type FieldMutationHook struct {
	Name     *string       `json:"name" anotherTag:"tag" database:"FieldMutationHookname"`
	Enum     *ExistingEnum `json:"enum" yetAnotherTag:"12" database:"FieldMutationHookenum"`
	NoVal    *string       `json:"noVal" yaml:"noVal" repeated:"true" database:"FieldMutationHooknoVal"`
	Repeated *string       `json:"repeated" someTag:"value" repeated:"true" database:"FieldMutationHookrepeated"`
}

type ImplArrayOfA struct {
	TrickyField        []*CDImplemented `json:"trickyField" database:"ImplArrayOfAtrickyField"`
	TrickyFieldPointer []*CDImplemented `json:"trickyFieldPointer" database:"ImplArrayOfAtrickyFieldPointer"`
}

func (ImplArrayOfA) IsArrayOfA() {}

type MissingInput struct {
	Name *string      `json:"name" database:"MissingInputname"`
	Enum *MissingEnum `json:"enum" database:"MissingInputenum"`
}

type MissingTypeNotNull struct {
	Name     string              `json:"name" database:"MissingTypeNotNullname"`
	Enum     MissingEnum         `json:"enum" database:"MissingTypeNotNullenum"`
	Int      MissingInterface    `json:"int" database:"MissingTypeNotNullint"`
	Existing ExistingType        `json:"existing" database:"MissingTypeNotNullexisting"`
	Missing2 MissingTypeNullable `json:"missing2" database:"MissingTypeNotNullmissing2"`
}

func (MissingTypeNotNull) IsMissingInterface() {}

func (MissingTypeNotNull) IsExistingInterface() {}

func (MissingTypeNotNull) IsMissingUnion() {}

func (MissingTypeNotNull) IsExistingUnion() {}

type MissingTypeNullable struct {
	Name     *string             `json:"name" database:"MissingTypeNullablename"`
	Enum     *MissingEnum        `json:"enum" database:"MissingTypeNullableenum"`
	Int      MissingInterface    `json:"int" database:"MissingTypeNullableint"`
	Existing *ExistingType       `json:"existing" database:"MissingTypeNullableexisting"`
	Missing2 *MissingTypeNotNull `json:"missing2" database:"MissingTypeNullablemissing2"`
}

func (MissingTypeNullable) IsMissingInterface() {}

func (MissingTypeNullable) IsExistingInterface() {}

func (MissingTypeNullable) IsMissingUnion() {}

func (MissingTypeNullable) IsExistingUnion() {}

type NotCyclicalA struct {
	FieldOne string `json:"FieldOne" database:"NotCyclicalAFieldOne"`
	FieldTwo int    `json:"FieldTwo" database:"NotCyclicalAFieldTwo"`
}

type NotCyclicalB struct {
	FieldOne string       `json:"FieldOne" database:"NotCyclicalBFieldOne"`
	FieldTwo NotCyclicalA `json:"FieldTwo" database:"NotCyclicalBFieldTwo"`
}

type Recursive struct {
	FieldOne   *Recursive `json:"FieldOne" database:"RecursiveFieldOne"`
	FieldTwo   *Recursive `json:"FieldTwo" database:"RecursiveFieldTwo"`
	FieldThree *Recursive `json:"FieldThree" database:"RecursiveFieldThree"`
	FieldFour  string     `json:"FieldFour" database:"RecursiveFieldFour"`
}

type RenameFieldTest struct {
	BadName    string `json:"badName" database:"RenameFieldTestbadName"`
	OtherField string `json:"otherField" database:"RenameFieldTestotherField"`
}

// TypeWithDescription is a type with a description
type TypeWithDescription struct {
	Name *string `json:"name" database:"TypeWithDescriptionname"`
}

func (TypeWithDescription) IsUnionWithDescription() {}

type FooBarr struct {
	Name string `json:"name" database:"_Foo_Barrname"`
}

func (FooBarr) IsFooBarer() {}

// EnumWithDescription is an enum with a description
type EnumWithDescription string

const (
	EnumWithDescriptionCat EnumWithDescription = "CAT"
	EnumWithDescriptionDog EnumWithDescription = "DOG"
)

var AllEnumWithDescription = []EnumWithDescription{
	EnumWithDescriptionCat,
	EnumWithDescriptionDog,
}

func (e EnumWithDescription) IsValid() bool {
	switch e {
	case EnumWithDescriptionCat, EnumWithDescriptionDog:
		return true
	}
	return false
}

func (e EnumWithDescription) String() string {
	return string(e)
}

func (e *EnumWithDescription) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnumWithDescription(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnumWithDescription", str)
	}
	return nil
}

func (e EnumWithDescription) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MissingEnum string

const (
	MissingEnumHello   MissingEnum = "Hello"
	MissingEnumGoodbye MissingEnum = "Goodbye"
)

var AllMissingEnum = []MissingEnum{
	MissingEnumHello,
	MissingEnumGoodbye,
}

func (e MissingEnum) IsValid() bool {
	switch e {
	case MissingEnumHello, MissingEnumGoodbye:
		return true
	}
	return false
}

func (e MissingEnum) String() string {
	return string(e)
}

func (e *MissingEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MissingEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MissingEnum", str)
	}
	return nil
}

func (e MissingEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
