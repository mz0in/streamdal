import { serviceSignal } from "../components/serviceMap/serviceSignal.ts";
import { audienceMetricsSignal } from "../components/serviceMap/customEdge.tsx";
import {
  MAX_TAIL_SIZE,
  tailSamplingSignal,
  tailSignal,
} from "../islands/drawer/tail.tsx";
import { Audience } from "streamdal-protos/protos/sp_common.ts";

import { serverErrorSignal } from "../components/serviceMap/serverErrorSignal.tsx";
import { demoHttpRequestSignal } from "../routes/demo/http/index.tsx";

export const getSocket = (path: string) => {
  const url = new URL(path, location.href);
  url.protocol = url.protocol.replace("http", "ws");
  url.pathname = path;
  return new WebSocket(url);
};

export const demoHttpRequestSocket = (path: string) => {
  const webSocket = getSocket(path);

  webSocket.addEventListener("open", () => {
    webSocket.send("ping");
  });

  webSocket.addEventListener("message", (event) => {
    if (event.data === "pong") {
      console.debug("got server pong");
      return;
    }

    try {
      demoHttpRequestSignal.value = JSON.parse(event.data);
    } catch (e) {
      console.error("error parsing demo http request data", e);
    }
  });
  return webSocket;
};

export const serviceMapSocket = (path: string) => {
  const webSocket = getSocket(path);

  webSocket.addEventListener("open", () => {
    webSocket.send("ping");
  });

  webSocket.addEventListener("message", (event) => {
    if (event.data === "pong") {
      console.debug("got server pong");
      return;
    }

    try {
      const parsedData = JSON.parse(event.data);
      serviceSignal.value = {
        ...parsedData,
        nodesMap: new Map(parsedData.nodesMap),
        edgesMap: new Map(parsedData.edgesMap),
        browserInitialized: true,
      };
    } catch (e) {
      console.error("error parsing serviceMap socket data", e);
    }
  });
  return webSocket;
};

export const audienceMetricsSocket = (path: string) => {
  const webSocket = getSocket(path);

  webSocket.addEventListener("open", () => {
    webSocket.send("ping");
  });

  webSocket.addEventListener("message", (event) => {
    if (event.data === "pong") {
      console.debug("got server pong");
      return;
    }

    try {
      const parsedData = JSON.parse(event.data);
      audienceMetricsSignal.value = {
        ...audienceMetricsSignal.value,
        ...parsedData,
      };
    } catch (e) {
      console.error("error parsing audience metrics socket data", e);
    }
  });
  return webSocket;
};

export const tailSocket = (path: string, audience: Audience) => {
  const webSocket = getSocket(path);

  webSocket.addEventListener("open", () => {
    webSocket.send("ping");
    webSocket.send(
      JSON.stringify({
        audience,
        ...tailSamplingSignal.value
          ? { sampling: tailSamplingSignal.value }
          : {},
      }),
    );
  });

  webSocket.addEventListener("message", (event) => {
    if (event.data === "pong") {
      console.debug("got server pong");
      return;
    }

    try {
      const parsedTail = JSON.parse(event.data);

      tailSignal.value = [
        ...(tailSignal.value || []).slice(-MAX_TAIL_SIZE),
        {
          timestamp: new Date(parsedTail.timestamp),
          data: parsedTail.data,
          originalData: parsedTail.originalData,
        },
      ];
    } catch (e) {
      console.error("error parsing tail data", e);
    }
  });
  return webSocket;
};

export const serverErrorSocket = (path: string) => {
  const webSocket = getSocket(path);
  webSocket.addEventListener("open", () => {
    webSocket.send("ping");
  });

  webSocket.addEventListener("message", (event) => {
    if (event.data === "pong") {
      console.debug("got server pong");
      return;
    }

    serverErrorSignal.value = event.data;
  });
  return webSocket;
};
