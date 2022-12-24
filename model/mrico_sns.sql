-- MySQL dump 10.13  Distrib 5.7.34, for osx10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: mrico_sns
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sns_account`
--

DROP TABLE IF EXISTS `sns_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_account` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账户名',
  `nicks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `pass` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：0-冻结/1-正常',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '软删除：0未删；1删除',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE COMMENT '账户名唯一',
  KEY `nick` (`nicks`) USING BTREE,
  KEY `is_delete` (`is_delete`) USING BTREE,
  KEY `state` (`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员账户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_account`
--

LOCK TABLES `sns_account` WRITE;
/*!40000 ALTER TABLE `sns_account` DISABLE KEYS */;
/*!40000 ALTER TABLE `sns_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_cate`
--

DROP TABLE IF EXISTS `sns_cate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_cate` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `alias` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '别名',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '介绍',
  `pid` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '上级节点',
  `level` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '层级：0一级(默认)',
  `icons` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `sorts` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：0-关闭/1-开启',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '软删除',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE COMMENT '名称唯一',
  KEY `state` (`state`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `sorts` (`sorts`) USING BTREE,
  KEY `is_delete` (`is_delete`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_cate`
--

LOCK TABLES `sns_cate` WRITE;
/*!40000 ALTER TABLE `sns_cate` DISABLE KEYS */;
INSERT INTO `sns_cate` VALUES (1,'测试','test','测试',0,0,'',0,1,0,'2022-12-21 22:16:03','2022-12-21 22:16:03',''),(2,'GO','go','GO版块',0,0,'',0,1,0,'2022-12-21 22:16:03','2022-12-21 22:16:03',''),(3,'PHP','php','PHP版块',0,0,'',0,1,0,'2022-12-21 22:16:03','2022-12-21 22:16:03',''),(4,'问答','answers','问答版块',0,0,'',0,1,0,'2022-12-21 22:16:04','2022-12-21 22:16:04','');
/*!40000 ALTER TABLE `sns_cate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_checkins`
--

DROP TABLE IF EXISTS `sns_checkins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_checkins` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `cumulative_days` bigint unsigned NOT NULL DEFAULT '0' COMMENT '累积签到(天)',
  `continuity_days` bigint unsigned NOT NULL DEFAULT '0' COMMENT '连续签到(天)',
  `last_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后签到时间',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户签到';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_checkins`
--

LOCK TABLES `sns_checkins` WRITE;
/*!40000 ALTER TABLE `sns_checkins` DISABLE KEYS */;
/*!40000 ALTER TABLE `sns_checkins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_follows`
--

DROP TABLE IF EXISTS `sns_follows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_follows` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `target_id` int unsigned NOT NULL COMMENT '被关注用户ID',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-关注/0-取消',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `target_id` (`target_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户关注';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_follows`
--

LOCK TABLES `sns_follows` WRITE;
/*!40000 ALTER TABLE `sns_follows` DISABLE KEYS */;
/*!40000 ALTER TABLE `sns_follows` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_integral_logs`
--

DROP TABLE IF EXISTS `sns_integral_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_integral_logs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `rewards` bigint unsigned NOT NULL DEFAULT '0' COMMENT '奖励积分',
  `mode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '获取方式',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='积分日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_integral_logs`
--

LOCK TABLES `sns_integral_logs` WRITE;
/*!40000 ALTER TABLE `sns_integral_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sns_integral_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_topics`
--

DROP TABLE IF EXISTS `sns_topics`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_topics` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `tags` json NOT NULL COMMENT '标签',
  `state` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态：0-草稿/1-发布',
  `type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '类型：0-默认/1-精华/2-置顶',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `md_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'markdown内容',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '软删除',
  `last_reply_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后回复时间',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `title` (`title`) USING BTREE,
  KEY `is_delete` (`is_delete`) USING BTREE,
  KEY `type` (`type`) USING BTREE,
  KEY `state` (`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='话题';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_topics`
--

LOCK TABLES `sns_topics` WRITE;
/*!40000 ALTER TABLE `sns_topics` DISABLE KEYS */;
/*!40000 ALTER TABLE `sns_topics` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_topics_users`
--

DROP TABLE IF EXISTS `sns_topics_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_topics_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `topics_id` int unsigned NOT NULL COMMENT '话题ID',
  `cate_id` int unsigned NOT NULL COMMENT '分类ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `reply_id` int unsigned NOT NULL COMMENT '最后回复者ID',
  `view_count` bigint unsigned NOT NULL COMMENT '浏览统计',
  `like_count` bigint unsigned NOT NULL COMMENT '喜欢统计',
  `comment_count` bigint unsigned NOT NULL COMMENT '评论统计',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `topics_id` (`topics_id`) USING BTREE,
  KEY `cate_id` (`cate_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `reply_id` (`reply_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='话题-用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_topics_users`
--

LOCK TABLES `sns_topics_users` WRITE;
/*!40000 ALTER TABLE `sns_topics_users` DISABLE KEYS */;
/*!40000 ALTER TABLE `sns_topics_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sns_users`
--

DROP TABLE IF EXISTS `sns_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sns_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `pass` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别：0-未知/1-男/2-女/3-保密',
  `city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '城市',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱/用户名',
  `phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机/用户名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `site` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '网站',
  `job` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '职业',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简介',
  `integral` bigint unsigned NOT NULL DEFAULT '0' COMMENT '个人积分',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正常/0-禁用',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '软删除',
  `level` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '等级：1-20',
  `created_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建时间',
  `updated_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间',
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `phone` (`phone`) USING BTREE,
  KEY `email` (`email`) USING BTREE,
  KEY `level` (`level`) USING BTREE,
  KEY `is_delete` (`is_delete`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='网站用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sns_users`
--

LOCK TABLES `sns_users` WRITE;
/*!40000 ALTER TABLE `sns_users` DISABLE KEYS */;
INSERT INTO `sns_users` VALUES (1,'test','123456',0,'','','','','','','',0,0,0,0,'2022-12-21 11:29:48','2022-12-21 11:29:48',''),(2,'test1','123456',0,'','','','','','','',0,0,0,0,'2022-12-21 11:42:19','2022-12-21 11:42:19',''),(8,'无言勻','8fc3a58446276025',0,'','','','','','','',0,0,0,0,'2022-12-21 15:25:24','2022-12-21 15:25:24',''),(9,'无言勻1','8fc3a58446276025',0,'','','','','','','',0,1,0,1,'2022-12-23 22:34:04','2022-12-23 22:34:04',''),(10,'无言勻2','8fc3a58446276025',0,'','','','','','','',0,1,0,1,'2022-12-23 22:43:50','2022-12-23 22:43:50',''),(11,'无言勻3','8fc3a58446276025',0,'','','','','','','',0,1,0,1,'2022-12-23 22:44:31','2022-12-23 22:44:31','');
/*!40000 ALTER TABLE `sns_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-24 18:26:18
