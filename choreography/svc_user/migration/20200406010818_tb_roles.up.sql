CREATE TABLE tb_roles (
	id_roles INT8 NOT NULL DEFAULT unique_rowid(),
	nm_roles VARCHAR(50) NULL,
	acl STRING NOT NULL,
	created_at DATE NULL DEFAULT now():::DATE,
	update_at DATE NOT NULL,
	deleted_at DATE NULL,
	CONSTRAINT tb_roles_pk PRIMARY KEY (id_roles ASC),
	FAMILY "primary" (id_roles, nm_roles, acl, created_at, update_at, deleted_at)
);