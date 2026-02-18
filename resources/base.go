package resources

import "github.com/rewritetoday/golang/rest"

// Base shares access to the low-level REST client.
type Base struct {
	Rest *rest.Client
}
