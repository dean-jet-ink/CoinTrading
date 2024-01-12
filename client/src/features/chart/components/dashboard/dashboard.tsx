"use client";

import { useEffect, useRef } from "react";

import { Candle } from "../../type";
import { initializeChart } from "../../initialize-option/initialize-option";
import { useGetDataframeCandleStream } from "../../api/get-dataframe-candle-stream";

const data: Candle[] = [
  {
    time: "2020-01-01T00:00:00",
    open: 100,
    high: 100,
    low: 100,
    close: 100,
    volume: 100,
  },
  {
    time: "2020-01-02T00:00:00",
    open: 101,
    high: 101,
    low: 101,
    close: 101,
    volume: 101,
  },
  {
    time: "2020-01-03T00:00:00",
    open: 102,
    high: 102,
    low: 102,
    close: 102,
    volume: 102,
  },
  {
    time: "2020-01-04T00:00:00",
    open: 103,
    high: 103,
    low: 103,
    close: 103,
    volume: 103,
  },
  {
    time: "2020-01-05T00:00:00",
    open: 104,
    high: 104,
    low: 104,
    close: 104,
    volume: 104,
  },
  {
    time: "2020-01-06T00:00:00",
    open: 105,
    high: 105,
    low: 105,
    close: 105,
    volume: 105,
  },
  {
    time: "2020-01-07T00:00:00",
    open: 106,
    high: 106,
    low: 106,
    close: 106,
    volume: 106,
  },
  {
    time: "2020-01-08T00:00:00",
    open: 107,
    high: 107,
    low: 107,
    close: 107,
    volume: 107,
  },
  {
    time: "2020-01-09T00:00:00",
    open: 108,
    high: 108,
    low: 108,
    close: 108,
    volume: 108,
  },
  {
    time: "2020-01-10T00:00:00",
    open: 109,
    high: 109,
    low: 109,
    close: 109,
    volume: 109,
  },
];

const dataframe = {
  candles: data,
};

const DashBoard = () => {
  const { sendMessage, close } = useGetDataframeCandleStream();

  useEffect(() => {
    // sendMessage();
    // setInterval(() => {
    //   sendMessage();
    // }, 3000);
    // return () => {
    //   close();
    // };
  }, []);

  const dashboard = useRef(null);

  useEffect(() => {
    if (dashboard.current) {
      initializeChart(dashboard.current, dataframe);
    }
  }, []);

  return <div className="w-2/3 min-h-[640px]" ref={dashboard}></div>;
};

export default DashBoard;
