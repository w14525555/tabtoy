local title = {['TypeID']=1,['GiftType']=2,['Rebate']=3}
local data = {
{ 1,5,{{1,2011401,2},{1,8000003,15},{1,2011201,2}} },
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
