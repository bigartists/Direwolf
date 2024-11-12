import useSWR from 'swr';
import { useMemo } from 'react';

import { endpoints, fetcher, postFetcher } from 'src/utils/axios';
import { IModelItem } from 'src/types/common';
import useSWRMutation from 'swr/mutation';
import { IChatProps } from 'src/sections/chat/type';

export function useGetModelList() {
  const URL = endpoints.maas.list;
  const { data, isLoading, error, isValidating } = useSWR(URL, fetcher);

  const memoizedValue = useMemo(
    () => ({
      modelList: (data?.payload as IChatProps[]) || [],
      modelListLoading: isLoading,
      modelListError: error,
      modelListValidating: isValidating,
      modelListEmpty: !isLoading && !data?.payload?.length,
    }),
    [data, isLoading, error, isValidating]
  );
  return memoizedValue;
}

export const useCreateModel = () => {
  const URL = endpoints.maas.create;
  async function createModel(
    url: string,
    {
      arg,
    }: {
      arg: IModelItem;
    }
  ) {
    return postFetcher(url, arg);
  }
  const { trigger, isMutating } = useSWRMutation(URL, createModel, {
    onSuccess: (data, variables) => console.log('onSuccess', data, variables),
    onError: (error, variables) => console.log('onError', error, variables),
  });
  return { trigger, isMutating };
};
