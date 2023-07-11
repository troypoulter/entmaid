package cmd

import (
	"fmt"
	"os"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func GenerateDiagram(schemaPath string, targetPath string) error {
	graph, err := entc.LoadGraph(schemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("failed to load schema graph from the path %s: %v", schemaPath, err)
	}
	// Generate the Mermaid code for the ERD diagram
	mermaidCode := generateMermaidCode(graph)

	// Save the Mermaid code to a file
	err = os.WriteFile(targetPath, []byte(mermaidCode), 0644)
	if err != nil {
		return fmt.Errorf("failed to write Mermaid file: %v", err)
	}

	fmt.Println("Mermaid file generated successfully.")

	return nil
}

// generateMermaidCode generates the Mermaid code for the ERD diagram based on the schema graph.
func generateMermaidCode(graph *gen.Graph) string {
	var builder strings.Builder

	builder.WriteString("```mermaid\n")
	builder.WriteString("erDiagram\n")

	for _, node := range graph.Nodes {
		entityName := node.Name
		builder.WriteString(fmt.Sprintf("\t%s {\n", entityName))

		if node.HasOneFieldID() {
			builder.WriteString(fmt.Sprintf("\t\t%s %s\n", formatType(node.ID.Type.String()), node.ID.Name))
		}

		for _, field := range node.Fields {
			builder.WriteString(fmt.Sprintf("\t\t%s %s\n", formatType(field.Type.String()), field.Name))
		}

		builder.WriteString("\t}\n\n")

		for _, edge := range node.Edges {
			if edge.IsInverse() {
				continue
			}

			builder.WriteString(fmt.Sprintf("  %s %s %s : %s-%s\n", node.Name, getEdgeRelationship(edge), edge.Type.Name, edge.Name, edge.Ref.Name))
		}
	}

	builder.WriteString("```")

	return builder.String()
}

func formatType(s string) string {
	return strings.NewReplacer(
		".", "-",
	).Replace(s)
}

func getEdgeRelationship(edge *gen.Edge) string {
	if edge.O2M() {
		return "|o--o|"
	}

	if edge.M2O() {
		return "|o--o{"
	}

	if edge.M2M() {
		return "}o--o|"
	}

	return "}o--o{"
}