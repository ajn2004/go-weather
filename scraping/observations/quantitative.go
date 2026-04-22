package observations

type WindSpeedValue struct {
	Str *string
	QV  *QuantitativeValue
}

type TemperatureValue struct {
	Int *int
	QV  *QuantitativeValue
}

type QuantitativeValue struct {
	Value          *float64 `json:"value,omitempty"`
	MaxValue       *float64 `json:"maxValue,omitempty"`
	MinValue       *float64 `json:"minValue,omitempty"`
	UnitCode       string   `json:"unitCode,omitempty"`
	QualityControl string   `json:"qualityControl,omitempty"`
}

type WindSpeedKind int

const (
	WindSpeedEmpty WindSpeedKind = iota
	WindSpeedString
	WindSpeedQV
)

func (w WindSpeedValue) Kind() WindSpeedKind {
	switch {
	case w.QV != nil:
		return WindSpeedQV
	case w.Str != nil:
		return WindSpeedString
	default:
		return WindSpeedEmpty
	}
}

type TemperatureKind int

const (
	TemperatureEmpty TemperatureKind = iota
	TemperatureInt
	TemperatureQV
)

func (t TemperatureValue) Kind() TemperatureKind {
	switch {
	case t.QV != nil:
		return TemperatureQV
	case t.Int != nil:
		return TemperatureInt
	default:
		return TemperatureEmpty
	}
}
