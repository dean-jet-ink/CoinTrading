import { BE_URL } from "@/config/constants";
import { Duration } from "../type";
import { useQuery } from "@tanstack/react-query";
import chartKeys from "./query-keys/chart-keys";

const useGetDurations = () => {
  const getDurations = async (): Promise<Duration[]> => {
    try {
      const res = await fetch(`${BE_URL}/durations`);
      const durations: Duration[] = await res.json();

      return durations;
    } catch (error) {
      console.error(error);
      throw error;
    }
  };

  const { data, isLoading } = useQuery({
    queryKey: chartKeys.durations,
    queryFn: getDurations,
  });

  return { data, isLoading };
};
