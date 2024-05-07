package database

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetProject(d neo4j.DriverWithContext, name string) *neo4j.EagerResult {
	ctx := context.Background()
	result, err := neo4j.ExecuteQuery(ctx, d, "MATCH (p:Project{name:$name}) return p", map[string]any{
		"name": name,
	}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}
	return result
}
