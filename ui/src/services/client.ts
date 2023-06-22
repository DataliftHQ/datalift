import type { AxiosError, AxiosResponse } from 'axios';
import axios from 'axios';

import type { DataliftError } from './errors.ts';
import { grpcResponseToError } from './errors.ts';

/**
 * HTTP response status.
 *
 * Responses are grouped in five classes:
 *  - Informational responses (100–199)
 *  - Successful responses (200–299)
 *  - Redirects (300–399)
 *  - Client errors (400–499)
 *  - Server errors (500–599)
 */
export interface HttpStatus {
  code: number;
  text: string;
}

const successInterceptor = (response: AxiosResponse) => {
  return response;
};

const errorInterceptor = (error: AxiosError): Promise<DataliftError> => {
  const response = error?.response as AxiosResponse;

  if (response === undefined) {
    const clientError = {
      status: {
        code: 500,
        text: 'Client Error',
      },
      message: error.message,
    } as DataliftError;
    return Promise.reject(clientError);
  }

  // This section handles authentication redirects.
  if (response?.status === 401) {
    // TODO: turn this in to silent refresh once refresh tokens are supported.
    const redirectUrl = window.location.pathname + window.location.search;
    window.location.href = `/auth/login?redirect_url=${encodeURIComponent(redirectUrl)}`;
  }

  // we are guaranteed to have a response object on the error from this point on
  // since we have already accounted for axios errors.
  const responseData = response?.data;

  // if the response data has a code on it we know it's a gRPC response.
  let err;
  if (responseData?.code !== undefined) {
    err = grpcResponseToError(error);
  } else {
    // TODO: test with non grpc errors
    const message =
      typeof error.response?.data === 'string' ? error.response.data : error?.message || error?.response?.statusText;
    err = {
      status: {
        code: error?.response?.status,
        text: error?.response?.statusText,
      } as HttpStatus,
      message,
      data: responseData,
    } as DataliftError;
  }
  return Promise.reject(err);
};

const createClient = () => {
  const axiosClient = axios.create({
    // n.b. the client will treat any response code >= 400 as an error and apply the error interceptor.
    validateStatus: (status) => {
      return status < 400;
    },
  });
  axiosClient.interceptors.response.use(successInterceptor, errorInterceptor);

  return axiosClient;
};

const client = createClient();

export { client as default, errorInterceptor, successInterceptor };
