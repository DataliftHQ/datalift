// @generated by protoc-gen-connect-es v0.10.1
// @generated from file application/v1/application.proto (package datalift.application.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateApplicationRequest, CreateApplicationResponse, DeleteApplicationRequest, DeleteApplicationResponse, GetApplicationRequest, GetApplicationResponse, ListApplicationsRequest, ListApplicationsResponse, UpdateApplicationRequest, UpdateApplicationResponse } from "./application_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service datalift.application.v1.ApplicationAPI
 */
export declare const ApplicationAPI: {
  readonly typeName: "datalift.application.v1.ApplicationAPI",
  readonly methods: {
    /**
     * @generated from rpc datalift.application.v1.ApplicationAPI.CreateApplication
     */
    readonly createApplication: {
      readonly name: "CreateApplication",
      readonly I: typeof CreateApplicationRequest,
      readonly O: typeof CreateApplicationResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc datalift.application.v1.ApplicationAPI.DeleteApplication
     */
    readonly deleteApplication: {
      readonly name: "DeleteApplication",
      readonly I: typeof DeleteApplicationRequest,
      readonly O: typeof DeleteApplicationResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc datalift.application.v1.ApplicationAPI.GetApplication
     */
    readonly getApplication: {
      readonly name: "GetApplication",
      readonly I: typeof GetApplicationRequest,
      readonly O: typeof GetApplicationResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc datalift.application.v1.ApplicationAPI.ListApplications
     */
    readonly listApplications: {
      readonly name: "ListApplications",
      readonly I: typeof ListApplicationsRequest,
      readonly O: typeof ListApplicationsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc datalift.application.v1.ApplicationAPI.UpdateApplication
     */
    readonly updateApplication: {
      readonly name: "UpdateApplication",
      readonly I: typeof UpdateApplicationRequest,
      readonly O: typeof UpdateApplicationResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

