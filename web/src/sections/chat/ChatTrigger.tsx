import { useEffect, useMemo, useState } from 'react';

import useChineseInput from 'src/hooks/useChineseInput';
import doubleStar from 'src/assets/images/v3/icon/double_star.svg';
import { DashboardContent } from 'src/layouts/dashboard';
import { Box, Button, Fab, Input, Stack, TextField } from '@mui/material';
import { Iconify } from 'src/components/iconify';

const defaultInputLen = 2000;

const ChatTrigger = ({ query, setQuery, trigger, newContext, isShowNewContext }: any) => {
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
    <div className=" flex-shrink-0 w-[768px]">
      <div className="flex justify-start mb-2 ">
        {isShowNewContext ? (
          <Button
            variant="outlined"
            // variant="contained"
            onClick={newContext}
            // startIcon={<Iconify icon="mingcute:add-line" />}
            startIcon={<Iconify icon="mingcute:chat-1-line" />}
            className="text-blue-500 border-blue-500 hover:bg-blue-500 hover:bg-opacity-10"
          >
            新建对话
          </Button>
        ) : null}
      </div>

      <Box
        // className="flex-1 py-3 pr-3 pl-6  w-full rounded-[28px] flex justify-between items-end "
        className="flex-1 py-3 pr-3 pl-6  w-full rounded-[8px] flex justify-between items-end "
        sx={{
          bgcolor: 'background.neutral',
        }}
      >
        <TextField
          variant="standard"
          value={query}
          onChange={(event: any) => setQuery(event.target.value)}
          onCompositionStart={handleCompositionStart}
          placeholder="请输入问题，可通过 Shift + 回车换行"
          multiline
          // minRows={1}
          maxRows={5}
          className="flex-1"
          InputProps={{
            disableUnderline: true,
          }}

          // 不显示输入框拉大缩小的按钮
          // onPressEnter={(e: any) => {
          //   if (e.key === 'Enter' && e.shiftKey) return;
          //   // eslint-disable-next-line

          //   if (e.keyCode !== 13 || e.isComposing) return;
          //   e.preventDefault();
          //   onEnter();
          // }}
        />

        <div className="cursor-pointer">
          <div onClick={onEnter}>
            {/* <div>
                {query.length}/{defaultInputLen}
              </div> */}

            {/* <Fab onClick={onEnter} color="inherit" size="small"> */}
            <img src="/assets/icons/apps/send.svg" className="w-[32px] h-[32px]" />
            {/* </Fab> */}
          </div>
        </div>
      </Box>
    </div>
  );
};

export default ChatTrigger;
