CREATE TABLE `Domains` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '域ID',
  `domain` varchar(32) NOT NULL COMMENT '域名称',
  `status` int NOT NULL COMMENT '状态',
  `created` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='域表'