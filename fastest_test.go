package play

import "testing"

func TestFastestWorker(t *testing.T) {
	res := make(chan int)

	State{50}.Work(res)

	fastest := <-res
	if 50 != fastest {
		t.Errorf("error", fastest)
	}
	/*
		cases := []struct {
			in, want string
		}{
			{"Hello, world", "dlrow ,olleH"},
			{"Hello, 世界", "界世 ,olleH"},
			{"", ""},
		}
		for _, c := range cases {
			got := Reverse(c.in)
			if got != c.want {
				t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	*/
}
