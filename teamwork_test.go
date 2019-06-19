package teamwork

import (
	"testing"
)


func TestEuUriOption(t *testing.T) {

	var opt ClientOption = EuUriOption()

	table := []struct{
		uri string
		exp string
	}{
		{
			uri:"https://test.teamwork.com",
			exp: "https://test.eu.teamwork.com",
		},
		{
			uri:"https://test.teamwork.com:443",
			exp: "https://test.eu.teamwork.com:443",
		},
	}

	for _,testCase := range table {

		client := &DefaultClient{
			uri: testCase.uri,
		}

		opt(client)

		if client.uri != testCase.exp {
			t.Errorf("Uri missmatch %s != %s", client.uri, testCase.exp)
		}
	}
}