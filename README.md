# asset-analysis
资产分析

1. asset save: 
    database is mysql.
    use two method to save data:
    a. http post form  
    b. spider
    filds is as below:
    利润表相关数据
    1. 营业收入
    2. 营业成本
    3. 毛利润(= 1 - 2)
    4. 营业费用(销售费用，管理费用，财务费用)
    5. 税金&资产减值
    6. 营业利润（=1 - 2 - 4 - 5）
    7. 营业外收入
    8. 营业外支出
    9. 税前利润
    9. 税
    10. 净利润(6 + 7 - 8 - 9)
2. asset analysis:
    compute by a server job
    analysis data is as below:
    1. 利润率（10年内）
3. asset view 
    view by web page, show asset's value and reliability list, click to view the detial.
    list is order by 利润率.
4. algrithm
    define some valuable parameter, different algrithm can save different value to get its best asset
    1. 毛利率 = 毛利润/营业收入（>40%），观察年限(10)
    2. 营业收入增长率
    3. 营业利润增长率
    4. 净利润增长率
5. data store
    save by file through json or encoding/binary

model:
1. data
2. algorithm
3. view
