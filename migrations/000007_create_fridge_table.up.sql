CREATE TABLE IF NOT EXISTS fridge(
	item_id INT,
	quantity REAL,
	purchase_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
	--CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(item_id));
	FOREIGN KEY(item_id) REFERENCES items(item_id) ON DELETE CASCADE);
