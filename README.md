# countYTB
计算YouTube视频的统计信息

![统计信息](https://img2020.cnblogs.com/blog/1681961/202103/1681961-20210308095705849-1452220001.png)

|参数|说明|
|:---:|:---:|
|Video ID / sCPN|每个视频独特的 独有的ID / 用于识别问题的字符串（开发人员适用）|
|Viewport / Frames|当前播放窗口的分辨率 / 视频帧数的变化情况（由于主机性能的原因导致的丢帧数）|
|Current / Optimal Res|视频的解析度 / 最佳解析度|
|Volume / Normalized|当前音量百分比 / 实际输出音量百分比（与YouTube标准音量的响度差距）|
|Codecs|视频类型 / 格式|
|Color|视频色域|
|Host|为您推播流媒体的YouTube服务器主机名|
|Connection Speed|视频的加载速度，也就是大家常参考的数字，这个数字并不是很准确，由于YouTube使用的是小数据包、高频发送次数的调度方案，所以在延迟服务器上这个数字会很大。高延迟速率的链接速率是明显偏小的。（状态条为蓝色，加载无压力）|
|Network Activity|网络连接速度（状态条为蓝色，连接无压力。已经缓存完全后，为黑色）|
|Buffer Health|已缓存的视频时长（假如断开网络，我们还能观看的时长）|
|Mystery Text|字母“ s”后的数字代表播放器的状态代码（例如：s:4 --> video paused）|