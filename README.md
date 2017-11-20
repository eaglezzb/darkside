# darkside


用户表
**UserModel**
    
    CREATE TABLE `user` (
      `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `username` varchar(16) DEFAULT '',
      `departname` varchar(20) DEFAULT '',
      `password` varchar(32) DEFAULT '',
      `sex` tinyint(1) unsigned zerofill DEFAULT '0',
      `userid` varchar(32) DEFAULT '',
      `phone` varchar(13) DEFAULT '',
      `phoneprefix` varchar(6) DEFAULT '',
      `createtime` int(10) unsigned NOT NULL,
      `updatetime` int(10) unsigned NOT NULL,
      `state` tinyint(1) DEFAULT '-1',
      `authtoken` char(32) DEFAULT '',
      `mail` varchar(40) DEFAULT '',
      `oldpassword` varchar(100) DEFAULT '',
      PRIMARY KEY (`uid`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1000047 DEFAULT CHARSET=utf8;
    
手机信息表 
**telephone**    
    
    CREATE TABLE `telephone` (
      `ncode` int(5) unsigned NOT NULL,
      `scount` int(11) unsigned DEFAULT '1',
      `mobile` varchar(13) NOT NULL DEFAULT '',
      PRIMARY KEY (`mobile`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    
    
发送的短信验证码记录表
**smstx**   
   
    CREATE TABLE `smstx` (
      `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `type` tinyint(1) NOT NULL DEFAULT '0',
      `message` varchar(200) NOT NULL,
      `result` int(6) DEFAULT '-1',
      `time` int(10) unsigned NOT NULL,
      `ext` varchar(200) DEFAULT '',
      `mobile` varchar(13) DEFAULT '',
      `ncode` int(5) unsigned NOT NULL,
      `errmsg` varchar(30) DEFAULT '',
      `sid` varchar(64) DEFAULT '',
      `fee` int(1) NOT NULL DEFAULT '0',
      `smscode` varchar(6) NOT NULL DEFAULT '',
      `status` int(1) DEFAULT '0',
      PRIMARY KEY (`uid`)
    ) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8;
    
**mailInfo**     
    
    CREATE TABLE `mailinfo` (
      `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `mail` varchar(30) NOT NULL DEFAULT '',
      `verifycode` char(6) NOT NULL DEFAULT '',
      `message` varchar(300) DEFAULT '',
      `createtime` int(11) NOT NULL,
      `type` tinyint(1) DEFAULT '0',
      `status` tinyint(1) DEFAULT '0',
      PRIMARY KEY (`uid`)
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
    
    
**用户关系表**    

    CREATE TABLE `user_friend_connect` (
      `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `userid_1` varchar(32) NOT NULL DEFAULT '',
      `userid_2` varchar(32) NOT NULL DEFAULT '',
      PRIMARY KEY (`uid`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;