CREATE TABLE tb_payment (
	id_payment INT8 NOT NULL DEFAULT unique_rowid(),
	uuid VARCHAR NOT NULL,
	uuid_order VARCHAR NOT NULL,
	uuid_user VARCHAR NOT NULL,
	id_payment_status INT8 NOT NULL,
    id_payment_model INT8 NOT NULL,
	inquiry_number VARCHAR NOT NULL,
	bank_account_number VARCHAR NOT NULL,
	nm_bank VARCHAR NOT NULL,
	payment_total INT8 NOT NULL,
	change_total INT8 NOT NULL,
    due_date DATE NOT NULL,
	payment_order INT8 NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT tb_payment_pk PRIMARY KEY (id_payment ASC),
	CONSTRAINT tb_payment_fk FOREIGN KEY (id_payment_status) REFERENCES tb_payment_status(id_payment_status) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT tb_payment_fk_model FOREIGN KEY (id_payment_model) REFERENCES tb_payment_model(id_payment_model) ON DELETE CASCADE ON UPDATE CASCADE,
	INDEX tb_payment_auto_index_tb_payment_fk (id_payment_status ASC),
	INDEX tb_payment_auto_index_tb_payment_fk_model (id_payment_model ASC),
	FAMILY "primary" (id_payment, uuid, uuid_order, uuid_user, id_payment_status, inquiry_number, bank_account_number, nm_bank, payment_total, change_total, created_at, updated_at, deleted_at, due_date, id_payment_model, payment_order)
);