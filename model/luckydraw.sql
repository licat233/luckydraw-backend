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

 Date: 27/02/2023 21:34:37
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of activity
-- ----------------------------
BEGIN;
INSERT INTO `activity` VALUES (1, '529de531256e41da8780a0fbf0f1d873', '植題3月活動', 0, '2023-02-23 10:19:31', '2023-02-23 10:19:55');
INSERT INTO `activity` VALUES (2, '36f10597f41048fbace8a1b1537ae6dc', '李中醫3月活動', 1, '2023-02-23 10:21:48', '2023-02-27 21:23:51');
COMMIT;

-- ----------------------------
-- Table structure for adminer
-- ----------------------------
DROP TABLE IF EXISTS `adminer`;
CREATE TABLE `adminer` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `access` varchar(255)  NOT NULL DEFAULT 'normal',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of adminer
-- ----------------------------
BEGIN;
INSERT INTO `adminer` VALUES (1, 'superadmin', '123456', '', 1, '2023-02-22 17:02:25', '2023-02-27 18:31:43');
INSERT INTO `adminer` VALUES (2, 'admin', '12345', '', 0, '2023-02-27 18:20:39', '2023-02-27 18:20:39');
INSERT INTO `adminer` VALUES (4, 'normaladmin', '123456', '', 0, '2023-02-27 21:23:17', '2023-02-27 21:23:17');
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
INSERT INTO `awards` VALUES (1, 2, '1cf09fc7ca1247e9b0ffa163407f737f', '一等獎', '燕窩蛋白凍+多效修護凍乾粉+積雪草安肌修复套装', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01xsbomP2BlB3KbnRhh_!!917298378.png', 4600.00, 100, 999, 0, 1, '2023-02-23 10:23:04', '2023-02-27 18:33:12');
INSERT INTO `awards` VALUES (2, 2, '99d8e56b42c44f80837e6fb352980963', '二等獎', '積雪草安肌修复套装+燕窩蛋白凍*2', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01nhKLND2BlB3EZ1Ue5_!!917298378.png', 3800.00, 100, 999, 0, 1, '2023-02-23 10:24:26', '2023-02-23 14:57:27');
INSERT INTO `awards` VALUES (3, 2, '662795660f304a41b9ddaf935133a119', '三等獎', '積雪草安肌修复套装+焕颜紧致活力眼霜', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01KbpHlp2BlB3Ut9BlI_!!917298378.png', 3200.00, 100, 999, 0, 1, '2023-02-23 10:24:57', '2023-02-23 14:57:30');
INSERT INTO `awards` VALUES (4, 2, '6bf252836a6943259e2032b8848f1280', '四等獎', '多效修護凍乾粉+潤澤水感面膜', 'https://img.alicdn.com/imgextra/i1/917298378/O1CN01SHo4sF2BlB3NSJMks_!!917298378.png', 2100.00, 100, 999, 0, 1, '2023-02-23 10:25:19', '2023-02-23 14:57:33');
INSERT INTO `awards` VALUES (5, 2, 'ce08ae497073436c932765656ef932eb', '五等獎', '莹润亮颜隔离霜', 'https://img.alicdn.com/imgextra/i2/917298378/O1CN013hWPpE2BlB3Qcvuik_!!917298378.png', 900.00, 100, 999, 0, 1, '2023-02-23 10:25:45', '2023-02-23 14:57:36');
INSERT INTO `awards` VALUES (6, 2, 'c4bc1f80c08646e89aba4fd4e3e1f13f', '特等獎', '神秘大禮', 'https://img.alicdn.com/imgextra/i4/917298378/O1CN01HWdHD22BlB3EZ0YRK_!!917298378.png', 9999.00, 100, 999, 0, 1, '2023-02-23 10:48:26', '2023-02-23 15:04:55');
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
INSERT INTO `users` VALUES (1, 2, '1', '客户A', 'anby03', 0, 99, '2023-02-23 11:00:19', '2023-02-24 19:06:39');
INSERT INTO `users` VALUES (2, 2, '7', '客戶B', 'anby001', 0, 0, '2023-02-23 16:27:20', '2023-02-23 17:41:38');
INSERT INTO `users` VALUES (3, 2, '8', '客戶C', 'anby002', 0, 2, '2023-02-23 16:27:20', '2023-02-24 21:12:41');
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
