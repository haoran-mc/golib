package timeutil

import (
	"fmt"
	"testing"
	"time"
)

// go test -v
func TestAverage(t *testing.T) {
	cases := []struct {
		solar string
		lunar string
	}{
		{"20010517", "辛巳蛇年四月廿五"},
		{"20250622", "乙巳蛇年五月廿七"},
	}
	for _, c := range cases {
		t.Run(c.solar, func(t *testing.T) {
			if ans := Lunar(c.solar); ans != c.lunar {
				t.Fatalf("%s expected %s, but %s got",
					c.solar, c.lunar, ans)
			}
		})
	}
	fmt.Println("Now lunar: ", Lunar(time.Now().Format("20060102")))
}
