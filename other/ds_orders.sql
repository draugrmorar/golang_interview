CREATE TABLE categories
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE products
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255)   NOT NULL,
    price       DECIMAL(10, 2) NOT NULL,
    category_id BIGINT REFERENCES categories (id)
);

CREATE TABLE orders
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT REFERENCES users (id),
    status     VARCHAR(50) NOT NULL CHECK (status IN ('new', 'paid', 'cancelled')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT REFERENCES orders (id),
    product_id BIGINT REFERENCES products (id),
    quantity   INT NOT NULL CHECK (quantity > 0)
);

-- Категории
INSERT INTO categories (name)
VALUES ('Электроника'),
       ('Бытовая техника'),
       ('Компьютеры'),
       ('Смартфоны');

-- Пользователи
INSERT INTO users (name)
VALUES ('Alice'),
       ('Bob'),
       ('Charlie');

-- Товары
INSERT INTO products (name, price, category_id)
VALUES ('iPhone 15', 999.99, 4),
       ('Samsung Galaxy S24', 899.99, 4),
       ('MacBook Pro', 2499.99, 3),
       ('Lenovo ThinkPad', 1299.99, 3),
       ('Samsung TV 55"', 799.99, 1),
       ('Sony Headphones', 249.99, 1),
       ('Холодильник LG', 699.99, 2),
       ('Стиральная машина Bosch', 599.99, 2),
       ('iPad Air', 749.99, 3),
       ('Apple Watch', 449.99, 1);

-- Заказы
INSERT INTO orders (user_id, status, created_at)
VALUES (1, 'paid', '2024-01-15 10:00:00'),
       (1, 'paid', '2024-02-20 14:30:00'),
       (1, 'cancelled', '2024-03-01 09:00:00'),
       (2, 'paid', '2024-01-10 11:00:00'),
       (2, 'paid', '2024-02-15 16:45:00'),
       (2, 'new', '2024-03-05 12:00:00'),
       (3, 'paid', '2024-01-25 13:15:00'),
       (3, 'paid', '2024-02-28 10:30:00');

-- Товары в заказах
INSERT INTO order_items (order_id, product_id, quantity)
VALUES
    -- Alice: заказ 1 (Электроника: Sony Headphones x2, Apple Watch x1)
    (1, 6, 2),  -- Sony Headphones
    (1, 10, 1), -- Apple Watch
    -- Alice: заказ 2 (Смартфоны: iPhone 15 x1, Samsung Galaxy x1; Компьютеры: MacBook Pro x1)
    (2, 1, 1),  -- iPhone 15
    (2, 2, 1),  -- Samsung Galaxy S24
    (2, 3, 1),  -- MacBook Pro
    -- Alice: заказ 3 (отменён, не должен учитываться)
    (3, 8, 1),  -- Стиральная машина Bosch

    -- Bob: заказ 4 (Бытовая техника: Холодильник x1, Стиральная машина x1)
    (4, 7, 1),  -- Холодильник LG
    (4, 8, 1),  -- Стиральная машина Bosch
    -- Bob: заказ 5 (Электроника: Samsung TV x1; Компьютеры: Lenovo ThinkPad x1)
    (5, 5, 1),  -- Samsung TV 55"
    (5, 4, 1),  -- Lenovo ThinkPad
    -- Bob: заказ 6 (новый, не должен учитываться)
    (6, 3, 1),  -- MacBook Pro

    -- Charlie: заказ 7 (Смартфоны: iPhone 15 x2; Электроника: Apple Watch x1)
    (7, 1, 2),  -- iPhone 15
    (7, 10, 1), -- Apple Watch
    -- Charlie: заказ 8 (Компьютеры: iPad Air x1, MacBook Pro x1)
    (8, 9, 1),  -- iPad Air
    (8, 3, 1); -- MacBook Pro


WITH user_category_stats AS (SELECT u.name                     AS user_name,
                                    c.name                     AS category_name,
                                    SUM(oi.quantity)           AS total_quantity,
                                    SUM(oi.quantity * p.price) AS total_spent
                             FROM users u
                                      JOIN orders o ON o.user_id = u.id
                                      JOIN order_items oi ON o.id = oi.order_id
                                      JOIN products p ON p.id = oi.product_id
                                      JOIN categories c ON p.category_id = c.id
                             WHERE o.status = 'paid'
                             GROUP BY u.id, u.name, c.id, c.nameb),
     RANKED AS (SELECT user_name,
                       category_name,
                       total_quantity,
                       total_spent,
                       RANK() OVER (PARTITION BY user_name ORDER BY total_quantity DESC) AS rnk
                FROM user_category_stats)
SELECT user_name,
       category_name,
       total_quantity,
       total_spent
FROM RANKED
WHERE rnk = 1
ORDER BY user_name, category_name;


WITH user_category_stats AS (SELECT c.name                     AS category_name,
                                    p.name                     AS product_name,
                                    SUM(oi.quantity)           AS total_quantity,
                                    SUM(oi.quantity * p.price) AS total_revenue
                             FROM users u
                                      JOIN orders o ON o.user_id = u.id
                                      JOIN order_items oi ON o.id = oi.order_id
                                      JOIN products p ON p.id = oi.product_id
                                      JOIN categories c ON p.category_id = c.id
                             WHERE o.status = 'paid'
                             GROUP BY c.id, c.name, p.id, p.name),
     RANKED AS (SELECT product_name,
                       category_name,
                       total_quantity,
                       total_revenue,
                       RANK() OVER (PARTITION BY category_name ORDER BY total_revenue DESC) AS rnk
                FROM user_category_stats)

SELECT product_name,
       category_name,
       total_quantity,
       total_revenue,
       rnk
FROM RANKED
WHERE rnk <= 3
ORDER BY category_name, product_name;


WITH order_totals AS (
    -- Шаг 1: считаем сумму каждого заказа
    SELECT o.user_id,
           o.id                       AS order_id,
           SUM(oi.quantity * p.price) AS order_amount
    FROM orders o
             JOIN order_items oi ON o.id = oi.order_id
             JOIN products p ON p.id = oi.product_id
    WHERE o.status = 'paid'
    GROUP BY o.user_id, o.id),
     user_stats AS (
         -- Шаг 2: считаем средний чек и количество заказов по пользователю
         SELECT ot.user_id,
                AVG(ot.order_amount) AS avg_order_amount,
                COUNT(ot.order_id)   AS total_orders
         FROM order_totals ot
         GROUP BY ot.user_id)
SELECT u.name AS user_name,
       us.avg_order_amount,
       us.total_orders,
       RANK()    OVER (ORDER BY us.avg_order_amount DESC) AS rnk
FROM user_stats us
         JOIN users u ON u.id = us.user_id
ORDER BY rnk;