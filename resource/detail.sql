WITH conversation_details AS (
    SELECT
        c.id as conversation_id,
        c.title as conversation_title,
        c.create_time as conversation_created_at
    FROM conversations c
    WHERE c.id = :conversation_id
      AND c.status = 'active'
),
     model_messages AS (
         SELECT
             m.model_id,
             md.name as model_name,
             md.description as model_description,
             JSON_ARRAYAGG(
                     JSON_OBJECT(
                             'message_id', COALESCE(q.id, m.id),
                             'message_type', COALESCE(q.message_type, m.message_type),
                             'content', COALESCE(q.content, m.content),
                             'sequence_number', COALESCE(q.sequence_number, m.sequence_number),
                             'created_at', COALESCE(q.created_at, m.created_at)
                     )
                         ORDER BY COALESCE(q.sequence_number, m.sequence_number)
             ) as messages
         FROM messages m
                  JOIN models md ON m.model_id = md.id
                  JOIN conversation_models cm ON
             cm.conversation_id = m.conversation_id
                 AND cm.model_id = m.model_id
                 AND cm.status = 'active'
                  LEFT JOIN messages q ON q.id = m.parent_question_id
         WHERE m.conversation_id = :conversation_id
           AND m.message_type = 'answer'
         GROUP BY m.model_id, md.name, md.description
     )
SELECT
    cd.conversation_id,
    cd.conversation_title,
    cd.conversation_created_at,
    JSON_ARRAYAGG(
            JSON_OBJECT(
                    'model_id', mm.model_id,
                    'model_name', mm.model_name,
                    'model_description', mm.model_description,
                    'messages', mm.messages
            )
    ) as models
FROM conversation_details cd
         CROSS JOIN model_messages mm
GROUP BY
    cd.conversation_id,
    cd.conversation_title,
    cd.conversation_created_at;