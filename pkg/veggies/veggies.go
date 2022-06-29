package veggies

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Vegetables []Vegetable

func UnmarshalVegetables(data []byte) (Vegetables, error) {
	var r Vegetables
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Vegetables) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Vegetable struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	PriceKg     float64     `json:"price_kg"`
	ShortID     string      `json:"short_id"`
	Description Description `json:"description"`
}

type Description struct {
	Componentes []string `json:"componentes"`
	Beneficios  string   `json:"beneficios"`
}

func LoadVeggies() ([]byte, error) {

	var veggies = Vegetables{
		{
			ID:      uuid.NewString(),
			Name:    "Ajo",
			Image:   "ajo",
			PriceKg: 50.50,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"agua y carbohidratos",
					"vitamina B y C",
					"calcio",
					"fósforo",
					"magnesio",
					"potasio",
				},
				Beneficios: "Alivia la gripa, calma dolor de oído y de muelas, es digestivo, protege la piel y reduce los niveles de colesterol malo.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Alcachofa",
			Image:   "alcachofa",
			PriceKg: 74.30,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"vitaminas del complejo B",
					"minerales como calcio",
					"cobre",
					"fibra",
					"hierro",
					"manganeso",
					"potasio",
					"compuestos antioxidantes",
				},
				Beneficios: "Estimula el apetito, facilita la digestión, protege la salud celular y equilibra el colesterol.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Brócoli",
			Image:   "brocoli",
			PriceKg: 40.60,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"vitaminas A, B, C, E y K",
					"betacaroteno",
					"luteína",
					"niacina",
					"riboflavina",
					"tiamina",
					"zeaxantina",
				},
				Beneficios: "Controla los niveles de colesterol, facilita el proceso digestivo, incrementa la resistencia a la insulina y mejora la función metabólica.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Camote",
			Image:   "camote",
			PriceKg: 25.50,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{

					"fibra",
					"vitaminas A, C y B6",
					"calcio",
					"fósforo",
					"magnesio",
					"potasio",
					"antioxidantes",
				},
				Beneficios: "Alivia los dolores de la artritis, estimula el sistema inmune, nivel el azúcar en la sangre y mejora la digestión.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Cebolla",
			Image:   "cebolla",
			PriceKg: 40.00,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"es rica en vitaminas A, B y C",
					"calcio",
					"magnesio",
					"fósforo",
					"hierro",
					"potasio",
				},
				Beneficios: "Alivia malestares estomacales, aumenta las defensas, calma la tos, combate infecciones bacterianas y fúngicas y fortalece la salud cutánea.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Espárragos",
			Image:   "esparragos",
			PriceKg: 120.20,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"bajo en calorías",
					"2.2 g de proteínas",
					"fibra",
					"vitaminas A, C, y E",
					"fósforo",
					"potasio",
				},
				Beneficios: "Aumenta la sensación de saciedad, estimula el desarrollo de bacterias intestinales benéficas, mantiene bajo control los niveles de presión arterial y combate el impacto de los radicales libres.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Espinacas",
			Image:   "espinacas",
			PriceKg: 60.50,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"rica en vitaminas A, E, y K",
					"minerales",
					"antioxidantes",
				},
				Beneficios: "Estimula los movimientos intestinales, provoca saciedad, fortalece la piel, protege la visión y regula los niveles de presión arterial.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Kale",
			Image:   "kale",
			PriceKg: 102.59,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"vitaminas A, B, C, y K",
					"calcio",
					"hierro",
					"potasio",
					"fibra dietética",
					"compuestos antioxidantes",
				},
				Beneficios: "Controla los niveles de azúcar en la sangre, disminuye la presión arterial, facilita el control de peso, fortalece la piel y el cabello y protege la salud gastrointestinal.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Jitomate",
			Image:   "jitomate",
			PriceKg: 90.9,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"rico en vitaminas A, B, C, E  y K",
					"cobre",
					"fósforo",
					"magnesio",
					"potasio",
					"gran contenido en agua"},
				Beneficios: "Evita el envejecimiento de la piel, reduce los niveles de colesterol, estimula la liberación de jugos gástricos y también favorece la absorción de nutrientes.",
			},
		},
		{
			ID:      uuid.NewString(),
			Name:    "Zanahoria",
			Image:   "zanahoria",
			PriceKg: 160.7,
			ShortID: uuid.NewString()[:5],
			Description: Description{
				Componentes: []string{
					"vitaminas A, C y D",
					"calcio",
					"hierro",
					"potasio",
				},
				Beneficios: "Calma problemas gastrointestinales, combate los efectos de los radicales libres, fortalece el tejido celular y además retrasa los síntomas del envejecimiento.",
			},
		},
	}

	vegBytes, err := veggies.Marshal()
	if err != nil {
		fmt.Println("error marshaling veggies: ", err)
		return nil, err
	}

	return vegBytes, nil
}
