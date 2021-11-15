package servicehelper

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	skip = false
)

var (
	NoParamErr = errors.New("no found param value")
)

const (
	maxLimit      = 500
	defaultLimit  = 20
	defaultOffset = 0
)

const (
	separated  = ","
	timeFormat = "2006-01-02 15:04:05"
)

type Param struct {
	values url.Values
	orders []string
	limit  int
	offset int
}

func NewParam(values url.Values) *Param {
	param := &Param{
		values: values,
		limit:  defaultLimit,
		offset: defaultOffset,
	}
	param.setPage()
	return param
}

func (p *Param) setPage() {
	limit, page, offset := p.Get("limit"), p.Get("page"), p.Get("offset")
	var err error
	if limit != "" {
		if p.limit, err = strconv.Atoi(limit); err != nil {
			p.limit = defaultLimit
		}
		if p.limit < 0 || p.limit > maxLimit {
			p.limit = defaultLimit
		}
	} else {
		p.limit = defaultLimit
	}

	if offset != "" {
		if p.offset, err = strconv.Atoi(offset); err != nil {
			p.offset = defaultOffset
		}
		if p.offset < 0 {
			p.offset = defaultOffset
		}
	} else {
		if _page, err := strconv.Atoi(page); err != nil {
			p.offset = defaultOffset
		} else {
			p.offset = (_page - 1) * p.limit
		}
		if p.offset < 0 {
			p.offset = defaultOffset
		}
	}
}

func (p Param) setOrders() {
	orders := p.Get("orders")
	if orders != "" {
		p.orders = strings.Split(orders, ",")
	}
}

func (p Param) Offset() int {
	return p.offset
}

func (p Param) Limit() int {
	return p.limit
}

func (p Param) Orders() []string {
	return p.orders
}

func (p Param) Has(name string) bool {
	_, has := p.values[name]
	return has
}

func (p Param) Get(name string) string {
	return p.values.Get(name)
}

func (p Param) GetInt(name string) (int, error) {
	s := p.Get(name)
	if name == "" {
		return 0, NoParamErr
	}
	return parseInt(s)
}

func (p Param) GetInts(name string) ([]int, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]int, 0, len(ss))
	for _, s := range ss {
		v, err := parseInt(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetInt8(name string) (int8, error) {
	s := p.Get(name)
	if name == "" {
		return 0, NoParamErr
	}
	return parseInt8(s)
}

func (p Param) GetInt8s(name string) ([]int8, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]int8, 0, len(ss))
	for _, s := range ss {
		v, err := parseInt8(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetInt16(name string) (int16, error) {
	s := p.Get(name)
	if name == "" {
		return 0, NoParamErr
	}
	return parseInt16(s)
}

func (p Param) GetInt16s(name string) ([]int16, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]int16, 0, len(ss))
	for _, s := range ss {
		v, err := parseInt16(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetInt32(name string) (int32, error) {
	s := p.Get(name)
	if name == "" {
		return 0, NoParamErr
	}
	return parseInt32(s)
}

func (p Param) GetInt32s(name string) ([]int32, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]int32, 0, len(ss))
	for _, s := range ss {
		v, err := parseInt32(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetInt64(name string) (int64, error) {
	s := p.Get(name)
	if name == "" {
		return 0, NoParamErr
	}
	return parseInt64(s)
}

func (p Param) GetInt64s(name string) ([]int64, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]int64, 0, len(ss))
	for _, s := range ss {
		v, err := parseInt64(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetFloat32(name string) (float32, error) {
	s := p.Get(name)
	if name == "" {
		return 0.0, NoParamErr
	}
	return parseFloat32(s)
}

func (p Param) GetFloat32s(name string) ([]float32, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]float32, 0, len(ss))
	for _, s := range ss {
		v, err := parseFloat32(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetFloat64(name string) (float64, error) {
	s := p.Get(name)
	if name == "" {
		return 0.0, NoParamErr
	}
	return parseFloat64(s)
}

func (p Param) GetFloat64s(name string) ([]float64, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]float64, 0, len(ss))
	for _, s := range ss {
		v, err := parseFloat64(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetTime(name string) (time.Time, error) {
	s := p.Get(name)
	if name == "" {
		return time.Now(), NoParamErr
	}
	return parseTime(s)
}

func (p Param) GetTimes(name string) ([]time.Time, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]time.Time, 0, len(ss))
	for _, s := range ss {
		v, err := parseTime(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetString(name string) (string, error) {
	s := p.Get(name)
	if name == "" {
		return "", NoParamErr
	}
	return parseString(s)
}

func (p Param) GetStrings(name string) ([]string, error) {
	s := p.Get(name)
	if name == "" {
		return nil, NoParamErr
	}
	ss := strings.Split(s, separated)
	vs := make([]string, 0, len(ss))
	for _, s := range ss {
		v, err := parseString(s)
		if err != nil {
			if skip {
				continue
			} else {
				return nil, err
			}
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (p Param) GetAny(name string) (string, error) {
	if !p.Has(name) {
		return "", NoParamErr
	}
	s := p.Get(name)
	return parseString(s)
}

func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int(i), err
}

func parseInt8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int8(i), err
}

func parseInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int16(i), err
}

func parseInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int32(i), err
}

func parseInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int64(i), err
}

func parseFloat32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 64)
	return float32(i), err
}

func parseFloat64(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	return float64(i), err
}

func parseTime(s string) (time.Time, error) {
	return time.Parse(timeFormat, s)
}

func parseString(s string) (string, error) {
	return s, nil
}
