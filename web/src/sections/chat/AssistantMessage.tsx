import { Box } from '@mui/material';
import { memo } from 'react';
// import { Markdown } from 'src/components/Markdown3';
import Markdown from 'src/components/Markdown';

interface AssistantMessageProps {
  content: string;
}

export const AssistantMessage = memo(({ content }: AssistantMessageProps) => (
  <Box
    className="overflow-hidden w-full  p-4 rounded-xl"
    sx={{
      // bgcolor: 'background.neutral',
      color: 'grey.800',
    }}
  >
    <Markdown content={content}></Markdown>
    {/* <Markdown children={content}></Markdown> */}
  </Box>
));
