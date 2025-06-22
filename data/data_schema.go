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
	deleted_at BIGINT,
	name VARCHAR(255)
);

CREATE TABLE messages (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	chat_id BIGINT NOT NULL,
	role VARCHAR(255) NOT NULL,
	created_at BIGINT NOT NULL,
	deleted_at BIGINT,
	in_reply_to BIGINT,
	text TEXT,
	images TEXT[],
	follow_ups TEXT[]
);

CREATE TABLE requests (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	chat_id BIGINT NOT NULL,
	started_at BIGINT NOT NULL,
	finished_at BIGINT NOT NULL,
	latency BIGINT NOT NULL,
	chunks BIGINT,
	errors BIGINT,
	language VARCHAR(255) NOT NULL,
	system_instruction TEXT NOT NULL,
	contents BIGINT[] NOT NULL,
	response BIGINT NOT NULL,
	finish_reason VARCHAR(255) NOT NULL,
	model VARCHAR(255) NOT NULL,
	cached_tokens BIGINT,
	non_cached_tokens BIGINT NOT NULL,
	tool_prompt_tokens BIGINT,
	thought_tokens BIGINT,
	response_tokens BIGINT NOT NULL,
	price DOUBLE PRECISION NOT NULL
);

*/
