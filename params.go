package gateway

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	core "github.com/shjp/shjp-core"
)

// Params defines interface that each model params should satisfy
type Params interface {
	GetID() string
	GenerateID()
	Pack(core.Intent, core.OperationType) (*core.Message, error)
	ReadParams(graphql.ResolveParams) error
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

	if date := p.Args["date"]; date != nil {
		dateStr := date.(string)
		e.Date = &dateStr
	}

	if length := p.Args["length"]; length != nil {
		e.Length = length.(int)
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

/*type user struct { core.User }

func (u *User) ReadParams(p graphql.ResolveParams) error {

}*/
