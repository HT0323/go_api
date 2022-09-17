package repositories_test

import (
	"testing"

	"github.com/HT0323/go_api/models"
	"github.com/HT0323/go_api/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestUpdateNiceNum(t *testing.T) {
	expectedNiceNum := 3
	articleId := 1
	err := repositories.UpdateNiceNum(testDB, articleId)
	if err != nil {
		t.Error(err)
	}

	article, err := repositories.SelectArticleDetail(testDB, articleId)
	if err != nil {
		t.Error(err)
	}

	if article.NiceNum != expectedNiceNum {
		t.Errorf("fail to update nice num: expected %d but got %d\n",
			expectedNiceNum,
			article.NiceNum)
	}

}
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertText",
		Contents: "testtest",
		UserName: "saki",
	}

	expectedArticleNum := 5
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}

	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := 4
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {

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
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)

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
