import { BE_URL } from "@/config/constants";
import { useQuery } from "@tanstack/react-query";
import { Durations } from "../type";
import chartKeys from "./query-keys/chart-keys";

const useGetDurations = () => {
  const getDurations = async (): Promise<Durations> => {
    try {
      const res = await fetch(`${BE_URL}/durations`, {
        cache: "force-cache",
      });
      const durations: Durations = await res.json();

      return durations;
    } catch (error) {
      console.error(error);
      throw error;
    }
  };

  const { data: durations, isLoading: isLoadingDurations } = useQuery({
    queryKey: chartKeys.durations,
    queryFn: getDurations,
  });

  return { durations, isLoadingDurations };
};

export default useGetDurations;
