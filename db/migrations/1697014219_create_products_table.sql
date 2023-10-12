CREATE TABLE `products` (
  `id` varchar(100) PRIMARY KEY,
  `name` varchar(255) UNIQUE NOT NULL,
  `daily_quota` int,
  `status` varchar(20) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
);
