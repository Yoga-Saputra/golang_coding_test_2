CREATE TABLE `members` (
  `id_member` bigint PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255),
  `gender` enum('P','L') DEFAULT 'L',
  `skintype` varchar(255),
  `skincolor` varchar(255)
);

CREATE TABLE `products` (
  `id_product` bigint PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `product` varchar(255),
  `price` decimal
);

CREATE TABLE `review_products` (
  `id_review` bigint PRIMARY KEY AUTO_INCREMENT,
  `id_product` bigint,
  `id_member` bigint,
  `desc_review` text
);

CREATE TABLE `like_reviews` (
  `id_review` bigint,
  `id_member` bigint
);

CREATE INDEX `products_index_0` ON `products` (`name`);

CREATE INDEX `products_index_1` ON `products` (`name`);

CREATE INDEX `review_products_index_2` ON `review_products` (`id_product`);

CREATE INDEX `review_products_index_3` ON `review_products` (`id_member`);

CREATE INDEX `review_products_index_4` ON `review_products` (`id_product`, `id_member`);

CREATE INDEX `like_reviews_index_5` ON `like_reviews` (`id_review`);

CREATE INDEX `like_reviews_index_6` ON `like_reviews` (`id_member`);

CREATE INDEX `like_reviews_index_7` ON `like_reviews` (`id_review`, `id_member`);

ALTER TABLE `review_products` ADD FOREIGN KEY (`id_product`) REFERENCES `products` (`id_product`);

ALTER TABLE `review_products` ADD FOREIGN KEY (`id_member`) REFERENCES `members` (`id_member`);

ALTER TABLE `like_reviews` ADD FOREIGN KEY (`id_review`) REFERENCES `review_products` (`id_review`);

ALTER TABLE `like_reviews` ADD FOREIGN KEY (`id_member`) REFERENCES `members` (`id_member`);
