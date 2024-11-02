'use client';

import { memo, useMemo } from 'react';
import { Collapse, Image, Typography, Tooltip, Popover } from 'antd';
import { RightOutlined } from '@ant-design/icons';

import ReactMarkdown from 'react-markdown';
import rehypeKatex from 'rehype-katex';
import remarkMath from 'remark-math';
import remarkGfm from 'remark-gfm';
import remarkBreaks from 'remark-breaks';
import rehypeRaw from 'rehype-raw';
import rehypeHighlight from 'rehype-highlight';
import remarkDirective from 'remark-directive';
import remarkDirectiveRehype from 'remark-directive-rehype';

import 'katex/dist/katex.min.css';
import 'highlight.js/styles/atom-one-light.css';
import CodeBlock from './CodeBlock';
import { useStyles } from './style';
import { escapeBrackets, escapeMhchem, fixMarkdownBold } from './utils';

export interface IMarkdownProps {
  content: string;
  isGenerating?: boolean;
  className?: string;
  style?: React.CSSProperties;
}

const Markdown = ({ content, className, style }: IMarkdownProps) => {
  const { styles, cx } = useStyles();

  const escapedContent = useMemo(() => {
    return fixMarkdownBold(escapeMhchem(escapeBrackets(content))).replace(/\\n/g, '\n');
  }, [content]);

  const components = useMemo(
    () => ({
      a: (props: any) => (
        <Tooltip
          arrow={false}
          overlayStyle={{ cursor: 'pointer' }}
          title={
            <span onClick={() => window.open(props?.href)}>
              {props?.title || props?.href}
              <RightOutlined style={{ marginLeft: '4px', fontSize: '12px' }} />
            </span>
          }
        >
          <Typography.Link rel="noopener noreferrer" target="_blank" {...props} />
        </Tooltip>
      ),
      details: (props: any) => <Collapse {...props} />,
      img: (props: any) => <Image {...props} />,
      pre: (props: any) => <CodeBlock {...props} />,
      citation: (props: any) => {
        console.log('citation props =>', props);
        return (
          <Popover
            arrow={false}
            overlayStyle={{ cursor: 'pointer' }}
            overlayInnerStyle={{ maxWidth: '300px', cursor: 'pointer' }}
            title={
              <Typography.Paragraph ellipsis={{ rows: 1 }} onClick={() => window.open(props?.href)}>
                {props?.title}
              </Typography.Paragraph>
            }
            content={
              <Typography.Paragraph ellipsis={{ rows: 8 }} onClick={() => window.open(props?.href)}>
                {props?.snippet}
              </Typography.Paragraph>
            }
          >
            <span className={styles.citation}>{props?.children}</span>
          </Popover>
        );
      },
    }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  if (!content) return null;

  return (
    <Typography className={className} style={style}>
      <ReactMarkdown
        className={styles.markdown}
        remarkPlugins={[
          remarkMath,
          remarkGfm,
          remarkBreaks,
          remarkDirective,
          remarkDirectiveRehype as any,
        ]}
        rehypePlugins={[
          rehypeKatex,
          rehypeRaw,
          [
            rehypeHighlight,
            {
              detect: false,
              ignoreMissing: true,
            },
          ],
        ]}
        components={components}
      >
        {escapedContent}
      </ReactMarkdown>
    </Typography>
  );
};

export default memo(Markdown);
