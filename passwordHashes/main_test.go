// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package main

import "testing"

func TestHash(t *testing.T) {
	s := "password"
	r := "5f4dcc3b5aa765d61d8327deb882cf99"
	h := PassHash(s)
	if h != r {
		t.Errorf("Expected hash of \"%s\" to be %s, but got %s", s, r, h)
	}

	s = "abc123"
	r = "e99a18c428cb38d5f260853678922e03"
	h = PassHash(s)
	if h != r {
		t.Errorf("Expected hash of \"%s\" to be %s, but got %s", s, r, h)
	}

	s = ""
	r = "d41d8cd98f00b204e9800998ecf8427e"
	h = PassHash(s)
	if h != r {
		t.Errorf("Expected hash of \"%s\" to be %s, but got %s", s, r, h)
	}
}
