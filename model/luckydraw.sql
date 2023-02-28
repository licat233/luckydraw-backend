/*
 Navicat MySQL Data Transfer

 Source Server         : localmysql
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : luckydraw

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 28/02/2023 15:49:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(254) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of activity
-- ----------------------------
BEGIN;
INSERT INTO `activity` VALUES (1, '1376ba46-c178-4baa-aab7-be6420b37bc1', '李中医3月份活动', 1, '2023-02-27 23:16:47', '2023-02-28 00:24:29');
COMMIT;

-- ----------------------------
-- Table structure for adminer
-- ----------------------------
DROP TABLE IF EXISTS `adminer`;
CREATE TABLE `adminer` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `access` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'normal',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of adminer
-- ----------------------------
BEGIN;
INSERT INTO `adminer` VALUES (1, 'licat', '131313', 'superAdmin', 1, '2023-02-22 17:02:25', '2023-02-28 15:23:06');
INSERT INTO `adminer` VALUES (2, 'admin', '12345', 'admin', 0, '2023-02-27 18:20:39', '2023-02-27 22:13:41');
INSERT INTO `adminer` VALUES (4, 'normaladmin', '123456', 'normal', 0, '2023-02-27 21:23:17', '2023-02-27 22:13:49');
INSERT INTO `adminer` VALUES (5, 'user', '123456', 'normal', 0, '2023-02-27 21:34:13', '2023-02-27 21:34:13');
COMMIT;

-- ----------------------------
-- Table structure for awards
-- ----------------------------
DROP TABLE IF EXISTS `awards`;
CREATE TABLE `awards` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `activity_id` int unsigned NOT NULL,
  `uuid` varchar(254) NOT NULL DEFAULT '',
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `image` varchar(255) NOT NULL DEFAULT '',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '奖品价格',
  `prob` int unsigned NOT NULL DEFAULT '0',
  `quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '总数量',
  `count` int unsigned NOT NULL DEFAULT '0' COMMENT '已抽数量',
  `is_win` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否中奖',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `activity_id` (`activity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of awards
-- ----------------------------
BEGIN;
INSERT INTO `awards` VALUES (1, 1, '1cf09fc7ca1247e9b0ffa163407f737f', '一等獎', '燕窩蛋白凍+多效修護凍乾粉+積雪草安肌修复套装', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01xsbomP2BlB3KbnRhh_!!917298378.png', 4600.00, 100, 999, 0, 1, '2023-02-23 10:23:04', '2023-02-28 00:24:37');
INSERT INTO `awards` VALUES (2, 1, '99d8e56b42c44f80837e6fb352980963', '二等獎', '積雪草安肌修复套装+燕窩蛋白凍*2', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01nhKLND2BlB3EZ1Ue5_!!917298378.png', 3800.00, 100, 999, 0, 1, '2023-02-23 10:24:26', '2023-02-28 00:24:39');
INSERT INTO `awards` VALUES (3, 1, '662795660f304a41b9ddaf935133a119', '三等獎', '積雪草安肌修复套装+焕颜紧致活力眼霜', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01KbpHlp2BlB3Ut9BlI_!!917298378.png', 3200.00, 100, 999, 0, 1, '2023-02-23 10:24:57', '2023-02-28 00:24:41');
INSERT INTO `awards` VALUES (4, 1, '6bf252836a6943259e2032b8848f1280', '四等獎', '多效修護凍乾粉+潤澤水感面膜', 'https://img.alicdn.com/imgextra/i1/917298378/O1CN01SHo4sF2BlB3NSJMks_!!917298378.png', 2100.00, 100, 999, 0, 1, '2023-02-23 10:25:19', '2023-02-28 00:24:42');
INSERT INTO `awards` VALUES (5, 1, 'ce08ae497073436c932765656ef932eb', '五等獎', '莹润亮颜隔离霜', 'https://img.alicdn.com/imgextra/i2/917298378/O1CN013hWPpE2BlB3Qcvuik_!!917298378.png', 900.00, 100, 999, 0, 1, '2023-02-23 10:25:45', '2023-02-28 00:24:44');
INSERT INTO `awards` VALUES (6, 1, 'c4bc1f80c08646e89aba4fd4e3e1f13f', '特等獎', '神秘大禮', 'https://img.alicdn.com/imgextra/i4/917298378/O1CN01HWdHD22BlB3EZ0YRK_!!917298378.png', 9999.00, 100, 999, 0, 1, '2023-02-23 10:48:26', '2023-02-28 00:24:46');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `activity_id` int unsigned NOT NULL,
  `available_awards` varchar(255) NOT NULL COMMENT '指定其可抽中的奖品',
  `name` varchar(255) NOT NULL DEFAULT '',
  `passport` varchar(255) NOT NULL DEFAULT 'none',
  `count` int NOT NULL DEFAULT '0' COMMENT '抽獎次數',
  `total` int NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `activity_id` (`activity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 1, '1', '客户A', 'anby03', 0, 99, '2023-02-23 11:00:19', '2023-02-28 00:24:59');
INSERT INTO `users` VALUES (2, 1, '7', '客戶B', 'anby001', 0, 0, '2023-02-23 16:27:20', '2023-02-28 00:25:01');
INSERT INTO `users` VALUES (3, 1, '8', '客戶C', 'anby002', 0, 2, '2023-02-23 16:27:20', '2023-02-28 00:25:02');
COMMIT;

-- ----------------------------
-- Table structure for winning_records
-- ----------------------------
DROP TABLE IF EXISTS `winning_records`;
CREATE TABLE `winning_records` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `award_id` int unsigned NOT NULL,
  `activity_id` int unsigned NOT NULL,
  `ip` varchar(255) NOT NULL DEFAULT '0.0.0.0',
  `platform` varchar(255) NOT NULL DEFAULT 'none',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `award_id` (`award_id`),
  KEY `activity_id` (`activity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
