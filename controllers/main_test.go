package controllers_test

import (
	"testing"

	"github.com/HT0323/go_api/controllers"
	"github.com/HT0323/go_api/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()

}
