import useSWR from 'swr';
import { useEffect, useMemo } from 'react';

import { endpoints, fetcher, postFetcher } from 'src/utils/axios';
import { IModelItem } from 'src/types/common';
import useSWRMutation from 'swr/mutation';
import { IChatProps } from 'src/sections/chat/type';
import { isArray } from 'lodash';
import { useMaasStore } from 'src/store/maas';

export function useGetMaasList() {
  const URL = endpoints.maas.list;
  const { setMaasList } = useMaasStore();
  const { data, isLoading, error, isValidating } = useSWR(URL, fetcher);

  useEffect(() => {
    if (isArray(data?.payload)) {
      setMaasList(data.payload);
    }
  }, [data, setMaasList]);

  const memoizedValue = useMemo(
    () => ({
      maasList: (data?.payload as IChatProps[]) || [],
      modelListLoading: isLoading,
      modelListError: error,
      modelListValidating: isValidating,
      modelListEmpty: !isLoading && !data?.payload?.length,
    }),
    [data, isLoading, error, isValidating]
  );
  return memoizedValue;
}

export const useImportMaas = () => {
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
