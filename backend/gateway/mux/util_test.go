package mux

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBrowser(t *testing.T) {
	testCases := []struct {
		expect bool
		header string
	}{
		{expect: true, header: "text/html"},
		{expect: true, header: "text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, */*;q=0.8"},
		{expect: true, header: "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		{expect: false, header: "*/*"},
		{expect: false, header: "application/json, text/plain, */*"},
	}

	for _, tc := range testCases {
		h := http.Header{}
		h.Add("Accept", tc.header)
		assert.Equal(t, tc.expect, isBrowser(h))
	}
}

// based on http.response private member.
type mockResponseWriter struct {
	http.ResponseWriter

	req *http.Request
}

func TestCopyHeadersFromrResponseWriter(t *testing.T) {
	headers := http.Header{}
	headers.Add("foo", "bar")
	headers.Add("Accept", "text/html")
	m := &mockResponseWriter{req: &http.Request{Header: headers}}

	ret := copyRequestHeadersFromResponseWriter(m)
	assert.EqualValues(t, headers, ret)
}
