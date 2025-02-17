import * as thrift from "thrift";
import { PingServiceClient } from "./gen-ts/PingService";

console.log('ping', PingServiceClient);

const connection = thrift.createConnection("localhost", 9090, {
  transport: thrift.TBufferedTransport,
  protocol: thrift.TBinaryProtocol,
});

connection.on("error", (err) => console.error(err));

const client = new PingServiceClient(connection);

connection.on("connect", async () => {
  try {
    const result = await client.ping();
    console.log("Server responded:", result);
    connection.end();
  } catch (err) {
    console.error(err);
  }
});