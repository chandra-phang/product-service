CREATE TABLE `daily_product_quotas` (
  `id` varchar(100) PRIMARY KEY,
  `product_id` varchar(100) NOT NULL,
  `daily_quota` int,
  `booked_quota` int,
  `date` date NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`product_id`) REFERENCES products(`id`),
  INDEX (`product_id`, `date`)
);
