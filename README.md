# Tradeblock-API-scraper
## This is still very much alpha!  
  
simple tool for querying Tradeblock's Bitcoin tx history (and mempool history, see below) into .csv   
sum block/tx size, sum txs, sum fees, and total blocks are available to be outputted at 1h, 2h, 6h, and 24h intervals with three column format:  
timestamp(unix), value, excel converted date ("=(A1/24/60/60)+DATE(1970,1,1)") 

Requires Go prog language (golang): https://golang.org/dl/   
  
## Usage:  
$ go run $GOPATH/SRC/github.com/test1/tbkcustom.go  
  
## available commands (self-explanatory usage):  
txsize_new, txes_new, txfee_new, blocks_new  
1h, 2h, 6h, 1d  
unix timestamp val (i.e. enter 1471327200 for 8/16/16 at 0600)


# Tradeblock mempool scraper
Gets the last week of mempool data, giving snapshots at 10 min intervals. Data is spit out into 8 columns and (typically) just under 1,000 rows. The columns aren't currently labeled, but they are: date/time (unix), tx bytes added to MP since last block found, tx bytes in MP since BEFORE last block, txs added to MP since last block, txs in MP since BEFORE last block, fees added to MP since last block found (in satoshis), fees added to MP since BEFORE last block, tx per min.     
So if you divide the sixth column (F in excel), by the second column (B in excel) you can get the fee per byte for all txs in MP that have been sent AFTER the last block was found. If you add columns F+G and B+C and do (F+G)/(B+C) the result is the fee per byte for all txs in the mempool at that moment. 
Basicaly this scraper is the data that is visualized here: https://tradeblock.com/bitcoin/. It can easily be modified to run API queries at weekly intervals if you want to just let it run in the background. 
Two points of caution: one, there are sometimes gaps in the data and the time will jump from 8/30 12:10 to 8/30 12:40 (or similar), it is rare but it does happen. Two, remember that Tradeblock is like any other node and will 'count' thier mempool differently. Some will hold on to txs for 72 hours with no min fee, others will only count txs as in the mempool fo 48 hours and only if the fee is >0. 

## Usage
$ go run $GOPATH/SRC/github.com/test1/tbkmempool.go (your location may vary of course)
Data is then automatically parsed out and, unlike the above tool, will actually display the data in your terminal rather than the "looks good" message.
