package messagedb

import (
	"testing"

	"github.com/influenzanet/messaging-service/pkg/types"
)

func TestOutgoingEmailsDB(t *testing.T) {
	t.Run("add outgoing emails", func(t *testing.T) {
		counter := 0
		for counter < 13 {
			_, err := testDBService.AddToOutgoingEmails(testInstanceID, types.OutgoingEmail{
				To:          []string{"test@example.org"},
				MessageType: "test",
				Subject:     "test",
				Content:     "<h1>test</h1>",
			})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			counter += 1
		}
	})

	t.Run("fetch outgoing emails", func(t *testing.T) {
		resp, err := testDBService.FetchOutgoingEmails(testInstanceID, 10)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if len(resp) != 10 {
			t.Errorf("unexpected number of emails found: %d", len(resp))
			return
		}
		// again:
		resp, err = testDBService.FetchOutgoingEmails(testInstanceID, 10)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if len(resp) != 3 {
			t.Errorf("unexpected number of emails found: %d", len(resp))
			return
		}
	})
}
