import { BE_URL } from "@/config/constants";
import { useQuery } from "@tanstack/react-query";
import { Symbols } from "../type";
import chartKeys from "./query-keys/chart-keys";

const useGetSymbols = () => {
  const getSymbols = async (): Promise<Symbols> => {
    try {
      const res = await fetch(`${BE_URL}/symbols`);
      const symbols: Symbols = await res.json();

      return symbols;
    } catch (error) {
      console.log(error);
      throw error;
    }
  };

  const { data: symbols, isLoading: isLoadingSymbols } = useQuery({
    queryKey: chartKeys.symbols,
    queryFn: getSymbols,
  });

  return { symbols, isLoadingSymbols };
};

export default useGetSymbols;
