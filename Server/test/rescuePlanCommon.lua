local title = {['ID']=1,['key']=2,['type']=3,['value1']=4,['value2']=5}
local data = {
{ 1,'MainViewOpenTimeText',2,0,'   1月17日-1月23日（曙光服）\n   2月5日-2月11日（全服）\n   2月5日之后的新服活动持续1周\n   每天的9:00-24:00开放玩法' },
{ 2,'MainViewRuleText',2,0,'    3-5人组成队伍前往前线，在限定时间内护送飞船运抵终点支援远古地球。协助别人也能获得奖励哦！' },
{ 3,'MainViewBGText',2,0,'    凛冬降临，沃灵冻土出现了神秘的"极能场域"，通过调查发现，“极能场域”链接到了20亿年前地球南极的某个坐标点上。\n    而当时的地球，正命悬一线，挣扎着离开太阳系。\n    如果不保护好多美祖先-远古地球，那么“时空悖论”将会彻底改写多美星系的历史，甚至摧毁掉这个生机焕发的星球。\n    两个星球穿越20亿年的时空，在多美国王“救援计划”的强烈号召下，彼此务必同舟共济！快行动吧，勇者们！' },
{ 4,'ObstructorRadius',1,2.4,'' },
{ 5,'ObstructorTypeID',1,20017404,'' },
}
local mt = {__index = function (table,key)
    local temp = title[key]
    if temp then
        return table[temp]
    end
    return nil
end}
for k, v in pairs(data) do
    setmetatable(v, mt)
end
return {data=data, title=title}
