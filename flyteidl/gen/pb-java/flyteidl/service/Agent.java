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
      ".service\032\032flyteidl/admin/agent.proto2\217\002\n" +
      "\021AsyncAgentService\022U\n\nCreateTask\022!.flyte" +
      "idl.admin.CreateTaskRequest\032\".flyteidl.a" +
      "dmin.CreateTaskResponse\"\000\022L\n\007GetTask\022\036.f" +
      "lyteidl.admin.GetTaskRequest\032\037.flyteidl." +
      "admin.GetTaskResponse\"\000\022U\n\nDeleteTask\022!." +
      "flyteidl.admin.DeleteTaskRequest\032\".flyte" +
      "idl.admin.DeleteTaskResponse\"\000B9Z7github" +
      ".com/flyteorg/flyteidl/gen/pb-go/flyteid" +
      "l/serviceb\006proto3"
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
          flyteidl.admin.Agent.getDescriptor(),
        }, assigner);
    flyteidl.admin.Agent.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
