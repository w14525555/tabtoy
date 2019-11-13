local title = {['configIndex']=1,['ShowBeginTime']=2,['TimeType']=3}
local data = {
{ 1,'"2019-1-29 5:00:00"',1 },
{ 2,'"2019-2-11 5:00:00"',1 },
{ 3,'"2019-2-11 5:00:00"',1 },
{ 4,'"11 5:00:00"',2 },
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
