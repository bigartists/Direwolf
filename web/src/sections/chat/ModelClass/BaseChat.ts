import { fetchEventSource } from '@echofly/fetch-event-source';
import { IChatProps, SAFTY_TEXT, SessionStatuTypes } from '../type';
import { cloneDeep } from 'lodash';
import { CONFIG } from 'src/config-global';

export interface IChat {
  // property

  id: number; // fk，模型id
  name: string; // 模型别名
  avatar: string;
  model: string;
  messages: Array<{ content: string; role: string }>;
  stream: boolean;
  session_id: string;
  abortController: AbortController;
  sessionStatus?: SessionStatuTypes;
  // function

  predict: (params: any, cb: (params?: any) => void) => Promise<any>;
  abortPredict: (params?: any) => void;
  getPrompt(messages: any): string;
  // 遗漏的一个问题， predict时的四种不同的状态处理；还有返回的对话的渲染逻辑；
}

export default class Chat implements IChat {
  id: number;

  model: string;

  maas_id: number;

  stream: boolean;

  name: string;

  session_id: string;

  avatar: string;

  messages: Array<any> = [];

  sessionStatus?: SessionStatuTypes;

  abortController: AbortController;

  constructor(props: IChatProps) {
    this.id = props.id;
    this.maas_id = props.id;
    this.stream = props.stream;
    this.session_id = props.session_id;
    this.model = props.model;
    this.name = props.name;
    this.avatar = props.avatar;
    this.messages = props.messages || [];
    this.sessionStatus = props.sessionStatus;
    this.abortController = new AbortController(); // 用于终止fetch请求
  }

  async getToken() {
    const token = sessionStorage.getItem(CONFIG.ACCESS_TOKEN);
    return token;
  }

  getPrompt(messages: any): string {
    const targetMessages = messages.slice(0, -1);
    const lastMessage = targetMessages[targetMessages.length - 1] || {};
    if (lastMessage.role === 'user') {
      return lastMessage.content;
    }
    return '';
  }

  async predict(params: any, cb: (params?: any, options?: any) => void): Promise<any> {
    const token = await this.getToken();
    const newParams = {
      maas_id: params.maas_id,
      prompt: this.getPrompt(params.messages),
      session_id: params.session_id,
      context: JSON.stringify(params.messages),
      params: {
        model: params.model,
        messages: params.messages.slice(0, -1),
        stream: params.stream || true,
      },
    };

    let xRequestId: any = '';
    fetchEventSource('/api/v1/maas/completions', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newParams),
      signal: this.abortController.signal,
      openWhenHidden: true,
      onopen: async (resp: any) => {
        // debugger;
        const contentype = resp.headers.get('content-type') || '';
        xRequestId = resp.headers.get('X-Request-Id');
        if (resp.ok && !contentype.includes('text/event-stream')) {
          const responseData = await resp.json();
          if (responseData.code !== 0) {
            const { messages, options } = errorItemMessage(
              responseData?.error?.message || '抱歉，暂无法回答该问题。',
              this
            );
            cb && cb(messages, options);
          }
        } else if (!resp.ok && !contentype.includes('text/event-stream')) {
          const responseData = await resp.json();
          const { messages, options } = errorItemMessage(
            responseData?.error?.message || '抱歉，暂无法回答该问题。',
            this,
            xRequestId
          );
          cb && cb(messages, options);
        }
      },
      onmessage: (msg: any) => {
        // debugger;
        const { data } = msg;
        if (data === ' [DONE]') {
          this.abortPredict();
          const messages: any = cloneDeep(this.messages);
          const lastMessage: any = messages[messages.length - 1];

          messages.splice(messages.length - 1, 1, lastMessage);
          cb &&
            cb(messages, {
              sessionStatus: SessionStatuTypes.ready,
            });
          return;
        }
        const messages: any = cloneDeep(this.messages);
        let lastMessage = messages[messages.length - 1] || {};

        const { choices } = JSON.parse(data || '{}');
        const { delta, finish_reason } = choices?.[0] || {};
        console.log('🚀 ~ Chat ~ finish_reason:', finish_reason);
        // ----------------- 以下为文本问答的逻辑 -----------------
        switch (finish_reason) {
          // 命中风控
          case 'content_filter':
            lastMessage = {
              ...lastMessage,
              content: SAFTY_TEXT,
              loading: false,
              renderFeedback: {
                context: JSON.stringify(this.messages),
                id: xRequestId,
              },
            };
            messages.splice(messages.length - 1, 1, lastMessage);
            cb && cb(messages);
            break;
          default:
            lastMessage = {
              ...lastMessage,
              content: `${lastMessage.content}${delta?.content}`,
              loading: false,
            };
            messages.splice(messages.length - 1, 1, lastMessage);
            // todo 除了react component内容的更新，其他接口不建议更新类的值
            cb && cb(messages);
        }
      },
      onerror: (err: any) => {
        const { messages, options } = errorItemMessage(
          err?.msg || '抱歉，暂无法回答该问题。',
          this
        );
        cb && cb(messages, options);
        console.log('eventSource error: ', `${err}`);
        throw err;
      },
      onclose: () => {
        // debugger;
        console.log('eventSource close');
        const messages: any = cloneDeep(this.messages);
        const lastMessage: any = messages[messages.length - 1];
        lastMessage.renderFeedback = {
          context: JSON.stringify(this.messages),
          id: xRequestId,
        };
        messages.splice(messages.length - 1, 1, lastMessage);
        cb &&
          cb(messages, {
            sessionStatus: SessionStatuTypes.ready,
          });
      },
    });
  }

  abortPredict(cb?: (params: any) => void) {
    this.abortController.abort();
  }
}

function errorItemMessage(text: string, chat: IChat, xRequestId?: string) {
  const messages: any = cloneDeep(chat.messages);
  const errorItem = {
    role: 'assistant',
    content: text,
    loading: false,
    type: 'error',
    id: xRequestId,
  };

  messages.splice(messages.length - 1, 1, errorItem);

  return {
    messages,
    options: {
      sessionStatus: SessionStatuTypes.ready,
    },
  };
}
