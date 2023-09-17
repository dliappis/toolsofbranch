package pkg

import "testing"

func TestEffectiveBranch(t *testing.T) {
	type args struct {
		gitBranch string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Main is treated as largest semver", args{gitBranch: "main"}, maxVer},
		{"Test valid major.minor", args{gitBranch: "8.1"}, "8.1"},
		{"Test valid major.minor.patch", args{gitBranch: "7.17.13"}, "7.17.13"},
		{"Test feature branch", args{gitBranch: "feature/branch/8.1"}, maxVer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EffectiveBranch(tt.args.gitBranch); got != tt.want {
				t.Errorf("EffectiveBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}
