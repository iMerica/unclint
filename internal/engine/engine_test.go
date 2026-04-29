package engine_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iMerica/unclint/internal/config"
	"github.com/iMerica/unclint/internal/engine"
	"github.com/iMerica/unclint/internal/output"
	"github.com/iMerica/unclint/internal/rules"
)

func TestGoldenFiles(t *testing.T) {
	allRules, err := rules.LoadDefaultRules()
	if err != nil {
		t.Fatalf("failed to load rules: %v", err)
	}
	matchers := rules.BuildMatchers(allRules)

	cfg := &config.Config{
		MaxScore:    0,
		MinSeverity: 1,
		Rules:       map[string]any{"corporate": true},
	}

	linter := engine.New(cfg, matchers)

	goldenDir := filepath.Join("..", "..", "testdata", "golden")
	entries, err := os.ReadDir(goldenDir)
	if err != nil {
		if os.IsNotExist(err) {
			t.Skip("No testdata/golden directory found")
		}
		t.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".txt" || !strings.HasSuffix(entry.Name(), ".input.txt") {
			continue
		}

		baseName := entry.Name()[:len(entry.Name())-10] // strip .input.txt
		t.Run(baseName, func(t *testing.T) {
			inputPath := filepath.Join(goldenDir, entry.Name())
			outputPath := filepath.Join(goldenDir, baseName+".output.txt")

			inputData, err := os.ReadFile(inputPath)
			if err != nil {
				t.Fatal(err)
			}

			res, err := linter.LintText(entry.Name(), string(inputData), true)
			if err != nil {
				t.Fatal(err)
			}

			var buf bytes.Buffer
			output.PrintText(&buf, []engine.Result{res}, cfg.MaxScore, true)

			expectedData, err := os.ReadFile(outputPath)
			if err != nil {
				if os.IsNotExist(err) {
					// Auto-generate if missing
					os.WriteFile(outputPath, buf.Bytes(), 0644)
					t.Logf("Generated golden file %s", outputPath)
					return
				}
				t.Fatal(err)
			}

			if string(expectedData) != buf.String() {
				t.Errorf("Mismatch in %s.\nExpected:\n%s\nGot:\n%s", baseName, string(expectedData), buf.String())
			}
		})
	}
}
