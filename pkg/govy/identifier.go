package govy

import "github.com/nobl9/govy/internal/uuid"

// instanceID is a composite identifier used to identify [PropertyRules] variations.
type instanceID struct {
	// generatedID is the auto-generated ID used when userSuppliedID is empty.
	// Internal slice and map rules leave it empty because their enclosing rule owns the identity.
	generatedID string
	// userSuppliedID overrides generatedID and is supplied by the user.
	userSuppliedID string
}

func newInstanceID() instanceID {
	return instanceID{generatedID: uuid.GenerateUUID()}
}

func (i instanceID) WithUserSuppliedID(id string) instanceID {
	if id == "" {
		return newInstanceID()
	}
	i.userSuppliedID = id
	return i
}

func (i instanceID) withNextGeneratedID() instanceID {
	if i.userSuppliedID != "" {
		return i
	}
	return newInstanceID()
}

func (i instanceID) GetID() string {
	if i.userSuppliedID != "" {
		return i.userSuppliedID
	}
	return i.generatedID
}
