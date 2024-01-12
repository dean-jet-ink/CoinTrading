package entities

type DataframeCandle struct {
	candles []*Candle
}

func NewDataframeCandle(candles []*Candle) *DataframeCandle {
	return &DataframeCandle{
		candles: candles,
	}
}

func (d *DataframeCandle) Candles() []*Candle {
	return d.candles
}

func (d *DataframeCandle) Closes() []float64 {
	closes := make([]float64, len(d.candles))

	for i, candle := range d.candles {
		closes[i] = candle.Close()
	}

	return closes
}