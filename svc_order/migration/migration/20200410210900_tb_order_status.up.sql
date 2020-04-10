CREATE TABLE tb_order_status (
	id_order_status INT8 NOT NULL DEFAULT unique_rowid(),
	nm_status_order VARCHAR NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT "primary" PRIMARY KEY (id_order_status ASC),
	FAMILY "primary" (id_order_status, nm_status_order, created_at, updated_at, deleted_at)
);
