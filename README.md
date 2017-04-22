# Tradeblock-API-scraper
## This is still very much alpha!  
  
simple tool for querying Tradeblock's Bitcoin tx history into .csv   
sum block/tx size, sum txs, sum fees, and total blocks are available to be outputted at 1h, 2h, 6h, and 24h intervals with three column format:  
timestamp(unix), value, excel converted date ("=(A1/24/60/60)+DATE(1970,1,1)"  

  
Requires Go! prog language (golang): https://golang.org/dl/  
Also requires some golang packages (which I'll specify later, sorry)  
  
# Usage:  
$ go run $GOPATH/SRC/github.com/test1/tbkcustom.go  
  
# available commands (self-explanatory usage):  
txsize_new, txes_new, txfee_new, blocks_new  
1h, 2h, 6h, 1d  
unix timestamp val (i.e. enter 1471327200 for 8/16/16 at 0600)
