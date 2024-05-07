package database

import (
	"context"
)

func GetIssue(name string, state ...string) (any, error) {
	ctx := context.Background()
	if state == nil {
		result, err := ReadQueryWithMapping("MATCH (task:Task{state:$state,name:$name}) return task", map[string]any{"state": state}, ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	result, err := ReadQueryWithMapping("MATCH (task:Task{name:$name}) return task", map[string]any{"name": name}, ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllIssues() (any, error) {
	ctx := context.Background()
	result, err := ReadQuery("MATCH (task:Task) return task", ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
