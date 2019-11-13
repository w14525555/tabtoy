local title = {['ActivityId']=1,['Index']=2,['ItemData']=3,['BeginTime']=4,['EndTime']=5}
local data = {
{ 10012,1,{1,8000005,1},'"1548259200"','"1548431999"' },
{ 10012,2,{1,8000005,1},'"1548432000"','"1548604799"' },
{ 10012,3,{1,8000005,1},'"1548604800"','"1548777599"' },
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
