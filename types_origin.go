package gotype

import (
	"go/ast"
)

func newTypeOrigin(v Type, ori ast.Node, info *info, doc, comment *ast.CommentGroup) Type {
	pkg := info.PkgPath
	goroot := info.Goroot
	if p := v.PkgPath(); p != "" {
		pkg = p
	}
	return &typeOrigin{
		Type:    v,
		pkgPath: pkg,
		goroot:  goroot || v.IsGoroot(),
		ori:     ori,
		doc:     doc,
		comment: comment,
	}
}

type typeOrigin struct {
	Type
	ori     ast.Node
	pkgPath string
	goroot  bool
	doc     *ast.CommentGroup
	comment *ast.CommentGroup
}

func (t *typeOrigin) Origin() ast.Node {
	return t.ori
}

func (t *typeOrigin) PkgPath() string {
	return t.pkgPath
}

func (t *typeOrigin) IsGoroot() bool {
	return t.goroot
}

func (t *typeOrigin) Doc() *ast.CommentGroup {
	return t.doc
}

func (t *typeOrigin) Comment() *ast.CommentGroup {
	return t.comment
}
