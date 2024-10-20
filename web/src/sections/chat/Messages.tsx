import type { Message } from 'ai';

import React from 'react';

import userAvartar from 'src/assets/images/v3/icon/user.svg';
import cuAvatar from 'src/assets/images/v3/icon/taichu2.svg';

import { classNames } from 'src/components/Markdown2/classNames';

import { UserMessage } from './UserMessage';
import { AssistantMessage } from './AssistantMessage';
import { Avatar, Box, Stack, Typography } from '@mui/material';

interface MessagesProps {
  id?: string;
  className?: string;
  isStreaming?: boolean;
  messages?: Message[];
  model?: string;
  brand?: string;
}

export const Messages = React.forwardRef<HTMLDivElement, MessagesProps>(
  (props: MessagesProps, ref) => {
    const { id, isStreaming = false, model, brand, messages = [] } = props;

    return (
      <div id={id} ref={ref} className={props.className}>
        {messages.length > 0
          ? messages.map((message, index) => {
              const { role, content } = message;
              const isUserMessage = role === 'user';
              const isAssistant = role === 'assistant';
              const isFirst = index === 0;
              const isLast = index === messages.length - 1;

              return (
                <div
                  key={index}
                  className={classNames('flex gap-4  w-full rounded-[calc(0.75rem-1px)]', {
                    'bg-bolt-elements-messages-background':
                      isUserMessage || !isStreaming || (isStreaming && !isLast),
                    'bg-gradient-to-b from-bolt-elements-messages-background from-30% to-transparent':
                      isStreaming && isLast,
                    'mt-4': !isFirst,
                  })}
                >
                  {/* {isUserMessage && (
                    <div className="flex items-center justify-center w-[34px] h-[34px] overflow-hidden bg-white text-gray-600 rounded-full shrink-0 self-start">
                      <Avatar alt="Remy Sharp" src={userAvartar} />
                    </div>
                  )} */}

                  <div className="grid grid-col-1 w-full">
                    {isUserMessage ? (
                      <UserMessage content={content} />
                    ) : (
                      <Stack direction="column" spacing={2}>
                        <Box>
                          <Stack direction="row" spacing={2} className="flex items-center">
                            {isAssistant && (
                              <div className="flex items-center justify-center w-[34px] h-[34px] overflow-hidden bg-white text-gray-600 rounded-full shrink-0 self-start">
                                <Avatar alt="Remy Sharp" src={brand} />
                              </div>
                            )}
                            <Typography
                              variant="body2"
                              noWrap
                              sx={{ color: 'var(--layout-nav-text-disabled-color)' }}
                            >
                              {model}
                            </Typography>
                          </Stack>
                        </Box>
                        <AssistantMessage content={content} />
                      </Stack>
                    )}
                  </div>
                </div>
              );
            })
          : null}
        {isStreaming && (
          <div className="text-center w-full text-bolt-elements-textSecondary i-svg-spinners:3-dots-fade text-4xl mt-4"></div>
        )}
      </div>
    );
  }
);
