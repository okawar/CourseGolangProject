package sql_test

import (
	"golang_pr/config"
	"golang_pr/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupDir(t *testing.T) {
	testDir := config.Get().Server.Workdir
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		err := os.Mkdir(testDir, 0755)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { os.RemoveAll(testDir) })
	}
}

func TestInitDB(t *testing.T) {
	TestSetupDir(t)
	assert.NotNil(t, sql.GetDB(), "Database should be initialized")
}

func TestAddMigrations(t *testing.T) {
	TestSetupDir(t)

	dummyMigration := func() {}

	sql.AddMigrations(dummyMigration)

	assert.Contains(t, sql.GetMigrations(), dummyMigration, "Migration should be added")
}

func TestMigrate(t *testing.T) {
	TestSetupDir(t)

	dummyMigrationExecuted := false
	dummyMigration := func() {
		dummyMigrationExecuted = true
	}

	sql.AddMigrations(dummyMigration)
	sql.Migrate()

	assert.True(t, dummyMigrationExecuted, "Migration should be executed")
}

func TestMain(m *testing.M) {
	TestSetupDir(nil)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetDB(t *testing.T) {
	TestSetupDir(t)
	assert.NotNil(t, sql.GetDB(), "Database instance should not be nil")
}

func TestGetMigrations(t *testing.T) {
	TestSetupDir(t)

	dummyMigration := func() {}

	sql.AddMigrations(dummyMigration)

	assert.Contains(t, sql.GetMigrations(), dummyMigration, "Migrations should contain the dummy migration")
}
