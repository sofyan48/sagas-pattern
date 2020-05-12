CREATE TABLE tb_login (
	id INT8 NOT NULL DEFAULT unique_rowid(),
	id_user INT8 NOT NULL,
	username VARCHAR NOT NULL,
	password STRING NOT NULL,
	id_roles INT8 NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	update_at TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	deleted_at TIMESTAMPTZ NULL,
	CONSTRAINT tb_login_pk PRIMARY KEY (id ASC),
	CONSTRAINT tb_login_fk FOREIGN KEY (id_user) REFERENCES tb_users(id_user) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT tb_login_fk_1 FOREIGN KEY (id_roles) REFERENCES tb_roles(id_roles) ON DELETE CASCADE ON UPDATE CASCADE,
	INDEX tb_login_auto_index_tb_login_fk (id_user ASC),
	INDEX tb_login_auto_index_tb_login_fk_1 (id_roles ASC),
	FAMILY "primary" (id, id_user, username, password, id_roles, created_at, update_at, deleted_at)
);
