package govy

import "github.com/nobl9/govy/internal/uuid"

// instanceID is a composite identifier used to identify [Validator] and [PropertyRules] variations.
type instanceID struct {
	// generatedID is always filled and generated upon creation of [instanceID].
	generatedID string
	// userSuppliedID overrides generatedID and is supplied by the user.
	userSuppliedID string
}

func newInstanceID() instanceID {
	return instanceID{generatedID: uuid.GenerateUUID()}
}

func (i instanceID) WithUserSuppliedID(id string) instanceID {
	i.userSuppliedID = id
	return i
}

func (i instanceID) HasUserSuppliedID() bool {
	return i.userSuppliedID != ""
}

func (i instanceID) GetUserSuppliedID() string {
	return i.userSuppliedID
}

func (i instanceID) GetGeneratedID() string {
	return i.generatedID
}
