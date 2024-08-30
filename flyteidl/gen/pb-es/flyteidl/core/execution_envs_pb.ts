// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/core/execution_envs.proto (package flyteidl.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Struct } from "@bufbuild/protobuf";

/**
 * ExecutionEnvAssignment is a message that is used to assign an execution environment to a set of
 * nodes.
 *
 * @generated from message flyteidl.core.ExecutionEnvAssignment
 */
export class ExecutionEnvAssignment extends Message<ExecutionEnvAssignment> {
  /**
   * node_ids is a list of node ids that are being assigned the execution environment.
   *
   * @generated from field: repeated string node_ids = 1;
   */
  nodeIds: string[] = [];

  /**
   * task_type is the type of task that is being assigned. This is used to override which Flyte
   * plugin will be used during execution.
   *
   * @generated from field: string task_type = 2;
   */
  taskType = "";

  /**
   * execution_env is the environment that is being assigned to the nodes.
   *
   * @generated from field: flyteidl.core.ExecutionEnv execution_env = 3;
   */
  executionEnv?: ExecutionEnv;

  constructor(data?: PartialMessage<ExecutionEnvAssignment>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ExecutionEnvAssignment";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "node_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "task_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "execution_env", kind: "message", T: ExecutionEnv },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExecutionEnvAssignment {
    return new ExecutionEnvAssignment().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExecutionEnvAssignment {
    return new ExecutionEnvAssignment().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExecutionEnvAssignment {
    return new ExecutionEnvAssignment().fromJsonString(jsonString, options);
  }

  static equals(a: ExecutionEnvAssignment | PlainMessage<ExecutionEnvAssignment> | undefined, b: ExecutionEnvAssignment | PlainMessage<ExecutionEnvAssignment> | undefined): boolean {
    return proto3.util.equals(ExecutionEnvAssignment, a, b);
  }
}

/**
 * ExecutionEnv is a message that is used to specify the execution environment.
 *
 * @generated from message flyteidl.core.ExecutionEnv
 */
export class ExecutionEnv extends Message<ExecutionEnv> {
  /**
   * name is a human-readable identifier for the execution environment. This is combined with the
   * project, domain, and version to uniquely identify an execution environment.
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * type is the type of the execution environment.
   *
   * @generated from field: string type = 2;
   */
  type = "";

  /**
   * environment is a oneof field that can be used to specify the environment in different ways.
   *
   * @generated from oneof flyteidl.core.ExecutionEnv.environment
   */
  environment: {
    /**
     * extant is a reference to an existing environment.
     *
     * @generated from field: google.protobuf.Struct extant = 3;
     */
    value: Struct;
    case: "extant";
  } | {
    /**
     * spec is a specification of the environment.
     *
     * @generated from field: google.protobuf.Struct spec = 4;
     */
    value: Struct;
    case: "spec";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * version is the version of the execution environment. This may be used differently by each
   * individual environment type (ex. auto-generated or manually provided), but is intended to
   * allow variance in environment specifications with the same ID.
   *
   * @generated from field: string version = 5;
   */
  version = "";

  constructor(data?: PartialMessage<ExecutionEnv>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ExecutionEnv";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "extant", kind: "message", T: Struct, oneof: "environment" },
    { no: 4, name: "spec", kind: "message", T: Struct, oneof: "environment" },
    { no: 5, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExecutionEnv {
    return new ExecutionEnv().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExecutionEnv {
    return new ExecutionEnv().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExecutionEnv {
    return new ExecutionEnv().fromJsonString(jsonString, options);
  }

  static equals(a: ExecutionEnv | PlainMessage<ExecutionEnv> | undefined, b: ExecutionEnv | PlainMessage<ExecutionEnv> | undefined): boolean {
    return proto3.util.equals(ExecutionEnv, a, b);
  }
}
