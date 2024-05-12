package database

import (
	"context"
	"time"
)

func GetUserById(id string) (any, error) {
	ctx := context.Background()
	result, err := ReadQueryWithMapping("MATCH (user:User{id:$id}) return user", map[string]any{"id": id}, ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func GetUserByName(name string) (any, error) {
	ctx := context.Background()
	result, err := ReadQueryWithMapping("MATCH(user:User{name:$name}) return user", map[string]any{"name": name}, ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func GetAllUsers() (any, error) {
	ctx := context.Background()
	result, err := ReadQuery("MATCH (user:User) return user", ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func CreateUser(name, lastName string, privilege any) error {
	ctx := context.Background()
	_, err := WriteQueryWithMapping("CREATE (user:User{name: $name,last_name: $last_name, privilege:$privilege,created_at: $created_at}) return user", map[string]any{
		"name":       name,
		"last_name":  lastName,
		"privilege":  privilege,
		"created_at": time.Now().String(),
	}, ctx)
	if err != nil {
		return err
	}
	return nil
}
