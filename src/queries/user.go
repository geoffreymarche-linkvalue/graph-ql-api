package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/graph-ql-api/src/types"
)

// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var users []types.User

			// ... Implémenter la logique de base de données ici

			return users, nil
		},
	}
}