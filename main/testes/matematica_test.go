package testes

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."
const erroPadraoMassa = "Valor esperado %v, mas o resultado encontrado foi %v, Massa: %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestMediaComDataSet(t *testing.T) {
	testes := []struct {
		notas         []float64
		valorEsperado float64
	}{
		{[]float64{7.2, 9.9, 6.1, 5.9}, 7.28},
		{[]float64{7.0, 6.0, 5.0, 0.0}, 4.50},
		{[]float64{7.0}, 7.0},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := Media(teste.notas...)

		if atual != teste.valorEsperado {
			t.Errorf(erroPadraoMassa, teste.valorEsperado, atual, teste)
		}
	}

}
