import { Loading } from "@/components/loading";
import Select, { Option } from "@/components/select/select";
import Image from "next/image";
import { useCallback } from "react";
import useGetSymbols from "../../api/get-symbols";
import useGetTradingConfig from "../../api/get-trading-config";
import useUpdateTradingConfig from "../../api/update-trading-config";
import { SymbolType } from "../../type";

type SymbolIcons = {
  [key in SymbolType]: string;
};

const symbolIcons: SymbolIcons = {
  "BTC/JPY": "/btc.svg",
  "ETH/JPY": "/eth.svg",
  "XRP/JPY": "/xrp.svg",
};

const Symbol = () => {
  const { symbols, isLoadingSymbols } = useGetSymbols();
  const { tradingConfig, isLoadingTradingConfig } = useGetTradingConfig();
  const { mutateTradingConfig, isPendingTradingConfig } =
    useUpdateTradingConfig();

  let symbolOptions: Option[] = [];

  if (symbols) {
    symbolOptions = symbols.symbols.map(({ name, value }) => {
      return {
        label: (
          <div className="flex gap-3 items-center">
            <Image
              src={symbolIcons[name]}
              alt="exchange icon"
              width={15}
              height={15}
              className="rounded-md"
            />
            <span>{name}</span>
          </div>
        ),
        value,
        selected: value === tradingConfig?.symbol.value,
      };
    });
  }

  const updateSymbol = useCallback((value: number) => {
    mutateTradingConfig({ symbol: value });
  }, []);

  if (isLoadingSymbols || isLoadingTradingConfig || isPendingTradingConfig) {
    return <Loading variant="sm" />;
  }

  return <Select options={symbolOptions} setOption={updateSymbol} />;
};

export default Symbol;
