CREATE TABLE tb_users (
	id_user INT8 NOT NULL DEFAULT unique_rowid(),
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NULL,
	handphone VARCHAR(14) NOT NULL,
	address STRING NOT NULL,
	site_profil VARCHAR(50) NULL,
	email VARCHAR(100) NOT NULL,
	city VARCHAR(100) NULL,
	province VARCHAR(100) NULL,
	created_at DATE NOT NULL,
	updated_at DATE NOT NULL,
	deleted_at DATE NULL,
	district VARCHAR(100) NULL,
	CONSTRAINT tb_users_pk PRIMARY KEY (id_user ASC),
	UNIQUE INDEX tb_users_un1 (site_profil ASC),
	FAMILY "primary" (id_user, first_name, last_name, handphone, address, site_profil, email, city, province, created_at, updated_at, deleted_at, district)
);
