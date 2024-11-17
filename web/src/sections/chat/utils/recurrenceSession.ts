import { groupBy, isArray } from 'lodash';
import { IConversation } from 'src/types';

export function recurrenceSession(conversation: IConversation): {
  messages: Record<string, { context: string; content: string }[]>;
  maas: number[];
  session: string;
} {
  const { session_id, maas, messages } = conversation;
  const maasret = isArray(maas) ? maas.map((m) => m.maas_id) : [];

  // group by maas_id
  const chats = isArray(messages) ? groupBy(messages, 'maas_id') : {};

  const newChats = messages.reduce((acc, message) => {
    const { maas_id } = message;
    if (!acc[maas_id]) {
      acc[maas_id] = [];
    }
    const value = acc[maas_id];
    if (value.length === 0) {
      acc[maas_id] = [message];
    }
    if (value.length === 1 && message.sequence_number > value[0].sequence_number) {
      acc[maas_id] = [message];
    }
    return acc;
  }, {} as any);
  return {
    messages: newChats,
    maas: maasret,
    session: session_id,
  };
}

export function buildMessages(target: { context: string; content: string }) {
  try {
    const context = JSON.parse(target.context);
    if (isArray(context)) {
      const last = context.pop();
      last.content = target.content;
      delete last.loading;
      context.push(last);
    }
    return context;
  } catch (error) {
    console.error(error);
  }
}
