package data

/*

CREATE TABLE users (
	id BIGSERIAL PRIMARY KEY,
	registered_at BIGINT NOT NULL
);

CREATE TABLE chats (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	created_at BIGINT NOT NULL,
	updated_at BIGINT NOT NULL,
	deleted_at BIGINT
);

CREATE TABLE messages (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	chat_id BIGINT NOT NULL,
	role VARCHAR(32) NOT NULL,
	created_at BIGINT NOT NULL,
	deleted_at BIGINT,
	in_reply_to BIGINT,
	text TEXT,
	images VARCHAR(64)[],
	structured_output JSONB
);

CREATE TABLE requests (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	chat_id BIGINT NOT NULL,
	started_at BIGINT NOT NULL,
	finished_at BIGINT NOT NULL,
	latency BIGINT NOT NULL,
	chunks BIGINT NOT NULL,
	attempts BIGINT NOT NULL,
	language VARCHAR(32) NOT NULL,
	system_instruction TEXT NOT NULL,
	contents BIGINT[] NOT NULL,
	response BIGINT NOT NULL,
	finish_reason VARCHAR(32) NOT NULL,
	model VARCHAR(64) NOT NULL,
	prompt_tokens BIGINT NOT NULL,
	response_tokens BIGINT NOT NULL,
	price DOUBLE PRECISION NOT NULL
);

*/
