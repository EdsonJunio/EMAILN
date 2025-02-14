package campaign

import (
	"emailn/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Bady",
		Emails:  []string{"test@test.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_create_SaveCampaign(t *testing.T) {

	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Bady",
		Emails:  []string{"test@test.com"},
	}
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
