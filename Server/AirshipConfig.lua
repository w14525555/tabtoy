local title = {['AirshipID']=1,['MapType']=2,['PathID']=3,['HorizontalSpeed']=4,['NextAirship']=5,['FinishQuestStep']=6,['AcceptQuestStep']=7}
local data = {
{ 1,6,1,15,2,0,0 },
{ 2,7,1,15,0,0,0 },
{ 9,7,9,15,0,0,0 },
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
