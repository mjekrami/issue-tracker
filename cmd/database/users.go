package database

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetUser(d neo4j.DriverWithContext, id string, name ...string) *neo4j.EagerResult {
	ctx := context.Background()
	if name != nil {
		result, err := neo4j.ExecuteQuery(ctx, d, "MATCH (user:User{name: $name,id:$id}) return user", map[string]any{
			"name": name,
			"id":   id,
		}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
		if err != nil {
			panic(err)
		}
		return result
	}
	result, err := neo4j.ExecuteQuery(ctx, d, "MATCH (user:User{id:$id}) return user", map[string]any{
		"id": id,
	}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}
	return result
}
