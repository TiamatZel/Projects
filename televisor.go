package main

import "fmt"

type Televisor struct {
	tipoDeEntrada int
	volumen       int
	canal         int
}

func subirVolumen(volumen, maximoVolumen int) bool {
	sePuedeSubir := false
	if volumen < maximoVolumen {
		volumen++
		sePuedeSubir = true
	}
	return sePuedeSubir
}

func bajarVolumen(volumen, maximoVolumen int) bool {
	sePuedeBajar := false
	if volumen > 0 {
		volumen--
		sePuedeBajar = true
	}
	return sePuedeBajar
}

/*func retrocederCanal(canal int) bool {
	if canal > 1 {
		canal--
	} else if canal == 1 {
		canal = getMaximoCanal()
	}
}*/

/*func avanzarCanal(canal int) bool {
	if canal < getMaximoCanal() {
		canal++
	} else if canal == getMaximoCanal() {
		canal = 1
	}
}*/

func (t Televisor) getMaximoCanal() int {
	maximoCanal := 0
	switch t.tipoDeEntrada {
	case 1:
		maximoCanal = 13
	case 2:
		maximoCanal = 100
	case 3:
		maximoCanal = 3
	}
	return maximoCanal
}

// Método equivalente a un Getter
func (t Televisor) Volumen() int {
	return t.volumen
}

// Método equivalente a un Setter
func (t *Televisor) SetVolumen(volumen int) {
	t.volumen = volumen
}

func (t Televisor) Canal() int {
	return t.canal
}

// Método equivalente a un Setter
func (t *Televisor) SetCanal(canal int) {
	t.canal = canal
}

func (t Televisor) TipoDeEntrada() string {
	rta := ""
	switch t.tipoDeEntrada {
	case 0:
		rta = "Apagado"
		t.canal = 0
		t.volumen = 0
	case 1:
		rta = "Antena"
	case 2:
		rta = "Cable"
	case 3:
		rta = "Auxiliar"
	}
	return rta
}

// Método equivalente a un Setter
func (t *Televisor) SetTipoDeEntrada(tipoDeEntrada int) {
	t.tipoDeEntrada = tipoDeEntrada
}

func main() {
	const maximoVolumen = 20
	const apagado = 0
	const antena = 1
	const cable = 2
	const auxiliar = 3
	/*tiposDeEntrada := [...]string{"apagado", "antena", "cable", "auxiliar"}
	fmt.Println(tiposDeEntrada)*/
	var t Televisor
	fmt.Scan(t.volumen)
	fmt.Scan(t.tipoDeEntrada)
	fmt.Scan(t.canal)
	fmt.Printf("%d\n", t.volumen)
	fmt.Printf("%d\n", t.tipoDeEntrada)
	fmt.Printf("%d\n", t.canal)
}
