local title = {['DialogID']=1,['ButtonID']=2,['Desc']=3,['ButtonType']=4,['ButtonParam']=5,['IconDefIndex']=6}
local data = {
{ 1001,4,'1001按钮4',2,'',0 },
{ 1001,5,'1001按钮5',4,'',0 },
{ 1001,6,'1001按钮6',1,'',0 },
{ 1001,7,'显示商店1',3,{StoreType=1},0 },
{ 1018,1,'打开界面',2,'',0 },
{ 1019,1,'了解玩法',2,'',0 },
{ 1019,2,'了解宣战夺宝',2,'',0 },
{ 1020,1,'离开副本',2,'',0 },
{ 1021,1,'我愿前往',2,'',0 },
{ 1022,1,'我愿前往',2,'',0 },
{ 1023,1,'我愿前往',2,'',0 },
{ 1024,1,'我愿前往',2,'',0 },
{ 1025,1,'',6,'',0 },
{ 1026,1,'',6,'',0 },
{ 1027,1,'',6,'',0 },
{ 1028,1,'',6,'',0 },
{ 1029,1,'',6,'',0 },
{ 1030,1,'',6,'',0 },
{ 1031,1,'',6,'',0 },
{ 1032,1,'',6,'',0 },
{ 1033,1,'',6,'',0 },
{ 1034,1,'',6,'',0 },
{ 1035,1,'',6,'',0 },
{ 1036,1,'',6,'',0 },
{ 1037,1,'',6,'',0 },
{ 1038,1,'',6,'',0 },
{ 1039,1,'',6,'',0 },
{ 1800,1,'领取联邦任务',4,{QuestID=102001},0 },
{ 2001,1,'虫族任务',2,'',0 },
{ 3001,1,'加入军团',2,'',0 },
{ 3001,2,'创建军团',2,'',0 },
{ 3001,3,'回到军团',2,'',0 },
{ 3002,1,'选择合团',2,'',0 },
{ 3002,2,'合团信息',2,'',0 },
{ 3003,1,'管理福利',2,'',0 },
{ 3003,2,'派对礼包',2,'',0 },
{ 3004,1,'军团事件',2,'',0 },
{ 3005,1,'军团任务',2,'',0 },
{ 3006,1,'军团技能',2,'',0 },
{ 3007,1,'资源研究站',2,'',0 },
{ 3008,1,'助力传承',2,'',0 },
{ 3011,1,'联邦试炼',2,'',0 },
{ 3101,1,'开启风暴试炼',2,'',0 },
{ 3201,1,'前往进化',2,'',0 },
{ 3201,2,'进化排行',2,'',0 },
{ 3201,3,'进化简介',2,'',0 },
{ 4000,1,'我要参加竞速',2,'',0 },
{ 4000,2,'竞速排行榜',2,'',0 },
{ 5001,1,'房间改名',2,'',0 },
{ 5001,2,'全部修理',2,'',0 },
{ 5001,3,'拍照分享',2,'',0 },
{ 5002,1,'升级别墅',2,'',0 },
{ 5100,1,'参加比赛',2,'',0 },
{ 5100,2,'比赛规则',2,'',0 },
{ 1004,1,'杂货商店',3,{StoreType=1},0 },
{ 6001,1,'送我去参赛',2,'',2 },
{ 6001,2,'查看赛事规则',2,'',0 },
{ 6001,4,'决赛对阵',2,'',0 },
{ 6001,5,'观看录像',2,'',0 },
{ 6002,1,'成立战队',2,'',0 },
{ 6002,2,'战队管理',2,'',0 },
{ 6007,1,'时空漩涡',2,'',0 },
{ 6007,2,'组队平台',2,'',0 },
{ 6008,1,'进入轮回走廊',2,'',0 },
{ 7001,1,'挑战别西卜',2,'',0 },
{ 8001,1,'前去战场',2,'',2 },
{ 8001,2,'了解玩法',2,'',0 },
{ 8001,3,'前往进化',2,'',0 },
{ 8002,1,'参与巅峰挑战',2,'',2 },
{ 8002,2,'了解玩法',2,'',0 },
{ 8002,3,'首席剑武',2,'',0 },
{ 8002,4,'首席异能',2,'',0 },
{ 8002,5,'首席格斗',2,'',0 },
{ 8002,6,'首席枪械',2,'',0 },
{ 8002,7,'首席治愈',2,'',0 },
{ 8003,1,'开始挑战',2,'',0 },
{ 8004,1,'挑战军团试炼',2,'',0 },
{ 8004,2,'玩法简介',2,'',0 },
{ 8004,3,'奖励查询',2,'',0 },
{ 8004,4,'排行榜',2,'',0 },
{ 8005,1,'领取道具1号',2,'',0 },
{ 8006,1,'领取道具2号',2,'',0 },
{ 8007,1,'领取道具3号',2,'',0 },
{ 8008,1,'领取道具4号',2,'',0 },
{ 8009,1,'领取道具5号',2,'',0 },
{ 8010,1,'与善良格鲁切磋',2,'',0 },
{ 8010,2,'消灭邪恶的格鲁',2,'',0 },
{ 8011,1,'进入星战界面',2,'',0 },
{ 8011,2,'星战流程图',2,'',0 },
{ 8011,3,'详细规则',2,'',0 },
{ 8011,4,'上次星战结果',2,'',0 },
{ 8012,1,'领取任务链',2,'',1 },
{ 8012,2,'放弃任务链',2,'',0 },
{ 8013,1,'开始战斗',2,'',0 },
{ 8013,2,'军团求助',2,'',0 },
{ 8014,1,'参加星斗场',2,'',0 },
{ 8015,1,'查看装备',2,'',0 },
{ 8015,2,'开始膜拜',2,'',0 },
{ 8016,1,'查看装备',2,'',0 },
{ 8017,1,'查看装备',2,'',0 },
{ 8020,1,'领取联邦通缉令',2,'',0 },
{ 8021,1,'进入副本',2,'',0 },
{ 8021,2,'对阵信息',2,'',0 },
{ 8021,3,'玩法介绍',2,'',0 },
{ 8021,4,'玩法指引',2,'',0 },
{ 8021,5,'战场预览',2,'',0 },
{ 8022,1,'进入比赛场',1,'',0 },
{ 9006,1,'(对战)捣蛋大作战',2,'',0 },
{ 9006,2,'(限时)跨年红包',2,'',0 },
{ 9008,1,'前往圣殿',2,'',0 },
{ 9503,1,'参与任务',2,'',0 },
{ 9503,2,'便捷组队',2,'',0 },
{ 17001,1,'幸运轮盘',2,'',0 },
{ 17010,1,'提交%s',2,{RelatedItemTypeID=1018001},0 },
{ 17010,2,'提交%s',2,{RelatedItemTypeID=1018002},0 },
{ 17010,3,'提交%s',2,{RelatedItemTypeID=1018003},0 },
{ 17010,4,'提交%s',2,{RelatedItemTypeID=1018004},0 },
{ 17010,5,'活动介绍',2,'',0 },
{ 17011,1,'继续提交',2,'',0 },
{ 18001,1,'感恩礼物',1,'',0 },
{ 19000,1,'许愿树',6,'',0 },
{ 99999,1,'世界领主',2,'',0 },
{ 20000,1,'幻影交锋',2,'',0 },
{ 20000,2,'便捷组队',2,'',0 },
{ 20010,1,'国庆冲鸭',2,'',0 },
{ 21010,1,'组队挑战',2,'',0 },
{ 99999,1,'世界领主',2,'',0 },
{ 61002803,1,'马上去救茉莉',7,{DialogID=20001},0 },
{ 20100,1,'参加捣蛋大作战',2,'',0 },
{ 20100,2,'了解玩法',2,'',0 },
{ 20101,1,'玩法详情',2,'',0 },
{ 16012,1,'参与活动',2,'',0 },
{ 7050,1,'我们要结义',2,'',0 },
{ 7050,2,'吸纳新成员',2,'',0 },
{ 7050,3,'我要退出结义',2,'',0 },
{ 7050,4,'发起踢出结义队员',2,'',0 },
{ 7050,5,'前往结义招募平台',2,'',0 },
{ 7050,6,'社交-结义',2,'',0 },
{ 7050,7,'结义说明',2,'',0 },
}
return {data, title}
