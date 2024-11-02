import { memo } from 'react';
// import { Markdown } from 'src/components/Markdown3';
import Markdown from 'src/components/Markdown';

interface AssistantMessageProps {
  content: string;
}

export const AssistantMessage = memo(({ content }: AssistantMessageProps) => (
  <div className="overflow-hidden w-full">
    <Markdown content={content}></Markdown>
    {/* <Markdown children={content}></Markdown> */}
  </div>
));
