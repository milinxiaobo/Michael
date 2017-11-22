package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xwb1989/sqlparser"
)

func Test_MySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:linxiaobo@tcp(10.5.232.78:3306)/linxiaobo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test")
	// db.Exec("show databases")
	// db.Exec("insert into test values (1, 1)")
	// tx, err := db.Begin()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// tx.Exec("select * from test")
	// tx.Commit()
	db.Exec("select * from test")
}

func Test_1(t *testing.T) {
	sql := `select ss from qq where a = 1 or b=3 and c = 4`
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		t.Error(err)
		return
		// Do something with the err
	}
	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		x := stmt
		if x.Where != nil {
			b := sqlparser.TrackedBuffer{}
			b.Buffer = new(bytes.Buffer)
			x.Where.Expr.Format(&b)
			t.Log(b.String())
			switch ss := x.Where.Expr.(type) {
			case *sqlparser.AndExpr:
				b := sqlparser.TrackedBuffer{}
				b.Buffer = new(bytes.Buffer)
				ss.Left.Format(&b)
				t.Log(b)
			}
		}
		b := sqlparser.TrackedBuffer{}
		b.Buffer = new(bytes.Buffer)
		// x.Where.Expr.Format(&b)
		// x.Format(&b)
		x.SelectExprs.Format(&b)
		t.Log(b.String())
		// x.WalkSubtree(visit)
	case *sqlparser.Insert:
		break
	}

}

func Test_2(t *testing.T) {
	sql := `select ss from qq where a = 1 or b=3 and c = 4`
	tokens := sqlparser.NewTokenizer(strings.NewReader(sql))
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			// t.Error(err)
			break
		}
		switch stmt := stmt.(type) {
		case *sqlparser.Select:
			b := sqlparser.TrackedBuffer{}
			b.Buffer = new(bytes.Buffer)
			stmt.Where.Format(&b)
			t.Log(b)
		}
		// Do something with stmt or err.
	}
}
