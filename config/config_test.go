package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfiguration tests loading the example config
func TestLoadConfiguration(t *testing.T) {
	exampleConfig, err := LoadConfiguration("../config-example.yaml")

	assert.NoError(t, err)

	expectedFeeds := []string{
		"https://aws.amazon.com/blogs/security/feed/",
		"https://www.phoronix.com/rss.php",
		"https://www.zdnet.com/topic/security/rss.xml",
	}

	assert.Equal(
		t,
		expectedFeeds,
		exampleConfig.Feeds,
		"Expected configuration does not match.",
	)

	// ExternalViewer should default to either 'xdg-open' on Linux,
	// or 'open' on macOS
	assert.Contains(
		t,
		[]string{"xdg-open", "open"},
		exampleConfig.ExternalViewer,
	)
}
