package main

import (
	"testing"
	"time"
)

func TestHumanDate (t *testing.T) {

	tests := [] struct {
		name string
		tm time.Time
		want string
	} {
		{
			name: "utc",
			tm: time.Date(2020, time.December,17,10, 0,0,0, time.UTC),
			want: "17 Dec 2020 at 10:00",
		},
		{
			name: "empty",
			tm: time.Time{},
			want: "",
		},
		{
			name: "cet",
			tm: time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Dec 2020 at 09:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("Want %q; got %q", tt.want, hd)
			}
		})
	}


	//tm := time.Date(2020, time.December,17,10, 0,0,0, time.UTC)
	//
	//hd := humanDate(tm)
	//
	//if hd != "17 Dec 2020 at 10:00" {
	//	t.Errorf("want %q; got %q", "17 Dec 2020 at 10:00", hd)
	//}
}