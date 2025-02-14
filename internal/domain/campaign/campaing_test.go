package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"edson@gmail.com", "lucas@gmail.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_CreateOnIsNotNoll(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValideteContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValideteContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}
