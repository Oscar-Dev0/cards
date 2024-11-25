package cards

import (
	"image/color"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)


type ColorType interface{}

// ResolvedColor resuelve y devuelve el valor numérico del color correspondiente.
//
// Parámetros:
//   - color: El color representado como un string, un array de int o un int.
//
// Devuelve:
//   - El valor numérico del color.
func ResolvedColor(color ColorType) int {
	switch v := color.(type) {
	case string:
		if v == "Random" {
			return rand.Intn(0xffffff + 1) // Genera un color aleatorio.
		} else if v == "Default" {
			return 0 // Retorna negro por defecto.
		} else if regexp.MustCompile(`^#?[\da-fA-F]{6}$`).MatchString(v) { // Verifica formato HEX.
			c, err := strconv.ParseInt(strings.TrimPrefix(v, "#"), 16, 32) // Convierte HEX a int.
			if err != nil {
				return 0 // Si hay un error, retorna negro.
			}
			return int(c)
		} else if val, ok := ColorsList[v]; ok { // Busca el color en ColorsList.
			return val
		} else {
			return 0 // Si no coincide, retorna negro.
		}
	case []int:
		if len(v) != 3 {
			return 0 // El array debe tener exactamente 3 valores (RGB).
		}
		return (v[0] << 16) + (v[1] << 8) + v[2] // Convierte RGB a entero.
	case int:
		if v < 0 || v > 0xffffff {
			return 0 // El valor debe estar en el rango de colores válidos.
		}
		return v
	default:
		return 0 // Si el tipo no es válido, retorna negro.
	}
}

// ResolvedRGB resuelve y devuelve el color en formato RGB.
//
// Parámetros:
//   - color: El color representado como un string, un array de int o un int.
//
func ResolvedRGB(colr ColorType) color.RGBA {
	hex := ResolvedColor(colr)
	return color.RGBA{
		R: uint8((hex >> 16) & 0xff),
		G: uint8((hex >> 8) & 0xff),
		B: uint8(hex & 0xff),
		A: 0xff, // Opacidad completa por defecto
	}
}