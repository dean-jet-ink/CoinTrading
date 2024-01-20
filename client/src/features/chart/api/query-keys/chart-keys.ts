type ChartKeys = {
  tradingConfig: string[];
  exchanges: string[];
  symbols: string[];
  durations: string[];
};

const chartKeys: ChartKeys = {
  tradingConfig: ["tradingConfig"] as const,
  exchanges: ["exchanges"] as const,
  symbols: ["symbols"] as const,
  durations: ["durations"] as const,
};

export default chartKeys;
