import { BE_URL } from "@/config/constants";
import { useQuery } from "@tanstack/react-query";
import { Exchanges } from "../type";

const useGetExchanges = () => {
  const getExchanges = async (): Promise<Exchanges> => {
    try {
      const res = await fetch(`${BE_URL}/exchanges`);
      const exchanges: Exchanges = await res.json();

      return exchanges;
    } catch (err) {
      console.log(err);
      throw err;
    }
  };

  const { data: exchanges, isLoading: isLoadingExchanges } = useQuery({
    queryKey: ["exchanges"],
    queryFn: getExchanges,
  });

  return { exchanges, isLoadingExchanges };
};

export default useGetExchanges;
