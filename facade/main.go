package facade

func Main() {
	theatre := HomeTheatre{
		Popper:    &Popper{},
		Light:     &Light{},
		Projector: &Projector{},
	}

	theatre.On()
	theatre.Off()
}
