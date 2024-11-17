export interface IConversation {
  session_id: string;
  title: string;
  maas: {
    id: number;
    conversation_session_id: string;
    maas_id: number;
    status: string;
    create_time: string;
    update_time: string;
  }[];
  last_message_at: string;
  messages: {
    id: number;
    conversation_session_id: string;
    message_type: string;
    content: string;
    context: string;
    expand: string;
    maas_id: number;
    sequence_number: number;
    create_time: string;
    parent_question_id: number | null;
  }[];
}
