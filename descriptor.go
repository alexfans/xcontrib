package xcontrib

import (
	"entgo.io/ent/entc/gen"
)

type Descriptor struct {
	Name         string    // *gen.Type.Name的冗余
	Package      string    // go语言的包名
	BasePackage  string    // 基本包名,如果demo/dao,则BasePackage中`demo`
	GenType      *gen.Type // ent生成
	Methods      int
	IsHardDelete bool
	Comments     string     // 备注
	FilterList   []filter   // list,count的中,通过参数进行过滤
	RestrictList []restrict // 全局过滤,是常量
	SortList     []sort     //排序
	Printers     []*Printer // 打印的方式
}

func (d *Descriptor) CreateFields() []*gen.Field {
	fields := make([]*gen.Field, 0)
	for _, field := range d.GenType.Fields {
		if !IN(field.Name, "created_at", "updated_at", "is_deleted", "deleted_at") {
			fields = append(fields, field)
		}
	}
	return fields
}

func (d *Descriptor) UpdateFields() []*gen.Field {
	fields := make([]*gen.Field, 0)
	for _, field := range d.GenType.Fields {
		if !IN(field.Name, "created_at", "updated_at", "is_deleted", "deleted_at") {
			fields = append(fields, field)
		}
	}
	return fields
}

func (d *Descriptor) ViewFields() []*gen.Field {
	fields := make([]*gen.Field, 0)
	fields = append(fields, d.GenType.ID)
	for _, field := range d.GenType.Fields {
		if field.Sensitive() {
			continue
		}
		if !IN(field.Name, "is_deleted", "deleted_at") {
			fields = append(fields, field)
		}
	}
	return fields
}

func (d *Descriptor) CanCreate() bool {
	return d.Methods&MethodCreate > 0
}

func (d *Descriptor) CanDelete() bool {
	return d.Methods&MethodDelete > 0
}

func (d *Descriptor) CanUpdate() bool {
	return d.Methods&MethodUpdate > 0
}

func (d *Descriptor) CanOne() bool {
	return d.Methods&MethodOne > 0
}

func (d *Descriptor) CanList() bool {
	return d.Methods&MethodList > 0
}

func (d *Descriptor) CanCount() bool {
	return d.Methods&MethodCount > 0
}

func (d *Descriptor) HasDeleteField() bool {
	var at, is bool
	for _, field := range d.GenType.Fields {
		if field.Name == "deleted_at" {
			at = true
		}
		if field.Name == "is_deleted" {
			is = true
		}
	}
	return at && is
}

func (d *Descriptor) HasUpdateField() bool {
	for _, field := range d.GenType.Fields {
		if field.Name == "updated_at" {
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

type EdgeWrapper struct {
	*Descriptor
	Edge         *gen.Edge
	Methods      int
	Comments     string     // 备注
	FilterList   []filter   // list,count的中,通过参数进行过滤
	RestrictList []restrict // 全局过滤,是常量
	SortList     []sort     //排序
}

func NewEdgeWrapper(d *Descriptor, edge *gen.Edge) *EdgeWrapper {
	w := &EdgeWrapper{
		Descriptor: d,
		Edge:       edge,
		Methods:    MethodAll,
	}
	serAnnot, err := extractMessageAnnotation(d.Name+"."+edge.Name, edge.Annotations)
	if err == nil && serAnnot != nil {
		w.Comments = serAnnot.Comments
		w.FilterList = serAnnot.FilterList
		w.RestrictList = serAnnot.RestrictList
		w.SortList = serAnnot.SortList
	}

	return w
}

func (ew *EdgeWrapper) CanCreate() bool {
	return ew.Methods&MethodCreate > 0
}

func (ew *EdgeWrapper) CanUpdate() bool {
	return ew.Methods&MethodUpdate > 0
}

func (ew *EdgeWrapper) CanView() bool {
	return ew.Methods&MethodOne > 0 && ew.Methods&MethodList > 0
}
