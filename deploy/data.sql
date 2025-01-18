SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for channel
-- ----------------------------
DROP TABLE IF EXISTS `channel`;
CREATE TABLE `channel` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `type` char(32) NOT NULL COMMENT '类型',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `ak` varchar(32) DEFAULT NULL COMMENT 'ak',
  `sk` varchar(32) DEFAULT NULL COMMENT 'sk',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `extra` tinytext COMMENT '扩展信息',
  `created_at` bigint unsigned DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `created_at` (`created_at`),
  KEY `updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='通知渠道信息';

-- ----------------------------
-- Records of channel
-- ----------------------------
BEGIN;
INSERT INTO `channel` (`id`, `type`, `name`, `ak`, `sk`, `status`, `extra`, `created_at`, `updated_at`) VALUES (1, 'email', '邮箱-860', '860808187@qq.com', 'fyudafxxxxxxxxhwbfbd', 1, '{\"host\":\"smtp.qq.com\",\"port\":25,\"name\":\"青城\"}', 1735892273, 1736241143);
INSERT INTO `channel` (`id`, `type`, `name`, `ak`, `sk`, `status`, `extra`, `created_at`, `updated_at`) VALUES (2, 'official_account', '微信公众号-860', 'wx7f869e909f30dc05', 'f9373b4c6xxxxxx2842fdcd82a', 1, '{}', 1736744452, 1736757542);
COMMIT;

-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `init` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gorm_init
-- ----------------------------
BEGIN;
INSERT INTO `gorm_init` (`id`, `init`) VALUES (1, 0);
COMMIT;

-- ----------------------------
-- Table structure for notify
-- ----------------------------
DROP TABLE IF EXISTS `notify`;
CREATE TABLE `notify` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `category_id` bigint unsigned NOT NULL COMMENT '分类ID',
  `keyword` char(32) NOT NULL COMMENT '标识',
  `name` varchar(128) NOT NULL COMMENT '名称',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `variable` varchar(512) NOT NULL COMMENT '变量',
  `send_mode` char(32) NOT NULL COMMENT '发送模式',
  `priority` tinyint NOT NULL COMMENT '发送优先级',
  `is_timely` tinyint(1) NOT NULL COMMENT '是否具有失效性',
  `expire` bigint unsigned DEFAULT NULL COMMENT '过期时间',
  `cache` bigint unsigned DEFAULT NULL COMMENT '缓存时间',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `description` varchar(512) NOT NULL COMMENT '描述',
  `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
  `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  KEY `priority` (`priority`,`created_at`),
  KEY `category_id` (`category_id`),
  KEY `deleted_at` (`deleted_at`),
  CONSTRAINT `notify_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `notify_category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='通知信息';

-- ----------------------------
-- Records of notify
-- ----------------------------
BEGIN;
INSERT INTO `notify` (`id`, `category_id`, `keyword`, `name`, `title`, `variable`, `send_mode`, `priority`, `is_timely`, `expire`, `cache`, `status`, `description`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 1, 'test', '测试用纸', '这是一个测试通知', 'captcha,minute', 'all', 0, 1, 10, 100, 1, '这是一个测试通知', 1736249259, 1736499848, 0);
COMMIT;

-- ----------------------------
-- Table structure for notify_category
-- ----------------------------
DROP TABLE IF EXISTS `notify_category`;
CREATE TABLE `notify_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `description` varchar(256) NOT NULL COMMENT '描述',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='分组信息';

-- ----------------------------
-- Records of notify_category
-- ----------------------------
BEGIN;
INSERT INTO `notify_category` (`id`, `name`, `description`, `created_at`) VALUES (1, '测试分组', '测试分组', 1735986622);
COMMIT;

-- ----------------------------
-- Table structure for send_log
-- ----------------------------
DROP TABLE IF EXISTS `send_log`;
CREATE TABLE `send_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `notify_id` bigint unsigned NOT NULL COMMENT '通知id',
  `channel_id` bigint unsigned NOT NULL COMMENT '渠道id',
  `user` varchar(128) NOT NULL COMMENT '通知用户',
  `variable` mediumtext NOT NULL COMMENT '通知变量',
  `content` mediumtext NOT NULL COMMENT '通知内容',
  `from_server` varchar(128) NOT NULL COMMENT '通知来源',
  `from_ip` varchar(128) NOT NULL COMMENT '来源ip',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `err_code` int NOT NULL DEFAULT '0' COMMENT '错误码',
  `reason` text COMMENT '错误原因',
  `cost` int NOT NULL COMMENT '耗时',
  `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `notify_id` (`notify_id`),
  KEY `channel_id` (`channel_id`),
  CONSTRAINT `send_log_ibfk_1` FOREIGN KEY (`notify_id`) REFERENCES `notify` (`id`) ON DELETE CASCADE,
  CONSTRAINT `send_log_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='模板信息';

-- ----------------------------
-- Records of send_log
-- ----------------------------
BEGIN;
INSERT INTO `send_log` (`id`, `notify_id`, `channel_id`, `user`, `variable`, `content`, `from_server`, `from_ip`, `status`, `err_code`, `reason`, `cost`, `created_at`, `updated_at`) VALUES (2, 1, 2, 'oavPm55S7trBviDM-LsBKfuvd1dk', '{\"captcha\":\"测试\"}', '{\"fields\":[{\"color\":\"#000000\",\"keyword\":\"thing7\",\"name\":\"用户名称\",\"value\":\"测试\"},{\"color\":\"#000000\",\"keyword\":\"amount2\",\"name\":\"提现金额\",\"value\":\"1.0\"},{\"color\":\"#000000\",\"keyword\":\"thing4\",\"name\":\"提现方式\",\"value\":\"网银\"},{\"color\":\"#000000\",\"keyword\":\"character_string6\",\"name\":\"账户编号\",\"value\":\"128821\"},{\"color\":\"#000000\",\"keyword\":\"time10\",\"name\":\"交易时间\",\"value\":\"2024-12-21\"}],\"templateId\":\"MrbCKbxeyl7K96chiaFAr0kahk4U7YLyQLT6SHsltA4\",\"title\":\"提现申请成功通知\",\"jumpType\":\"url\",\"url\":\"http://www.baidu.com\"}', '', '127.0.0.1', 1, 0, '', 1, 1736856957, 1736856957);
INSERT INTO `send_log` (`id`, `notify_id`, `channel_id`, `user`, `variable`, `content`, `from_server`, `from_ip`, `status`, `err_code`, `reason`, `cost`, `created_at`, `updated_at`) VALUES (3, 1, 2, 'oavPm55S7trBviDM-LsBKfuvd1dk', '{\"captcha\":\"hhhh\"}', '{\"fields\":[{\"color\":\"#000000\",\"keyword\":\"thing7\",\"name\":\"用户名称\",\"value\":\"hhhh\"},{\"color\":\"#000000\",\"keyword\":\"amount2\",\"name\":\"提现金额\",\"value\":\"1.0\"},{\"color\":\"#000000\",\"keyword\":\"thing4\",\"name\":\"提现方式\",\"value\":\"网银\"},{\"color\":\"#000000\",\"keyword\":\"character_string6\",\"name\":\"账户编号\",\"value\":\"128821\"},{\"color\":\"#000000\",\"keyword\":\"time10\",\"name\":\"交易时间\",\"value\":\"2024-12-21\"}],\"templateId\":\"MrbCKbxeyl7K96chiaFAr0kahk4U7YLyQLT6SHsltA4\",\"title\":\"提现申请成功通知\",\"jumpType\":\"url\",\"url\":\"http://www.baidu.com\"}', 'system-test', '', 1, 0, '', 1, 1736861990, 1736861990);
INSERT INTO `send_log` (`id`, `notify_id`, `channel_id`, `user`, `variable`, `content`, `from_server`, `from_ip`, `status`, `err_code`, `reason`, `cost`, `created_at`, `updated_at`) VALUES (4, 1, 2, 'oavPm55S7trBviDM-LsBKfuvd1dk', '{\"captcha\":\"hhhh\"}', '{\"fields\":[{\"color\":\"#000000\",\"keyword\":\"thing7\",\"name\":\"用户名称\",\"value\":\"hhhh\"},{\"color\":\"#000000\",\"keyword\":\"amount2\",\"name\":\"提现金额\",\"value\":\"1.0\"},{\"color\":\"#000000\",\"keyword\":\"thing4\",\"name\":\"提现方式\",\"value\":\"网银\"},{\"color\":\"#000000\",\"keyword\":\"character_string6\",\"name\":\"账户编号\",\"value\":\"128821\"},{\"color\":\"#000000\",\"keyword\":\"time10\",\"name\":\"交易时间\",\"value\":\"2024-12-21\"}],\"templateId\":\"MrbCKbxeyl7K96chiaFAr0kahk4U7YLyQLT6SHsltA4\",\"title\":\"提现申请成功通知\",\"jumpType\":\"url\",\"url\":\"http://www.baidu.com\"}', 'system-test', '127.0.0.1', 1, 0, '', 1, 1736862315, 1736862315);
COMMIT;

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `notify_id` bigint unsigned NOT NULL COMMENT '通知id',
  `channel_id` bigint unsigned NOT NULL COMMENT '渠道id',
  `content` mediumtext NOT NULL COMMENT '通知内容',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `weight` bigint NOT NULL COMMENT '权重',
  `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `notify_id` (`notify_id`),
  KEY `channel_id` (`channel_id`),
  CONSTRAINT `template_ibfk_1` FOREIGN KEY (`notify_id`) REFERENCES `notify` (`id`) ON DELETE CASCADE,
  CONSTRAINT `template_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='模板信息';

-- ----------------------------
-- Records of template
-- ----------------------------
BEGIN;
INSERT INTO `template` (`id`, `notify_id`, `channel_id`, `content`, `status`, `weight`, `created_at`, `updated_at`) VALUES (1, 1, 1, '<div style=\"margin:auto;max-width: 700px;\">\n    <div style=\"padding:10px 0; margin:0 10px; display: flex;align-items: center; \">\n        <image\n            src=\"http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image\"\n            width=\'35px\' height=\'35px\' />\n        <span style=\"margin: 0 10px;font-size: 20px; font-weight: 700;\">统一应用管理平台</span>\n    </div>\n    <div\n        style=\"box-sizing: border-box;min-height: 100px;padding:15px;margin:0 10px 25px 10px; line-height: 24px; text-align: justify; border-top: 2px solid #1e90ff; background-color: #fefefe;box-shadow: 0 1px 8px rgba(0,0,0,0.1);\">\n        <div style=\"font-size:14px;white-space: normal;margin: 0px;padding: 0px;box-sizing: border-box;\">\n            您的邮箱验证码为： ${captcha}，该验证码在${minute}分钟内有效，为了账号安全，请勿向其他人泄漏验证码信息。\n        </div>\n    </div>\n    <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\">\n        <tbody>\n            <tr>\n                <td align=\"center\" style=\"color:#76808E;font-size:13px;\">开放协作，拥抱未来</td>\n            </tr>\n            <tr>\n                <td align=\"center\" style=\"color:#76808E;font-size:13px;\">\n                    Copyright © 2024 qlime.cn. All rights reserved.\n                </td>\n            </tr>\n        </tbody>\n    </table>\n</div>', 1, 0, 1736250271, 1736499861);
INSERT INTO `template` (`id`, `notify_id`, `channel_id`, `content`, `status`, `weight`, `created_at`, `updated_at`) VALUES (2, 1, 2, '{\"fields\":[{\"color\":\"#000000\",\"keyword\":\"thing7\",\"name\":\"用户名称\",\"value\":\"${captcha}\"},{\"color\":\"#000000\",\"keyword\":\"amount2\",\"name\":\"提现金额\",\"value\":\"1.0\"},{\"color\":\"#000000\",\"keyword\":\"thing4\",\"name\":\"提现方式\",\"value\":\"网银\"},{\"color\":\"#000000\",\"keyword\":\"character_string6\",\"name\":\"账户编号\",\"value\":\"128821\"},{\"color\":\"#000000\",\"keyword\":\"time10\",\"name\":\"交易时间\",\"value\":\"2024-12-21\"}],\"templateId\":\"MrbCKbxeyl7K96chiaFAr0kahk4U7YLyQLT6SHsltA4\",\"title\":\"提现申请成功通知\",\"jumpType\":\"url\",\"url\":\"http://www.baidu.com\"}', 1, 1, 1736830850, 1736834895);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
