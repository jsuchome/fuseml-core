package fuseml

import (
	"github.com/pkg/errors"

	codeset "github.com/fuseml/fuseml-core/gen/codeset"
)

type CodesetStore struct {
	// FIXME this is just internal representation, it should go away
	items map[string]*codeset.Codeset
}

var (
	codesetStore = CodesetStore{items: make(map[string]*codeset.Codeset)}
)

// FIXME check git content, not internal map
func (cs *CodesetStore) FindCodeset(project, name string) *codeset.Codeset {
	return cs.items[name]
}

// FIXME for all projects and labels, return codeset element
func (cs *CodesetStore) GetAllCodesets(project string, label *string) (result []*codeset.Codeset) {
	result = make([]*codeset.Codeset, 0, len(cs.items))
	for _, r := range cs.items {
		if project != "all" && r.Project != project {
			continue
		}
		if label != nil && *r.Label != *label {
			continue
		}
		result = append(result, r)
	}
	return
}

// 1. create org + new repo
// 2. register in some other store ???
func (cs *CodesetStore) AddCodeset(c *codeset.Codeset) (*codeset.Codeset, error) {
	csc, err := NewCodesetClient()
	if err != nil {
		return nil, errors.Wrap(err, "Creating codeset client failed")
	}
	err = csc.PrepareRepo(c)
	if err != nil {
		return nil, errors.Wrap(err, "Preparing Repository failed")
	}
	// Code itself needs to be pushed from client, here we could do some additional registration
	cs.items[c.Name] = c
	return c, nil
}
