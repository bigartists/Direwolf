import { createStyles } from 'antd-style'

export const useStyles = createStyles(({ token, css, cx, prefixCls }: any) => {
  return {
    markdown: cx(
      `${prefixCls}-markdown`,
      css`
        color: ${token.colorText};
        position: relative;
        /* overflow: auto;
        border-radius: ${token.borderRadius}px; */
        
        h1,
        h2,
        h3,
        h4,
        h5 {
          font-weight: 600;
        }

        h1 { font-size: 26px; }
        h2 { font-size: 21px; }
        h3 { font-size: 17px; }
        h4 { font-size: 15px; }
        h5 { font-size: 14px; }

        p {
          margin-block: 0 0;
          font-size: inherit;
          line-height: 1.57;
          color: ${token.colorText};

          + * {
            margin-block-end: 1em;
          }

          &:not(:last-child) {
            margin-block-end: 1.5em;
          }
        }

        blockquote {
          margin-block: 16px;
          margin-inline: 0;
          padding-block: 0;
          padding-inline: 12px;

          p {
            font-style: italic;
            color: ${token.colorTextDescription};
          }
        }

        a {
          color: ${token.colorLink};

          &:hover {
            color: ${token.colorLinkHover};
          }

          &:active {
            color: ${token.colorLinkActive};
          }
        }

        img {
          max-width: 100%;
          width: 300px !important;
        }

        table {
          border-spacing: 0;

          width: 100%;
          margin-block: 1em 1em;
          margin-inline: 0 0;
          padding: 8px;

          border: 1px solid ${token.colorBorderSecondary};
          border-radius: ${token.borderRadius}px;

          code {
            display: inline-flex;
          }
        }

        th,
        td {
          padding-block: 10px 10px;
          padding-inline: 16px 16px;
        }

        thead {
          tr {
            th {
              background: ${token.colorFillTertiary};

              &:first-child {
                border-start-start-radius: ${token.borderRadius}px;
                border-end-start-radius: ${token.borderRadius}px;
              }

              &:last-child {
                border-start-end-radius: ${token.borderRadius}px;
                border-end-end-radius: ${token.borderRadius}px;
              }
            }
          }
        }

        /* https://github.com/orgs/remarkjs/discussions/1262 */
        li p {
          display: inline;
        }

        > ul > li {
          list-style-type: disc;
        }

        > ol,
        > ul {
          > li {
            margin-inline-start: 2em;
            line-height: 1.8;

            &::marker {
              color: ${token.cyan10A} !important;
            }
          }
        }

        ol,
        ul {
          > li::marker {
            color: ${token.colorTextDescription};
          }
        }

        details {
          margin-block-end: 1em;
          padding-block: 12px;
          padding-inline: 16px;

          background: ${token.colorFillTertiary};
          border: 1px solid ${token.colorBorderSecondary};
          border-radius: ${token.borderRadiusLG}px;

          transition: all 400ms ${token.motionEaseOut};
        }

        details[open] {
          summary {
            padding-block-end: 12px;
            border-block-end: 1px solid ${token.colorBorder};
          }
        }
      `,
    ),
    citation: cx(
      css`
        cursor: pointer;
        user-select: none;
        box-sizing: border-box;
        position: relative;
        display: inline-block;
        border-radius: 50px;
        padding: 0 4.5px;
        top: -2px;
        background: #4f5866;
        color: #fff;
        font-size: 10px;
        line-height: 15px;
        font-weight: 600;
        height: 15px;
        min-width: 15px;
        margin-left: 4px;
      `,
    )
  }
})
