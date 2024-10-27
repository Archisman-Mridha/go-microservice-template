-- name: GetChats :many
SELECT
  chat_participants_left.chat_id AS id,
  chat_participants_right.user_id AS with_user_id,
  latest_message.id AS last_message_id,
  latest_message.message AS last_message,
  latest_message.sender_id AS last_message_sender_id,
  latest_message.sent_at AS last_message_sent_at
FROM
  chat_participants chat_participants_left
INNER JOIN
  chat_participants chat_participants_right ON
    chat_participants_left.chat_id = chat_participants_right.chat_id AND
    chat_participants_left.user_id != chat_participants_right.user_id
INNER JOIN
  LATERAL (
    SELECT
      messages.chat_id,
      messages.id,
      messages.message,
      messages.sender_id,
      messages.sent_at
    FROM
      messages
    WHERE
      chat_id = chat_participants_left.chat_id
    ORDER BY
      sent_at DESC
    LIMIT 1
  ) latest_message ON
    latest_message.chat_id = chat_participants_left.chat_id
WHERE
  chat_participants_left.user_id = $1
ORDER BY
  latest_message.sent_at DESC
LIMIT $2 OFFSET $3;

-- name: GetMessages :many
SELECT
  id,
  message,
  sender_id,
  sent_at
FROM
  messages
WHERE
  messages.chat_id = $1
ORDER BY
  sent_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateChat :one
INSERT INTO
  chats
DEFAULT VALUES
RETURNING id;

-- name: CreateChatParticipant :exec
INSERT INTO
  chat_participants(chat_id, user_id)
VALUES
  ($1, $2);

-- name: CreateMessage :exec
INSERT INTO
  messages(id, chat_id, message, sender_id, sent_at)
VALUES
  ($1, $2, $3, $4, $5);
