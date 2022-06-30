package utils

import (
	"testing"
)

func TestIsOnion(t *testing.T) {

	// proper onion v3 address
	test := IsOnion("testq3qwbiwdtitdydemcez56oenag25gdym4xmhraoqbsmrsrqnagad.onion")
	if !test {
		t.Errorf("%s should have matched!", "testq3qwbiwdtitdydemcez56oenag25gdym4xmhraoqbsmrsrqnagad.onion")
	}

	// proper onion v3 address with subdomains
	test = IsOnion("subdomain.lots.of.test3uasslvk65oxtcbrane7bqvuqbv5wuaomu4gmf3wjxi7c5kxgiad.onion")
	if !test {
		t.Errorf("%s should have matched!", "subdomain.lots.of.test3uasslvk65oxtcbrane7bqvuqbv5wuaomu4gmf3wjxi7c5kxgiad.onion")
	}

	// address with charset outside of specified encoding
	test = IsOnion("thisis9notavalidonionbecauseitcontainsinvalidenc1letters.onion")
	if test {
		t.Errorf("%s should not have matched!", "thisis9notavalidonionbecauseitcontainsinvalidenc1letters.onion")
	}

	// does not end in .onion
	test = IsOnion("maschqwtr4lqt2pj.com")
	if test {
		t.Errorf("%s should not have matched!", "maschqwtr4lqt2pj.com")
	}

	// address too long
	test = IsOnion("www.thisisnotanonionitistoolongwwwwwwwwwwwwwwwwwwwwwwwwww.onion")
	if test {
		t.Errorf("%s should not have matched!", "www.thisisnotanonionitistoolongwwwwwwwwwwwwwwwwwwwwwwwwww.onion")
	}

}
