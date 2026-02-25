package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "palindrome odd",
			x:    121,
			want: true,
		},
		{
			name: "negative number",
			x:    -121,
			want: false,
		},
		{
			name: "ends with zero",
			x:    10,
			want: false,
		},
		{
			name: "zero",
			x:    0,
			want: true,
		},
		{
			name: "palindrome even",
			x:    1221,
			want: true,
		},
		{
			name: "not palindrome",
			x:    123,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.x)
			if got != tt.want {
				t.Fatalf("IsPalindrome(%d) = %t, want %t", tt.x, got, tt.want)
			}
		})
	}
}
