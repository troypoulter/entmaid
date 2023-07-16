package cmd

import (
	"fmt"
	"os"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func GenerateDiagram(schemaPath string, targetPath string, outputType OutputType, startPattern string, endPattern string) error {
	graph, err := entc.LoadGraph(schemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("failed to load schema graph from the path %s: %v", schemaPath, err)
	}
	// Generate the Mermaid code for the ERD diagram
	mermaidCode, err := generateMermaidCode(graph)
	if err != nil {
		return err
	}

	mermaidCode = addMermaidToType(mermaidCode, outputType)

	err = insertMultiLineString(targetPath, mermaidCode, startPattern, endPattern)
	if err != nil {
		return fmt.Errorf("failed to insert Mermaid code into the file: %v", err)
	}

	fmt.Println("Mermaid file generated successfully.")

	return nil
}

// generateMermaidCode generates the Mermaid code for the ERD diagram based on the schema graph.
func generateMermaidCode(graph *gen.Graph) (string, error) {
	var builder strings.Builder

	builder.WriteString("erDiagram\n")

	for _, node := range graph.Nodes {
		builder.WriteString(fmt.Sprintf(" %s {\n", node.Name))

		if node.HasOneFieldID() {
			builder.WriteString(fmt.Sprintf("  %s %s\n", formatType(node.ID.Type.String()), node.ID.Name))
		}

		for _, field := range node.Fields {
			builder.WriteString(fmt.Sprintf("  %s %s\n", formatType(field.Type.String()), field.Name))
		}

		builder.WriteString(" }\n\n")
	}

	for _, node := range graph.Nodes {
		for _, edge := range node.Edges {
			if edge.IsInverse() {
				continue
			}

			_, err := builder.WriteString(fmt.Sprintf(" %s %s %s : %s%s\n", node.Name, getEdgeRelationship(edge), edge.Type.Name, edge.Name, getEdgeRefName(edge.Ref)))
			if err != nil {
				return "", fmt.Errorf("failed to write string: %v", err)
			}
		}
	}

	return builder.String(), nil
}

func addMermaidToType(mermaidCode string, outputType OutputType) string {
	switch outputType {
	case Markdown:
		return fmt.Sprintf("```mermaid\n%s\n```", mermaidCode)
	case Plain:
		return mermaidCode
	default:
		return mermaidCode
	}
}

func formatType(s string) string {
	return strings.NewReplacer(
		".", "-",
	).Replace(s)
}

func getEdgeRelationship(edge *gen.Edge) string {
	if edge.O2M() {
		return "|o--o{"
	}

	if edge.M2O() {
		return "}o--o|"
	}

	if edge.M2M() {
		return "}o--o{"
	}

	return "|o--o|"
}

func getEdgeRefName(ref *gen.Edge) string {
	if ref == nil {
		return ""
	}

	return fmt.Sprintf("-%s", ref.Name)
}

func insertMultiLineString(filePath string, multiLineString string, startPattern string, endPattern string) error {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Convert the content to a string
	fileContent := string(content)

	// Find the starting and ending strings
	startIndex := strings.Index(fileContent, startPattern)
	endIndex := strings.Index(fileContent, endPattern)

	// Check if the starting and ending strings are found
	if startIndex == -1 || endIndex == -1 {
		return fmt.Errorf("starting (%s) or ending (%s) string not found in the file", startPattern, endPattern)
	}

	// Construct the updated content with the generated multi-line string
	updatedContent := fileContent[:startIndex+len(startPattern)+1] + multiLineString + "\n" + fileContent[endIndex:]

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
