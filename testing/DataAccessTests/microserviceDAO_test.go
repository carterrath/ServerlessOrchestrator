package dataaccesstests

import (
	"database/sql"
	"log"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	// Other imports as necessary
)

type DataAccessTestSuite struct {
	suite.Suite
	DB         *sql.DB
	DataAccess *dataaccess.MicroservicesDAO
}

func (suite *DataAccessTestSuite) SetupSuite() {
	// Create a new database connection for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	suite.DB = db

	// Initialize your DataAccess object with the test database
	suite.DataAccess = &dataaccess.MicroservicesDAO{DB: db}

	// Create the same schema as your actual database
	createMicroservicesTable := `CREATE TABLE IF NOT EXISTS microservices (
		"id" INTEGER PRIMARY KEY, 
		"name" TEXT UNIQUE,
		"service_hook" TEXT, 
		"build_script" TEXT,
		"place_holder" TEXT
	);`
	_, err = db.Exec(createMicroservicesTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func (suite *DataAccessTestSuite) TearDownSuite() {
	// Close the database connection after tests
	suite.DB.Close()
}

// TestOpenDBConnection tests the OpenDBConnection function
func TestOpenDBConnection(t *testing.T) {
	//Test successful database connection.
	//Test failure to connect to the database (e.g., wrong path or permissions).
	dao := &dataaccess.MicroservicesDAO{} // Assuming that MicroservicesDAO is properly defined somewhere

	// Attempt to open the database connection
	db, err := dao.OpenDBConnection()

	// If an error occurs, the test should fail
	if err != nil {
		t.Errorf("OpenDBConnection() error = %v, wantErr %v", err, false)
	}

	// If db is nil, then the connection was not successful
	if db == nil {
		t.Errorf("OpenDBConnection() db is nil, want non-nil")
	}

	// Assuming that the db is not nil, it should be possible to ping the database
	if err := db.Ping(); err != nil {
		t.Errorf("OpenDBConnection() db cannot be pinged: %v", err)
	}

	// Close the database connection after the test
	defer db.Close()
}

// TestCloseDBConnection tests the CloseDBConnection function
func TestCloseDBConnection(t *testing.T) {
	//Test closing an open database connection.
	//Test closing a nil database connection to ensure no panic.
	dao := &dataaccess.MicroservicesDAO{} // Assuming that MicroservicesDAO is properly defined somewhere

	// Set up by opening a database connection which will be closed later
	db, err := dao.OpenDBConnection()
	if err != nil {
		t.Fatalf("OpenDBConnection() error = %v, wantErr %v", err, false)
	}
	if db == nil {
		t.Fatalf("OpenDBConnection() db is nil, want non-nil")
	}

	// Defer the CloseDBConnection call to ensure it runs even if something fails
	defer func() {
		dao.CloseDBConnection(db)
		// After closing, we should get an error if we try to ping the database
		if err := db.Ping(); err == nil {
			t.Errorf("DB connection should be closed, but Ping did not return an error")
		}
	}()

	// Normal operation here, the connection should be open and pinging should not return an error
	if err := db.Ping(); err != nil {
		t.Fatalf("DB connection should be open, but Ping returned an error: %v", err)
	}
	// Actual test for CloseDBConnection, however, this will be effectively checked by the deferred function
}

// TestGetAllMicroservices tests the GetAllMicroservices function
func TestGetAll(t *testing.T) {
	//Test retrieving a list of microservices successfully.
	//Test retrieving an empty list if there are no microservices.
	//Test handling SQL query errors.
}

// TestGetMicroserviceByName tests the GetMicroserviceByName function
func TestGetMicroserviceByName(t *testing.T) {
	//Test retrieving a single microservice by name successfully.
	//Test behavior when no microservice is found by that name.
	//Test handling SQL query errors.
}

// TestGetMicroserviceByID tests the GetMicroserviceByID function
func TestGetMicroserviceByID(t *testing.T) {
	//Test retrieving a single microservice by ID successfully.
	//Test behavior when no microservice is found by that ID.
	//Test handling SQL query errors.
}

// TestInsertMicroservice tests the InsertMicroservice function
func TestInsertMicroservice(t *testing.T) {
	//Test inserting a new microservice successfully.
	//Test handling SQL execution errors.
}

// TestUpdateMicroservice tests the UpdateMicroservice function
func TestUpdateMicroservice(t *testing.T) {
	//Test updating an existing microservice successfully.
	//Test handling SQL execution errors.
}

// TestDeleteMicroservice tests the DeleteMicroservice function
func TestDeleteMicroservice(t *testing.T) {
	//Test deleting an existing microservice successfully.
	//Test handling SQL execution errors.
}
