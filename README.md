A Chinese Tour of Go
====================

安裝與執行請先參考這篇文章：<a href="http://imazole.wordpress.com/2013/12/03/go-lang-part1/" target="_blank">Go lang 初體驗 – 安裝與執行</a>

在自己看了一些範例後，覺得 go 像是 C, C++, Java/C#, Javascript 等語言的綜合體，然後再加上一些新的做法，很有趣。        
1. go 有 pointer，但沒有 pointer 運算。       
2. go 有 closure。           
3. go 有 garbage collection，但做法又略有不同。         
4. go 有 interface，但也跟 java/C# 的不一樣，它不需要特別聲明實作，但又能有編譯時的靜態檢查。        
5. go 有 package，但沒有命名空間；go 有 struct，但沒有 class，可是他的 struct 又可以掛載方法，但沒有繼承。      
6. go 有錯誤處理，是透過一個預先定義的 error interface，而非傳統意義上的異常機制。    
7. go 有 concurrency，但它共享共存的方式是透過溝通。    
> The Go Approach: Don't communicate by sharing memory, share memory by communicating.


學習材料：      
1. <a href="http://tour.golang.org/" target="_blank">A Tour of Go</a>     
2. <a href="http://www.slideshare.net/c9s/happy-gopart1" target="_blank">Happy Go Programming Part I</a>      
3. <a href="http://www.slideshare.net/c9s/happy-go-programming-part-2" target="_blank">Happy Go Programming Part II</a>      
4. <a href="http://tour.golang.org/#72" target="_blank">官方Tour最後提供的資源</a>  

中文學習材料：      
1. <a href="http://blog.jobbole.com/36480/" target="_blank">Go在谷歌：以软件工程为目的的语言设计</a>     


資源：    
1. <a href="http://golang.org/pkg/strings/" target="_blank">strings</a>      
2. <a href="http://www.gorillatoolkit.org/" target="_blank">Gorilla</a>       
3. <a href="https://github.com/codegangsta/martini" target="_blank">Martini</a>      
4. <a href="https://github.com/paulbellamy/mango" target="_blank">Mango</a>    
5. <a href="http://golang.org/cmd/cgo/" target="_blank">cgo</a> 

