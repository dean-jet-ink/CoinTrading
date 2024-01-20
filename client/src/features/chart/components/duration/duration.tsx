import { Loading } from "@/components/loading";
import useGetDurations from "../../api/get-durations";
import useGetTradingConfig from "../../api/get-trading-config";
import useUpdateTradingConfig from "../../api/update-trading-config";

const Duration = () => {
  const { durations, isLoadingGetDurations } = useGetDurations();
  const { tradingConfig, isLoadingGetTradingConfig } = useGetTradingConfig();
  const { mutateTradingConfig, isPendingTradingConfig } =
    useUpdateTradingConfig();

  if (isLoadingGetDurations || isLoadingGetTradingConfig) {
    return <Loading variant="sm" />;
  }

  return (
    <div className="flex gap-3">
      {durations?.durations.map(({ value, display_value }) => (
        <div
          key={value}
          className={`cursor-pointer hover:text-orange-500 ${
            tradingConfig?.duration.value == value && "text-sub hover:text-sub"
          }`}
          onClick={() => {
            if (tradingConfig?.duration.value == value) return;
            mutateTradingConfig({ duration: value });
          }}
        >
          {display_value}
        </div>
      ))}
    </div>
  );
};

export default Duration;
