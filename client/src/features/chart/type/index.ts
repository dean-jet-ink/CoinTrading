export type Candle = {
  time: string;
  open: number;
  close: number;
  high: number;
  low: number;
  volume: number;
};

export type Dataframe = {
  candles: Candle[];
};
