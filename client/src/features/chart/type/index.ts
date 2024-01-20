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

export type Duration = {
  value: string;
  display_value: string;
};

export type Durations = {
  durations: Duration[];
};

export type ExchangeType = "Bitflyer";

export type Exchange = {
  name: ExchangeType;
  value: number;
};

export type Exchanges = {
  exchanges: Exchange[];
};

export type SymbolType = "BTC/JPY" | "ETH/JPY" | "XRP/JPY";

export type Symbol = {
  name: SymbolType;
  value: number;
};

export type Symbols = {
  symbols: Symbol[];
};

export type TradingConfig = {
  exchange: Exchange;
  symbol: Symbol;
  duration: Duration;
};

export type TradingConfigParams = {
  exchange?: number;
  symbol?: number;
  duration?: string;
};
