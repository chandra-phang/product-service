CREATE TABLE products (
  `id` varchar(100) PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `daily_quota` int,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL
);
