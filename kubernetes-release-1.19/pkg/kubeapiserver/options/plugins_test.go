package options

import (
	"strings"
	"testing"
)

func TestAdmissionPluginOrder(t *testing.T) {
	// Ensure the last four admission plugins listed are webhooks, quota, and deny
	allplugins := strings.Join(AllOrderedPlugins, ",")
	expectSuffix := ",MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,AlwaysDeny"
	if !strings.HasSuffix(allplugins, expectSuffix) {
		t.Fatalf("AllOrderedPlugins must end with ...%s", expectSuffix)
	}
}
