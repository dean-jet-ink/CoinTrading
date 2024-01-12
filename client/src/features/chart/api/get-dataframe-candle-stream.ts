import { Dataframe } from "../type";
import { WS_URL } from "@/config/constants";

export const useGetDataframeCandleStream = () => {
  if (!WS_URL) throw new Error("WS_URL is not defined");

  const websocket = new WebSocket(`${WS_URL}/candles`);

  const onMessage = (e: MessageEvent<Dataframe>) => {};

  websocket.addEventListener("message", onMessage);

  const sendMessage = () => {
    websocket.send(JSON.stringify({}));
  };

  const { close } = websocket;

  return {
    sendMessage,
    close,
  };
};
