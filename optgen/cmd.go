package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/go/packages"
)

func main() {
	err := runMain()
	if err != nil {
		panic(err)
	}
}

type specValue struct {
	typ     *ast.TypeSpec
	comment *ast.Comment
}

const directive = "+troptgen"
const directiveLine = "// +troptgen"

func runMain() error {

	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedImports | packages.NeedDeps | packages.NeedTypes |
			packages.NeedSyntax | packages.NeedTypesInfo,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.ParseComments)
		},
	}

	pkgs, err := packages.Load(&cfg, "./")
	if err != nil {
		return fmt.Errorf("cannot load %q: %w", "./", err)
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("packages.Load returned %d packages, not 1", len(pkgs))
	}

	pkg := pkgs[0]

	specs := make(map[string][]specValue)

	for _, syntax := range pkg.Syntax {
		s := getTaggedComments(syntax, directive)
		for _, sp := range s {
			filename := pkg.Fset.Position(syntax.Package).Filename
			fn, ok := specs[filename]
			if !ok {
				fn = []specValue{}
				specs[filename] = fn
			}
			specs[filename] = append(specs[filename], sp)
		}
	}

	for filename, fnspecs := range specs {
		dir, file := filepath.Split(filename)
		newfile := filepath.Join(dir, fmt.Sprintf("%s_gen%s", strings.TrimSuffix(file, filepath.Ext(filename)), filepath.Ext(filename)))

		f := jen.NewFilePath(pkg.PkgPath)
		f.PackageComment("Code generated by generator, DO NOT EDIT.")

		optionsbuilder := map[string][]jen.Code{}

		for _, stype := range fnspecs {
			obj := pkg.Types.Scope().Lookup(stype.typ.Name.Name)
			if obj == nil {
				continue
			}

			// root, get, set, delete, refresh
			directiveCmd := strings.TrimSpace(strings.TrimPrefix(stype.comment.Text, directiveLine))
			UCDirectiveCMD := cases.Title(language.Und).String(directiveCmd)
			UCDirectiveCMDOptional := cases.Title(language.Und).String(directiveCmd)
			if directiveCmd == "root" {
				UCDirectiveCMDOptional = ""
			}

			// only named interface types are supported
			namedType, ok := obj.Type().(*types.Named)
			if !ok {
				return fmt.Errorf("only interface types are supported: %s", obj.String())
			}

			interfaceType, ok := namedType.Underlying().(*types.Interface)
			if !ok {
				return fmt.Errorf("only interface types are supported: %s", obj.String())
			}

			_, optbok := optionsbuilder[directiveCmd]
			if !optbok {
				optionsbuilder[directiveCmd] = []jen.Code{}
				optionsbuilder[directiveCmd] = append(optionsbuilder[directiveCmd],
					jen.Type().Id(fmt.Sprintf("%sOptionBuilder", UCDirectiveCMD)).
						Add(FromTypeParams(namedType.TypeParams())).
						Struct(
							jen.Id(fmt.Sprintf("%sOptionBuilderBase", UCDirectiveCMD)),
						))
				optionsbuilder[directiveCmd] = append(optionsbuilder[directiveCmd],
					jen.Func().Id(fmt.Sprintf("%sOpt", UCDirectiveCMD)).
						Add(FromTypeParams(namedType.TypeParams())).
						Params().
						Id(fmt.Sprintf("*%sOptionBuilder", UCDirectiveCMD)).Add(CallFromTypeParams(namedType.TypeParams())).
						Block(
							jen.Return(
								jen.Id(fmt.Sprintf("&%sOptionBuilder", UCDirectiveCMD)).Add(CallFromTypeParams(namedType.TypeParams())).Values(jen.Dict{}),
							),
						),
				)
			}

			for i := 0; i < interfaceType.NumMethods(); i++ {
				// generate a "With" function for each interface method
				fsig := interfaceType.Method(i).Type().(*types.Signature)
				methodName := fmt.Sprintf("With%s%s", UCDirectiveCMDOptional, strings.TrimPrefix(interfaceType.Method(i).Name(), "Opt"))
				f.Func().Id(methodName).
					Add(FromTypeParams(namedType.TypeParams())).
					Add(FromParams(fsig.Params(), fsig.Variadic())).
					Qual("github.com/RangelReale/trcache", fmt.Sprintf("%sOption", UCDirectiveCMD)).
					Block(
						jen.Return(
							jen.Qual("github.com/RangelReale/trcache", fmt.Sprintf("%sOptionFunc", UCDirectiveCMD)).Call(
								jen.Func().
									Params(jen.Id("o").Id("any")).
									Bool().
									BlockFunc(func(g *jen.Group) {
										g.Switch(jen.Id("opt").Op(":=").Id("o.(type)").Block(
											jen.Case(QualFromType(namedType)).Block(
												jen.Id("opt").Dot(interfaceType.Method(i).Name()).Add(CallFromParams(fsig.Params(), fsig.Variadic())),
											),
											jen.Return(jen.True()),
										))

										g.Return(jen.False())
									})),
						),
					)

				// generate an "OptionsBuilder" method for each interface method
				optionsbuilder[directiveCmd] = append(optionsbuilder[directiveCmd],
					jen.Func().
						Params(jen.Id("ob").Id(fmt.Sprintf("*%sOptionBuilder", UCDirectiveCMD)).Add(CallFromTypeParams(namedType.TypeParams()))).
						Id(methodName).
						Add(FromParams(fsig.Params(), fsig.Variadic())).
						Id(fmt.Sprintf("*%sOptionBuilder", UCDirectiveCMD)).Add(CallFromTypeParams(namedType.TypeParams())).
						Block(
							jen.Id("ob").Dot("AppendOptions").Call(
								jen.Id(methodName).Add(CallFromTypeParams(namedType.TypeParams()).
									Add(CallFromParams(fsig.Params(), fsig.Variadic())),
								),
							),
							jen.Return(
								jen.Id("ob"),
							),
						),
				)

			}

		}

		// generate an options builder for each interface method
		for _, d := range []string{"root", "get", "set", "delete", "refresh"} {
			ob, ok := optionsbuilder[d]
			if !ok {
				continue
			}
			for _, obi := range ob {
				f.Add(obi)
			}
		}

		err = f.Save(newfile)
		if err != nil {
			return err
		}
	}

	return nil
}

// getTaggedComments walks the AST and returns types which have directive comment
// returns a map of TypeSpec to directive
func getTaggedComments(pkg ast.Node, directive string) []specValue {
	var specs []specValue

	ast.Inspect(pkg, func(n ast.Node) bool {
		g, ok := n.(*ast.GenDecl)

		// is it a type?
		// http://golang.org/pkg/go/ast/#GenDecl
		if !ok || g.Tok != token.TYPE {
			// never mind, move on
			return true
		}

		if g.Lparen == 0 {
			// not parenthesized, copy GenDecl.Doc into TypeSpec.Doc
			g.Specs[0].(*ast.TypeSpec).Doc = g.Doc
		}

		for _, s := range g.Specs {
			t := s.(*ast.TypeSpec)

			if c := findAnnotation(t.Doc, directive); c != nil {
				specs = append(specs, specValue{
					typ:     t,
					comment: c,
				})
			}
		}

		// no need to keep walking, we don't care about TypeSpec's children
		return false
	})

	return specs
}

// findDirective return the first line of a doc which contains a directive
// the directive and '//' are removed
func findAnnotation(doc *ast.CommentGroup, directive string) *ast.Comment {
	if doc == nil {
		return nil
	}

	// check lines of doc for directive
	for _, c := range doc.List {
		l := c.Text
		// does the line start with the directive?
		t := strings.TrimLeft(l, "/ ")
		if !strings.HasPrefix(t, directive) {
			continue
		}

		// remove the directive from the line
		t = strings.TrimPrefix(t, directive)

		// must be eof or followed by a space
		if len(t) > 0 && t[0] != ' ' {
			continue
		}

		return c
	}

	return nil
}