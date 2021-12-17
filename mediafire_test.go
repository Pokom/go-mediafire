package go_mediafire

import "testing"

func Test_findFileName(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Base Case",
			args: args{
				content: `attachment; filename="1997-12-11 - Rochester War Memorial - Rochester, NY.rar"`,
			},
			want: "1997-12-11 - Rochester War Memorial - Rochester, NY.rar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFileName(tt.args.content); got != tt.want {
				t.Errorf("findFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
