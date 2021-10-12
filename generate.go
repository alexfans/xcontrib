package xcontrib

import (
	"embed"
	"entgo.io/ent/entc/gen"
	"errors"
	"fmt"
	"github.com/alexfans/xcontrib/utils"
	"go.uber.org/multierr"
	"os"
	"path/filepath"
	"strings"
)

var (
	//go:embed helper/*
	helperFS embed.FS
)

func generateHelper(baseDir string, basePackage string) {
	baseDir = filepath.Join(baseDir, "helper")
	helperWalk("helper", baseDir, basePackage)
}

func helperWalk(root string, path string, basePackage string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return
	}

	files, err := helperFS.ReadDir(root)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		if f.IsDir() {
			helperWalk(filepath.Join(root, f.Name()), filepath.Join(path, f.Name()), basePackage)
		} else {
			writeHelper(filepath.Join(root, f.Name()), filepath.Join(path, f.Name()), basePackage)
		}
	}
}

func writeHelper(root string, base string, basePackage string) error {
	f, err := os.OpenFile(base, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	b, err := helperFS.ReadFile(root)
	if err != nil {
		fmt.Println(" ", err)
		return err
	}
	content := string(b)
	content = strings.Replace(content, "xcontrib/rename", basePackage, -1)
	_, err = f.WriteString(content)
	return err
}

func Generate(g *gen.Graph, assign string) error {
	// 加载
	adapter, err := LoadAdapter(g)
	if err != nil {
		return fmt.Errorf("xcontrib: failed parsing ent graph: %w", err)
	}

	// 检查
	var errs error
	for _, schema := range g.Schemas {
		name := schema.Name
		_, err := adapter.GetDescriptor(name)
		if err != nil && !errors.Is(err, ErrSchemaSkipped) {
			errs = multierr.Append(errs, err)
		}
	}
	if errs != nil {
		return fmt.Errorf("xcontrib: failed parsing some schemas: %w", errs)
	}

	// 写入
	baseDir := filepath.Join(g.Config.Target, "..")
	generateHelper(baseDir, filepath.Dir(g.Config.Package))

	initTemplates()
	descriptors := adapter.AllDescriptor()
	for _, descriptor := range descriptors {
		if assign != "" {
			if descriptor.Name != assign {
				continue
			}
		}
		for _, printer := range descriptor.Printers {
			generate(printer, baseDir, descriptor)
		}
	}
	return nil
}

func generate(printer *Printer, baseDir string, descriptor *Descriptor) {
	fullPath := printer.GetFilePath(baseDir, utils.Snake(descriptor.Name))
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return
	}
	f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	err = printer.Exec(descriptor, printer.Name(), f)
	if err != nil {
		fmt.Println(err)
	}
}
