//
// Autogenerated by Thrift Compiler (0.21.0)
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//

import Int64 = require('node-int64');

import { ping } from "./ping_types";


export declare class PingServiceClient {
  input: Thrift.TJSONProtocol;
  output: Thrift.TJSONProtocol;
  seqid: number;

  constructor(input: Thrift.TJSONProtocol, output?: Thrift.TJSONProtocol);

  ping(): string;

  ping(callback: (data: string)=>void): void;
}
