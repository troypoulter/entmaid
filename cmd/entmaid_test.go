package cmd

import (
	"fmt"
	"os"
	"testing"
)

type testCase struct {
	schemaPath     string
	targetPath     string
	expectedOutput string
	startPattern   string
	endPattern     string
}

const (
	defaultStartPattern = "<!-- #start:entmaid -->"
	defaultEndPattern   = "<!-- #end:entmaid -->"
)

func TestGenerateDiagram(t *testing.T) {
	testCases := []testCase{
		{
			schemaPath:     "../examples/start/schema",
			targetPath:     "../examples/start/readme.md",
			expectedOutput: "../examples/start/readme-expected.md",
			startPattern:   defaultStartPattern,
			endPattern:     defaultEndPattern,
		},
		{
			schemaPath:     "../examples/m2m2types/schema",
			targetPath:     "../examples/m2m2types/readme.md",
			expectedOutput: "../examples/m2m2types/readme-expected.md",
			startPattern:   defaultStartPattern,
			endPattern:     defaultEndPattern,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("SchemaPath=%s, TargetPath=%s", tc.schemaPath, tc.targetPath), func(t *testing.T) {
			// Generate the diagram
			err := GenerateDiagram(tc.schemaPath, tc.targetPath, Markdown, tc.startPattern, tc.endPattern)
			if err != nil {
				t.Fatalf("Failed to generate diagram: %v", err)
			}

			// Compare the generated file with the expected output file
			if !compareFiles(tc.targetPath, tc.expectedOutput) {
				t.Errorf("Generated file does not match the expected output")
			}
		})
	}
}

func compareFiles(file1, file2 string) bool {
	content1, err := os.ReadFile(file1)
	if err != nil {
		fmt.Printf("Failed to read file %s: %v\n", file1, err)
		return false
	}

	content2, err := os.ReadFile(file2)
	if err != nil {
		fmt.Printf("Failed to read file %s: %v\n", file2, err)
		return false
	}

	return string(content1) == string(content2)
}
