export enum RoleTypes {
  user = 'user',
  assistant = 'assistant',
}

export enum SessionStatuTypes {
  running = 'running',
  pending = 'add',
  ready = 'finish',
  error = 'error',
  risk = 'risk',
}

export const SAFTY_TEXT =
  '非常抱歉，这个问题我暂时无法回答，如果您有其他的问题咨询，我非常乐意帮助你';

export const ERR_TEXT = '抱歉，暂无法回答该问题。';

export interface IChat {
  // property

  id: number; // fk，模型id
  chatId: string; // 会话id
  name: string; // 模型别名
  model: string;
  api_key: string;
  base_url: string;
  avatar: string;
  messages: Array<{ content: string; role: string }>;
  stream: boolean;
  abortController: AbortController;
  sessionStatus?: SessionStatuTypes;
  // function
  predict: (params: any, cb: (params?: any) => void) => Promise<any>;
  abortPredict: (params?: any) => void;
  // 遗漏的一个问题， predict时的四种不同的状态处理；还有返回的对话的渲染逻辑；
}

export interface IChatState {
  maxChatAdmittance: number;
  curChat: IChatProps;
  chatList: IChatProps[];
}

export type IChatProps = Omit<Omit<Omit<IChat, 'predict'>, 'abortPredict'>, 'abortController'>;

export type IMultiChatProps = IChatProps & { context: string };

/**
 * 场景1： 当选择一个模型并发起一个对话问答时
 *
 * 1. 跳转到对话页；
 * 2. 基于被选择的模型，基于IChat接口 初始化 Chat 对象；
 * 3. 通过 Chat 对象，发起对话请求；对话结果存储至messages对象中；
 * 4. 基于 message对象，渲染对话视口，基于curChat，渲染对话控制台参数界面；
 * 5. 当前离开对话页时，确认提醒，销毁predict请求；清空chatList；
 *
 *
 * 场景2： 当选择同一modelModalType下多个模型并发起一个对话问答时
 *
 * 1. 跳转到对话页；
 * 2. 基于被选择的模型，基于IChat接口 初始化多个 Chat 对象；
 * 3. 通过 Chat 对象，发起对话请求；对话结果存储至自身的messages对象中；
 * 4. 基于 message对象，渲染对话视口，基于curChat，渲染对话控制台参数界面；
 * 5. 当前离开对话页时，确认提醒，销毁所有的predict请求；清空所有的chatList；
 */
