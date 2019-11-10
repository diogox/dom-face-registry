package person

import (
	"github.com/google/uuid"

	pb "github.com/diogox/dom-face-recognizer/internal/pb"
)

type Converter struct {}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) PersonAsResponse(res Person) *pb.Person {
	return &pb.Person{
		Id:        res.ID.String(),
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Roles:     c.mapRolesToResponse(res.Roles),
	}
}

func (c *Converter) PersonAsRequest(req pb.Person) Person {
	return Person{
		ID:        uuid.MustParse(req.Id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Roles:     c.mapRolesToRequest(req.Roles),
	}
}

func (c *Converter) PeopleAsResponse(req []Person) []*pb.Person {
	res := make([]*pb.Person, 0, len(req))

	for _, p := range req {
		res = append(res, c.PersonAsResponse(p))
	}

	return res
}

func (c *Converter) PeopleAsRequest(res []pb.Person) []Person {
	req := make([]Person, 0, len(res))

	for _, p := range res {
		req = append(req, c.PersonAsRequest(p))
	}

	return req
}

func (c *Converter) mapRolesToResponse(req []string) []pb.PersonRole {
	roles := make([]pb.PersonRole, 0, len(req))

	for _, r := range req {
		switch r {
		case RoleFamilyMember:
			roles = append(roles, pb.PersonRole_FAMILY_MEMBER)
		case RoleInhabitant:
			roles = append(roles, pb.PersonRole_INHABITANT)
		case RoleGuest:
			roles = append(roles, pb.PersonRole_GUEST)
		case RoleTransient:
			roles = append(roles, pb.PersonRole_TRANSIENT)
		}
	}

	return roles
}

func (c *Converter) mapRolesToRequest(res []pb.PersonRole) []string {
	roles := make([]string, 0, len(res))

	for _, r := range res {
		switch r {
		case pb.PersonRole_FAMILY_MEMBER:
			roles = append(roles, RoleFamilyMember)
		case pb.PersonRole_INHABITANT:
			roles = append(roles, RoleInhabitant)
		case pb.PersonRole_GUEST:
			roles = append(roles, RoleGuest)
		case pb.PersonRole_TRANSIENT:
			roles = append(roles, RoleTransient)
		}
	}

	return roles
}
