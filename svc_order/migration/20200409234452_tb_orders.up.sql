CREATE TABLE tb_orders (
	id_order INT8 NOT NULL DEFAULT unique_rowid(),
	uuid VARCHAR NOT NULL,
	order_number VARCHAR NOT NULL,
	user_uuid VARCHAR NOT NULL,
	id_order_type INT8 NOT NULL,
	id_payment_model INT8 NOT NULL,
	created_at DATE NOT NULL DEFAULT now():::DATE,
	updated_at DATE NOT NULL DEFAULT now():::DATE,
	deleted_at DATE NULL,
	CONSTRAINT tb_orders_pk PRIMARY KEY (id_order ASC),
	CONSTRAINT tb_orders_fk FOREIGN KEY (id_order_type) REFERENCES tb_order_type(id_order_type) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT tb_orders_fk_payment_model FOREIGN KEY (id_payment_model) REFERENCES tb_payment_model(id_payment_model) ON DELETE CASCADE ON UPDATE CASCADE,
	UNIQUE INDEX tb_orders_un_uuid (uuid ASC),
	INDEX tb_orders_auto_index_tb_orders_fk (id_order_type ASC),
	INDEX tb_orders_auto_index_tb_orders_fk_payment_model (id_payment_model ASC),
	FAMILY "primary" (id_order, uuid, order_number, user_uuid, id_order_type, id_payment_model, created_at, updated_at, deleted_at)
);