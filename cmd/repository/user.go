package repository

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type userRepo struct {
}

func NewUserRepo(d neo4j.DriverWithContext) *userRepo {
	return &userRepo{}
}
