package schema

//import (
//	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
//
//	"github.com/graphql-go/graphql"
//	"github.com/graphql-go/graphql/gqlerrors"
//)
//
//var requiresAuthDirective = &graphql.Directive{
//	Name:        "requiresAuth",
//	Description: "Indicates that a field requires authentication.",
//	Locations: []string{
//		graphql.DirectiveLocationField,
//	},
//}
//
//func authMiddleware(resolve graphql.FieldResolveFn) graphql.FieldResolveFn {
//	return func(p graphql.ResolveParams) (interface{}, error) {
//		// Check if the "requiresAuth" directive is present
//		if _, ok := p.Info.Directives["requiresAuth"]; ok {
//			// Check if the user is authenticated
//			if !isAuthenticated(p.Context) {
//				return nil, gqlerrors.NewFormattedError("User is not authenticated")
//			}
//		}
//
//		// Call the original resolver
//		return resolve(p)
//	}
//}
//
//func loggedInQuery(resolver resolvers.Resolver) *graphql.Object {
//	query := graphql.NewObject(graphql.ObjectConfig{
//		Name: "something",
//		Fields: graphql.Fields{
//			"profile": &graphql.Field{
//				Type:        userType,
//				Description: "Get user profile",
//				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//
//					return nil, nil
//				},
//			},
//		},
//	})
//
//	return query
//}
