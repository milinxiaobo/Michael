package mysqlparse

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/xwb1989/sqlparser"
)

// Parse blabla
func Parse(sql string) {
	tokens := sqlparser.NewTokenizer(strings.NewReader(sql))
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			// t.Error(err)
			break
		}
		fmt.Println(parseStmt(stmt))
	}
}

func parseStmt(stmt sqlparser.Statement) string {
	switch stmt := (stmt).(type) {
	case *sqlparser.Select:
		return parseSelect(stmt)
	case *sqlparser.Insert:

	default:
		return ""
	}
	return ""
}

func parseSelect(stmt *sqlparser.Select) string {
	ret := " select "
	ret += getStr(stmt.SelectExprs)
	ret += " from " + getStr(stmt.From)
	if stmt.Where != nil {
		ret += " where "
		ret += parseExpr(stmt.Where.Expr)
	}
	if stmt.GroupBy != nil {

	}
	return ret
}

func visit(node sqlparser.SQLNode) (kontinue bool, err error) {
	return false, nil
}

func parseExpr(expr sqlparser.Expr) string {
	expr.WalkSubtree(visit)
	switch expr := expr.(type) {
	case *sqlparser.AndExpr:
		return parseExpr(expr.Left) + " and " + parseExpr(expr.Right)
	case *sqlparser.OrExpr:
		fmt.Println("OrExpr")
		return parseExpr(expr.Left) + " or " + parseExpr(expr.Right)
	case *sqlparser.NotExpr:
		fmt.Println("NotExpr")
	case *sqlparser.ParenExpr:
		fmt.Println("ParenExpr")
	case *sqlparser.ComparisonExpr:
		fmt.Println("ComparisonExpr")
		return parseExpr(expr.Left) + expr.Operator + parseExpr(expr.Right)
	case *sqlparser.RangeCond:
		fmt.Println("RangeCond")
	case *sqlparser.IsExpr:
		fmt.Println("IsExpr")
		return getStr(expr)
	case *sqlparser.ExistsExpr:
		fmt.Println("ExistsExpr")
	case *sqlparser.SQLVal:
		fmt.Println("SQLVal")
		return "?"
	case *sqlparser.NullVal:
		fmt.Println("NullVal")
	case sqlparser.BoolVal:
		fmt.Println("BoolVal")
	case *sqlparser.ColName:
		fmt.Println("ColName")
		return expr.Name.String()
	case sqlparser.ValTuple:
		fmt.Println("ValTuple")
	case *sqlparser.Subquery:
		fmt.Println("Subquery")
	case sqlparser.ListArg:
		fmt.Println("ListArg")
	case *sqlparser.BinaryExpr:
		fmt.Println("BinaryExpr")
	case *sqlparser.UnaryExpr:
		fmt.Println("UnaryExpr")
	case *sqlparser.IntervalExpr:
		fmt.Println("IntervalExpr")
	case *sqlparser.CollateExpr:
		fmt.Println("CollateExpr")
	case *sqlparser.FuncExpr:
		fmt.Println("FuncExpr")
	case *sqlparser.CaseExpr:
		fmt.Println("CaseExpr")
	case *sqlparser.ValuesFuncExpr:
		fmt.Println("ValuesFuncExpr")
	case *sqlparser.ConvertExpr:
		fmt.Println("ConvertExpr")
	case *sqlparser.ConvertUsingExpr:
		fmt.Println("ConvertUsingExpr")
	case *sqlparser.MatchExpr:
		fmt.Println("MatchExpr")
	case *sqlparser.GroupConcatExpr:
		fmt.Println("GroupConcatExpr")
	case *sqlparser.Default:
		fmt.Println("Default")
	default:
	}
	return ""
}

func getStr(node sqlparser.SQLNode) string {
	buff := sqlparser.TrackedBuffer{}
	buff.Buffer = new(bytes.Buffer)
	node.Format(&buff)
	return buff.String()
}
