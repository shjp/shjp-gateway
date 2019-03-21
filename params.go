package gateway

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	core "github.com/shjp/shjp-core"
)

// Params defines interface that each API payload format should satisfy
type Params interface {
	core.Entity
	ReadParams(graphql.ResolveParams) error
}

// ModelParams defines interface that each model params should satisfy
type ModelParams interface {
	Params
	GetID() string
	GenerateID()
}

type group struct{ core.Group }

func (g *group) ReadParams(p graphql.ResolveParams) error {
	if id := p.Args["id"]; id != nil {
		g.ID = id.(string)
	}

	if name := p.Args["name"]; name != nil {
		g.Name = name.(string)
	} else {
		return errors.New("Name is required for group")
	}

	if description := p.Args["description"]; description != nil {
		g.Description = description.(string)
	}

	if imageURL := p.Args["image_url"]; imageURL != nil {
		g.ImageURL = imageURL.(string)
	}

	return nil
}

/*type role struct{ core.Role }

func (r *role) ReadParams(p graphql.ResolveParams) error {
	if name := p.Args["name"]; name != nil {
		r.Name = name.(string)
	}

	if group := p.Args["group"]; group != nil {
		r.Group = group.(string)
	}

	return nil
}*/

type announcement struct{ core.Announcement }

func (a *announcement) ReadParams(p graphql.ResolveParams) error {
	if id := p.Args["id"]; id != nil {
		a.ID = id.(string)
	}

	if name := p.Args["name"]; name != nil {
		a.Name = name.(string)
	} else {
		return errors.New("Name is required for announcement")
	}

	if authorID := p.Args["authorId"]; authorID != nil {
		a.AuthorID = authorID.(string)
	}

	if content := p.Args["content"]; content != nil {
		a.Content = content.(string)
	}

	return nil
}

type event struct{ core.Event }

func (e *event) ReadParams(p graphql.ResolveParams) error {
	if id := p.Args["id"]; id != nil {
		e.ID = id.(string)
	}

	if name := p.Args["name"]; name != nil {
		e.Name = name.(string)
	} else {
		return errors.New("Name is required for event")
	}

	if start := p.Args["start"]; start != nil {
		startStr := start.(string)
		e.Start = &startStr
	}

	if end := p.Args["end"]; end != nil {
		endStr := end.(string)
		e.End = &endStr
	}

	if creator := p.Args["creator"]; creator != nil {
		creatorStr := creator.(string)
		e.Creator = &creatorStr
	}

	if deadline := p.Args["deadline"]; deadline != nil {
		deadlineStr := deadline.(string)
		e.Deadline = &deadlineStr
	}

	if allowMaybe := p.Args["allow_maybe"]; allowMaybe != nil {
		e.AllowMaybe = allowMaybe.(bool)
	}

	if description := p.Args["description"]; description != nil {
		e.Description = description.(string)
	}

	if location := p.Args["location"]; location != nil {
		locationStr := location.(string)
		e.Location = &locationStr
	}

	if locationDescription := p.Args["location_description"]; locationDescription != nil {
		ldStr := locationDescription.(string)
		e.LocationDescription = &ldStr
	}

	for _, gid := range p.Args["group_ids"].([]interface{}) {
		e.GroupIDs = append(e.GroupIDs, gid.(string))
	}

	return nil
}

type user struct{ core.User }

func (u *user) ReadParams(p graphql.ResolveParams) error {
	if id := p.Args["id"]; id != nil {
		u.ID = id.(string)
	}

	if name := p.Args["name"]; name != nil {
		nameStr := name.(string)
		u.Name = &nameStr
	}

	if accountType := p.Args["accountType"]; accountType != nil {
		accountTypeStr := accountType.(string)
		u.AccountType = &accountTypeStr
	}

	if password := p.Args["password"]; password != nil {
		passwordStr := password.(string)
		u.Password = &passwordStr
	}

	if baptismalName := p.Args["baptismal_name"]; baptismalName != nil {
		bnStr := baptismalName.(string)
		u.BaptismalName = &bnStr
	}

	if birthday := p.Args["birthday"]; birthday != nil {
		bStr := birthday.(string)
		u.Birthday = &bStr
	}

	if feastday := p.Args["feastday"]; feastday != nil {
		fdStr := feastday.(string)
		u.Feastday = &fdStr
	}

	if lastActive := p.Args["lastActive"]; lastActive != nil {
		laStr := lastActive.(string)
		u.LastActive = &laStr
	}

	if email := p.Args["email"]; email != nil {
		emailStr := email.(string)
		u.Email = &emailStr
	}

	if password := p.Args["password"]; password != nil {
		passwordStr := password.(string)
		u.Password = &passwordStr
	}

	return nil
}

type groupMembership struct{ core.GroupMembership }

func (m *groupMembership) ReadParams(p graphql.ResolveParams) error {
	if p.Args["group_id"] == nil {
		return errors.New("GroupID is required for group membership request")
	}
	m.GroupID = p.Args["group_id"].(string)

	if p.Args["user_id"] == nil {
		return errors.New("UserID is required for group membership request")
	}
	m.UserID = p.Args["user_id"].(string)

	if p.Args["role_id"] == nil {
		return errors.New("RoleID is required for group membership request")
	}
	m.RoleID = p.Args["role_id"].(string)

	if p.Args["status"] == nil {
		return errors.New("Status is required for group membership request")
	}
	m.Status = core.MembershipStatus(p.Args["status"].(string))

	return nil
}

type eventRSVP struct{ core.EventRSVP }

func (r *eventRSVP) ReadParams(p graphql.ResolveParams) error {
	if p.Args["event_id"] == nil {
		return errors.New("event_id is required for event RSVP request")
	}
	r.EventID = p.Args["event_id"].(string)

	if p.Args["user_id"] == nil {
		return errors.New("user_id is required for event RSVP request")
	}
	r.UserID = p.Args["user_id"].(string)

	if p.Args["rsvp"] == nil {
		return errors.New("rsvp is required for event RSVP request")
	}
	r.RSVP = p.Args["rsvp"].(string)

	return nil
}
