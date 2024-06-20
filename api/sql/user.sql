/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pass` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '密码',
  `email` varchar(128) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'email',
  `phone` varchar(20) COLLATE utf8mb4_bin DEFAULT '' COMMENT '手机号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='平台账号表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `player` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account_id` bigint NOT NULL COMMENT '账户ID',
  `player_id` bigint NOT NULL COMMENT '玩家ID',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gender` int DEFAULT '1' COMMENT '性别',
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;
