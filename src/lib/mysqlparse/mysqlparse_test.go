package mysqlparse

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	sql := "select * from t1 where a=1 and b=2 or c=3 or d is null"
	fmt.Println(sql)
	Parse(sql)
}
