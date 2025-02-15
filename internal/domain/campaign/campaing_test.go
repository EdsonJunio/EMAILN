package campaign

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign X"
	content  = "Body Hi!"
	contacts = []string{"edson@gmail.com", "lucas@gmail.com"}
	fake     = faker.New()
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

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValideteContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValideteContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(2025), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValideteContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required min 1", err.Error())
}

func Test_NewCampaign_MustValideteContactsMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
