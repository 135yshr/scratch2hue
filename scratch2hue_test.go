package scratch2hue

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestScaratch2Hue(t *testing.T) {
	Describe(t, "Scratch2Hue TestCase", func() {
		Context("connect to hue", func() {
			It("create new conneciton.", func() {
				actual, err := NewConnection("localhost")
				Expect(err).To(NotExist)
				Expect(actual).To(Exist)
			})
		})
	})
}
