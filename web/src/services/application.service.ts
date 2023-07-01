import { GetApplicationRequest, GetApplicationResponse } from '../api/application/v1/application_pb';
import client from './client.ts';

export class ApplicationService {
  public get(request: GetApplicationRequest): Promise<GetApplicationResponse> {
    return client.get(`/api/v1/application/${request.id}`).then((res) => {
      return res.data as GetApplicationResponse;
    });
  }
}
