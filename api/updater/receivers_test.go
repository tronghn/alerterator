package updater

import (
	"testing"

	"github.com/nais/alerterator/api/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestReceivers(t *testing.T) {
	t.Run("Validerer at Receivers blir opprettet riktig", func(t *testing.T) {
		receiver := createReceiver(fixtures.AlertResource)
		assert.Equal(t, fixtures.AlertResource.Name, receiver.Name)
		assert.Len(t, receiver.EmailConfigs, 1)
		assert.Len(t, receiver.SlackConfigs, 1)

		receivers := fixtures.AlertResource.Spec.Receivers
		assert.Equal(t, receivers.Email.To, receiver.EmailConfigs[0].To)
		assert.Equal(t, receivers.Email.SendResolved, receiver.EmailConfigs[0].SendResolved)

		assert.Equal(t, receivers.Slack.Channel, receiver.SlackConfigs[0].Channel)
		assert.Equal(t, receivers.Slack.PrependText, fixtures.AlertResource.Spec.Receivers.Slack.PrependText)
	})

	t.Run("Valider at send_resolved for email blir beholdt", func(t *testing.T) {
		alert := fixtures.AlertResource
		alert.Spec.Receivers.Email.SendResolved = true
		receiver := createReceiver(alert)
		assert.True(t, receiver.EmailConfigs[0].SendResolved)
	})
}
