package prose

import "testing"

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{[]string{}, ""},
		{[]string{"my parents"}, "my parents"},
		{[]string{"apple", "orange"}, "apple and orange"},
		{[]string{"apple", "orange", "pear"}, "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}
