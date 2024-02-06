// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/agent.proto

package flyteidl.service;

public final class Agent {
  private Agent() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\034flyteidl/service/agent.proto\022\020flyteidl" +
      ".service\032\034google/api/annotations.proto\032\032" +
      "flyteidl/admin/agent.proto2\314\003\n\021AsyncAgen" +
      "tService\022U\n\nCreateTask\022!.flyteidl.admin." +
      "CreateTaskRequest\032\".flyteidl.admin.Creat" +
      "eTaskResponse\"\000\022L\n\007GetTask\022\036.flyteidl.ad" +
      "min.GetTaskRequest\032\037.flyteidl.admin.GetT" +
      "askResponse\"\000\022U\n\nDeleteTask\022!.flyteidl.a" +
      "dmin.DeleteTaskRequest\032\".flyteidl.admin." +
      "DeleteTaskResponse\"\000\022a\n\016GetTaskMetrics\022%" +
      ".flyteidl.admin.GetTaskMetricsRequest\032&." +
      "flyteidl.admin.GetTaskMetricsResponse\"\000\022" +
      "X\n\013GetTaskLogs\022\".flyteidl.admin.GetTaskL" +
      "ogsRequest\032#.flyteidl.admin.GetTaskLogsR" +
      "esponse\"\0002\360\001\n\024AgentMetadataService\022k\n\010Ge" +
      "tAgent\022\037.flyteidl.admin.GetAgentRequest\032" +
      " .flyteidl.admin.GetAgentResponse\"\034\202\323\344\223\002" +
      "\026\022\024/api/v1/agent/{name}\022k\n\nListAgents\022!." +
      "flyteidl.admin.ListAgentsRequest\032\".flyte" +
      "idl.admin.ListAgentsResponse\"\026\202\323\344\223\002\020\022\016/a" +
      "pi/v1/agentsB?Z=github.com/flyteorg/flyt" +
      "e/flyteidl/gen/pb-go/flyteidl/serviceb\006p" +
      "roto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.api.AnnotationsProto.getDescriptor(),
          flyteidl.admin.AgentOuterClass.getDescriptor(),
        }, assigner);
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.api.AnnotationsProto.http);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.api.AnnotationsProto.getDescriptor();
    flyteidl.admin.AgentOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}