/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50540
Source Host           : localhost:3306
Source Database       : goblog

Target Server Type    : MYSQL
Target Server Version : 50540
File Encoding         : 65001

Date: 2015-01-26 16:44:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog
-- ----------------------------
DROP TABLE IF EXISTS `blog`;
CREATE TABLE `blog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `context` text,
  `updatetime` datetime DEFAULT NULL,
  `createtime` datetime DEFAULT NULL,
  `viewnum` int(11) DEFAULT NULL,
  `intro` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=251 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of blog
-- ----------------------------
INSERT INTO `blog` VALUES ('247', 'API接口', '# 用户接口\r\n\r\n## 接口地址 /user\r\n\r\n### get方法\r\n\r\n*获取用户信息*\r\n\r\n### post方法\r\n\r\n*用户注册*\r\n\r\n### put方法\r\n\r\n*更新信息*\r\n\r\n### delete方法\r\n\r\n*删除信息*\r\n\r\n# 会话接口\r\n\r\n## 接口地址 /api/session\r\n\r\n### post方法\r\n\r\n*登录*\r\n\r\n### delete 方法\r\n\r\n*删除会话(退出)*\r\n\r\n# 博客接口\r\n\r\n## 接口地址 /api/blog/\r\n\r\n### post方法\r\n\r\n*添加*\r\n\r\n### delete 方法\r\n\r\n*删除*\r\n\r\n### put 方法\r\n\r\n修改\r\n\r\n### get方法\r\n\r\n查找\r\n\r\n/api/blog/ 返回列表;\r\n\r\n/api/blog/:id 返回内容.\r\n\r\n\r\n', '2015-01-19 17:10:10', '2015-01-19 17:10:10', '14', ' ');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fid` int(11) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `uid` int(11) DEFAULT NULL,
  `context` text,
  `createtime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `uid` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL COMMENT '0=女;1=男',
  `password` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `avatar` varchar(50) DEFAULT NULL,
  `qq` varchar(20) DEFAULT NULL,
  `createtime` datetime DEFAULT NULL,
  `updatetime` datetime DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('12', 'demo', '1', '1313c725d1af896e8464cbbabadb1084', '444186377@qq.com', '', 'http://goblog.b0.upaiyun.com/avatar/12.png!avatar', '', '2015-01-24 00:43:00', '0000-00-00 00:00:00');
