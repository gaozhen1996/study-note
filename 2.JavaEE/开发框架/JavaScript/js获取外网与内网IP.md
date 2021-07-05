# 1、外网IP与局域网IP
1. 公网ip具有世界范围的唯一性，而内网ip只在局域网内部具有唯一性。并且，一个局域网里所有电脑的内网IP是互不相同的,但共用一个外网IP。就像我们前面所说的你所在学校的校名在整个世界上只有一个，但是你学校里面的A栋大楼3层3号教室只有在你的校园内部才具有唯一性。别的学校也有A栋大楼3层3号教室。你只能跟快递小哥说请帮我把包裹送到xx大学，而不能说请帮我把包裹送到A栋大楼3层3号教室。

2. 在局域网中，每台电脑都可以自己分配自己的IP，但是这个IP只在局域网中有效。而如果你将电脑连接到互联网，你的网络提供商的服务器会为你分配一个IP地址，这个IP地址才是你在外网的IP。两个IP同时存在，一个对内，一个对外。

3. 互联网上的IP（即外网IP）地址统一由一个叫“IANA”(InternetAssigned NumbersAuthority，互联网网络号分配机构)的组织来管理。由于分配不合理以及IPv4协议本身存在的局限，现在互联网的IP地址资源越来越紧张。IANA将A、B、C类IP地址的一部分保留下来，留作局域网使用。具体如下

**IP地址空间：**
- a类网
10.0.0.0~10.255.255.255
- b类网
172.16.0.0~172.31.255.255
- c类网
192.168.0.0~192.168.255.255


# 2、获取局域网IP
```
function getUserIP(onNewIP) { //  onNewIp - your listener function for new IPs
      //compatibility for firefox and chrome
      var myPeerConnection = window.RTCPeerConnection || window.mozRTCPeerConnection || window.webkitRTCPeerConnection;
      var pc = new myPeerConnection({
         iceServers: []
     }),
     noop = function() {},
     localIPs = {},
     ipRegex = /([0-9]{1,3}(\.[0-9]{1,3}){3}|[a-f0-9]{1,4}(:[a-f0-9]{1,4}){7})/g,
     key;
 
     function iterateIP(ip) {
         if (!localIPs[ip]) onNewIP(ip);
         localIPs[ip] = true;
    }
 
      //create a bogus data channel
     pc.createDataChannel("");
 
     // create offer and set local description
     pc.createOffer().then(function(sdp) {
         sdp.sdp.split('\n').forEach(function(line) {
             if (line.indexOf('candidate') < 0) return;
             line.match(ipRegex).forEach(iterateIP);
         });
         
         pc.setLocalDescription(sdp, noop, noop);
     }).catch(function(reason) {
         // An error occurred, so handle the failure to connect
     });
 
     //sten for candidate events
     pc.onicecandidate = function(ice) {
         if (!ice || !ice.candidate || !ice.candidate.candidate || !ice.candidate.candidate.match(ipRegex)) return;
         ice.candidate.candidate.match(ipRegex).forEach(iterateIP);
     };
}
 
// Usage
 
getUserIP(function(ip){
     alert("Got IP! :" + ip);
});
```

# 2、获取外网IP
调用搜狐的接口，返回真实的IP

```
http://pv.sohu.com/cityjson?ie=utf-8
```
