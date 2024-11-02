import { fetchEventSource } from '@echofly/fetch-event-source';
import { IChat, IChatProps, SAFTY_TEXT, SessionStatuTypes } from '../type';
import { cloneDeep } from 'lodash';
import { CONFIG } from 'src/config-global';

export default class Chat implements IChat {
  id: number;

  model_id: number;

  chatId: string;

  model: string;

  api_key: string;

  base_url: string;

  stream: boolean;

  brand: string;

  messages: Array<any> = [];

  sessionStatus?: SessionStatuTypes;

  abortController: AbortController;

  constructor(props: IChatProps) {
    this.id = props.id;
    this.model_id = props.id;
    this.chatId = props.chatId;
    this.model = props.model;
    this.api_key = props.api_key;
    this.base_url = props.base_url;
    this.stream = props.stream;
    this.brand = props.brand;
    this.messages = props.messages || [];
    this.sessionStatus = props.sessionStatus;
    this.abortController = new AbortController(); // 用于终止fetch请求
  }

  async getToken() {
    const token = sessionStorage.getItem(CONFIG.ACCESS_TOKEN);
    return token;
  }

  async predict(params: any, cb: (params?: any, options?: any) => void): Promise<any> {
    const token = await this.getToken();
    // const newParms = cloneDeep(params);
    const newParams = {
      base_url: params.base_url,
      api_key: params.api_key,
      model: params.model,
      prompt: '该字段没啥用，后面优化掉', // 该字段没啥用，后面优化掉；
      params: {
        model: params.model,
        messages: params.messages.slice(0, -1),
        stream: params.stream || true,
      },
    };

    let xRequestId: any = '';
    fetchEventSource('/api/v1/model/invoke', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newParams),
      signal: this.abortController.signal,
      openWhenHidden: true,
      onopen: async (resp: any) => {
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
        const { data } = msg;
        if (data === ' [DONE]') {
          this.abortPredict();
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
          // 默认显示
          default:
            lastMessage = {
              ...lastMessage,
              content: `${lastMessage.content}${delta?.content}`,
              loading: false,
              // id: xRequestId,
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
