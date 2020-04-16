CREATE TABLE tb_payment_status (
	id_payment_status INT8 NOT NULL DEFAULT unique_rowid(),
	nm_payment_status VARCHAR NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT tb_payment_status_pk PRIMARY KEY (id_payment_status ASC),
	FAMILY "primary" (id_payment_status, nm_payment_status, created_at, updated_at, deleted_at)
);