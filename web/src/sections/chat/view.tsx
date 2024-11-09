import { useCallback, useEffect, useMemo, useState } from 'react';

import Box from '@mui/material/Box';

import useScrollToBottom from 'src/hooks/useScrollToBottom';

import { DashboardContent } from 'src/layouts/dashboard';

import Chat from './ModelClass/BaseChat';
import { IChat, IChatProps, SessionStatuTypes } from './type';
import { Iconify } from 'src/components/iconify';

import ChatTrigger from './ChatTrigger';
import { Messages } from './Messages';
// import { modelList } from 'src/.api-key';
import { Avatar, ButtonBase, Paper, Typography } from '@mui/material';
import { useGetModelList } from 'src/actions/model';

// ----------------------------------------------------------------------

type Props = {
  title?: string;
};

export function MultiLLMChat({ title = 'Blank' }: Props) {
  const MaxChatAdmittance = 3;
  const [chats, setChats] = useState<any>([]);
  const [models, selectModels] = useState<number[]>([]);
  const [query, setQuery] = useState<string>('');

  const [shouldChatInfer, setShouldChatInfer] = useState<boolean>(false);

  const { scrollRef: scrollRef1 } = useScrollToBottom();
  const { scrollRef: scrollRef2 } = useScrollToBottom();
  const { scrollRef: scrollRef3 } = useScrollToBottom();

  const { modelList } = useGetModelList();

  const buildChats = useCallback(() => {
    // 根据models构建chats
    const chatInstance = models.map(
      (model: number) => modelList?.find((item: any) => item.id === model),
      []
    );
    const isModelExist = (model: number) => chats.some((chat: any) => chat.id === model);

    const newChats = chatInstance.map((chat: any) => {
      if (isModelExist(chat.id)) {
        const ret = chats.find((c: any) => c.id === chat.id);
        return ret as IChat;
      }

      const chatParams = {
        id: chat.id,
        chatId: chat.id,
        model: chat.model,
        api_key: chat.api_key,
        base_url: chat.base_url,
        brand: chat.brand,
        stream: chat.stream,
        name: chat.name,
        avatar: chat.avatar,
        messages: [],
        // 刷入额外参数
      };

      const chatItem: Chat = new Chat(chatParams as IChatProps);

      return chatItem;
    });
    if (newChats && newChats.length) {
      setChats(newChats);
    } else {
      setChats([]);
    }
  }, [chats, modelList, models]);

  const handleUserQuery = useCallback(
    (query: string) => {
      // eslint-disable-next-line no-debugger
      const isNotReady = chats.some(
        (chat: any) => chat.sessionStatus === SessionStatuTypes.running
      );
      if (isNotReady) {
        alert('模型回复问题中，请等待所有模型就绪后再发起新一轮问答');
        return;
      }
      chats.forEach((chat: any) => {
        const params: any = {
          ...chat,
        };

        console.log('🚀 ~ params:', params);

        // 添加系统人设
        if (params?.system_prompt) {
          // if (!isCurSystemPrompt(params?.system_prompt)) {
          //   params.messages.push({
          //     role: 'system',
          //     content: params.system_prompt,
          //   })
          // }
          if (params.messages.length > 0) {
            if (params.messages[0].role !== 'system') {
              params.messages.unshift({
                role: 'system',
                content: params.system_prompt,
              });
            } else {
              params.messages[0].content = params.system_prompt;
            }
          }
        }

        // Add user query to messages
        params.messages.push({
          // id: creatUuid(),
          content: query,
          role: 'user',
          // picture: params.curPicture,
        });
        // Add empty message for assistant to fill

        params.messages.push({
          // id: creatUuid(),
          // type: chat.model_type,
          // picture: chat.picture,
          content: '',
          role: 'assistant',
          loading: true,
          // question: query,
          // context: JSON.stringify(chat.messages),
        });

        chat.predict(params, (response: any, options?: any) => {
          setChats(
            chats.map((chatItem: any) => {
              // 清除图片

              if (chatItem.model === chat.model) {
                chatItem.messages = response;
                if (options) {
                  const { sessionStatus } = options;
                  chatItem.sessionStatus = sessionStatus;
                }
                return chatItem;
              }
              return chatItem;
            })
          );
        });

        // 将sessionStatus 设置为进行中
        setChats(
          chats.map((chatItem: any) => {
            if (chatItem.model === chat.model) {
              chatItem.sessionStatus = SessionStatuTypes.running;
              return chatItem;
            }
            return chatItem;
          })
        );

        // 重置query
        setQuery('');
      });
    },
    [chats]
  );

  const userQueryAdapter = useCallback(
    (queryString: string, pic?: string) => {
      setQuery(queryString);
      buildChats();
      setShouldChatInfer(true);
      // handleUserQuery(queryString);
    },
    [buildChats]
  );

  useEffect(() => {
    if (shouldChatInfer && query) {
      handleUserQuery(query);
      setShouldChatInfer(false); // Reset the flag
    }
  }, [setShouldChatInfer, query, handleUserQuery, shouldChatInfer]);

  const getChats = useMemo(
    () =>
      chats.map((chat: any, index: number) => {
        let scrollRef: any = null;
        if (index === 0) {
          scrollRef = scrollRef1;
        } else if (index === 1) {
          scrollRef = scrollRef2;
        } else if (index === 2) {
          scrollRef = scrollRef3;
        }

        return (
          <div ref={scrollRef} key={chat.chatId} className="flex-1 flex-shrink-0 max-w-screen-lg">
            <Messages
              className="flex flex-col w-full flex-1 max-w-chat  gap-4 mx-auto z-1"
              messages={chat.messages}
              model={chat.model}
              brand={chat.avatar}
              isStreaming
            />

            {chat.sessionStatus === SessionStatuTypes.running && (
              <div
                className="flex justify-center items-center mt-6 cursor-pointer"
                onClick={() => {
                  if (chat.abortPredict) {
                    chat.abortPredict();
                  }
                  setChats(
                    chats.map((chatItem: { model: string; sessionStatus: SessionStatuTypes }) => {
                      if (chatItem.model === chat.model) {
                        chat.abortController = new AbortController();
                        chatItem.sessionStatus = SessionStatuTypes.ready;
                        return chatItem;
                      }
                      return chatItem;
                    })
                  );
                }}
              >
                <div className="w-[108px] h-[36px]  rounded flex items-center justify-between py-4 px-4">
                  <div>停止回答</div>
                </div>
              </div>
            )}
          </div>
        );
      }),
    [chats, scrollRef1, scrollRef2, scrollRef3]
  );

  const triggerToolkit = useMemo(
    () => ({
      query,
      setQuery,
      trigger: userQueryAdapter,
      isShowNewContext: (() => {
        let flag = false;
        for (let i = 0; i < chats.length; i++) {
          if (chats[i].messages && chats[i].messages.length > 0) {
            flag = true;
            break;
          }
        }
        return flag;
      })(),

      newContext: () => {
        // 重置所有chats 的messages
        setChats(
          chats.map((chat: any) => {
            chat.messages = [];
            // 刷新图文模型的默认选项
            chat.picture = '';
            chat.curPicture = '';
            return chat;
          })
        );
      },
    }),
    [chats, query, userQueryAdapter]
  );

  return (
    <DashboardContent maxWidth="xl">
      <Box
        id="chatContainer"
        className=" flex flex-col  w-full
          justify-center items-center gap-4
        "
        sx={{
          height: 'calc(100vh - 114px)',
        }}
      >
        {chats?.length > 0 ? (
          <>
            <div className="w-full flex  gap-4  overflow-y-auto justify-center flex-1">
              {chats?.length > 0 ? (
                <div className="flex flex-1 gap-2   overflow-x-auto w-full p-2 justify-center ">
                  {getChats}
                </div>
              ) : null}
            </div>
            {/* <ChatTrigger {...triggerToolkit} /> */}
          </>
        ) : (
          <>
            {chats?.length === 0 && (
              <Box
                sx={{
                  width: '100%',
                  display: 'flex',
                  flexDirection: 'column',
                  justifyContent: 'center',
                  alignItems: 'center',
                  gap: 4,
                  color: 'var(--layout-nav-text-primary-color)',
                }}
              >
                <Typography
                  variant="h4"
                  noWrap
                  sx={{ color: 'var(--layout-nav-text-primary-color)' }}
                >
                  选择模型开始对话体验
                </Typography>

                <Box gap={2} display="grid" gridTemplateColumns="repeat(4, 1fr)">
                  {modelList.map((item) => (
                    <Paper
                      component={ButtonBase}
                      variant="outlined"
                      key={item.model}
                      onClick={() =>
                        selectModels((ids: number[]) => {
                          if (ids.includes(item.id)) {
                            return ids.filter((id) => id !== item.id);
                          }
                          return [...ids, item.id];
                        })
                      }
                      sx={{
                        py: 1.5,
                        px: 4,
                        borderRadius: 1,
                        typography: 'subtitle2',
                        flexDirection: 'column',
                        borderWidth: 2,
                        ...(models.includes(item.id) && {
                          borderColor: 'text.primary',
                          borderWidth: 2,
                        }),
                      }}
                    >
                      <Box>
                        <Avatar src={item.avatar}></Avatar>
                      </Box>
                      {item.model}
                    </Paper>
                  ))}
                </Box>
              </Box>
            )}
          </>
        )}
        <ChatTrigger {...triggerToolkit} />
      </Box>
    </DashboardContent>
  );
}
