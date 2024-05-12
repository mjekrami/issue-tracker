package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mjekrami/issue-tracker/cmd/auth"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	db *neo4j.DriverWithContext

	ErrQueryDB      = errors.New("could not run the query")
	ErrConnectionDB = errors.New("could not connect to the database")
	ErrAuthDB       = errors.New("could not authenticate to the database")
)

func New(host, port string, authdb auth.DBAuth) error {
	c, err := neo4j.NewDriverWithContext(fmt.Sprintf("bolt://%s:%s", host, port), neo4j.BasicAuth(authdb.Username, authdb.Password, ""))
	if err != nil {
		return err
	}
	err = verifyConnection(c)
	if err != nil {
		return err
	}
	db = &c
	return nil
}
func verifyConnection(d neo4j.DriverWithContext) error {
	ctx := context.Background()
	err := d.VerifyConnectivity(ctx)
	if err != nil {
		return ErrConnectionDB
	}
	err = d.VerifyAuthentication(ctx, nil)
	if err != nil {
		return ErrAuthDB
	}
	return nil
}
func ReadQuery(cypher string, ctx context.Context) (any, error) {
	session := (*db).NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	res, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, _ := tx.Run(ctx, cypher, nil)
		return result.Collect(ctx)
	}, neo4j.WithTxTimeout(time.Second*5), neo4j.WithTxMetadata(map[string]any{"appName": "IssueTracker"}))
	if err != nil {
		return nil, ErrQueryDB
	}
	return res, nil
}

func ReadQueryWithMapping(cypher string, mapping map[string]any, ctx context.Context) (any, error) {
	session := (*db).NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	res, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, cypher, mapping)
		if err != nil {
			return nil, err
		}
		return result.Collect(ctx)
	}, neo4j.WithTxTimeout(time.Second*5), neo4j.WithTxMetadata(map[string]any{"appName": "IssueTracker"}))
	if err != nil {
		return nil, ErrQueryDB
	}
	return res, nil
}

func WriteQuery(cypher string, ctx context.Context) (any, error) {
	session := (*db).NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	res, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, _ := tx.Run(ctx, cypher, nil)
		return result.Collect(ctx)
	}, neo4j.WithTxTimeout(time.Second*5), neo4j.WithTxMetadata(map[string]any{"appName": "IssueTracker"}))
	if err != nil {
		return nil, ErrQueryDB
	}
	return res, nil
}

func WriteQueryWithMapping(cypher string, mapping map[string]any, ctx context.Context) (any, error) {
	session := (*db).NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	res, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, cypher, mapping)
		if err != nil {
			return nil, err
		}
		return result.Collect(ctx)
	}, neo4j.WithTxTimeout(time.Second*5), neo4j.WithTxMetadata(map[string]any{"appName": "IssueTracker"}))
	if err != nil {
		return nil, ErrQueryDB
	}
	return res, nil
}
