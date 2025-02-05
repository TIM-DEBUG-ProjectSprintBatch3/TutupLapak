-- Hash index untuk sku
CREATE INDEX IF NOT EXISTS idx_sku_hash ON products USING hash (sku);

-- Hash index untuk product_id
CREATE INDEX IF NOT EXISTS idx_product_id_hash ON products USING hash (id);

-- Hash index untuk category
CREATE INDEX IF NOT EXISTS idx_category_hash ON products USING hash (category);