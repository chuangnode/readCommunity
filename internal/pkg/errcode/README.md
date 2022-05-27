#错误码实现
***Learned***: [go项目开发实战](https://github.com/marmotedu/iam)  
##Code 设计规范
<hr>
Code代码从100101开始，1000以下为errors保留code<br/>
错误代码说明：100101<br>
1. 10:服务<br>
2. 01:模块<br>
3. 01:模块下的错误码序号,每个模块可以注册100个错误<br>

##服务和模块说明
<table>
  <tr>
    <th>服务</th>
    <th>模块</th>
    <th>说明</th>
  </tr>
  <tr>
    <th>10</th>
    <th>00</th>
    <th>通用-基本错误</th>
  </tr>
<tr>
    <th>10</th>
    <th>01</th>
    <th>通用-数据库类错误</th>
  </tr>
<tr>
    <th>10</th>
    <th>02</th>
    <th>通用-认证授权错误</th>
  </tr>
<tr>
    <th>10</th>
    <th>03</th>
    <th>通用-加解码类错误</th>
  </tr>
<tr>
    <th>11</th>
    <th>00</th>
    <th>用户相关(模块)错误</th>
  </tr>
<tr>
    <th>12</th>
    <th>00</th>
    <th>图书相关(模块)错误</th>
  </tr>
</table>


##错误记录规范
* 只在错误产生的最初位置打印日志，其他地方直接返回错误  
* 当代码调用第三方包的函数时,第三方包函数出错时打印错误信息







