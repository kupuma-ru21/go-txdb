package graph

// Basic imports
import (
	"context"
	"database/sql"
	"go-txdb/graph/db"
	"go-txdb/graph/model"
	"log"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ExampleTestSuite struct {
	suite.Suite
	txDB *sql.DB
	*client.Client
}

var link db.Link

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.txDB = openTestDB(suite.T())
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	suite.setUpEnvForTest()
	operation := `
		mutation UpdateLink{
			updateLink(
				input: {
					id: "1",
					title: "something",
					address: "somewhere"
				}){
				title,
				address,
				id,
			}
		}`

	var resp struct {
		Link model.Link `json:"updateLink"`
	}

	err := suite.Post(
		operation,
		&resp,
		client.Var("input", model.UpdateLinkInput{}),
	)

	suite.Assert().NoError(err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (suite *ExampleTestSuite) setUpEnvForTest() {
	link = db.Link{
		Title:   null.StringFrom("Title"),
		Address: null.StringFrom("Address"),
	}
	err := link.Insert(context.TODO(), suite.txDB, boil.Infer())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func openTestDB(t *testing.T) (db *sql.DB) {
	// dsn serves as an unique identifier for connection pool
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}

func init() {
	// we register an sql driver named "txdb"
	txdb.Register("txdb", "mysql", "root@/txdb_test")
}
