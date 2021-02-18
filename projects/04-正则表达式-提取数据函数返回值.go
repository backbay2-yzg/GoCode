package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `<div class="main_father clearfix d-flex justify-content-center" style="height:100%;">
<div class="container clearfix" id="mainBox">
	<div  class='space_container'></div>
	<main>
	<div>
	<div>
	<div>
	<div>
	<h1 class="title-article">可视化工具D3教程</h1>
	</div>
	<div>
	<div>
	<!--文章类型-->
	<span class="article-type type-1 float-left">原创</span>                                                                                                                                            <a class="follow-nickName" href="https://me.csdn.net/qq_31052401" target="_blank" rel="noopener">Lelliam</a>
	<span class="time">最后发布于2019-06-27 20:24:29                    </span>
	<span class="read-count">阅读数 3893</span>
	<a id='blog_detail_zk_collection' data-report-click='{"mod":"popu_823"}'>
	<svg class="icon">
	<use xlink:href="#icon-csdnc-Collection-G" ></use>
</svg>
	收藏
	</a>
	</div>
	<div>发布于2019-06-27 20:24:29</div>
	<div>
	<div>
	<span class="creativecommons">
	<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"></a>
	<span>
		版权声明：本文为博主原创文章，遵循<a href="http://creativecommons.org/licenses/by-sa/4.0/" target="_blank" rel="noopener"> CC 4.0 BY-SA </a>版权协议，转载请附上原文出处链接和本声明。                            </span>
	<div>
1231121
4564465
		本文链接：<a href="https://blog.csdn.net/qq_31052401/article/details/93786425">https://blog.csdn.net/qq_31052401/article/details/93786425</a>
	</div>
	</span>
	</div>
	</div>
	<div>
	<a class="href-article-edit slide-toggle">展开</a>
	</div>
	</div>
	</div>
	</div>`
	//ret:=regexp.MustCompile(`a[^0-9a-z]c`)
	//ret:=regexp.MustCompile(`<div>(.*)</div>`)
	ret := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	all := ret.FindAllStringSubmatch(str, -1)
	//fmt.Println("all is : ",all)
	for _, list := range all {
		fmt.Println("1***", list[0])
		fmt.Println("2***", list[1])
	}
}
