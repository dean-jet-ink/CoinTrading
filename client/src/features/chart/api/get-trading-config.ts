import { BE_URL } from "@/config/constants";
import { useQuery } from "@tanstack/react-query";
import { TradingConfig } from "../type";
import chartKeys from "./query-keys/chart-keys";

const useGetTradingConfig = () => {
  const getTradingConfig = async (): Promise<TradingConfig> => {
    try {
      const res = await fetch(`${BE_URL}/trading-config`, {
        cache: "no-store",
      });
      const tradingConfig: TradingConfig = await res.json();

      return tradingConfig;
    } catch (error) {
      console.log(error);
      throw error;
    }
  };

  const { data: tradingConfig, isLoading: isLoadingTradingConfig } = useQuery({
    queryKey: chartKeys.tradingConfig,
    queryFn: getTradingConfig,
  });

  return { tradingConfig, isLoadingTradingConfig };
};

export default useGetTradingConfig;
