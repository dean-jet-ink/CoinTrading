"use client";

import { useEffect, useRef, useState } from "react";

import { useDataframeParamsStore } from "@/store/dataframe-params";
import useWebSocketStore from "@/store/web-socket/web-socket";
import { useGetDataframeCandleStream } from "../../api/get-dataframe-candle-stream";
import useGetTradingConfig from "../../api/get-trading-config";
import { drawChart } from "../../draw-chart/draw-chart";
import { Dataframe } from "../../type";
import Duration from "../duration/duration";
import Exchange from "../exchange/exchange";
import Symbol from "../symbol/symbol";

const DashBoard = () => {
  const { webSocket, setWebSocket } = useWebSocketStore();
  const { onMessage, sendMessage } = useGetDataframeCandleStream();
  const { params } = useDataframeParamsStore();
  const [dataframe, setDataframe] = useState<Dataframe>({ candles: [] });
  const { tradingConfig } = useGetTradingConfig();

  useEffect(() => {
    setWebSocket("candles");
  }, []);

  useEffect(() => {
    let timer: NodeJS.Timeout;

    const setSendMessageTimer = (ws: WebSocket) => {
      if (ws.readyState === WebSocket.OPEN) {
        sendMessage(ws, params);

        timer = setInterval(() => {
          sendMessage(ws, params);
        }, 3000);
      } else {
        setTimeout(() => {
          setSendMessageTimer(ws);
        }, 500);
      }
    };

    if (webSocket) {
      webSocket.onmessage = onMessage(setDataframe);
      setSendMessageTimer(webSocket);
    }

    return () => {
      clearInterval(timer);
    };
  }, [webSocket, tradingConfig]);

  const dashboard = useRef(null);

  useEffect(() => {
    if (!dashboard.current) return;
    drawChart(dashboard.current, dataframe);
  }, [dataframe]);

  return (
    <div className="w-2/3 p-10 bg-zinc-900">
      <div className="flex justify-end items-center gap-10 relative z-50">
        <Exchange />
        <Symbol />
        <Duration />
      </div>
      <div className="min-h-[640px] rounded-lg" ref={dashboard}></div>
    </div>
  );
};

export default DashBoard;
