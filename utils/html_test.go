package utils

import (
	"testing"
)

func TestStripHTMLTags(t *testing.T) {
	testCases := []struct {
		htmlInput string
		expect    string
	}{
		{htmlInput: "<html><body<h1>Hello World!</h1></body></html>",
			expect: "Hello World!"},
		{htmlInput: "<h1>Hello World!</h1>",
			expect: "Hello World!"},
		{htmlInput: "<h1 style='color: #5e9ca0';>Hello World!</h1>",
			expect: "Hello World!"},
		{htmlInput: "<td><img style='margin: 1px 15px;' src='images/smiley.png' alt='laughing' width='40' height='16' /><strong>Hello World!</strong></td>", expect: "Hello World!"},
	}
	for _, tc := range testCases {
		got := StripHTMLTags(tc.htmlInput)
		expect := tc.expect

		if got != expect {
			t.Errorf("Expected %s, got %s", expect, got)
		}
	}
}
