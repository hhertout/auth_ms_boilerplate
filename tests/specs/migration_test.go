package specs

import (
	"auth_ms/migrations"
	"os"
	"testing"
)

func TestGetMigrationFiles(t *testing.T) {
	m := migrations.NewMigration("")

	workingDir, _ := os.Getwd()

	files := []string{"file1.sql", "file2.sql", "file3.txt"}
	baseDir := workingDir + "/../data/migrations"

	for _, file := range files {
		_, err := os.Create(baseDir + "/" + file)
		if err != nil {
			t.Errorf("Erreur lors de la création du fichier de test: %s", err)
		}
	}

	result, err := m.GetMigrationFiles("/../data/")
	if err != nil {
		t.Errorf("Erreur rencontrée lors de l'exécution de GetMigrationFiles: %s", err)
	}

	expected := []string{"file1.sql", "file2.sql"}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Résultat incorrect à l'indice %d: obtenu %s, attendu %s", i, result[i], expected[i])
		}
	}

	for _, file := range files {
		err := os.Remove(baseDir + "/" + file)
		if err != nil {
			t.Errorf("Erreur lors de la suppression du fichier de test: %s", err)
		}
	}
}
