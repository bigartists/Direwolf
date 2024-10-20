import { useState, useRef, memo } from 'react';
import { Button, Tooltip } from 'antd';
import { Copy, Check, ChevronDown, ChevronRight } from 'lucide-react';

import 'katex/dist/katex.min.css';
import 'highlight.js/styles/atom-one-light.css';
import { copyToClipboard } from '../copyToClipboard';
import { useStyles } from './style';

const CodeBlock = ({ children }: any) => {
  const ref = useRef<HTMLPreElement>(null);
  const [isCopied, setIsCopied] = useState(false);
  const [expand, setExpand] = useState(true);
  const { styles } = useStyles();

  if (!children) return;
  const { className } = children.props;

  const lang = className?.replace('language-', '') || 'text';

  const handleCopyClick = () => {
    setIsCopied(true);
    setTimeout(() => {
      setIsCopied(false);
    }, 2000);
  };

  const CopyIcon = (props: any) =>
    isCopied ? <Check color="#16a34a" {...props} /> : <Copy color="#777" {...props} />;
  const ExpandIcon = expand ? ChevronDown : ChevronRight;

  return (
    <div className={styles.codeBlock}>
      <div className={styles.codeBlockHeader}>
        <Button
          type="text"
          icon={
            <ExpandIcon
              size={14}
              color="#777"
              style={{
                verticalAlign: '-0.125em',
              }}
            />
          }
          size="small"
          onClick={() => setExpand(!expand)}
        />
        <span>{lang.replace('hljs', '')}</span>
        <Tooltip placement="left" title={isCopied ? '已复制' : '复制'}>
          <Button
            type="text"
            icon={
              <CopyIcon
                size={14}
                style={{
                  verticalAlign: '-0.125em',
                }}
              />
            }
            size="small"
            onClick={() => {
              copyToClipboard(ref.current?.innerText || '');
              handleCopyClick();
            }}
          />
        </Tooltip>
      </div>
      <pre style={expand ? {} : { height: 0, overflow: 'hidden' }} ref={ref}>
        {children}
      </pre>
    </div>
  );
};

export default CodeBlock;
