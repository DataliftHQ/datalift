import type { AxiosError } from 'axios';

import { DataliftError, DataliftErrorDetails, Help, isDataliftErrorDetails } from './errors.ts';
import { grpcResponseToError, isHelpDetails } from './errors.ts';

describe('datalift error', () => {
  const axiosError = {
    response: {
      status: 404,
      statusText: 'Not Found',
      data: {
        code: 5,
        message: 'Could not find resource',
      },
    },
  } as AxiosError;

  describe('returns a basic DataliftError object', () => {
    let err: DataliftError;
    beforeAll(() => {
      err = grpcResponseToError(axiosError);
    });

    it('with a error code', () => {
      expect(err.code).toBe(5);
    });

    it('with a error messsage', () => {
      expect(err.message).toBe('Could not find resource');
    });

    it('with a status code', () => {
      expect(err.status.code).toBe(404);
    });

    it('with a status text', () => {
      expect(err.status.text).toBe('Not Found');
    });

    it('without details', () => {
      expect(err.details).toBeUndefined();
    });
  });

  // describe('returns a detailed DataliftError object', () => {
  //   let err: DataliftError;
  //   beforeAll(() => {
  //     const complexAxiosError = { ...axiosError };
  //     complexAxiosError.response.data.details = [
  //       {
  //         '@type': 'types.googleapis.com/google.rpc.Help',
  //         links: [
  //           {
  //             description: 'This is a link',
  //             url: 'https://www.clutch.sh',
  //           },
  //         ],
  //       },
  //     ];
  //     err = grpcResponseToError(complexAxiosError);
  //   });

  //   it('with a list of details', () => {
  //     expect(err.details).toHaveLength(1);
  //   });

  //   it('with correct typing', () => {
  //     const helpDetails = err.details[0] as Help;
  //     expect(helpDetails.links).toHaveLength(1);
  //   });
  // });
});

describe('isHelpDetails', () => {
  it('returns true for help details', () => {
    const details = {
      _type: 'types.googleapis.com/google.rpc.Help',
      links: [
        {
          description: 'Please file a ticket here for more help.',
          url: 'https://www.example.com',
        },
      ],
    } as Help;

    expect(isHelpDetails(details)).toBe(true);
  });

  it('returns false for non-help details', () => {
    const details = {
      type: 'unknownType',
      something: [
        {
          key: 'value',
        },
      ],
    };
    expect(isHelpDetails(details)).toBe(false);
  });
});

describe('isDataliftErrorDetails', () => {
  it('returns true for Datalift specific error details', () => {
    const details = {
      _type: 'type.googleapis.com/datalift.api.v1.ErrorDetails',
      wrapped: [
        {
          code: 2,
          message: 'core-staging-0: yikes',
        },
        {
          code: 16,
          message: 'core-staging-1: nono',
        },
      ],
    } as DataliftErrorDetails;

    expect(isDataliftErrorDetails(details)).toBe(true);
  });

  it('returns false for non-Datalift specific error details', () => {
    const details = {
      type: 'unknownType',
      something: [
        {
          key: 'value',
        },
      ],
    };
    expect(isDataliftErrorDetails(details)).toBe(false);
  });
});
