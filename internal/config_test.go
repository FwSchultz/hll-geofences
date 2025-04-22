package internal_test

import (
	"github.com/floriansw/go-hll-rcon/rconv2/api"
	"github.com/floriansw/hll-geofences/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log/slog"
	"os"
)

var _ = Describe("Config", func() {
	Describe("Persistence", func() {
		It("persists config change of server", func() {
			l := slog.New(slog.NewTextHandler(os.Stdout, nil))
			f, err := os.CreateTemp(os.TempDir(), "config")
			Expect(err).ToNot(HaveOccurred())
			defer os.Remove(f.Name())
			Expect(os.WriteFile(f.Name(), []byte("{}"), 0655)).ToNot(HaveOccurred())
			c, err := internal.NewConfig(f.Name(), l)
			Expect(err).ToNot(HaveOccurred())

			Expect(c.Save()).ToNot(HaveOccurred())
			c, err = internal.NewConfig(f.Name(), l)
			Expect(err).ToNot(HaveOccurred())

			Expect(c.Save()).ToNot(HaveOccurred())

			c, err = internal.NewConfig(f.Name(), l)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Fence", func() {
		It("returns false when not includes", func() {
			Expect(internal.Fence{
				X:       Pointer("G"),
				Y:       Pointer(4),
				Numpads: []int{4, 5, 6},
			}.Includes(api.Grid{X: "H", Y: 1, Numpad: 7})).To(BeFalse())
		})

		It("includes fence when direct match", func() {
			Expect(internal.Fence{
				X:       Pointer("G"),
				Y:       Pointer(4),
				Numpads: []int{4, 5, 6},
			}.Includes(api.Grid{X: "G", Y: 4, Numpad: 5})).To(BeTrue())
		})

		It("includes fence when matching whole line X-axis", func() {
			Expect(internal.Fence{
				X:       Pointer("G"),
				Numpads: []int{4, 5, 6},
			}.Includes(api.Grid{X: "G", Y: 8, Numpad: 5})).To(BeTrue())
		})

		It("includes fence when matching whole line Y-axis", func() {
			Expect(internal.Fence{
				Y:       Pointer(5),
				Numpads: []int{4, 5, 6},
			}.Includes(api.Grid{X: "A", Y: 5, Numpad: 5})).To(BeTrue())
		})

		It("includes fence when no numpad specified", func() {
			Expect(internal.Fence{
				X: Pointer("G"),
				Y: Pointer(5),
			}.Includes(api.Grid{X: "G", Y: 5, Numpad: 7})).To(BeTrue())
		})
	})
})

func Pointer[T any](v T) *T {
	return &v
}
