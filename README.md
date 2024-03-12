#  CRAFTBIT
<img src="docs/images/logo.png" width="100" height="100" />
<br/>
Bitcoin Swiss Army Knife 🪛 🌕 🔧

## Abstract
This software is a lightweight Command Line Interface (CLI) containing multiple utilities designed for interacting with the Bitcoin ecosystem.  
These utilities are reusable and located within the `pkg` folder.  
Most of these utilities either serve as wrappers for [*btcd*](https://github.com/btcsuite/btcd) libraries or make calls to the [*mempool.space REST APIs*](https://mempool.space/docs/api/rest).  

> [!WARNING] 
> As certain functions leverage the APIs of the public instance of **mempool.space**, this tool may be suboptimal from a privacy perspective and inadvertently expose personally identifiable information (PII) such as transactions or addresses.


