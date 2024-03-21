CREATE TABLE Products 
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE Shelves 
(
    id SERIAL PRIMARY KEY,
    location VARCHAR(255) NOT NULL
);

CREATE TABLE Orders 
(
    id SERIAL PRIMARY KEY,
    order_date DATE NOT NULL
);

CREATE TABLE ProductShelves 
(
    product_id INTEGER,
    shelf_id INTEGER,
    FOREIGN KEY (product_id) REFERENCES Products(id),
    FOREIGN KEY (shelf_id) REFERENCES Shelves(id),
    is_main BOOLEAN NOT NULL,
    PRIMARY KEY (product_id, shelf_id)
);

CREATE TABLE OrderItems 
(
    order_id INTEGER,
    product_id INTEGER,
    FOREIGN KEY (order_id) REFERENCES Orders(id),
    FOREIGN KEY (product_id) REFERENCES Products(id),
    count INTEGER,
    PRIMARY KEY (order_id, product_id)
);

INSERT INTO Products (name) VALUES
    ('Ноутбук'),
    ('Телевизор'),
    ('Телефон'),
    ('Системный блок'),
    ('Часы'),
    ('Микрофон');

INSERT INTO Shelves (location) VALUES
    ('Стеллаж А'),
    ('Стеллаж Б'),
    ('Стеллаж В'),
    ('Стеллаж Ж'),
    ('Стеллаж З');

INSERT INTO Orders (id ,order_date) VALUES
    (10,'2024-03-20'),
    (11,'2024-03-20'),
    (14,'2024-03-20'),
    (15,'2024-03-20');

-- Продукт Ноутбук
INSERT INTO ProductShelves (product_id, shelf_id, is_main) VALUES
    (1, 1, TRUE); -- Стеллаж А

-- Продукт Телевизор
INSERT INTO ProductShelves (product_id, shelf_id, is_main) VALUES
    (2, 1, TRUE); -- Стеллаж А

-- Продукт Телефон
INSERT INTO ProductShelves (product_id, shelf_id, is_main) VALUES
    (3, 2, TRUE), -- Стеллаж Б
    (3, 3, FALSE), -- Доп. Стеллаж В
    (3, 5, FALSE); -- Доп. Стеллаж Ж

-- Продукт Системный блок
INSERT INTO ProductShelves (product_id, shelf_id, is_main) VALUES
    (4, 4, TRUE); -- Стеллаж Ж

-- Продукт Часы
INSERT INTO ProductShelves (product_id, shelf_id, is_main) VALUES
    (5, 4, TRUE), -- Стеллаж Ж
    (5, 1, FALSE); -- Доп. Стеллаж А

-- Продукт Микрофон
INSERT INTO ProductShelves (product_id, shelf_id, is_main) VALUES
    (6, 4, TRUE); -- Стеллаж Ж

INSERT INTO OrderItems (order_id, product_id, count)
VALUES
    (10, 1, 2),  -- Заказ 10, 2 шт. Ноутбук
    (10, 3, 1),  -- Заказ 10, 1 шт. Телефон
    (11, 2, 3),  -- Заказ 11, 3 шт. Телевизор
    (14, 1, 3),  -- Заказ 14, 3 шт. Ноутбук
    (14, 4, 4),  -- Заказ 14, 4 шт. Системный блок
    (15, 5, 1),  -- Заказ 15, 1 шт. Часы
    (10, 6, 1);  -- Заказ 10, 1 шт. Микрофон