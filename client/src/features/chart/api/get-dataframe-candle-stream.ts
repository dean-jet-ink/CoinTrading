import { Dataframe } from "../type";
import { DataframeParams } from "@/store/dataframe-params";

export const useGetDataframeCandleStream = () => {
  const onMessage = (setState: (df: Dataframe) => void) => {
    return (e: MessageEvent<string>) => {
      const dataframe: Dataframe = JSON.parse(e.data);

      setState(dataframe);
    };
  };

  const sendMessage = (ws: WebSocket, message: DataframeParams) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify(message));
      console.log("Sent message: ", message);
    } else {
      console.log("WebSocket not connected");
    }
  };

  return {
    onMessage,
    sendMessage,
  };
};
