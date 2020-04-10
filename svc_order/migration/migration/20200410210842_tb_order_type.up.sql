CREATE TABLE tb_order_type (
	id_order_type INT8 NOT NULL DEFAULT unique_rowid(),
	order_type VARCHAR NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT "primary" PRIMARY KEY (id_order_type ASC),
	FAMILY "primary" (id_order_type, order_type, created_at, updated_at, deleted_at)
);