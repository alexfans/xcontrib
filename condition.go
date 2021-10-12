package xcontrib

import "fmt"

type predicate string

type operator string

type kind string

type paramType string

const (
	Empty        predicate = ""
	EQ           predicate = "EQ"
	NEQ          predicate = "NEQ"
	GT           predicate = "GT"
	GTE          predicate = "GTE"
	LT           predicate = "LT"
	LTE          predicate = "LTE"
	In           predicate = "In"
	NotIN        predicate = "NotIN"
	Contains     predicate = "Contains"
	ContainsFold predicate = "ContainsFold"
	HasPrefix    predicate = "HasPrefix"
	HasSuffi     predicate = "HasSuffi"
)

const (
	And operator = "And"
	Or  operator = "Or"
)

const (
	VarKind      kind = "var"
	ConstKind    kind = "const"
	OperatorKind kind = "operator"
)

const (
	Int8     paramType = "Int8"
	Int16    paramType = "Int16"
	Int32    paramType = "Int32"
	Int64    paramType = "Int64"
	Int      paramType = "Int"
	Float32  paramType = "Float32"
	Float64  paramType = "Float64"
	String   paramType = "String"
	Time     paramType = "Times"
	Int8s    paramType = "Int8s"
	Int16s   paramType = "Int16s"
	Int32s   paramType = "Int32s"
	Int64s   paramType = "Int64s"
	Ints     paramType = "Ints"
	Float32s paramType = "Float32s"
	Float64s paramType = "Float64s"
	Strings  paramType = "Strings"
	Times    paramType = "Times"
	Any      paramType = "Any"
)

func (s predicate) String() string {
	return string(s)
}

func (o operator) String() string {
	return string(o)
}

func (k kind) String() string {
	return string(k)
}

func (s paramType) String() string {
	return string(s)
}

func KindIN(find kind, kinds ...kind) bool {
	for _, kind := range kinds {
		if kind == find {
			return true
		}
	}
	return false
}

type Condition struct {
	Kind       kind        // var, const, operator
	Column     string      // var, const
	Predicate  predicate   // var, const
	Value      string      // const
	Operator   operator    // operator
	Conditions []Condition // operator
}

func (c Condition) GetKind() kind {
	return c.Kind
}

func (c Condition) GetColumn() string {
	if KindIN(c.GetKind(), VarKind, ConstKind) {
		return c.Column
	} else {
		panic(fmt.Sprintf("%s cannot get Column\n", c.GetKind().String()))
	}
}

func (c Condition) GetPredicate() predicate {
	if KindIN(c.GetKind(), VarKind, ConstKind) {
		return c.Predicate
	} else {
		panic(fmt.Sprintf("%s cannot get Column\n", c.GetKind().String()))
	}
}

func (c Condition) GetValue() string {
	if KindIN(c.GetKind(), ConstKind) {
		return c.Value
	} else {
		panic(fmt.Sprintf("%s cannot get Column\n", c.GetKind().String()))
	}
}

func (c Condition) GetOperator() operator {
	if KindIN(c.GetKind(), OperatorKind) {
		return c.Operator
	} else {
		panic(fmt.Sprintf("%s cannot get Column\n", c.GetKind().String()))
	}
}

func (c Condition) GetConditions() []Condition {
	if KindIN(c.GetKind(), OperatorKind) {
		return c.Conditions
	} else {
		panic(fmt.Sprintf("%s cannot get Column\n", c.GetKind().String()))
	}
}

func Var(column string, p predicate) Condition {
	return Condition{
		Kind:      VarKind,
		Column:    column,
		Predicate: p,
	}
}

func Const(column string, p predicate, value string) Condition {
	return Condition{
		Kind:      ConstKind,
		Column:    column,
		Predicate: p,
		Value:     value,
	}
}

func Operator(o operator, conditions ...Condition) Condition {
	return Condition{
		Kind:       OperatorKind,
		Operator:   o,
		Conditions: conditions,
	}
}

func Conditions(conditions ...Condition) []Condition {
	return conditions
}

type filter struct {
	Param     string
	ParamType paramType
	Condition Condition
	Comments  string
}

func Filter(param string, typ paramType, condition Condition) filter {
	return filter{
		Param:     param,
		ParamType: typ,
		Condition: condition,
	}
}

func (f filter) Comment(c string) filter {
	fp := &f
	fp.Comments = c
	return f
}

type restrict struct {
	Condition Condition
}

func Restrict(condition Condition) restrict {
	return restrict{
		Condition: condition,
	}
}
