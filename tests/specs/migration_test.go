package specs

import (
	"auth_ms/migrations"
	"os"
	"testing"
)

func TestNewMigration(t *testing.T) {
	basePath := "/path/to/base"
	m := migrations.NewMigration(basePath)

	if m == nil {
		t.Error("Expected non-nil value for Migration, got nil")
	}
}

func TestGetMigrationFiles(t *testing.T) {
	m := migrations.NewMigration("")

	workingDir, _ := os.Getwd()

	files := []string{"file1.sql", "file2.sql", "file3.txt"}
	baseDir := workingDir + "/../data/migrations"

	for _, file := range files {
		_, err := os.Create(baseDir + "/" + file)
		if err != nil {
			t.Errorf("Failed to create test files: %s", err)
		}
	}

	result, err := m.GetMigrationFiles("/../data/")
	if err != nil {
		t.Errorf("Failed to GetMigrationFiles: %s", err)
	}

	expected := []string{"file1.sql", "file2.sql"}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Files are incorrect")
		}
	}

	for _, file := range files {
		err := os.Remove(baseDir + "/" + file)
		if err != nil {
			t.Errorf("Erreur lors de la suppression du fichier de test: %s", err)
		}
	}
}
