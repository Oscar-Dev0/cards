package cards

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/webp"
)

// Función para dibujar un rectángulo redondeado
func drawRoundedRect(dc *gg.Context, x, y, w, h, r float64) {
	dc.MoveTo(x+r, y)
	dc.LineTo(x+w-r, y)
	dc.QuadraticTo(x+w, y, x+w, y+r)
	dc.LineTo(x+w, y+h-r)
	dc.QuadraticTo(x+w, y+h, x+w-r, y+h)
	dc.LineTo(x+r, y+h)
	dc.QuadraticTo(x, y+h, x, y+h-r)
	dc.LineTo(x, y+r)
	dc.QuadraticTo(x, y, x+r, y)
	dc.ClosePath()
}

func loadImage(url *string) image.Image {
	if url == nil {
		return nil
	}
	resp, err := getImageBufferFromURL(*url)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	var img image.Image
	img, _, err = image.Decode(resp)
	if err != nil {
		img, err = webp.Decode(resp)
		if err != nil {
			return nil
		}
	}

	return img
}

// Función para dibujar una imagen circular
func drawCircleImage(dc *gg.Context, img image.Image, x, y, radius float64, borderColor color.RGBA, borderWidth float64) {
	// Redimensionar la imagen para que se ajuste al diámetro del círculo
	img = resizeImageToFit(img, int(radius*2), int(radius*2))

	// Dibujar el círculo de borde (antes del clip)
	dc.SetRGB(float64(borderColor.R), float64(borderColor.G), float64(borderColor.B))

	dc.SetLineWidth(borderWidth) // Configurar el ancho del borde
	dc.DrawCircle(x+radius, y+radius, radius +10)
	dc.Stroke()

	// Dibujar el contenido de la imagen dentro del círculo
	dc.DrawCircle(x+radius, y+radius, radius-borderWidth/2) // Ajustar el círculo al borde
	dc.Clip()
	dc.DrawImageAnchored(img, int(x+radius), int(y+radius), 0.5, 0.5)

	// Restablecer el clip
	dc.ResetClip()
}

// Función para cargar una fuente
func loadFont(fontPath string, size float64) *font.Face {
	font, err := gg.LoadFontFace(fontPath, size)
	if err != nil {
		return nil
	}
	return &font
}

// Función para guardar la imagen como archivo PNG
func SavePNG(buffer bytes.Buffer, filepath string) {
	outFile, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error al guardar la imagen: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()
	outFile.Write(buffer.Bytes())
}

func getImageBufferFromURL(url string) (*bytes.Buffer, error) {
	// Crear un cliente HTTP con un tiempo de espera
	client := &http.Client{
		Timeout: 10 * time.Second, // Tiempo de espera configurable
	}

	// Hacer la solicitud HTTP
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verificar el código de estado HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta HTTP inválida: %d %s", resp.StatusCode, resp.Status)
	}

	// Verificar que el tipo de contenido sea una imagen
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" || !isImageContentType(contentType) {
		return nil, fmt.Errorf("el contenido no es una imagen válida: %s", contentType)
	}

	const maxImageSize = 500 * 1024 * 1024
	limitedReader := io.LimitReader(resp.Body, maxImageSize)

	// Leer el cuerpo de la respuesta en un buffer
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, limitedReader)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	return &buffer, nil
}

// isImageContentType verifica si el tipo de contenido es de una imagen
func isImageContentType(contentType string) bool {
	return contentType == "image/jpeg" ||
		contentType == "image/png" ||
		contentType == "image/gif" ||
		contentType == "image/webp"
}

func resizeImageToFit(img image.Image, width, height int) image.Image {
	if img == nil {
		return nil
	}

	// Redimensionar la imagen para que ocupe todo el espacio disponible en el lienzo
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

func GGToBuffer(dc gg.Context) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, dc.Image())
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func setFont(dc *gg.Context, defaul string, size float64) {

	fontFace := loadFont(defaul, size)

	if fontFace != nil {
		dc.SetFontFace(*fontFace)
	}
}


func DrawStringAnchoredShadow(dc *gg.Context, text string, x, y float64, colorP color.RGBA, font string, size float64) {
	// Configurar el color del sombreado
	shadowColor := color.RGBA{0, 0, 0, 128} // Color del sombreado (negro semi-transparente)
	shadowOffset := 3.45                    // Desplazamiento del sombreado en píxeles

	// Configurar la fuente y el tamaño del texto
	setFont(dc, font, size)

	// Función para convertir color.RGBA a valores entre 0.0 y 1.0
	toRGBA := func(c color.RGBA) (r, g, b, a float64) {
		return float64(c.R) / 255.0, float64(c.G) / 255.0, float64(c.B) / 255.0, float64(c.A) / 255.0
	}

	// Dibujar el sombreado del texto
	r, g, b, a := toRGBA(shadowColor)
	dc.SetRGBA(r, g, b, a)
	dc.DrawStringAnchored(text, x+shadowOffset, y+shadowOffset, 0.5, 0.5)

	dc.SetRGB(float64(colorP.R), float64(colorP.G), float64(colorP.B))
	dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
}


