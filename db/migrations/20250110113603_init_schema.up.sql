CREATE TABLE users (
   id BIGSERIAL PRIMARY KEY,  -- автоинкрементируемый идентификатор пользователя
   username VARCHAR(255) NOT NULL,  -- имя пользователя
   password VARCHAR(255) NOT NULL,  -- пароль пользователя (хранить в зашифрованном виде!)
   role VARCHAR(50) CHECK(role IN ('User', 'Admin')) NOT NULL,  -- роль пользователя
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- дата создания
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- дата последнего обновления
);

CREATE TABLE products (
  id BIGSERIAL PRIMARY KEY,  -- автоинкрементируемый идентификатор продукта
  name VARCHAR(255) NOT NULL,  -- название продукта
  price DECIMAL(10, 2) NOT NULL,  -- цена продукта
  quantity INT NOT NULL  -- количество продукта
);

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,  -- автоинкрементируемый идентификатор заказа
    customer_name VARCHAR(255) NOT NULL,  -- имя клиента
    status VARCHAR(50) CHECK(status IN ('pending', 'confirmed', 'cancelled')) NOT NULL,  -- статус заказа
    total_price DECIMAL(10, 2) NOT NULL,  -- общая стоимость заказа
    product_id BIGINT REFERENCES products(id),  -- внешний ключ на таблицу продуктов
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- дата создания заказа
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- дата последнего обновления
    is_deleted BOOLEAN DEFAULT FALSE  -- флаг "мягкого" удаления
);


CREATE INDEX idx_users_username ON users(username);

CREATE INDEX idx_users_role ON users(role);

CREATE INDEX idx_users_created_at ON users(created_at);

-- Индекс на поле status
CREATE INDEX idx_orders_status ON orders(status);

-- Индекс на поле total_price для фильтрации заказов по цене
CREATE INDEX idx_orders_total_price ON orders(total_price);

-- Комбинированный индекс на статус и total_price для фильтрации по этим двум полям
CREATE INDEX idx_orders_status_total_price ON orders(status, total_price);

-- Индекс на поле created_at для сортировки заказов по дате
CREATE INDEX idx_orders_created_at ON orders(created_at);





