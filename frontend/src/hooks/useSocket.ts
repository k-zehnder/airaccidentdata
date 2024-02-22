import { useEffect, useState, useRef } from "react";
import io, { Socket } from "socket.io-client";

export const useSocket = <T>(url: string, eventName: string): T[] => {
  const [data, setData] = useState<T[]>([]);
  const socketRef = useRef<Socket | null>(null);

  useEffect(() => {
    // Initialize the socket only if it doesn't already exist
    if (!socketRef.current) {
      socketRef.current = io(url);

      socketRef.current.on("connect", () => {
        console.log("Connected to socket server");
      });

      socketRef.current.on(eventName, (newData: T) => {
        setData((prevData) => [newData, ...prevData]);
      });

      socketRef.current.on("echo", (message: string) => {
        console.log("Client got message from server:", message);
      });
    }

    return () => {
      if (socketRef.current) {
        socketRef.current.off(eventName);
        socketRef.current.disconnect();
        socketRef.current = null;
      }
    };
  }, [url, eventName]);

  return data;
};
