package grand

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestColor(t *testing.T) {
	_assert := assert.New(t)

	// RGB
	rgb := Color(ColorRGB)
	matchRgb, _ := regexp.MatchString(
		`rgb\(\d{1,3},\d{1,3},\d{1,3}\)`,
		rgb)
	fmt.Println(rgb)
	_assert.True(matchRgb)

	// RGBA
	rgba := Color(ColorRGBA)
	matchRgba, _ := regexp.MatchString(
		`rgba\(\d{1,3},\d{1,3},\d{1,3},[1|0]\.\d{1,2}\)`,
		rgba)
	fmt.Println(rgba)
	_assert.True(matchRgba)

	// HEX
	matchHex, _ := regexp.MatchString(
		`#[a-zA-Z0-9]{3,6}`,
		Color(ColorHEX))
	_assert.True(matchHex)

	// HSL
	matchHsl, _ := regexp.MatchString(
		`hsl\(\d{1,3},\d{1,3}%,\d{1,3}%\)`,
		Color(ColorHSL))
	_assert.True(matchHsl)

	// HSLA
	matchHsla, _ := regexp.MatchString(
		`hsla\(\d{1,3},\d{1,3}%,\d{1,3}%,[1|0]\.\d{1,2}\)`,
		Color(ColorHSLA))
	_assert.True(matchHsla)
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color(ColorRGB)
		Color(ColorRGBA)
		Color(ColorHEX)
		Color(ColorNAME)
	}
}
