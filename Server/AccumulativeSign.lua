local title = {['TypeID']=1,['AccumulativeCount']=2,['LevelCap']=3,['Reward']=4}
local data = {
{ 1,2,59,{{2011391,1,100}} },
{ 2,5,59,{{2011392,1,100}} },
{ 3,10,59,{{2011393,1,100}} },
{ 4,20,59,{{2011394,1,100}} },
{ 5,2,69,{{2011391,1,100}} },
{ 6,5,69,{{2011392,1,100}} },
{ 7,10,69,{{2011393,1,100}} },
{ 8,20,69,{{2011394,1,100}} },
{ 9,2,79,{{2011391,1,100}} },
{ 10,5,79,{{2011392,1,100}} },
{ 11,10,79,{{2011393,1,100}} },
{ 12,20,79,{{2011394,1,100}} },
{ 13,2,89,{{2011391,1,100}} },
{ 14,5,89,{{2011392,1,100}} },
{ 15,10,89,{{2011393,1,100}} },
{ 16,20,89,{{2011394,1,100}} },
{ 17,2,99,{{2011391,1,100}} },
{ 18,5,99,{{2011392,1,100}} },
{ 19,10,99,{{2011393,1,100}} },
{ 20,20,99,{{2011394,1,100}} },
{ 21,2,109,{{2011391,1,100}} },
{ 22,5,109,{{2011392,1,100}} },
{ 23,10,109,{{2011393,1,100}} },
{ 24,20,109,{{2011394,1,100}} },
{ 25,2,119,{{2011391,1,100}} },
{ 26,5,119,{{2011392,1,100}} },
{ 27,10,119,{{2011393,1,100}} },
{ 28,20,119,{{2011394,1,100}} },
{ 29,2,129,{{2011391,1,100}} },
{ 30,5,129,{{2011392,1,100}} },
{ 31,10,129,{{2011393,1,100}} },
{ 32,20,129,{{2011394,1,100}} },
{ 33,2,139,{{2011391,1,100}} },
{ 34,5,139,{{2011392,1,100}} },
{ 35,10,139,{{2011393,1,100}} },
{ 36,20,139,{{2011394,1,100}} },
{ 37,2,149,{{2011391,1,100}} },
{ 38,5,149,{{2011392,1,100}} },
{ 39,10,149,{{2011393,1,100}} },
{ 40,20,149,{{2011394,1,100}} },
{ 41,2,159,{{2011391,1,100}} },
{ 42,5,159,{{2011392,1,100}} },
{ 43,10,159,{{2011393,1,100}} },
{ 44,20,159,{{2011394,1,100}} },
{ 45,2,169,{{2011391,1,100}} },
{ 46,5,169,{{2011392,1,100}} },
{ 47,10,169,{{2011393,1,100}} },
{ 48,20,169,{{2011394,1,100}} },
{ 49,2,179,{{2011391,1,100}} },
{ 50,5,179,{{2011392,1,100}} },
{ 51,10,179,{{2011393,1,100}} },
{ 52,20,179,{{2011394,1,100}} },
{ 53,2,189,{{2011391,1,100}} },
{ 54,5,189,{{2011392,1,100}} },
{ 55,10,189,{{2011393,1,100}} },
{ 56,20,189,{{2011394,1,100}} },
{ 57,2,199,{{2011391,1,100}} },
{ 58,5,199,{{2011392,1,100}} },
{ 59,10,199,{{2011393,1,100}} },
{ 60,20,199,{{2011394,1,100}} },
{ 61,2,209,{{2011391,1,100}} },
{ 62,5,209,{{2011392,1,100}} },
{ 63,10,209,{{2011393,1,100}} },
{ 64,20,209,{{2011394,1,100}} },
{ 65,2,219,{{2011391,1,100}} },
{ 66,5,219,{{2011392,1,100}} },
{ 67,10,219,{{2011393,1,100}} },
{ 68,20,219,{{2011394,1,100}} },
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