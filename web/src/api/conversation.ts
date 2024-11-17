import useSWR, { mutate } from 'swr';
import { useCallback, useMemo, useState } from 'react';

import { endpoints, fetcher, postFetcher } from 'src/utils/axios';
import { IModelItem } from 'src/types/common';
import useSWRMutation from 'swr/mutation';
import { IChatProps } from 'src/sections/chat/type';

export function useGetConversations() {
  const URL = endpoints.conversation.list;
  const { data, isLoading, error, isValidating } = useSWR(URL, fetcher);

  const memoizedValue = useMemo(
    () => ({
      conversations: (data?.payload as any[]) || [],
      conversationsLoading: isLoading,
      conversationsError: error,
      conversationsValidating: isValidating,
      conversationsEmpty: !isLoading && !data?.payload?.length,
    }),
    [data, isLoading, error, isValidating]
  );
  return memoizedValue;
}

export const useCreateConversation = () => {
  const URL = endpoints.conversation.create;
  async function createConversation(
    url: string,
    {
      arg,
    }: {
      arg: {
        title: string;
        maas_ids: number[];
      };
    }
  ) {
    return postFetcher(url, arg);
  }
  const { trigger, isMutating } = useSWRMutation(URL, createConversation, {
    onSuccess: (data, variables) => console.log('onSuccess', data, variables),
    onError: (error, variables) => console.log('onError', error, variables),
  });
  return { trigger, isMutating };
};

export const useGetConversationById = (params?: { session_id: string }) => {
  const { session_id } = params || {};
  const shouldFetch = !!session_id;
  const URL = shouldFetch && endpoints.conversation.get(session_id);

  const { data, isLoading, error, isValidating, mutate } = useSWR(
    URL ? [URL, { params: '' }] : null,
    fetcher
  );

  const memoizedValue = useMemo(
    () => ({
      conversation: data?.payload || {},
      conversationLoading: isLoading,
      conversationError: error,
      conversationValidating: isValidating,
      refetch: () => mutate(),
    }),
    [data, isLoading, error, isValidating, mutate]
  );

  return memoizedValue;
};
