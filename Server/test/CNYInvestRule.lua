local title = {['ActivityId']=1,['RuleTitle']=2,['ContentList']=3}
local data = {
{ 10012,'投资规则说明',{{"投资时间","每轮投资均有一定投资日期，在时间到达时，可选择相应投资计划进行投资"},{"投资消耗","每次投资仅需消耗1枚五级劳尔晶体，即可获得丰厚收益"},{"投资收益","每次投资成功后，奖励池相应收益将会亮起，第一次投资可激活前3个奖励，第二次激活2个，第三次激活1个"},{"奖励领取","投资结束后,可在领奖期间领取投资收益，奖励多多，如果没在期间登录的玩家也不用担心，未领取的奖励将会为你发入邮件中"}} },
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
