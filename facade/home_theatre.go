package facade

import "fmt"

type Popper struct{}

func (p *Popper) On() {
	fmt.Println("pop on")
}

func (p *Popper) Off() {
	fmt.Println("pop of")
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("light on")
}

func (l *Light) Off() {
	fmt.Println("light off")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("projector on")
}

func (p *Projector) Off() {
	fmt.Println("projector off")
}

type HomeTheatre struct {
	*Popper
	*Light
	*Projector
}

func (t *HomeTheatre) On() {
	t.Popper.On()
	t.Light.Off()
	t.Projector.On()
	t.Popper.Off()
}

func (t *HomeTheatre) Off() {
	t.Light.On()
	t.Projector.Off()
}
