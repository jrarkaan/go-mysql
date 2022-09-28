package go_db

import (
	"context"
	"fmt"
	"testing"
)

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "select * from general_db_id.pr_m_subdepart pms;"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Result: ", rows)
	defer rows.Close()

}
