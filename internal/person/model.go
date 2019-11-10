package person

import "github.com/google/uuid"

const (
	RoleGuest        = "guest"
	RoleInhabitant   = "inhabitant"
	RoleTransient    = "transient"
	RoleFamilyMember = "family_member"
)

type Person struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Roles     []string  `json:"roles"`
}
