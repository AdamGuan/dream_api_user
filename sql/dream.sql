-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.1.73 - Source distribution
-- 服务器操作系统:                      redhat-linux-gnu
-- HeidiSQL 版本:                  9.1.0.4867
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 dream_api_user 的数据库结构
DROP DATABASE IF EXISTS `dream_api_user`;
CREATE DATABASE IF NOT EXISTS `dream_api_user` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `dream_api_user`;


-- 导出  表 dream_api_user.t_config_pkg 结构
DROP TABLE IF EXISTS `t_config_pkg`;
CREATE TABLE IF NOT EXISTS `t_config_pkg` (
  `F_pkg` varchar(250) NOT NULL COMMENT '包名',
  `F_app_name` varchar(250) NOT NULL COMMENT '包对应的应用名字',
  `F_app_id` varchar(250) NOT NULL COMMENT 'leancloud对应的app id',
  `F_app_key` varchar(250) NOT NULL COMMENT 'leancloud对应的app key',
  `F_app_master_key` varchar(250) NOT NULL COMMENT 'leancloud对应的master key',
  `F_app_msm_template` varchar(250) NOT NULL COMMENT 'leancloud对应的短信模板名',
  UNIQUE KEY `F_pkg` (`F_pkg`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='包相关信息';

-- 正在导出表  dream_api_user.t_config_pkg 的数据：8 rows
DELETE FROM `t_config_pkg`;
/*!40000 ALTER TABLE `t_config_pkg` DISABLE KEYS */;
INSERT INTO `t_config_pkg` (`F_pkg`, `F_app_name`, `F_app_id`, `F_app_key`, `F_app_master_key`, `F_app_msm_template`) VALUES
	('cn.dream.android.shuati.debug', '刷题有道debug', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('cn.dream.android.shuati', '刷题有道', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('com.dream.phone.wenba.debug', '搜作业debug', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('com.dream.phone.wenba', '搜作业', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('abc', '测试', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('com.team.englishsquare', '英语广场', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('com.team.englishsquare.debug', '英语广场debug', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'test2'),
	('webdream', '追梦网站', 'ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'template1');
/*!40000 ALTER TABLE `t_config_pkg` ENABLE KEYS */;


-- 导出  表 dream_api_user.t_config_response 结构
DROP TABLE IF EXISTS `t_config_response`;
CREATE TABLE IF NOT EXISTS `t_config_response` (
  `F_response_no` smallint(5) NOT NULL COMMENT '响应code',
  `F_response_msg` varchar(50) NOT NULL COMMENT '响应信息'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='api的响应配置';

-- 正在导出表  dream_api_user.t_config_response 的数据：14 rows
DELETE FROM `t_config_response`;
/*!40000 ALTER TABLE `t_config_response` DISABLE KEYS */;
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES
	(0, '成功'),
	(-1, '失败'),
	(-2, '已注册'),
	(-3, '密码不符合规则'),
	(-4, '没有注册'),
	(-5, '用户名或密码错误'),
	(-6, '签名错误'),
	(-7, '包名不存在'),
	(-8, '现有密码错误'),
	(-9, '密码错误'),
	(-10, '参数错误'),
	(-16, '手机号错误'),
	(-17, '没有数据'),
	(-23, '新手机号码无效，已被注册');
/*!40000 ALTER TABLE `t_config_response` ENABLE KEYS */;


-- 导出  表 dream_api_user.t_sms_action_valid 结构
DROP TABLE IF EXISTS `t_sms_action_valid`;
CREATE TABLE IF NOT EXISTS `t_sms_action_valid` (
  `F_action` char(32) NOT NULL COMMENT '动作，(md5(phone+pkg+sms))',
  UNIQUE KEY `F_action` (`F_action`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='记录每个动作对应的短信验证码，用于安全验证。暂时的，会改为redis';

-- 正在导出表  dream_api_user.t_sms_action_valid 的数据：0 rows
DELETE FROM `t_sms_action_valid`;
/*!40000 ALTER TABLE `t_sms_action_valid` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_sms_action_valid` ENABLE KEYS */;


-- 导出  表 dream_api_user.t_sms_rate 结构
DROP TABLE IF EXISTS `t_sms_rate`;
CREATE TABLE IF NOT EXISTS `t_sms_rate` (
  `F_action` char(32) NOT NULL COMMENT '动作，(md5(phone+pkg))',
  `F_last_timestamp` datetime NOT NULL COMMENT '时间',
  UNIQUE KEY `F_action` (`F_action`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='记录短信发送的频率，用于限制短信的频繁发送，暂时的，会改为redis';

-- 正在导出表  dream_api_user.t_sms_rate 的数据：0 rows
DELETE FROM `t_sms_rate`;
/*!40000 ALTER TABLE `t_sms_rate` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_sms_rate` ENABLE KEYS */;


-- 导出  表 dream_api_user.t_user 结构
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE IF NOT EXISTS `t_user` (
  `F_user_name` varchar(50) NOT NULL COMMENT '用户ID',
  `F_user_password` char(40) NOT NULL COMMENT '用户密码',
  `F_user_phone` varchar(15) NOT NULL COMMENT '手机号码',
  `F_crate_datetime` datetime NOT NULL COMMENT '创建时间',
  `F_modify_datetime` datetime NOT NULL COMMENT '修改时间',
  UNIQUE KEY `F_user_name` (`F_user_name`),
  UNIQUE KEY `F_user_phone` (`F_user_phone`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 正在导出表  dream_api_user.t_user 的数据：0 rows
DELETE FROM `t_user`;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
