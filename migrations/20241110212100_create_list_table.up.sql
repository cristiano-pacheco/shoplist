CREATE TABLE IF NOT EXISTS category (
	id bigserial PRIMARY KEY,
	user_id bigint NOT NULL,
	name text NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT category_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS list (
	id bigserial PRIMARY KEY,
	user_id bigint NOT NULL,
	name text NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT list_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS list_item (
	id bigserial PRIMARY KEY,
	user_id bigint NOT NULL,
	list_id bigint NOT NULL,
	category_id bigint NOT NULL,
	name text NOT NULL,
	current_quantity integer NOT NULL,
	minimum_quantity integer NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT list_item_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT list_item_list_id_foreign FOREIGN KEY (list_id) REFERENCES list(id) ON DELETE CASCADE,
	CONSTRAINT list_item_category_id_foreign FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);