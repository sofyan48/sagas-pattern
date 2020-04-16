CREATE TABLE tb_payment_model (
	id_payment_model INT8 NOT NULL DEFAULT unique_rowid(),
	nm_payment_model VARCHAR NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT "primary" PRIMARY KEY (id_payment_model ASC),
	FAMILY "primary" (id_payment_model, nm_payment_model, created_at, updated_at, deleted_at)
);