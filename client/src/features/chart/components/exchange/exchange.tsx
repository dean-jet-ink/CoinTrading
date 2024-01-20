import Image from "next/image";

import { Loading } from "@/components/loading";
import Select, { Option } from "@/components/select/select";
import { useCallback } from "react";
import useGetExchanges from "../../api/get-exchanges";
import useGetTradingConfig from "../../api/get-trading-config";
import useUpdateTradingConfig from "../../api/update-trading-config";
import { ExchangeType } from "../../type";

type ExchangeIcons = {
  [key in ExchangeType]: string;
};

const exchangeIcons: ExchangeIcons = {
  Bitflyer: "/bitflyer.png",
};

const Exchange = () => {
  const { tradingConfig, isLoadingTradingConfig } = useGetTradingConfig();
  const { exchanges, isLoadingExchanges } = useGetExchanges();
  const { mutateTradingConfig, isPendingTradingConfig } =
    useUpdateTradingConfig();

  let exchangeOptions: Option[] = [];

  if (exchanges) {
    exchangeOptions = exchanges?.exchanges.map(({ name, value }) => {
      return {
        label: (
          <div className="flex gap-3 items-center">
            <Image
              src={exchangeIcons[name]}
              alt="exchange icon"
              width={22}
              height={22}
              className="rounded-md"
            />
            <span>{name}</span>
          </div>
        ),
        value: name as string,
        selected: value === tradingConfig?.exchange.value,
      };
    });
  }

  const updateExchange = useCallback((value: number) => {
    mutateTradingConfig({ exchange: value });
  }, []);

  if (isLoadingTradingConfig || isLoadingExchanges || isPendingTradingConfig) {
    return <Loading variant="sm" />;
  }

  return <Select options={exchangeOptions} setOption={updateExchange} />;
};

export default Exchange;
