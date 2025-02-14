package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	//arrange organizacao do código
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"edson@gmail.com", "lucas@gmail.com"}

	//action acao
	campaign := NewCampaign(name, content, contacts)

	// assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_CreateOnIsNotNoll(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"edilson@gmail.com", "pucas@gmail.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}
