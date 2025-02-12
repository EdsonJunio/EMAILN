package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"edson@gmail.com", "lucas@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "2" {
		t.Errorf("Expected 1")
	} else if campaign.Name != name {
		t.Errorf("Expected correct name ")
	} else if campaign.Content != content {
		t.Errorf("Expected correct content")
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected correct contacts")
	}

}
