local title = {['EquipLevel']=1,['Order']=2}
local data = {
{ 0,1 },
{ 1,1 },
{ 2,2 },
{ 3,2 },
{ 4,3 },
{ 5,3 },
{ 6,4 },
{ 7,4 },
{ 8,5 },
{ 9,5 },
{ 10,6 },
{ 11,6 },
{ 12,7 },
{ 13,7 },
{ 14,8 },
{ 15,8 },
{ 16,9 },
{ 17,9 },
{ 18,10 },
{ 19,10 },
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
