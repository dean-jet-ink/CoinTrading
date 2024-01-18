import { create } from "zustand";

type CandleParams = {
  message: string;
};

export type DataframeParams = {
  candle_params: CandleParams;
};

type State = {
  params: DataframeParams;
  setCandleParams: (params: CandleParams) => void;
};

export const useDataframeParamsStore = create<State>((set) => ({
  params: {
    candle_params: {
      message: "",
    },
  },
  setCandleParams: (candleParams: CandleParams) => {
    set(({ params }) => ({
      params: {
        ...params,
        candle_params: candleParams,
      },
    }));
  },
}));
