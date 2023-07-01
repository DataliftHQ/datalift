import { ApplicationService } from './application.service.ts';

export interface Services {
  application: ApplicationService;
}

export const services: Services = {
  application: new ApplicationService(),
};
