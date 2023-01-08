package rdbreposity

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	testingDB *gorm.DB
)

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	testingDB = db
	testingDB.AutoMigrate(&model.Manufacturer{})
	testingDB.AutoMigrate(&model.Product{})

	m.Run()
	os.Exit(0)
}

var testGetUsersCases = []struct {
	name       string
	testCount  int
	testParams model.GetProductsParams
	wantCount  int
	wantError  bool
}{
	{
		name:      "zero path - GetUsers",
		testCount: 0,
		wantCount: 0,
		wantError: false,
	},
	{
		name:      "happy path - GetUsers",
		testCount: 10,
		wantCount: 10,
		wantError: false,
	},
	{
		name:      "error path",
		wantCount: 10,
		wantError: true,
	},
	{
		name: "filter path - fullname contains 1",
		testParams: model.GetProductsParams{
			Name: model.StringToPointer("1"),
		},
		testCount: 20,
		wantCount: 3, // 1, 10, 11
		wantError: false,
	},
	{
		name: "filter path - fullname contains test",
		testParams: model.GetProductsParams{
			Name: model.StringToPointer("test"),
		},
		testCount: 20,
		wantCount: 20,
		wantError: false,
	},
}

func Test_reposity_Get(t *testing.T) {
	for _, c := range testGetUsersCases {

		t.Run(c.name, func(t *testing.T) {
			createRandomUserToDB(c.testCount)
			testParams := c.testParams
			repo := repository{db: testingDB}

			prodct, err := repo.Get(context.Background(), testParams)

			if err != nil && !c.wantError {
				t.Errorf("got error %v", err)
				return
			}
			if len(prodct) != c.wantCount {
				t.Errorf("Expected %d users, got %d", c.wantCount, len(prodct))
			}
		})
	}
}

func createRandomUserToDB(numbers int) {
	for i := 0; i < numbers; i++ {
		product := &model.Product{
			ID:   strconv.Itoa(i),
			Name: "test" + strconv.Itoa(i),
		}
		testingDB.Create(product)
	}
}

var testGetUserIdCases = []struct {
	name      string
	id        string
	testCount int
	want      string
	wantError bool
}{
	{
		name:      "happy - get user id",
		id:        "0",
		testCount: 1,
		wantError: false,
	},
	{
		name:      "happy - get user id 2",
		id:        "1",
		testCount: 2,
		wantError: false,
	},
	{
		name:      "happy - get user id 100",
		id:        "99",
		testCount: 100,
		wantError: false,
	},
	{
		name:      "error - not create",
		testCount: 0,
		id:        "1",
		wantError: true,
	},
	{
		name:      "error - random string",
		testCount: 20,
		id:        "gfeawgeawgew",
		wantError: true,
	},
}

func Test_repository_GetByID(t *testing.T) {
	for _, c := range testGetUserIdCases {

		t.Run(c.name, func(t *testing.T) {
			createRandomUserToDB(c.testCount)
			repo := repository{db: testingDB}

			prodct, err := repo.GetByID(context.Background(), c.id)

			if err != nil && !c.wantError {
				t.Errorf("got error %v", err)
				return
			}

			if prodct.ID == c.id {
				t.Errorf("Expected %s, got %s", c.id, prodct.ID)
			}
		})
	}

}
