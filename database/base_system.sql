-- name: create-admin
INSERT INTO `admin` (`id`, `login_name`, `real_name`, `password`, `role_ids`, `phone`, `email`, `avatar`, `salt`, `last_ip`, `last_login`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`, `level`)
VALUES
	(1,'admin','admin','8ebf088c456f925ce22b881eff76bec7','0','','','/static/admin/images/avatar.jpg','BtBFNwXINC','::1','2020-10-15 17:07:29',1,0,0,'2020-09-27 17:44:35','2020-10-15 17:07:29',99),
	(2,'ceshi','测试账号','fa3fb5825c2e64bc764f29245dd1ec7a','1,2','13988009988','abc@188.com',NULL,'i8Nf','','2020-01-01 15:01:01',1,1,0,NULL,'2020-10-09 11:38:42',1);

-- name: create-auth
INSERT INTO `auth` (`id`, `auth_name`, `auth_url`, `user_id`, `pid`, `sort`, `icon`, `is_show`, `status`,`power_type`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
    (2, '权限管理', '/', 1, 0, 3, 'layui-icon layui-icon-survey', 1, 1, 0, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (3, '管理员', '/system/admin/list', 1, 2, 1, 'fa-user-o', 1, 1, 1, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (4, '角色管理', '/system/role/list', 1, 2, 2, 'fa-user-circle-o', 1, 1, 1, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (5, '新增', '/system/admin/add', 1, 3, 1, '', 0, 1, 2, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (6, '修改', '/system/admin/edit', 1, 3, 2, '', 0, 1, 2, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (7, '删除', '/system/admin/delete', 1, 3, 3, '', 0, 1, 2, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (8, '新增', '/system/role/add', 1, 4, 1, '', 1, 1, 2, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (9, '修改', '/system/role/edit', 1, 4, 2, '', 0, 1, 2, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (10, '删除', '/system/role/delete', 1, 4, 3, '', 0, 1, 2, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (11, '权限因子', '/system/auth/list', 1, 2, 3, 'fa-list', 1, 1, 1, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (12, '新增', '/system/auth/edit', 1, 11, 1, '', 0, 1, 2, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (13, '修改', '/system/auth/edit', 1, 11, 2, '', 0, 1, 2, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (14, '删除', '/system/auth/delete', 1, 11, 3, '', 0, 1, 2, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (15, '个人中心', 'profile/edit', 1, 0, 4, 'layui-icon layui-icon-user', 1, 1, 0, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (20, '基础设置', '/', 1, 0, 2, 'layui-icon layui-icon-auz', 1, 1, 0, 1, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (24, '资料修改', '/system/user/edit', 1, 15, 1, 'fa-edit', 1, 1, 1, 0, 1, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (45, '权限树', '/system/auth/nodes', 0, 11, 4, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (46, '单个权限获取', '/system/auth/node', 0, 11, 5, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (47, '管理员列表', '/system/admin/json', 0, 3, 4, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (48, '角色列表', '/system/role/json', 0, 4, 4, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (49, '站点设置', '/system/site/edit', 0, 20, 1, 'layui-icon layui-icon-home', 1, 1, 1, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (50, '上传图片', '/system/upload', 0, 49, 1, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (51, '更新站点设置', '/system/site/edit', 0, 49, 2, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (55, '更新头像', '/system/user/avatar', 0, 24, 2, '', 0, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (56, '修改密码', '/system/user/pwd', 0, 15, 2, '', 1, 1, 2, 0, 0, '2020-10-14 17:04:30', '2020-10-14 17:04:30'),
     (59, '工作空间', '/', 0, 0, 1, 'layui-icon layui-icon-console', 1, 1, 0, 0, 0, '2021-05-21 15:16:44', '2021-05-21 15:16:44'),
     (60, '后台首页', '/system/main', 0, 59, 1, 'layui-icon layui-icon-rate', 1, 1, 1, 0, 0, '2021-05-21 15:22:36', '2021-05-21 15:22:36'),
     (61, '文件迁移', '/system/backup', 0, 0, 5, 'layui-icon layui-icon-upload-drag', 1, 1, 0, 0, 0, '2021-09-26 15:16:06.8134987+08:00', '2021-09-26 15:16:06.8134987+08:00'),
     (62, '任务管理', '/system/backup/', 0, 61, 1, 'layui-icon ', 1, 1, 1, 0, 0, '2021-09-26 15:19:04.9333835+08:00', '2021-09-26 15:19:04.9333835+08:00'),
     (63, '服务器管理', '/system/server', 0, 61, 2, 'layui-icon ', 1, 1, 1, 0, 0, '2021-09-26 15:19:44.6511204+08:00', '2021-09-26 15:19:44.6511204+08:00'),
     (64, '列表', '/system/backup/list', 0, 62, 1, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:22:30.3483503+08:00', '2021-09-26 15:22:30.3483503+08:00'),
     (65, '数据', '/system/backup/json', 0, 62, 2, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:23:23.0356+08:00', '2021-09-26 15:23:23.0356+08:00'),
     (66, '新增', '/system/backup/add', 0, 62, 3, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:23:54.3020752+08:00', '2021-09-26 15:23:54.3020752+08:00'),
     (67, '修改', '/system/backup/edit', 0, 62, 4, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:24:35.2751281+08:00', '2021-09-26 15:24:35.2751281+08:00'),
     (68, '删除', '/system/backup/delete', 0, 62, 5, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:25:32.0431483+08:00', '2021-09-26 15:25:32.0431483+08:00'),
     (69, '列表', '/system/server/list', 0, 63, 1, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:27:06.6511277+08:00', '2021-09-26 15:27:06.6511277+08:00'),
     (70, '数据', '/system/server/json', 0, 63, 2, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:27:57.5155823+08:00', '2021-09-26 15:27:57.5155823+08:00'),
     (71, '新增', '/system/server/add', 0, 63, 3, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:28:14.1551906+08:00', '2021-09-26 15:28:14.1551906+08:00'),
     (72, '修改', '/system/server/edit', 0, 63, 4, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:28:52.4433335+08:00', '2021-09-26 15:28:52.4433335+08:00'),
     (73, '删除', '/system/server/delete', 0, 63, 5, 'layui-icon ', 0, 1, 2, 0, 0, '2021-09-26 15:29:09.611474+08:00', '2021-09-26 15:29:09.611474+08:00');

-- name: create-role
INSERT INTO `role` (`id`, `role_name`, `detail`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
    (1,'管理员','拥有管理权限',1,0,0,'2020-09-28 14:00:10','2020-09-28 14:00:06'),
	(2,'客服','拥有客服权限',1,0,2,'2020-09-28 13:59:53','2020-09-28 14:00:03');


-- name: create-role-auth
INSERT INTO `role_auth` (`role_id`, `auth_id`)
VALUES
	(1,16),
	(1,17),
	(1,18),
	(1,19),
	(2,0),
	(2,1),
	(2,15),
	(2,20),
	(2,21),
	(2,22),
	(2,23),
	(2,24);

-- name: create-pear-config
INSERT INTO `pear_config`(`id`, `created_at`, `updated_at`, `config_type`, `config_data`, `config_status`)
VALUES 
    (1, '2021-05-28 10:48:35', '2021-05-28 10:48:35', 'pear-config', '{"colors":[{"color":"#2d8cf0","id":"1"},{"color":"#5FB878","id":"2"},{"color":"#1E9FFF","id":"3"},{"color":"#FFB800","id":"4"},{"color":"darkgray","id":"5"}],"header":{"message":"/static/admin/data/message.json"},"links":[{"href":"http://www.pearadmin.com","icon":"layui-icon layui-icon-auz","title":"官方网站"},{"href":"http://www.pearadmin.com","icon":"layui-icon layui-icon-auz","title":"开发文档"},{"href":"https://gitee.com/Jmysy/Pear-Admin-Layui","icon":"layui-icon layui-icon-auz","title":"开源地址"}],"logo":{"image":"/static/admin/images/logo.png","title":"Pear Admin"},"menu":{"accordion":true,"control":false,"data":"/system/menu","method":"GET","select":"60"},"other":{"autoHead":false,"keepLoad":100},"tab":{"index":{"href":"/system/main","id":"60","title":"首页"},"keepState":true,"muiltTab":true,"tabMax":30},"theme":{"allowCustom":true,"defaultColor":"2","defaultMenu":"dark-theme"}}', 1);
