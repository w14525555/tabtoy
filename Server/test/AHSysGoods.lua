local title = {['GoodsType']=1,['TypeID']=2,['Count']=3,['Price']=4}
local data = {
{ 1,3010101,1,40000 },
{ 1,3010102,1,40000 },
{ 1,3010103,1,40000 },
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
