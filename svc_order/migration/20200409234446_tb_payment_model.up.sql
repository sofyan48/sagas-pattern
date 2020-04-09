CREATE TABLE tb_payment_model (
	id_payment_model INT8 NOT NULL DEFAULT unique_rowid(),
	payment_model_name VARCHAR NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT "primary" PRIMARY KEY (id_payment_model ASC),
	FAMILY "primary" (id_payment_model, payment_model_name, created_at, updated_at, deleted_at)
);