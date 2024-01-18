"use client";

import { useEffect, useRef, useState } from "react";

import { Dataframe } from "../../type";
import { drawChart } from "../../draw-chart/draw-chart";
import { useGetDataframeCandleStream } from "../../api/get-dataframe-candle-stream";
import { useDataframeParamsStore } from "@/store/dataframe-params";
import useWebSocketStore from "@/store/web-socket/web-socket";
import { Button } from "@/components/button";
import Duration from "../duration/duration";

const DashBoard = () => {
  const { webSocket, setWebSocket } = useWebSocketStore();
  const { onMessage, sendMessage } = useGetDataframeCandleStream();
  const { params } = useDataframeParamsStore();
  const [dataframe, setDataframe] = useState<Dataframe>({ candles: [] });

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
  }, [webSocket]);

  const dashboard = useRef(null);

  useEffect(() => {
    if (!dashboard.current) return;
    drawChart(dashboard.current, dataframe);
  }, [dataframe]);

  return (
    <div className="w-2/3 min-h-[640px] p-10 bg-zinc-900 rounded-lg flex flex-col">
      <div className="flex justify-end">
        <Duration />
      </div>
      <div className="flex-grow" ref={dashboard}></div>
    </div>
  );
};

export default DashBoard;
