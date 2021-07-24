package bungie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"root/destinyhome/bungie/responses"
)

type call struct {
	s         *service
	method    string
	url       string
	component int
}

func (c *call) Component(comp int) *call {
	return c
}

func (c *call) Do() (*responses.OuterRes, error) {

	const operation = "call.Do"

	// Add the component to the url.
	c.url += strconv.Itoa(c.component)

	// Create the request.
	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		return nil, errors.Wrap(err, operation)
	}

	var res responses.OuterRes

	if c.component >= 200 && c.component < 300 {

	} else if c.component >= 300 && c.component < 400 {
	} else {
	}

	// Send the request.
	if err := c.send(req, &res); err != nil {
		return nil, errors.Wrap(err, operation)
	}

	return nil, nil
}

func (c *call) send(req *http.Request, v1, v2 interface{}) error {

	const operation = "destiny2Service.sendRequest"

	// Set headers.
	req.Header.Set("Content-Type", "application/json")

	// Do the request.
	httpRes, err := c.s.client.Do(req)
	if err != nil {
		return errors.Wrap(err, operation)
	}
	defer httpRes.Body.Close()

	// Decode the response body.
	var body json.RawMessage
	err = json.NewDecoder(httpRes.Body).Decode(body)

	return errors.Wrap(err, operation)
}

type destiny2Service struct {
	s *service
}

func (d2s *destiny2Service) GetCharacter(memType, memID, characterID string) *call {
	return &call{
		s:      d2s.s,
		method: "GET",
		url: fmt.Sprintf("%sDestiny2/%s/Profile/%s/Character/%s",
			d2s.s.basePath,
			memType,
			memID,
			characterID),
	}
}

func (d2s *destiny2Service) GetDestinyEntityDefinition(entityType, hashIdentifier string) *call {
	return &call{
		s:      d2s.s,
		method: "GET",
		url: fmt.Sprintf("%sDestiny2/Manifest/%s/%s",
			d2s.s.basePath,
			entityType,
			hashIdentifier),
	}
}
