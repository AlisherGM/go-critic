package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(&dupBranchBodyChecker{}, attrExperimental)
}

type dupBranchBodyChecker struct {
	checkerBase
}

func (c *dupBranchBodyChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects duplicated branch bodies inside conditional statements"
	d.Before = `
if cond {
	println("cond=true")
} else {
	println("cond=true")
}`
	d.After = `
if cond {
	println("cond=true")
} else {
	println("cond=false")
}`
}

func (c *dupBranchBodyChecker) VisitStmt(stmt ast.Stmt) {
	// TODO(quasilyte): extend to check switch statements as well.
	// Should be very careful with type switches.

	if stmt, ok := stmt.(*ast.IfStmt); ok {
		c.checkIf(stmt)
	}
}

func (c *dupBranchBodyChecker) checkIf(stmt *ast.IfStmt) {
	thenBody := stmt.Body
	elseBody, ok := stmt.Else.(*ast.BlockStmt)
	if ok && astequal.Stmt(thenBody, elseBody) {
		c.warnIf(stmt)
	}
}

func (c *dupBranchBodyChecker) warnIf(cause ast.Node) {
	c.ctx.Warn(cause, "both branches in if statement has same body")
}
