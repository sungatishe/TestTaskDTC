-- Удаляем индексы для пользователей
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_created_at;

-- Удаляем индексы для заказов
DROP INDEX IF EXISTS idx_orders_status;
DROP INDEX IF EXISTS idx_orders_total_price;
DROP INDEX IF EXISTS idx_orders_status_total_price;
DROP INDEX IF EXISTS idx_orders_created_at;

-- Удаляем таблицу заказов
DROP TABLE IF EXISTS orders;

-- Удаляем таблицу продуктов
DROP TABLE IF EXISTS products;

-- Удаляем таблицу пользователей
DROP TABLE IF EXISTS users;
