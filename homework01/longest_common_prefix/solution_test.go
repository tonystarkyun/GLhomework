package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "common prefix exists",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "no common prefix",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "single string",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "contains empty string",
			strs: []string{"", "abc"},
			want: "",
		},
		{
			name: "empty input",
			strs: []string{},
			want: "",
		},
		{
			name: "longer shared prefix",
			strs: []string{"interspecies", "interstellar", "interstate"},
			want: "inters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Fatalf("LongestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
