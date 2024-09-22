CREATE TABLE IF NOT EXISTS account_confirmation (
	user_id bigint NOT NULL UNIQUE,
	token text NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT account_confirmation_pkey PRIMARY KEY (user_id),
	CONSTRAINT account_confirmation_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);