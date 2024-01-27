package tzloc

import (
	"testing"
	"time"
)

func TestValidLocation(t *testing.T) {
	type args struct {
		loc string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid location",
			args: args{loc: "Asia/Shanghai"},
			want: true,
		},
		{
			name: "invalid location",
			args: args{loc: "Asia\\Hong_kong"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidLocation(tt.args.loc); got != tt.want {
				t.Errorf("ValidLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocationList(t *testing.T) {
	locations := GetLocationList()
	for _, loc := range locations {
		_, err := time.LoadLocation(loc)
		if err != nil {
			t.Errorf("%v is not a valid location", loc)
		}
	}
}
