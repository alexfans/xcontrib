package xcontrib

import (
	"entgo.io/ent/entc/gen"
)

type Descriptor struct {
	Name         string     // *gen.Type.Name的冗余
	Package      string     // go语言的包名
	BasePackage  string     // 基本包名,如果demo/dao,则BasePackage中`demo`
	GenType      *gen.Type  // ent生成
	Comments     string     // 备注
	FilterList   []filter   // list,count的中,通过参数进行过滤
	RestrictList []restrict // 全局过滤,是常量
	SortList     []sort     //排序
	Printers     []*Printer // 打印的方式
}

type EdgeWrapper struct {
	*Descriptor
	Edge         *gen.Edge
	Comments     string     // 备注
	FilterList   []filter   // list,count的中,通过参数进行过滤
	RestrictList []restrict // 全局过滤,是常量
	SortList     []sort     //排序
}

func NewEdgeWrapper(d *Descriptor, edge *gen.Edge) *EdgeWrapper {
	w := &EdgeWrapper{
		Descriptor: d,
		Edge:       edge,
	}
	serAnnot, err := extractMessageAnnotation(d.Name+"."+edge.Name, edge.Annotations)
	if err == nil {
		w.Comments = serAnnot.Comments
		w.FilterList = serAnnot.FilterList
		w.RestrictList = serAnnot.RestrictList
		w.SortList = serAnnot.SortList
	}

	return w
}

func (d *Descriptor) CreateFields() []*gen.Field {
	// TODO
	fields := make([]*gen.Field, 0)
	for _, field := range d.GenType.Fields {
		fieldName := field.Name
		if fieldName != "created_at" && fieldName != "updated_at" && fieldName != "is_deleted" && fieldName != "deleted_at" {
			fields = append(fields, field)
		}
	}
	return fields
}

func (d *Descriptor) UpdateFields() []*gen.Field {
	// TODO
	fields := make([]*gen.Field, 0)
	for _, field := range d.GenType.Fields {
		fieldName := field.Name
		if fieldName != "created_at" && fieldName != "updated_at" && fieldName != "is_deleted" && fieldName != "deleted_at" {
			fields = append(fields, field)
		}
	}
	return fields
}

func (d *Descriptor) ViewFields() []*gen.Field {
	// TODO
	fields := make([]*gen.Field, 0)
	fields = append(fields, d.GenType.ID)
	for _, field := range d.GenType.Fields {
		fieldName := field.Name
		if fieldName != "is_deleted" && fieldName != "deleted_at" {
			fields = append(fields, field)
		}
	}
	return fields
}

func (d *Descriptor) CanDelete() bool {
	for _, field := range d.GenType.Fields {
		fieldName := field.Name
		if fieldName == "is_deleted" {
			return true
		}
		if fieldName == "deleted_at" {
			return true
		}
	}
	return false
}

func (d *Descriptor) CanUpdate() bool {
	for _, field := range d.GenType.Fields {
		fieldName := field.Name
		if fieldName == "updated_at" {
			return true
		}
	}
	return false
}

func (d *Descriptor) ToEdges() []*EdgeWrapper {
	edges := make([]*EdgeWrapper, 0)

	for _, edge := range d.GenType.Edges {
		if edge.Inverse == "" {
			edges = append(edges, NewEdgeWrapper(d, edge))
		}
	}
	return edges
}

func (d *Descriptor) FromEdges() []*EdgeWrapper {
	edges := make([]*EdgeWrapper, 0)

	for _, edge := range d.GenType.Edges {
		if edge.Inverse != "" {
			edges = append(edges, NewEdgeWrapper(d, edge))
		}
	}
	return edges
}

func (d *Descriptor) Edges() []*EdgeWrapper {
	edges := make([]*EdgeWrapper, 0)

	for _, edge := range d.GenType.Edges {
		edges = append(edges, NewEdgeWrapper(d, edge))
	}
	return edges
}

type ConditionWrapper struct {
	Condition
	GenName string
}
