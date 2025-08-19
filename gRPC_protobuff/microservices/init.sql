CREATE DATABASE IF NOT EXISTS `order`;
CREATE DATABASE IF NOT EXISTS `payment`;
CREATE DATABASE IF NOT EXISTS `shipping`;

USE `order`;

CREATE TABLE IF NOT EXISTS `inventory_items` (
	id INT AUTO_INCREMENT PRIMARY KEY,
	product_code VARCHAR(64) NOT NULL UNIQUE,
	name VARCHAR(128) NOT NULL,
	description TEXT,
	unit_price FLOAT NOT NULL
);

INSERT INTO `inventory_items` (product_code, name, description, unit_price) VALUES
	('ITEM001', 'Widget A', 'Basic widget', 10.00),
	('ITEM002', 'Widget B', 'Advanced widget', 15.00),
	('ITEM003', 'Gadget X', 'Multi-purpose gadget', 23.00),
	('ITEM004', 'Gadget Y', 'High-end gadget', 45.00),
	('ITEM005', 'Tool Z', 'Essential tool', 5.00)
ON DUPLICATE KEY UPDATE product_code=product_code;