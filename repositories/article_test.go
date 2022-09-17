package repositories_test

import (
	"testing"

	"github.com/HT0323/go_api/models"
	"github.com/HT0323/go_api/repositories"
	"github.com/HT0323/go_api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestUpdateNiceNum(t *testing.T) {
	expectedNiceNum := 3
	articleId := 1
	beforeNiceNum := testdata.ArticleTestData[0].NiceNum

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

	t.Cleanup(func() {
		const sqlStr = `
			update articles set nice = ?
			where article_id = ?
		`
		testDB.Exec(sqlStr, beforeNiceNum, articleId)
	})
}
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertText",
		Contents: "testtest",
		UserName: "saki",
	}

	expectedArticleNum := 3
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
	expectedNum := len(testdata.ArticleTestData)
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
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subTest2",
			expected:  testdata.ArticleTestData[1],
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
				t.Errorf("title: get %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("contents: get %s but want %s\n", got.Contents, test.expected.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("username: get %s but want %s\n", got.UserName, test.expected.UserName)
			}

			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("nicenum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}
