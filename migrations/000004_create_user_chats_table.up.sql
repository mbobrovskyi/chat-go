CREATE TABLE IF NOT EXISTS user_chats (
    user_id BIGINT NOT NULL REFERENCES users ("id") ON UPDATE CASCADE ON DELETE CASCADE,
    chat_id BIGINT NOT NULL REFERENCES chats ("id") ON UPDATE CASCADE ON DELETE CASCADE,
    PRIMARY KEY ("user_id", "chat_id")
);