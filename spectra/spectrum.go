package spectra

import (
	"fmt"
	"reflect"
)

type Spectra struct {
	data []Spectrum
}

func NewSpectrumFromArray(data []string) (*Spectrum, error) {
	if len(data) < 292 {
		return nil, fmt.Errorf("insufficient data: got %d, need 292", len(data))
	}
	s := &Spectrum{}
	v := reflect.ValueOf(s).Elem()

	// Iterate over teh remaining fields to populate spectrum
	for i := 2; i < v.NumField() && i < len(data); i++ {
		// TODO Capture all items into type
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}
		dataValue := reflect.ValueOf(data[i])
		if dataValue.Type().ConvertibleTo(field.Type()) {
			field.Set(dataValue.Convert(field.Type()))
		}

	}
	return s, nil
}

type Spectrum struct {
	Name     string
	Dilution int
	W220     float32
	W222     float32
	W224     float32
	W226     float32
	W228     float32
	W230     float32
	W232     float32
	W234     float32
	W236     float32
	W238     float32
	W240     float32
	W242     float32
	W244     float32
	W246     float32
	W248     float32
	W250     float32
	W252     float32
	W254     float32
	W256     float32
	W258     float32
	W260     float32
	W262     float32
	W264     float32
	W266     float32
	W268     float32
	W270     float32
	W272     float32
	W274     float32
	W276     float32
	W278     float32
	W280     float32
	W282     float32
	W284     float32
	W286     float32
	W288     float32
	W290     float32
	W292     float32
	W294     float32
	W296     float32
	W298     float32
	W300     float32
	W302     float32
	W304     float32
	W306     float32
	W308     float32
	W310     float32
	W312     float32
	W314     float32
	W316     float32
	W318     float32
	W320     float32
	W322     float32
	W324     float32
	W326     float32
	W328     float32
	W330     float32
	W332     float32
	W334     float32
	W336     float32
	W338     float32
	W340     float32
	W342     float32
	W344     float32
	W346     float32
	W348     float32
	W350     float32
	W352     float32
	W354     float32
	W356     float32
	W358     float32
	W360     float32
	W362     float32
	W364     float32
	W366     float32
	W368     float32
	W370     float32
	W372     float32
	W374     float32
	W376     float32
	W378     float32
	W380     float32
	W382     float32
	W384     float32
	W386     float32
	W388     float32
	W390     float32
	W392     float32
	W394     float32
	W396     float32
	W398     float32
	W400     float32
	W402     float32
	W404     float32
	W406     float32
	W408     float32
	W410     float32
	W412     float32
	W414     float32
	W416     float32
	W418     float32
	W420     float32
	W422     float32
	W424     float32
	W426     float32
	W428     float32
	W430     float32
	W432     float32
	W434     float32
	W436     float32
	W438     float32
	W440     float32
	W442     float32
	W444     float32
	W446     float32
	W448     float32
	W450     float32
	W452     float32
	W454     float32
	W456     float32
	W458     float32
	W460     float32
	W462     float32
	W464     float32
	W466     float32
	W468     float32
	W470     float32
	W472     float32
	W474     float32
	W476     float32
	W478     float32
	W480     float32
	W482     float32
	W484     float32
	W486     float32
	W488     float32
	W490     float32
	W492     float32
	W494     float32
	W496     float32
	W498     float32
	W500     float32
	W502     float32
	W504     float32
	W506     float32
	W508     float32
	W510     float32
	W512     float32
	W514     float32
	W516     float32
	W518     float32
	W520     float32
	W522     float32
	W524     float32
	W526     float32
	W528     float32
	W530     float32
	W532     float32
	W534     float32
	W536     float32
	W538     float32
	W540     float32
	W542     float32
	W544     float32
	W546     float32
	W548     float32
	W550     float32
	W552     float32
	W554     float32
	W556     float32
	W558     float32
	W560     float32
	W562     float32
	W564     float32
	W566     float32
	W568     float32
	W570     float32
	W572     float32
	W574     float32
	W576     float32
	W578     float32
	W580     float32
	W582     float32
	W584     float32
	W586     float32
	W588     float32
	W590     float32
	W592     float32
	W594     float32
	W596     float32
	W598     float32
	W600     float32
	W602     float32
	W604     float32
	W606     float32
	W608     float32
	W610     float32
	W612     float32
	W614     float32
	W616     float32
	W618     float32
	W620     float32
	W622     float32
	W624     float32
	W626     float32
	W628     float32
	W630     float32
	W632     float32
	W634     float32
	W636     float32
	W638     float32
	W640     float32
	W642     float32
	W644     float32
	W646     float32
	W648     float32
	W650     float32
	W652     float32
	W654     float32
	W656     float32
	W658     float32
	W660     float32
	W662     float32
	W664     float32
	W666     float32
	W668     float32
	W670     float32
	W672     float32
	W674     float32
	W676     float32
	W678     float32
	W680     float32
	W682     float32
	W684     float32
	W686     float32
	W688     float32
	W690     float32
	W692     float32
	W694     float32
	W696     float32
	W698     float32
	W700     float32
	W702     float32
	W704     float32
	W706     float32
	W708     float32
	W710     float32
	W712     float32
	W714     float32
	W716     float32
	W718     float32
	W720     float32
	W722     float32
	W724     float32
	W726     float32
	W728     float32
	W730     float32
	W732     float32
	W734     float32
	W736     float32
	W738     float32
	W740     float32
	W742     float32
	W744     float32
	W746     float32
	W748     float32
	W750     float32
	W752     float32
	W754     float32
	W756     float32
	W758     float32
	W760     float32
	W762     float32
	W764     float32
	W766     float32
	W768     float32
	W770     float32
	W772     float32
	W774     float32
	W776     float32
	W778     float32
	W780     float32
	W782     float32
	W784     float32
	W786     float32
	W788     float32
	W790     float32
	W792     float32
	W794     float32
	W796     float32
	W798     float32
	W800     float32
}
