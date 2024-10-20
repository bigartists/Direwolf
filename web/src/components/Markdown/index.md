---
nav:
  title: 组件
  path: /components
  order: 0
---

# Markdown

Markdown 渲染

```tsx
import { Markdown } from '@wair/components';

const text = '# 我是一级标题 \n ## 我是二级标题 \n ### 我是三级标题 \n #### 我是四级标题 \n ##### 我是五级标题 :citation[2]{href=https://baidu.com title=湖北唯一!黄陂木兰湖获批国家级旅游度假区 snippet=detail} \n\n > 找到了第 1篇资料: [武汉天气](https://baidu.com "https://baidu.com") \n\n >  :citation[1]{href=https://baidu.com title=湖北唯一!黄陂木兰湖获批国家级旅游度假区国家级旅游度假区国家级旅游度假区国家级旅游度假区国家级旅游度假区 snippet=电影《困兽》讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后讲述了千禧年前后，虎城动荡不安，政府为了重振秩序，降低犯罪率，决定铲除腐败的地下娱乐业...} \n\n ![image](https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/assets/txt2img_1.png)  \n\n "name=$2" 你可以使用以下的shell命令来检测特定的Docker容器名是否存在并且正在运行，如果满足条件，则停止并删除该容器：\n\n```bash\n#!/bin/bash\n\n# 容器名称\ncontainer_name=mycontainer\n\n# 检查容器是否存在且正在运行\nif docker ps -aqf "name=$2" >/dev/null; then\n  # 容器存在, 停止容器\n  docker stop $1\n  # 删除容器\n  docker rm $1\n  echo "Container \'$1\' has been stopped and removed."\nelse\n  echo "Container \'$1\' does not exist or is not running."\nfi\n```\n\n确保将`mycontainer`替换为你要检查的容器名称。\n\n这段脚本首先检查名为`mycontainer`的容器是否存在且正在运行。如果容器存在，脚本会先停止容器，然后再删除它。如果容器不存在或未运行，脚本会输出一条信息告知用户。\n\n在执行此脚本之前，请确保你有权限执行这些操作，并且理解这将永久删除容器及其所有数据。\n\n$$\n\\int_{-\\infty}^{\\infty} e^{-x^2} \\, dx = \\sqrt{\\pi}\n$$\n\n这段代码会在你的文档中产生一个著名的伽马函数积分结果 \n\n 好的，这里有一段 LaTeX 数学公式的示例，它展示了著名的爱因斯坦质能方程：\n\n\\[ E = mc^2 \\]\n\n以下是这段公式的详细解释：\n\n- \\( E \\): 能量\n- \\( m \\): 质量\n- \\( c \\): 光速\n\n这个公式表达了质量和能量之间的关系，表明质量和能量是可以相互转换的。\n\n如果你需要更复杂的数学公式，下面是一个展示二次方程求根公式的例子：\n\n\\[ x = \\frac{-b \\pm \\sqrt{b^2 - 4ac}}{2a} \\]\n\n这段公式用于求解标准形式为 \\( ax^2 + bx + c = 0 \\) 的二次方程，其中：\n\n- \\( a \\), \\( b \\), \\( c \\) 是常数\n- \\( x \\) 是未知数 \n\n ```$1 + 1 = 2$```'

export default () => <Markdown content={text} isGenerating={false} />;
```


### 自定义指令

```
:citation[index]{href=xxx title=xxx snippet=xxx}
```