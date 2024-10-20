import { useEffect, useMemo, useState } from 'react';

import useChineseInput from 'src/hooks/useChineseInput';
import doubleStar from 'src/assets/images/v3/icon/double_star.svg';

import { Button, Fab, Input, Stack, TextField } from '@mui/material';

const defaultInputLen = 2000;

const ChatTrigger = ({ query, setQuery, trigger, newContext }: any) => {
  const { curValue, handleChange, handleCompositionStart, handleCompositionEnd } = useChineseInput(
    query,
    (value: string) => {
      setQuery(value);
    }
  );

  const onEnter = () => {
    if (!query) {
      alert('请输入问题');
      return;
    }

    if (trigger) {
      trigger(query);
    }
  };

  return (
    <div className="w-full flex-shrink-0">
      <div className="flex justify-start mb-2">
        <Button
          variant="outlined"
          onClick={newContext}
          className="text-blue-500 border-blue-500 hover:bg-blue-500 hover:bg-opacity-10"
        >
          新建对话
        </Button>
      </div>
      <div>
        <div className="relative flex-1">
          <TextField
            value={query}
            onChange={(event: any) => setQuery(event.target.value)}
            onCompositionStart={handleCompositionStart}
            placeholder="请输入问题，可通过 Shift + 回车换行"
            multiline
            minRows={2}
            maxRows={5}
            className="w-full rounded-lg"

            // 不显示输入框拉大缩小的按钮
            // onPressEnter={(e: any) => {
            //   if (e.key === 'Enter' && e.shiftKey) return;
            //   // eslint-disable-next-line

            //   if (e.keyCode !== 13 || e.isComposing) return;
            //   e.preventDefault();
            //   onEnter();
            // }}
          />

          <div className="absolute bottom-5 right-4">
            <Stack direction="row" spacing={2}>
              {/* <div>
                {query.length}/{defaultInputLen}
              </div> */}

              <Fab onClick={onEnter} color="inherit" size="small">
                <img src="/assets/icons/apps/send.svg" />
              </Fab>
            </Stack>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ChatTrigger;
