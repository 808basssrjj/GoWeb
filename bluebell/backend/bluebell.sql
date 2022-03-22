/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : bluebell

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-03-22 20:12:57
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `comment_id` bigint(20) unsigned NOT NULL,
  `content` text NOT NULL,
  `post_id` bigint(20) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_comment_id` (`comment_id`),
  KEY `idx_author_Id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) unsigned NOT NULL,
  `community_name` varchar(128) NOT NULL,
  `introduction` varchar(256) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of community
-- ----------------------------
INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community` VALUES ('3', '3', 'PUBG', '大吉大利，今晚吃鸡。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community` VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL COMMENT '帖子id',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `content` varchar(8192) NOT NULL COMMENT '内容',
  `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
  `community_id` bigint(20) NOT NULL COMMENT '所属社区',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES ('1', '27621935483457536', 'v1.18', '泛型', '27395531407888384', '4', '1', '2022-03-18 13:19:42', '2022-03-19 16:44:36');
INSERT INTO `post` VALUES ('2', '27640998045683712', '琪亚娜', '元素女皇', '27395531407888384', '4', '1', '2022-03-18 14:35:27', '2022-03-18 14:44:31');
INSERT INTO `post` VALUES ('6', '27642983645974528', 'dasd', 'asdas', '27395531407888384', '3', '1', '2022-03-18 14:43:20', '2022-03-18 14:43:20');
INSERT INTO `post` VALUES ('7', '28013251275001856', '英雄联盟', '哈哈哈哈哈哈哈哈哈哈哈哈哈', '27395531407888384', '4', '1', '2022-03-19 15:14:39', '2022-03-19 15:14:39');
INSERT INTO `post` VALUES ('9', '28369034420424704', 'gin', 'wwwwwwwwwwwwwwwwwwwwwwwwwwwwww', '27247875285061632', '1', '1', '2022-03-20 14:48:24', '2022-03-20 14:48:24');
INSERT INTO `post` VALUES ('11', '28370021805723648', '二叉树', '水水水水水水水水水水水水水水水水水水水', '27247875285061632', '2', '1', '2022-03-20 14:52:20', '2022-03-20 14:52:20');
INSERT INTO `post` VALUES ('12', '28370125161762816', 'gggggggggggggggggg', 'ggggggggggggggggggggggggggggggg', '27247875285061632', '1', '1', '2022-03-20 14:52:44', '2022-03-20 14:52:44');
INSERT INTO `post` VALUES ('15', '28698229876985856', 'golang', 'kllllllllllllllllllllllllllllllll', '27395531407888384', '1', '1', '2022-03-21 12:36:31', '2022-03-21 12:36:31');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `email` varchar(64) DEFAULT NULL,
  `gender` int(4) DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '27247875285061632', '路飞', '3132333435366dcbca412e7463585a81d28878a06747', null, '0', '2022-03-17 12:33:19', '2022-03-17 12:33:19');
INSERT INTO `user` VALUES ('2', '27395531407888384', '索隆', '3132333435366dcbca412e7463585a81d28878a06747', null, '0', '2022-03-17 22:20:03', '2022-03-17 22:20:03');
