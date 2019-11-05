local title = {['BadgeID']=1,['Name']=2,['UnlockLevel']=3,['WearLevel']=4}
local data = {
{ 1,'勇武',200,10 },
{ 2,'睿智',260,12 },
{ 3,'坚韧',280,13 },
{ 4,'慷慨',300,14 },
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
