package generator

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/packages"
)

var errIsNotSlice = errors.New("it is not slice")

// Named Capture Group support since Go 1.22.
// When we remove support for Go versions below 1.22, we will be able to use code like
//
// var (
//
//	importPackageMask             = regexp.MustCompile(`(?<pkgName>[\w_\-\.\d]+)(\/v\d+)?$`)
//	importPackageMaskPkgNameIndex = importPackageMask.SubexpIndex("pkgName")
//
// )

var importPackageMask = regexp.MustCompile(`([\w_\-\.\d]+)(\/v\d+)?$`)

const importPackageMaskPkgNameIndex = 1

func formatComment(comment string) string {
	if comment == "" {
		return ""
	}

	buf := bytes.NewBuffer(nil)

	lines := strings.Split(comment, "\n")
	for i := range lines {
		// Last line contains an empty string.
		if lines[i] == "" && i == len(lines)-1 {
			continue
		}

		if i != 0 {
			buf.WriteString("\n")
		}

		buf.WriteString("// ")
		buf.WriteString(lines[i])
	}

	return buf.String()
}

func findStructTypeParamsAndFields(packages map[string]*ast.Package, typeName string) (*ast.File, []*ast.Field, []*ast.Field, bool) { //nolint:lll
	for _, pkgObj := range packages {
		for _, fileObj := range pkgObj.Files {
			for _, decl := range fileObj.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}

				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec) //nolint:varnamelen
					if !ok {
						continue
					}

					if typeSpec.Name.Name != typeName {
						continue
					}

					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok {
						continue
					}

					return fileObj, extractFields(typeSpec.TypeParams), extractFields(structType.Fields), true
				}
			}
		}
	}

	return nil, nil, nil, false
}

func extractFields(fl *ast.FieldList) []*ast.Field {
	if fl == nil {
		return nil
	}

	return fl.List
}

func isPublic(fieldName string) bool {
	char, _ := utf8.DecodeRuneInString(fieldName)

	return char != utf8.RuneError && unicode.IsUpper(char)
}

func checkDefaultValue(fieldType string, tag string) error {
	var err error
	switch fieldType {
	case "int", "int8", "int16", "int32", "int64":
		_, err = strconv.ParseInt(tag, 10, 64)

	case "uint", "uint8", "uint16", "uint32", "uint64":
		_, err = strconv.ParseUint(tag, 10, 64)

	case "float32", "float64":
		_, err = strconv.ParseFloat(tag, 64)

	case "time.Duration":
		_, err = time.ParseDuration(tag)

	case "bool":
		if tag != "true" && tag != "false" {
			return fmt.Errorf("bool type only supports true/false")
		}

	case "string":
		// As is.

	default:
		return fmt.Errorf("unsupported type `%s`", fieldType)
	}

	if err != nil {
		return fmt.Errorf("bad default value %w %s", err, tag)
	}

	return nil
}

func normalizeTypeName(typeName string) string {
	if idx := strings.LastIndex(typeName, "."); idx > -1 {
		typeName = typeName[idx+1:]
	}

	return strings.TrimPrefix(strings.TrimPrefix(typeName, "[]"), "*")
}

// extractSliceElemType will find the element type for given slice.
func extractSliceElemType(
	workDir string,
	fset *token.FileSet,
	curFile *ast.File,
	expr ast.Expr,
) (string, error) {
	switch expr := expr.(type) {
	default:
		return "", errIsNotSlice
	case *ast.SelectorExpr:
		// Extract package name and type name
		pkgIdent, ok := expr.X.(*ast.Ident)
		if !ok {
			return "", errors.New("unsupported selector")
		}

		pkgName := pkgIdent.Name
		typeName := expr.Sel.Name

		importPath, alias := findImportPath(curFile.Imports, pkgName)
		if importPath == "" {
			return "", errors.New("import path not found")
		}

		pkg, err := loadPkg(fset, importPath, workDir)
		if err != nil {
			return "", errors.New("unable to load package")
		}

		lookupType := pkg.Types.Scope().Lookup(typeName)
		if expr, ok := lookupType.(*types.TypeName); ok { //nolint:nestif
			if expr, ok := expr.Type().(*types.Named); ok {
				if expr, ok := expr.Underlying().(*types.Slice); ok {
					switch expr := expr.Elem().(type) {
					case *types.Named:
						if importPath == expr.Obj().Pkg().Path() {
							return alias + "." + expr.Obj().Name(), nil
						}

						return expr.Obj().Pkg().Name() + "." + expr.Obj().Name(), nil
					case *types.Basic:
						return expr.Name(), nil
					}
				}

				return "", errIsNotSlice
			}
		}

		return "", errors.New("lookup type not found")
	case *ast.ArrayType:
		return types.ExprString(expr.Elt), nil
	case *ast.Ident:
		if expr.Obj == nil {
			return "", errIsNotSlice
		}

		switch expr := expr.Obj.Decl.(type) {
		default:
			return "", errors.New("unsupported ident expression")
		case *ast.TypeSpec:
			return extractSliceElemType(workDir, fset, curFile, expr.Type)
		}
	}
}

// findImportPath return full package name and alias if presented.
func findImportPath(imports []*ast.ImportSpec, pkgName string) (string, string) {
	for _, imp := range imports {
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}

		// If the import has an alias, check that
		if imp.Name != nil {
			aliasName, err := strconv.Unquote(imp.Name.Name)
			if err == nil && aliasName == pkgName {
				return importPath, aliasName
			}
		} else {
			// Otherwise, check if the base package name matches
			match := importPackageMask.FindStringSubmatch(importPath)
			if len(match) > 0 && match[importPackageMaskPkgNameIndex] == pkgName {
				return importPath, pkgName
			}
		}
	}

	return "", ""
}

// loadPkg loads a package by full import path.
func loadPkg(fset *token.FileSet, pkgName, dirPath string) (*packages.Package, error) {
	cfg := &packages.Config{ //nolint:exhaustruct
		Mode: packages.NeedName |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypesInfo,
		Dir:  dirPath,
		Fset: fset,
	}

	pkgs, err := packages.Load(cfg, pkgName)
	if err != nil {
		return nil, fmt.Errorf("loading package: %w", err)
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found")
	}

	return pkgs[0], nil
}
