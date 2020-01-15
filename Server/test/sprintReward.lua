local title = {}
local data = {
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
