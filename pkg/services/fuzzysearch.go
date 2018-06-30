package services

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"gitlab.com/paulsevere/dff/pkg/util"
)

var cli *client.Client

type NameIDPair struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (pair *NameIDPair) Display() string {
	return pair.Name
}

func iterMatch(services []swarm.Service, query string) (nameIDMatch *NameIDPair) {
	var nameIDs []*NameIDPair = nil
	for _, s := range services {
		name := s.Spec.Name
		if fuzzy.MatchFold(query, name) {
			nameIDs = append(nameIDs, &NameIDPair{Name: name, ID: s.ID})
		}
	}
	var selection *NameIDPair

	if len(nameIDs) > 1 {
		choices := make([]util.ListOption, len(nameIDs))
		for i, item := range nameIDs {
			// _, ok := item.(util.ListOption)
			choices[i] = item
		}
		answer := util.SelectFromList("Your query has matched multiple services, please select: ", choices)

		selection, _ = answer.(*NameIDPair)

	} else {
		selection = nameIDs[0]
	}

	return selection

}

func ByName(partialName string) (sID string) {
	cli = util.Client()
	services, _ := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	match := iterMatch(services, partialName)
	sID = match.ID
	return
}
