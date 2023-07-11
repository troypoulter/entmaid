package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func GenerateDiagram(schemaPath string) error {
	graph, err := entc.LoadGraph(schemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("failed to load schema graph from the path %s: %v", schemaPath, err)
	}
	// Generate the Mermaid code for the ERD diagram
	mermaidCode := generateMermaidCode(graph)

	// Save the Mermaid code to a file
	err = ioutil.WriteFile("erd.md", []byte(mermaidCode), 0644)
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

	// Process each node in the graph
	for _, node := range graph.Nodes {
		entityName := node.Name
		builder.WriteString(fmt.Sprintf("   %s {\n", entityName))

		// Process each field in the node
		for _, field := range node.Fields {
			builder.WriteString(fmt.Sprintf("    %s %s\n", formatType(field.Type.String()), field.Name))
		}

		builder.WriteString("  }\n\n")

		// Process the edges/relationships
		// for _, edge := range node.Edges {
		// 	builder.WriteString(fmt.Sprintf("  %s }|--|{ %s\n", edge.Name, edge.Ref.Name))
		// }
	}

	builder.WriteString("```")

	return builder.String()
}

func formatType(s string) string {
	return strings.NewReplacer(
		".", "DOT",
		"*", "STAR",
		"[", "LBRACK",
		"]", "RBRACK",
	).Replace(s)
}
