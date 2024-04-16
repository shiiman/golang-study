CREATE TABLE `users` (
  `id` varchar(255) NOT NULL COMMENT 'ユーザID',
  `username` varchar(255) NOT NULL COMMENT '名前',
  `email` varchar(255) NOT NULL COMMENT 'メール',
  `created_at` datetime NOT NULL COMMENT '作成日時',
  `updated_at` datetime NOT NULL COMMENT '更新日時',
  `deleted_at` datetime DEFAULT NULL COMMENT '削除日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザテーブル';
