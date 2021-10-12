package xcontrib

const (
	Asc  order = "Asc"
	Desc order = "Desc"
)

type order string

func (s order) String() string {
	return string(s)
}

type sort struct {
	Param    string
	Order    order
	Columns  []string
	Comments string
}

func Sort(param string, o order, columns ...string) sort {
	return sort{
		Param:   param,
		Order:   o,
		Columns: columns,
	}
}

func (s sort) Comment(c string) sort {
	sp := &s
	sp.Comments = c
	return s
}
