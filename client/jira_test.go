package client

import (
	"github.com/innogames/slack-bot/bot/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJira(t *testing.T) {
	t.Run("no credentials", func(t *testing.T) {
		cfg := config.Jira{
			Host: "https://jira.example.com",
		}
		client, err := GetJiraClient(cfg)

		assert.Nil(t, err)
		assert.Equal(t, "jira.example.com", client.GetBaseURL().Host)
	})

	t.Run("with credentials", func(t *testing.T) {
		cfg := config.Jira{
			Host:     "https://jira.example.com",
			Username: "foo",
			Password: "bar",
		}
		client, err := GetJiraClient(cfg)

		assert.Nil(t, err)
		assert.False(t, client.Authentication.Authenticated())
	})
}
