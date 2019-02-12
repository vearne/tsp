

### 概述
使用遗传算法解决TSP问题
详细代码将example

经过试验发现，选择算法对结果的优化明显。
随机遍历抽样法效果最好，其次是竞标赛算法   

遗传算法的发现较优解的速度非常快，可以用来解决某些NP问题


![image_1d3dl5h72fps47oavmb2v5na9.png-164.4kB](http://static.zybuluo.com/woshiaotian/r2zrn5kvthi8r94zx1c1z9dr/image_1d3dl5h72fps47oavmb2v5na9.png)
### 1. 初始化种群

### 2. 评估种群中个体的适应度？
### 3. 选择
选择算法
1. 轮盘选择算法 roulette wheel selection
2. 锦标赛选择算法 tournament selection
3. 排序选择算法  rank selection
4. 档次轮盘法 Rank Based Wheel Selection

### 4. 交叉
交叉算法
1. 单点杂交
2. 多点杂交
3. 均匀杂交
4. 洗牌杂交

### 5. 变异
重复 步骤2 ~ 5 


### 路径序列与2进制基因码互相转换

把`路径`转换为`基因序列`
1. B-C-D 
2. 项对应的编号
3. 转换成数值
4. 转换成2进制(形如: "010")


把`基因序列`转换为`路径`
1. 2进制(形如: "010")
2. 转换为数值
3. 转换为项对应的编号
4. B-C-D 

test2.go中的结果可以直接使用R语言绘制路径图
### R语言生成图表 
```
data = read.table("/tmp/1.csv",header=F, sep=",");
traveller_way = read.table("/tmp/1.csv",header=TRUE, sep=",");
m <- ggplot(traveller_way, aes(posX, posY));
m + geom_path();
```
