package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/HT0323/go_api/models"
	"github.com/HT0323/go_api/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// DBへの接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subTest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		}, {
			testTitle: "subTest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)

			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("ID: get %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("ID: get %s but want %s\n", got.Contents, test.expected.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("ID: get %s but want %s\n", got.UserName, test.expected.UserName)
			}

			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("ID: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}
