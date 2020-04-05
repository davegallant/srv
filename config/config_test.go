package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfiguration tests loading the example config
func TestLoadConfiguration(t *testing.T) {
	exampleConfig := LoadConfiguration("../config-example.yaml")

	expectedFeeds := []string{
		"https://news.ycombinator.com/rss",
		"https://www.reddit.com/r/golang/.rss",
		"https://www.reddit.com/r/linux/.rss",
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
