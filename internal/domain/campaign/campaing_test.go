package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	//arrange organizacao do código
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"edson@gmail.com", "lucas@gmail.com"}

	//action acao
	campaign := NewCampaign(name, content, contacts)

	// assert
	assert.Equal(campaign.ID, "2")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}
