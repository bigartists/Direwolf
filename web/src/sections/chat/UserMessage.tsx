import { Box } from '@mui/material';
import { Markdown } from 'src/components/Markdown2';

interface UserMessageProps {
  content: string;
}

export const MODIFICATIONS_TAG_NAME = 'bolt_file_modifications';

export const modificationsRegex = new RegExp(
  `^<${MODIFICATIONS_TAG_NAME}>[\\s\\S]*?<\\/${MODIFICATIONS_TAG_NAME}>\\s+`,
  'g'
);

export function UserMessage({ content }: UserMessageProps) {
  return (
    <Box
      className="overflow-hidden p-4 rounded-xl"
      sx={{
        bgcolor: { md: 'background.neutral' },
      }}
    >
      <Markdown limitedMarkdown>{sanitizeUserMessage(content)}</Markdown>
    </Box>
  );
}

function sanitizeUserMessage(content: string) {
  return content.replace(modificationsRegex, '').trim();
}
