package xcontrib

import (
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

const ServiceAnnotation = "ServiceAnnotation"

const (
	Echo       = 1 << 0
	Tcp        = 1 << 1
	Udp        = 1 << 2
	WebSocket  = 1 << 3
	Doc        = 1 << 4
	Vue        = 1 << 5
	VueRequest = 1 << 6
	PostMan    = 1 << 7
)
const (
	BaseService = 0xFFFFFFFF
)

type service struct {
	Options      int
	Comments     string
	FilterList   []filter // 通过参数来过来条件
	RestrictList []restrict
	SortList     []sort
}

func (service) Name() string {
	return ServiceAnnotation
}

func Service(options int) *service {
	return &service{
		Options:      options,
		FilterList:   make([]filter, 0),
		RestrictList: make([]restrict, 0),
		SortList:     make([]sort, 0),
	}
}

func (s *service) Comment(c string) *service {
	s.Comments = c
	return s
}

func (s *service) Filters(filters ...filter) *service {
	s.FilterList = filters
	return s
}

func (s *service) Restricts(restricts ...restrict) *service {
	s.RestrictList = restricts
	return s
}

func (s *service) Sorts(sorts ...sort) *service {
	s.SortList = sorts
	return s
}

func extractMessageAnnotation(name string, annotations gen.Annotations) (*service, error) {
	annot, ok := annotations[ServiceAnnotation]
	if !ok {
		return nil, fmt.Errorf("xcontrib: schema %q does not have an xcontrib.ServiceAnnotation annotation", name)
	}

	var out service
	err := mapstructure.Decode(annot, &out)
	if err != nil {
		return nil, fmt.Errorf("xcontrib: unable to decode xcontrib.ServiceAnnotation annotation for schema %q: %w",
			name, err)
	}

	return &out, nil
}
