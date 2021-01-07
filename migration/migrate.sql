CREATE TABLE `t_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '账号',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `created_on` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `is_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 1 启用 0 禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB ;

