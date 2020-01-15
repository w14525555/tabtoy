local title = {['RewardIndex']=1,['Day']=2,['FirstDayReward']=3,['BagOccupy']=4}
local data = {
{ 1,0,{{3,1,1,1,1},{1,9000204,1,1},{1,7000027,1,1,1},{1,8000005,5,1},{1,2011403,1,1}},5 },
{ 2,1,{{3,2,1,1,1},{1,14030101,3,1},{1,1003012,3,1},{1,7300001,3,1}},4 },
{ 3,2,{{1,2011902,1,1,1},{1,14090102,1,1},{1,1000022,10,1},{1,1000005,10,1,1}},4 },
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
