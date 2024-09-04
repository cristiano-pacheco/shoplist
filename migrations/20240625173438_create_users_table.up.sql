CREATE TABLE IF NOT EXISTS  users (
	id bigserial NOT NULL,
	"name" text NOT NULL,
	email text NOT NULL,
	password_hash text NOT NULL,
	is_activated bool NOT NULL DEFAULT false,
	reset_password_token text NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS user_email_idx ON users (email);