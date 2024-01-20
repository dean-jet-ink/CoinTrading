import { BE_URL } from "@/config/constants";
import queryClient from "@/lib/react_query";
import { useMutation } from "@tanstack/react-query";
import { TradingConfig, TradingConfigParams } from "../type";
import chartKeys from "./query-keys/chart-keys";

const useUpdateTradingConfig = () => {
  const updateTradingConfig = async (params: TradingConfigParams) => {
    try {
      const res = await fetch(`${BE_URL}/trading-config`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(params),
      });
      const tradingConfig: TradingConfig = await res.json();

      return tradingConfig;
    } catch (error) {
      console.log(error);
      throw error;
    }
  };

  const { mutate: mutateTradingConfig, isPending: isPendingTradingConfig } =
    useMutation({
      mutationKey: chartKeys.tradingConfig,
      mutationFn: updateTradingConfig,
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: chartKeys.tradingConfig,
        });
      },
    });

  return { mutateTradingConfig, isPendingTradingConfig };
};

export default useUpdateTradingConfig;
