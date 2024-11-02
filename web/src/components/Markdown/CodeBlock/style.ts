import { createStyles } from 'antd-style';

export const useStyles = createStyles(({ token, css, cx, prefixCls }: any) => {
  return {
    codeBlock: cx(
      `${prefixCls}-codeBlock`,
      css`
        background-color: rgba(0, 0, 0, 0.03);
        border-radius: 8px;
        pre {
          margin: 0;
          padding: 0;
          overflow-x: auto;
          background: transparent;
          border: none;
        }

        code {
          display: block !important;
          padding: 1em !important;
        }
      `
    ),
    codeBlockHeader: cx(
      `${prefixCls}-codeBlockHeader`,
      css`
        padding-block: 4px;
        padding-inline: 1em;
        box-sizing: border-box;
        background: rgba(0, 0, 0, 0.015);
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        span {
          font-size: 14px;
          color: #999;
        }
      `
    ),
  };
});
