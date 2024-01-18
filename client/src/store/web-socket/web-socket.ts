import { WS_URL } from "@/config/constants";
import { create } from "zustand";

type State = {
  webSocket: WebSocket | null;
  setWebSocket: (path: string) => void;
  clearWebSocket: () => void;
};

const useWebSocketStore = create<State>((set) => ({
  webSocket: null,
  setWebSocket: (path: string) =>
    set(() => {
      const ws = new WebSocket(`${WS_URL}/${path}`);

      ws.onopen = () => {
        console.log("Connected to websocket");
      };

      ws.onerror = (e) => {
        console.log(`Websocket error: `, e);
        ws.close();

        return { websocket: null };
      };

      return { webSocket: ws };
    }),
  clearWebSocket: () =>
    set((state) => {
      state.webSocket?.close();

      return { webSocket: null };
    }),
}));

export default useWebSocketStore;
