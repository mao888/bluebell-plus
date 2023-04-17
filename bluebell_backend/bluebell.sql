/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : bluebell

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 03/03/2022 10:27:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) UNSIGNED NOT NULL,
  `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_community_id`(`community_id`) USING BTREE,
  UNIQUE INDEX `idx_community_name`(`community_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of community
-- ----------------------------
INSERT INTO `community` VALUES (1, 1, 'Go', 'Golang', '2022-02-12 17:25:26', '2022-02-12 17:25:28');
INSERT INTO `community` VALUES (2, 2, 'leetcode', '刷题刷题刷题', '2022-02-12 17:25:38', '2022-02-12 17:25:40');
INSERT INTO `community` VALUES (3, 3, 'Java', 'springboot', '2022-02-12 17:25:46', '2022-02-12 17:26:20');
INSERT INTO `community` VALUES (4, 4, 'LOL', '欢迎来到英雄联盟!', '2022-02-12 17:25:53', '2022-02-12 17:25:55');

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL COMMENT '帖子id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
  `community_id` bigint(20) NOT NULL COMMENT '所属社区',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_post_id`(`post_id`) USING BTREE,
  INDEX `idx_author_id`(`author_id`) USING BTREE,
  INDEX `idx_community_id`(`community_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES (1, 508837878562817, '学习', '12321131231xuexi', 95519300911105, 1, 1, '2022-02-12 20:14:51', '2022-02-12 20:14:51');
INSERT INTO `post` VALUES (2, 526606024048641, '超哥', '超好超超超', 95519300911105, 1, 1, '2022-02-12 23:11:21', '2022-02-12 23:11:21');
INSERT INTO `post` VALUES (3, 780736906919937, '投票功能', '投票功能真的很棒呢', 95519300911105, 2, 1, '2022-02-14 17:15:55', '2022-02-14 17:15:55');
INSERT INTO `post` VALUES (4, 803331253469185, '传感帖子', '贴子内容', 95519300911105, 2, 1, '2022-02-14 21:00:22', '2022-02-14 21:00:22');
INSERT INTO `post` VALUES (5, 4, 'test', 'test123', 1, 1, 1, '2022-02-19 17:03:49', '2022-02-19 17:03:49');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `gender` tinyint(4) NOT NULL DEFAULT 0,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 95519300911105, '胡超', '313233f5d77a10ae47e3738837865e6a831793', NULL, 0, '2022-02-09 23:48:53', '2022-02-09 23:48:53');
INSERT INTO `user` VALUES (2, 166698938269697, '胡超超', '313233f5d77a10ae47e3738837865e6a831793', NULL, 0, '2022-02-10 11:36:00', '2022-02-10 11:36:00');

SET FOREIGN_KEY_CHECKS = 1;
