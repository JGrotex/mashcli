package publichostnames

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/mmussett/mashcli/cli/app/mashcli"
	"github.com/stretchr/testify/assert"
	"testing"
)

var config = []byte(`
{
  "userid": "mmussett",
  "password": "f1rest0rm",
  "apikey": "7rn8vgpty6nywruhgc755qh6",
  "apikeysecret": "2q34GBG3nx",
  "name": "Emealocal1",
  "area": "c7e8e2d5-ff91-42eb-9885-10f2aa2cc3f5",
  "tm": "emealocal1.api.mashery.com",
  "ccurl": "https://emealocal1.admin.mashery.com/control-center"
}
`)

const (
	DOMAINS_ID = "e8f91b8b-7cc7-4203-bf45-c3498066be06"
)

var mp = &MethodParams{}

func TestGetCollection(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	a.Nil(err)

	var m *[]PublicHostnames

	m, err = GetCollection(tok, &mashcli.Params{Fields: PUBLICHOSTNAMES_ALL_FIELDS})

	a.Nil(err)

	spew.Dump(m)

}

func TestShowAllPublicHostnames(t *testing.T) {

	assert := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()

	if err != nil {
		assert.Error(err)
	}

	err = ShowAllPublicHostnames(tok)

	if err != nil {
		assert.Error(err)
	}

}

func TestExportAll(t *testing.T) {

	a := assert.New(t)

	var c mashcli.Config
	err := json.Unmarshal(config, &c)
	tok, err := c.FetchOAuthToken()
	a.Nil(err)

	err = ExportAll(tok, "/Users/markmussett/Desktop")
	a.Nil(err)
}