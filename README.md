# stressContainerController
在容器云的环境下，我们有了另一种构建压力源的形式，即把压力机放到云上，以容器为单位生成压力（一个容器可以只包含一个压力线程，也可以包含一个压力线程组）。 最近试着用go语言写了一款容器级的压测场景调度器：stressContainerController。功能类似于LoadRunner controller组件的scenario schedule：
![image](https://github.com/wkqyxyh/stressContainerController/blob/master/images/1.jpg)
stressContainerController使用方法很简单，在配置文件中定义docker镜像名、增压时长、持续时长、减压时长、目标压力值：
![image](https://github.com/wkqyxyh/stressContainerController/blob/master/images/2.jpg)

之后运行可执行文件./stressContainerController即可。stressContainerController会以容器为单位，逐步启动压力机容器直到形成足够的压力，持续一定时间后，再逐步停止压力。
