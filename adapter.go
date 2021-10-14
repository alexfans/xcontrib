package xcontrib

import (
	"entgo.io/ent/entc/gen"
	"errors"
	"fmt"
	"path/filepath"
)

var (
	ErrSchemaSkipped = errors.New("xcontrib: schema not annotated with Generate=true")
)

// Adapter facilitates the transformation of ent gen.Type to desc.FileDescriptors
type Adapter struct {
	graph       *gen.Graph
	descriptors map[string]*Descriptor
	errors      map[string]error
}

// LoadAdapter takes a *gen.Graph and parses it into protobuf file descriptors
func LoadAdapter(graph *gen.Graph) (*Adapter, error) {
	a := &Adapter{
		graph:       graph,
		descriptors: make(map[string]*Descriptor),
		errors:      make(map[string]error),
	}
	if err := a.parse(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Adapter) GetDescriptor(schemaName string) (*Descriptor, error) {
	if err, ok := a.errors[schemaName]; ok {
		return nil, err
	}
	d, ok := a.descriptors[schemaName]
	if !ok {
		return nil, fmt.Errorf("xcontrib: could not find file descriptor for schema %s", schemaName)
	}
	return d, nil
}

func (a *Adapter) parse() error {
	for _, genType := range a.graph.Nodes {
		descriptor, err := a.toDescriptor(genType)
		if err != nil {
			a.errors[genType.Name] = err
			continue
		}
		a.descriptors[genType.Name] = descriptor
	}
	return nil
}

func (a *Adapter) toDescriptor(genType *gen.Type) (*Descriptor, error) {
	serAnnot, err := extractMessageAnnotation(genType.Name, genType.Annotations)
	if err != nil {
		return nil, ErrSchemaSkipped
	}
	msg := &Descriptor{
		Name:         genType.Name,
		BasePackage:  filepath.Dir(genType.Config.Package),
		GenType:      genType,
		Comments:     serAnnot.Comments,
		Printers:     loadDescriptorPoint(serAnnot),
		Methods:      serAnnot.Methods,
		IsHardDelete: serAnnot.IsHardDelete,
		FilterList:   serAnnot.FilterList,
		RestrictList: serAnnot.RestrictList,
		SortList:     serAnnot.SortList,
	}

	return msg, nil
}

func (a *Adapter) AllDescriptor() map[string]*Descriptor {
	return a.descriptors
}

func loadDescriptorPoint(serAnnot *service) []*Printer {
	printers := make([]*Printer, 0, 8)
	for _, printer := range allPointer {
		if printer.MatchOptionCode(serAnnot.Options) {
			printers = append(printers, printer)
		}
	}
	return printers
}
