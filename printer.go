package xcontrib

import (
	"bytes"
	"entgo.io/ent/entc/gen"
	"go/format"
	"golang.org/x/tools/imports"
	"io"
	"io/ioutil"
	"path/filepath"
)

const (
	NoSrcDir = "others"
)

type PointerExecFn func(descriptor *Descriptor, tmpl *gen.Template, w io.WriteCloser) error

type FilePathFn func(name string, suffix string) string

type Printer struct {
	name        string
	packageName string
	fileSuffix  string
	isSrc       bool
	optionCode  int
	filePathFn  FilePathFn
	execFn      PointerExecFn
}

func (p *Printer) Name() string {
	return p.name
}

func (p *Printer) PackageName() string {
	return p.packageName
}

func (p *Printer) FileSuffix() string {
	return p.fileSuffix
}

func (p *Printer) IsSrc() bool {
	return p.isSrc
}

func (p *Printer) OptionCode() int {
	return p.optionCode
}

func (p *Printer) ExecFn() PointerExecFn {
	return p.execFn
}

func (p *Printer) GetFilePath(base string, name string) string {
	var fn FilePathFn

	if p.IsSrc() {
		base = filepath.Join(base, p.PackageName())
	} else {
		base = filepath.Join(base, NoSrcDir, p.PackageName())
	}
	if p.filePathFn != nil {
		fn = p.filePathFn
	} else {
		fn = func(name string, suffix string) string {
			return name + suffix
		}
	}

	path := filepath.Join(base, fn(name, p.fileSuffix))

	return path
}

func (p *Printer) Exec(descriptor *Descriptor, tmpl string, w io.ReadWriteCloser) error {
	descriptor.Package = p.PackageName()
	buf := new(bytes.Buffer)
	err := templates.ExecuteTemplate(buf, tmpl, descriptor)
	if err != nil {
		return err
	}
	bs, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	if p.IsSrc() {
		bs, err = formatGOCode(bs)
		if err != nil {
			return err
		}
	}
	_, err = w.Write(bs)
	return err
}

func formatGOCode(bs []byte) ([]byte, error) {
	var err error
	bs, err = format.Source(bs)
	if err != nil {
		return nil, err
	}
	bs, err = imports.Process("", bs, nil)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (p *Printer) MatchOptionCode(option int) bool {
	if p.optionCode == BaseService {
		return true
	}
	if option&p.optionCode > 0 {
		return true
	}
	return false
}

var (
	allPointer = []*Printer{
		{
			name:        "vm",
			packageName: "vm",
			fileSuffix:  ".go",
			isSrc:       true,
			optionCode:  BaseService,
		},
		{
			name:        "dao",
			packageName: "dao",
			fileSuffix:  ".go",
			isSrc:       true,
			optionCode:  BaseService,
		},
		{
			name:        "service",
			packageName: "service",
			fileSuffix:  ".go",
			isSrc:       true,
			optionCode:  BaseService,
		},
		{
			name:        "echo",
			packageName: "echohandle",
			fileSuffix:  ".go",
			optionCode:  Echo,
			isSrc:       true,
		},
		{
			name:        "tcp",
			packageName: "tcphandle",
			fileSuffix:  ".go",
			isSrc:       true,
			optionCode:  Tcp,
		},
		{
			name:        "udp",
			packageName: "udphandle",
			fileSuffix:  ".go",
			isSrc:       true,
			optionCode:  Udp,
		},
		{
			name:        "websocket",
			packageName: "websockethandle",
			fileSuffix:  ".go",
			isSrc:       true,
			optionCode:  WebSocket,
		},
		{
			name:        "vue",
			packageName: "vue",
			fileSuffix:  ".vue",
			isSrc:       false,
			optionCode:  Vue,
			filePathFn: func(name string, suffix string) string {
				return filepath.Join(name, "index"+suffix)
			},
		},
		{
			name:        "vuerequest",
			packageName: "api",
			fileSuffix:  ".js",
			isSrc:       false,
			optionCode:  VueRequest,
		},
		{
			name:        "doc",
			packageName: "doc",
			fileSuffix:  ".md",
			isSrc:       false,
			optionCode:  Doc,
		},
		{
			name:        "postman",
			packageName: "postman",
			fileSuffix:  ".json",
			isSrc:       false,
			optionCode:  PostMan,
		},
	}
)
