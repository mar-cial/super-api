package meat

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

const (
	Wagyu      string = "Wagyu"
	Prime      string = "Prime"
	Choice     string = "Choice"
	Select     string = "Select"
	Commercial string = "Commercial"
)

type Meats []Meat

func UnmarshalMeats(data []byte) (Meats, error) {
	var r Meats
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Meats) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Meat struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Grades      []string `json:"grade"`
	PriceKG     float64  `json:"price_kg"`
	ShortID     string   `json:"short_id"`
	Description string   `json:"description"`
}

func LoadMeats() ([]byte, error) {
	var meats = Meats{
		{
			ID:          uuid.NewString(),
			Name:        "Rib-eye",
			Grades:      []string{Wagyu, Prime, Choice},
			PriceKG:     500.00,
			ShortID:     uuid.NewString()[:5],
			Description: "Si te has preguntado cuál es el corte de carne más suave y jugoso, el Rib-Eye es uno de ellos. Este corte se obtiene del costillar de la res y no tiene hueso, ideal para disfrutar a la parrilla. Su punto perfecto de cocción es tres cuartos o término medio, de esta manera al morderlo sentirás la suavidad y jugosidad que lo caracterizan.",
		},
		{
			ID:          uuid.NewString(),
			Name:        "Tomahawk",
			Grades:      []string{Prime, Choice, Select},
			PriceKG:     600.50,
			ShortID:     uuid.NewString()[:5],
			Description: "Entre los cortes de carnes más ricos sin dudas se encuentra el Tomahawk. Esta pieza es sacada del mismo lugar que se obtiene el Rib-Eye, la diferencia es que a este se le deja un hueso al que se le quita la grasa en su totalidad para obtener la forma que lo ha hecho tan conocido. La mejor forma de cocinarlo es a la parrilla a fuego indirecto por unos 10 minutos.",
		},
		{
			ID:          uuid.NewString(),
			Name:        "T-bone",
			Grades:      []string{Wagyu, Prime, Choice},
			PriceKG:     700.25,
			ShortID:     uuid.NewString()[:5],
			Description: "Muchos se preguntan cuáles son los cortes americanos, pues el anterior y este están en la lista. El T-Bone es un filete en el que vas a disfrutar de una parte de solomillo y otra de entrecot al ser este un corte de carne de res que se obtiene de la parte media del lomo.  Además, este es una de las piezas más jugosas y por tanto, de las que mejor quedan en una parrilla de pallets.",
		},
		{
			ID:          uuid.NewString(),
			Name:        "New York",
			Grades:      []string{Wagyu, Prime, Choice, Select},
			PriceKG:     550.50,
			ShortID:     uuid.NewString()[:5],
			Description: "Esta pieza es alargada y suele venderse en trozos de más de dos centímetros de espesor por su maravilloso marmoleo y concentrado sabor. Lo encuentras en el lomo de la res, por lo que su concentración de grasa es alta. Estos cortes de carne son perfectos para el asador por su suavidad",
		},
		{
			ID:          uuid.NewString(),
			Name:        "Arrachera",
			Grades:      []string{Prime, Choice, Select, Commercial},
			PriceKG:     250.00,
			ShortID:     uuid.NewString()[:5],
			Description: "Esta es la excepción a la regla de los cortes de carne que van en el asador pues proviene de la falda de la res, es decir, de la pancita que cuelga. La arrachera corresponde a una parte fibrosa por lo que debe marinarse y ablandarse para que sepa rico y tenga buena textura. Hay dos opciones: la outside o capa exterior de la caja toráxica, que está en las entrañas de la res y la que tiene un poco más de grasa y la que se ubica más abajo llamada inside; es más fibrosa y por eso se marina.",
		},
		{
			ID:          uuid.NewString(),
			Name:        "Picaña",
			Grades:      []string{Choice, Select, Commercial},
			PriceKG:     300.00,
			ShortID:     uuid.NewString()[:5],
			Description: "Justo donde termina el lomo y comienzan lo que los carniceros llaman cuartos traseros hay una pieza magra cubierta por un lado de una capa de grasa. Conocida en Estados Unidos como Top Sirloin o tri trap -aunque con algunas variaciones en su presentación- la picaña es una parte inferior del lomo que bien vale la pena ponerla en el asador. Da click al video y pruébala",
		},
		{
			ID:          uuid.NewString(),
			Name:        "Cowboy",
			Grades:      []string{Prime, Choice, Select},
			PriceKG:     400.00,
			ShortID:     uuid.NewString()[:5],
			Description: "Tiene el mismo origen que el Tomahawk pero se diferencia en el largo del hueso. Este corte de carne es magro con mucho marmoleo y grasita sabrosa que se obtiene con todo y hueso. La magia de ambos consiste en la concentración de jugos que causa cocinarlo a la vuelta y vuelta aunque también tenemos un secreto para eso: cualquier corte debe reposar antes de cortarlo por lo menos el 30% del tiempo que estuvo en cocción para que no se deshidrate.",
		},
	}

	meatBytes, err := meats.Marshal()
	if err != nil {
		fmt.Println("could not marshal meats")
		return nil, err
	}

	return meatBytes, nil
}
