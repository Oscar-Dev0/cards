package cards

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

var (
	canvasWidth int  = 1260
	canvasHeight int = 620
	defaultFont string = "./resources/fonts/Arial.ttf"
	defaultTitle string = "¡Bienvenido al Servidor!"
	defaultUser string = "Usuario123"
	defaultDesc string  = "¡Gracias por unirte a nuestra comunidad!"
	defaultAvatarURL string = "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/9c64cfe3-bb3b-4ae8-b5a6-d2f39d21ff87/d3jme6i-8c702ad4-4b7a-4763-9901-99f8b4f038b0.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzljNjRjZmUzLWJiM2ItNGFlOC1iNWE2LWQyZjM5ZDIxZmY4N1wvZDNqbWU2aS04YzcwMmFkNC00YjdhLTQ3NjMtOTkwMS05OWY4YjRmMDM4YjAucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.oQC1FIUxsmeyLHm6qNdoRb8wzoMdKI1p49kNBstoU-w"
	defaultBGURL string    = "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/9c64cfe3-bb3b-4ae8-b5a6-d2f39d21ff87/d3jme6i-8c702ad4-4b7a-4763-9901-99f8b4f038b0.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzljNjRjZmUzLWJiM2ItNGFlOC1iNWE2LWQyZjM5ZDIxZmY4N1wvZDNqbWU2aS04YzcwMmFkNC00YjdhLTQ3NjMtOTkwMS05OWY4YjRmMDM4YjAucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.oQC1FIUxsmeyLHm6qNdoRb8wzoMdKI1p49kNBstoU-w"
)

type MemberCard struct {
	Title       *string
	Description *string
	Background  *string
	Box         *bool
	User        *string
	UserAvatar  *string
	CountMember *int
	Colors      Colors
}

type Colors struct {
	Title       color.RGBA
	Description color.RGBA
	User        color.RGBA
	CountMember color.RGBA
	Box         color.RGBA
}



func NewMemberCard() *MemberCard {
	return &MemberCard{
		Title:       stringPtr(defaultTitle),
		Description: stringPtr(defaultDesc),
		Background:  stringPtr(defaultBGURL),
		User:        stringPtr(defaultUser),
		UserAvatar:  stringPtr(defaultAvatarURL),
		Box:         boolPtr(false),
		Colors: Colors{
			Title:       color.RGBA{1, 1, 1, 255},
			Description: color.RGBA{1, 1, 1, 255},
			User:        color.RGBA{1, 1, 1, 255},
			CountMember: color.RGBA{1, 1, 1, 255},
			Box:         color.RGBA{0, 0, 0, 128}, // Transparente
		},
	}
}

func (c *MemberCard) SetTitle(title string, color ColorType) *MemberCard {
	c.Title = &title
	if color != nil {
		c.Colors.Title = ResolvedRGB(color)
	}
	return c
}

func (c *MemberCard) SetUser(user string, color ColorType) *MemberCard {
	c.User = &user
	if color != nil {
		c.Colors.User = ResolvedRGB(color)
	}
	return c
}

func (c *MemberCard) SetDescription(description string, color ColorType, font *string) *MemberCard {
	c.Description = &description
	if color != nil {
		c.Colors.Description = ResolvedRGB(color)
	}
	return c
}

func (c *MemberCard) SetCountMember(count int, color ColorType, font *string) *MemberCard {
	c.CountMember = &count
	if color != nil {
		c.Colors.CountMember = ResolvedRGB(color)
	}
	return c
}

func (c *MemberCard) SetBackground(background string) *MemberCard {
	c.Background = &background
	return c
}

func (c *MemberCard) SetBox(box bool, color ColorType) *MemberCard {
	c.Box = &box
	if color != nil {
		c.Colors.Box = ResolvedRGB(color)
	}
	return c
}

func (c *MemberCard) Buffer() (*bytes.Buffer, error) {
	// Crear el canvas
	dc := gg.NewContext(canvasWidth, canvasHeight)

	// Fondo redondeado y clip
	radius := 30.0
	drawRoundedRect(dc, 0, 0, float64(canvasWidth), float64(canvasHeight), radius)
	dc.Clip()

	// Dibujar imagen de fondo
	backgroundImg := loadImage(c.Background)
	if backgroundImg == nil {
		backgroundImg = loadImage(&defaultBGURL)
	}
	if backgroundImg != nil {
		backgroundImg = resizeImageToFit(backgroundImg, canvasWidth, canvasHeight)
		dc.DrawImage(backgroundImg, 0, 0)
	}

	if c.Box != nil && *c.Box {
		dc.SetRGBA(float64(c.Colors.Box.R),float64(c.Colors.Box.G) , float64(c.Colors.Box.B), 0.5) // Color y opacidad

		drawRoundedRect(dc, 63, 50, 1134, 520, 10)
		dc.Fill()
		if c.CountMember != nil && *c.CountMember > 0 {
			dc.SetRGB(float64(c.Colors.Title.R), float64(c.Colors.Title.G), float64(c.Colors.Title.B))
			setFont(dc, defaultFont, 35)
			dc.DrawStringAnchored(fmt.Sprintf("#%v", c.CountMember), 1150, 120, 1, 0.5)
		}
	} else {
		if c.CountMember != nil && *c.CountMember > 0 {

			dc.SetRGB(float64(c.Colors.Title.R), float64(c.Colors.Title.G), float64(c.Colors.Title.B))
			setFont(dc, defaultFont, 35)
			dc.DrawStringAnchored(fmt.Sprintf("#%v", c.CountMember), 1150, 120, 1, 0.5)
		}
	}

	// Dibujar avatar
	var avarI image.Image
	if c.UserAvatar != nil && *c.UserAvatar != "" {
		avarI = loadImage(c.UserAvatar)
	}
	if avarI == nil {
		avarI = loadImage(&defaultBGURL)
	}

	if avarI != nil {
		drawCircleImage(dc, avarI, 520, 110, 110, c.Colors.User, 10)
	}

	// Dibujar título
	dc.SetRGB(float64(c.Colors.Title.R), float64(c.Colors.Title.G), float64(c.Colors.Title.B))                                            // Blanco
	setFont(dc,  defaultFont, 80) // Fuente grande
	title := defaultTitle
	if c.Title != nil && *c.Title != "" {
		title = *c.Title
	}
	dc.DrawStringAnchored(title, 642, 430, 0.5, 0.5)

	// Dibujar nombre de usuario
	dc.SetRGB(float64(c.Colors.User.R), float64(c.Colors.User.G), float64(c.Colors.User.B))
	setFont(dc,  defaultFont, 45)
	username := defaultUser // Nombre de usuario
	if c.User != nil && *c.User != "" {
		username = *c.User
	}

	dc.DrawStringAnchored(username, 642, 490, 0.5, 0.5)

	// Dibujar descripción)
	dc.SetRGB(float64(c.Colors.Description.R), float64(c.Colors.Description.G), float64(c.Colors.Description.B))
	setFont(dc,  defaultFont, 35)
	description := defaultDesc // Descripción
	if c.Description != nil && *c.Description != "" {
		description = *c.Description
	}
	dc.DrawStringAnchored(description, 642, 540, 0.5, 0.5)

	return GGToBuffer(*dc)
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
