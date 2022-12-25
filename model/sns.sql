CREATE TABLE `sns_account`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户名',
  `nicks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `pass` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `state` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：0-冻结/1-正常',
  `is_delete` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '软删除：0未删；1删除',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name`(`name`) USING BTREE COMMENT '账户名唯一',
  INDEX `nick`(`nicks`) USING BTREE,
  INDEX `is_delete`(`is_delete`) USING BTREE,
  INDEX `state`(`state`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员账户';

CREATE TABLE `sns_cate`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `alias` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '别名',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '介绍',
  `pid` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级节点',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级：0一级(默认)',
  `icons` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `sorts` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `state` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：0-关闭/1-开启',
  `is_delete` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '软删除',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name`(`name`) USING BTREE COMMENT '名称唯一',
  INDEX `state`(`state`) USING BTREE,
  INDEX `created_at`(`created_at`) USING BTREE,
  INDEX `sorts`(`sorts`) USING BTREE,
  INDEX `is_delete`(`is_delete`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分类表';

CREATE TABLE `sns_checkins`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
  `cumulative_days` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '累积签到(天)',
  `continuity_days` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '连续签到(天)',
  `last_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后签到时间',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户签到';

CREATE TABLE `sns_follows`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
  `target_id` int UNSIGNED NOT NULL COMMENT '被关注用户ID',
  `state` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：1-关注/0-取消',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `target_id`(`target_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户关注';

CREATE TABLE `sns_integral_logs`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
  `rewards` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '奖励积分',
  `mode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '获取方式',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '积分日志';

CREATE TABLE `sns_topics`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
  `state` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0-草稿/1-发布',
  `type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0-默认/1-精华/2-置顶',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `md_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'markdown内容',
  `is_delete` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '软删除',
  `last_reply_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后回复时间',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  INDEX `title`(`title`) USING BTREE,
  INDEX `is_delete`(`is_delete`) USING BTREE,
  INDEX `type`(`type`) USING BTREE,
  INDEX `state`(`state`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '话题';

CREATE TABLE `sns_topics_users`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `topics_id` int UNSIGNED NOT NULL COMMENT '话题ID',
  `cate_id` int UNSIGNED NOT NULL COMMENT '分类ID',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
  `reply_id` int UNSIGNED NOT NULL COMMENT '最后回复者ID',
  `view_count` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览统计',
  `like_count` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '喜欢统计',
  `comment_count` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论统计',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `topics_id`(`topics_id`) USING BTREE,
  INDEX `cate_id`(`cate_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `reply_id`(`reply_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '话题-用户';

CREATE TABLE `sns_users`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `pass` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别：0-未知/1-男/2-女/3-保密',
  `city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '城市',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱/用户名',
  `phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机/用户名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `site` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网站',
  `job` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '职业',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简介',
  `integral` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '个人积分',
  `state` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：1-正常/0-禁用',
  `is_delete` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '软删除',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '等级：1-20',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name`(`name`) USING HASH,
  INDEX `phone`(`phone`) USING BTREE,
  INDEX `email`(`email`) USING BTREE,
  INDEX `level`(`level`) USING BTREE,
  INDEX `is_delete`(`is_delete`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '网站用户';

