import * as echarts from "echarts/core";
import {
  TitleComponentOption,
  TooltipComponent,
  TooltipComponentOption,
  GridComponent,
  GridComponentOption,
  LegendComponent,
  LegendComponentOption,
  DataZoomComponent,
  DataZoomComponentOption,
  MarkLineComponent,
  MarkLineComponentOption,
  MarkPointComponent,
  MarkPointComponentOption,
} from "echarts/components";
import {
  CandlestickChart,
  CandlestickSeriesOption,
  LineChart,
  LineSeriesOption,
  BarChart,
  BarSeriesOption,
} from "echarts/charts";
import { UniversalTransition } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";

import { Candle, Dataframe } from "../type";

type EchartOption = echarts.ComposeOption<
  | TitleComponentOption
  | TooltipComponentOption
  | GridComponentOption
  | LegendComponentOption
  | DataZoomComponentOption
  | MarkLineComponentOption
  | MarkPointComponentOption
  | CandlestickSeriesOption
  | LineSeriesOption
  | BarSeriesOption
>;

const splitData = (
  candles: Candle[]
): {
  times: string[];
  values: number[][];
  volumes: number[];
} => {
  let times: string[] = [];
  let values: number[][] = [];
  let volumes: number[] = [];

  for (let i = 0; i < candles.length; i++) {
    times.push(candles[i].time);
    values.push([
      candles[i].close,
      candles[i].open,
      candles[i].high,
      candles[i].low,
    ]);
    volumes.push(candles[i].volume);
  }

  return { times, values, volumes };
};

export const drawChart = (chartDom: HTMLElement, dataframe: Dataframe) => {
  echarts.use([
    TooltipComponent,
    GridComponent,
    LegendComponent,
    DataZoomComponent,
    MarkLineComponent,
    MarkPointComponent,
    CandlestickChart,
    LineChart,
    BarChart,
    CanvasRenderer,
    UniversalTransition,
  ]);

  const upColor = "#60a9a6";
  const downColor = "#ab4c4c";
  const volumeColor = "#535353";
  const textColor = "#d4d4d8";

  const { times, values, volumes } = splitData(dataframe.candles);

  const option: EchartOption = {
    tooltip: {
      trigger: "axis",
      axisPointer: {
        type: "cross",
      },
      padding: 20,
      position: function (pos, params, el, elRect, size) {
        const obj: Record<string, number> = {
          top: 10,
        };
        obj[["left", "right"][+(pos[0] < size.viewSize[0] / 2)]] = 30;
        return obj;
      },
    },
    axisPointer: {
      link: [
        {
          xAxisIndex: "all",
        },
      ],
      label: {
        backgroundColor: "#777",
      },
    },
    grid: [
      {
        left: "10%",
        right: "8%",
        height: "50%",
      },
      {
        left: "10%",
        right: "8%",
        top: "63%",
        height: "16%",
      },
    ],
    xAxis: [
      {
        type: "category",
        data: times,
        boundaryGap: false,
        min: "dataMin",
        max: "dataMax",
        axisLabel: {
          color: textColor,
        },
      },
      {
        type: "category",
        gridIndex: 1,
        data: times,
        boundaryGap: false,
        splitLine: { show: false },
        axisLabel: { show: false },
        axisTick: { show: false },
        axisLine: { lineStyle: { color: "#777" } },
        min: "dataMin",
        max: "dataMax",
      },
    ],
    yAxis: [
      {
        scale: true,
        axisLabel: {
          color: textColor,
        },
      },
      {
        scale: true,
        gridIndex: 1,
        splitNumber: 2,
        axisLabel: { show: false },
        axisLine: { show: false },
        axisTick: { show: false },
        splitLine: { show: false },
      },
    ],
    dataZoom: [
      {
        type: "inside",
        xAxisIndex: [0, 1],
        start: 50,
        end: 100,
      },
      {
        show: true,
        type: "slider",
        xAxisIndex: [0, 1],
        top: "85%",
        start: 50,
        end: 100,
      },
    ],
    series: [
      {
        name: "currency",
        type: "candlestick",
        data: values,
        itemStyle: {
          color: upColor,
          color0: downColor,
          borderColor: upColor,
          borderColor0: downColor,
          borderWidth: 2,
        },
      },
      {
        name: "volume",
        type: "bar",
        data: volumes,
        xAxisIndex: 1,
        yAxisIndex: 1,
        itemStyle: {
          color: volumeColor,
        },
      },
    ],
  };

  const myChart = echarts.init(chartDom, null, {
    renderer: "canvas",
    useDirtyRect: false,
  });

  option && myChart.setOption(option);
};
