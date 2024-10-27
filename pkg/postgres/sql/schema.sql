CREATE TABLE chats (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT(now( ) AT TIME ZONE 'UTC'),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT(now( ) AT TIME ZONE 'UTC')
);

CREATE TABLE chat_participants (
  chat_id INTEGER NOT NULL REFERENCES chats(id),
  user_id INTEGER NOT NULL,

  /* This ensures that a user cannot join a chat more than once. */
  PRIMARY KEY (user_id, chat_id)
);
CREATE INDEX chat_pariticipants_user_id_idx ON chat_participants (user_id);
CREATE INDEX chat_pariticipants_chat_id_idx ON chat_participants (chat_id);

CREATE TABLE messages (
  id UUID PRIMARY KEY,
  chat_id INTEGER NOT NULL REFERENCES chats(id),
  message VARCHAR(2000) NOT NULL,
  sender_id INTEGER NOT NULL,
  sent_at TIMESTAMP WITH TIME ZONE DEFAULT(now( ) AT TIME ZONE 'UTC')
);
CREATE INDEX messages_chat_id_sent_at_idx ON messages (chat_id, sent_at);
