// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/api/Resource.proto

package org.wso2.gateway.discovery.api;

public interface OperationOrBuilder extends
    // @@protoc_insertion_point(interface_extends:wso2.discovery.api.Operation)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string method = 1;</code>
   * @return The method.
   */
  java.lang.String getMethod();
  /**
   * <code>string method = 1;</code>
   * @return The bytes for method.
   */
  com.google.protobuf.ByteString
      getMethodBytes();

  /**
   * <code>repeated .wso2.discovery.api.SecurityList security = 2;</code>
   */
  java.util.List<org.wso2.gateway.discovery.api.SecurityList> 
      getSecurityList();
  /**
   * <code>repeated .wso2.discovery.api.SecurityList security = 2;</code>
   */
  org.wso2.gateway.discovery.api.SecurityList getSecurity(int index);
  /**
   * <code>repeated .wso2.discovery.api.SecurityList security = 2;</code>
   */
  int getSecurityCount();
  /**
   * <code>repeated .wso2.discovery.api.SecurityList security = 2;</code>
   */
  java.util.List<? extends org.wso2.gateway.discovery.api.SecurityListOrBuilder> 
      getSecurityOrBuilderList();
  /**
   * <code>repeated .wso2.discovery.api.SecurityList security = 2;</code>
   */
  org.wso2.gateway.discovery.api.SecurityListOrBuilder getSecurityOrBuilder(
      int index);

  /**
   * <code>string tier = 3;</code>
   * @return The tier.
   */
  java.lang.String getTier();
  /**
   * <code>string tier = 3;</code>
   * @return The bytes for tier.
   */
  com.google.protobuf.ByteString
      getTierBytes();
}
